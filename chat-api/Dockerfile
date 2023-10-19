FROM golang:1.21-alpine

ARG environment=docker

WORKDIR /

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download


COPY . .

COPY .env.${environment} .env

RUN go build -o /chat-api /cmd/chat/

CMD [ "/chat-api" ]