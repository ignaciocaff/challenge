FROM golang:1.21-alpine

ARG environment=docker

WORKDIR /

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY .env.${environment} .env

RUN go build -o /chat-api /cmd/chat/

CMD [ "/chat-api" ]
