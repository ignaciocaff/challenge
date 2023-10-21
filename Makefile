.PHONY: run-bot
run-bot:
	cd ./bot && go run ./cmd/bot.go

.PHONY: run-chat-api
run-chat-api:
	cd ./chat-api && go run ./cmd/chat/main.go

.PHONY: run-chat-app
run-chat-app:
	cd ./chat-app && npm run dev

.PHONY: install-app
install-app:
	cd ./chat-app && npm i

.PHONY: test-services-ws
test-services-ws:
	cd ./chat-api && go test ./services/ws