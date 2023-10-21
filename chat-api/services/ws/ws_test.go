package ws

import (
	"chatjobsity/env"
	"chatjobsity/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/mongo"
)

var mockMessage = "{\"id\":\"testroom\",\"text\":\"hello testing\",\"sender\":{\"username\":\"testclient\",\"id\":\"testclient\"},\"roomId\":\"testroom\",\"date\":\"2023-10-20T17:52:42.9212761-03:00\"}"

var envApp = env.EnvApp{
	MongoDbName: "jobsity",
	BotQueue:    "bot_queue",
}

/*
This test checks behavior of the room with clients with and without messages
*/
func TestRoomBehaviorAndConnections(t *testing.T) {
	// Create room
	room := NewRoom("testroom", nil, envApp)
	go room.Start()

	// Create ws server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgradeToWebSocket(w, r)
		if err != nil {
			t.Errorf("Error upgrading to websocket: %v", err)
			return
		}
		setInitialStage(conn, "testclient", room)
	}))
	// Join 3 client to the websocket server without sending any message
	for i := 0; i < 3; i++ {
		_, _, err := websocket.DefaultDialer.Dial(strings.Replace(server.URL, "http://", "ws://", 1), nil)
		if err != nil {
			t.Errorf("Error trying top connect: %v", err)
			return
		}
	}

	// Join 1 client to the websocket server and send a message
	conn, _, err := websocket.DefaultDialer.Dial(strings.Replace(server.URL, "http://", "ws://", 1), nil)
	if err != nil {
		t.Errorf("Error trying to connect: %v", err)
		return
	}
	message := []byte(mockMessage)
	err = conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		t.Errorf("Error sending the message: %v", err)
		return
	}
	if len(room.clients) != 4 {
		t.Errorf("Expected 4 clients, but got %d", len(room.clients))
		return
	}
}

func setInitialStage(conn *websocket.Conn, clientId string, room *Room) {
	client := NewClient(clientId, conn, room, MongoClient(), envApp)
	room.join <- client
	go client.Write()
	client.Read()
}

func upgradeToWebSocket(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	// Upgrade a WebSocket
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

/*
This test checks the integrity of a WebSocket message, verifying that it is correctly transmitted and received,
including its text, sender information, room ID, and timestamp.
*/
func TestMessageIntegrity(t *testing.T) {
	srv := createWsServer(http.HandlerFunc(handlerToBeTested))

	conn, err := createWsClient(srv.URL)
	if err != nil {
		t.Fatalf("Failed to establish WebSocket connection: %v", err)
	}
	parsedTime, _ := time.Parse(time.RFC3339Nano, "2023-10-20T17:52:42.9212761-03:00")
	messageData := models.Message{
		Text: "hello testing",
		Id:   "testroom",
		Sender: models.Sender{
			Id:       "testclient",
			Username: "testclient",
		},
		RoomId: "testroom",
		Date:   parsedTime,
	}
	sendMessageAndValidateResponse(t, conn, messageData, mockMessage)
}

func createWsServer(handler http.HandlerFunc) *httptest.Server {
	srv := httptest.NewServer(handler)
	return srv
}

func createWsClient(serverURL string) (*websocket.Conn, error) {
	u, _ := url.Parse(serverURL)
	u.Scheme = "ws"
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	return conn, err
}

func sendMessageAndValidateResponse(t *testing.T, conn *websocket.Conn, messageData interface{}, expectedResponse string) {
	messageBytes, err := json.Marshal(messageData)
	if err != nil {
		t.Fatalf("Failed to serialize JSON message: %v", err)
	}

	err = conn.WriteMessage(websocket.BinaryMessage, messageBytes)
	if err != nil {
		t.Fatalf("Failed to write the message: %v", err)
	}

	_, p, err := conn.ReadMessage()
	if err != nil {
		t.Fatalf("Failed to read the message: %v", err)
	}

	if string(p) != expectedResponse {
		t.Fatalf("Unexpected response. Expected %q, but received %q", expectedResponse, string(p))
	}
}

func handlerToBeTested(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade the connection", http.StatusInternalServerError)
		return
	}
	room := NewRoom("testroom", nil, envApp)
	go room.Start()
	client := NewClient("testclient", conn, room, MongoClient(), envApp)
	room.join <- client
	defer func() { room.leave <- client }()
	go client.Write()
	client.Read()
}

func MongoClient() *mongo.Client {
	return &mongo.Client{}
}
