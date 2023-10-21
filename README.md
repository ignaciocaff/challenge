# Table of Contents

- [General Considerations](#general-considerations)``
- [Dockerization](#dockerization)
- [Bonus](#bonus)
- [Backend](#backend)
- [Bot](#bot)
- [Frontend](#frontend)

# General Considerations

### Important

* The easiest way to set up the entire project is with Docker. If you don't have it, you can skip that section entirely.

* No user or room creation were developed. For the sake of the challenge, I did not consider them necessary due to their low complexity. That's why they are inserted automatically through a script when the MongoDB container starts. If you're not using Docker to start the MongoDB and Rabbit containers and you have a local instance, you must insert the documents into 'users' and 'rooms' collecitions manually both of them in the 'jobsity' database or de MONGO_DB_NAME you set in the .env file.

* The bot is not a user, it is a service that is listening to the RabbitMQ queue. It is not necessary to create it as a user in the database.

* There is no log files in the project. The logs are printed in the console.

* If you use /stock command, the bot will send a message to the chat room with the stock information. If the stock does not exist, the bot will send a message to the chat room indicating that the stock does not exist and all the room users would receive the message. You will not see the message /stock but you will see the bot response

* The /stock command is case-sensitive. If you use it in uppercase, it won't work, and this is intentional. I could have validated or used 'to lower,' but the challenge was explicit with '/stock'.

* It was decided to handle authentication with the session + cookie approach, as opposed to using tokens. This way, the frontend is completely agnostic to the login process. All it needs to do is make requests to the /auth and /me URLs, but apart from that, it doesn't know what's happening. A simple in-memory session was implemented to identify if it's valid or not. For the challenge's purposes, the implementation was simplified

* Due to time constraints, message encryption before storing them in the database was not implemented. It is a feature that should be implemented in a real-world scenario 

* Password encryption at the time of logging in was not implemented. It is a feature that should be implemented in a real-world scenario

- Frontend
    - Technologies used: Angular 16, TailwindCSS, DaisyUI, Socket.io-client
    
- API
    - Technologies used: Go 1.21, MongoDB, RabbitMQ, Websockets

- Bot
    - Technologies used: Go 1.21, RabbitMQ

# Dockerization

Dockerfiles have been included in each of the projects to facilitate their execution in a local development environment. To run the containerized application, you should execute the following command at the project's root:

```bash
docker-compose up --build -d
```

1. It will start the necessary MongoDB and RabbitMQ services to run the project

2. It will automatically insert users and chat rooms into MongoDB.
Which will raise the api, the bot front containers on ports 8080 and 80 respectively.

3. Furthermore, it will start the front-end, API, and bot.

After executing the command, if everything started correctly, we should simply navigate to localhost on port 80

# Bonus

You can try the application running at the following link -

It was deployed on an AWS EC2 instance.

# Backend

## Usage

You have several ways to run it

### Makefile

The following command should always be executed at the root of the repository

To run:

```bash
make run-chat-api
```

The chat-api will be available at  http://localhost:3000

### Manual

The following commands should always be executed at the root of the repository

To run:

```bash
cd chat-api
go run ./cmd/chat/main.go
```

The chat-api will be available at  http://localhost:3000

## Test

You have several ways to run it

### Makefile

The following command should always be executed at the root of the repository

To run:

```bash
make test-services-ws
```

### Manual

The following commands should always be executed at the root of the repository

To run:

```bash
cd chat-api
go test ./services/ws
```

## Structure

Below, I provide an incomplete but representative overview of the bot structure

```bash
chat-api
├── cmd                 #  entry point of the application.                 
├── database            #  Database configuration and migrations.
├── env                 #  Environment variables for the application.     
├── handlers            #  Handlers for the application.
├── services            #  Services that are used in the application WS, HTTP and RabbitMQ.
├── utils               #  Session manager utility.
├── server              #  Main configuration of the server and middlewares
├── Dockerfile          #  Instructions to build a Docker image.    
```

# Bot

## Usage

You have several ways to run it

### Makefile

The following command should always be executed at the root of the repository

To run:

```bash
make run-bot
```

The bot will be running in the background or in the console where you execute the command. The message that indicates it has started successfully is as follows:

```bash
[*] Waiting for messages. To exit, press CTRL+C.
```

### Manual

The following commands should always be executed at the root of the repository

To run:

```bash
cd bot
go run ./cmd/bot.go
```

The bot will be running in the background or in the console where you execute the command. The message that indicates it has started successfully is as follows:

```bash
[*] Waiting for messages. To exit, press CTRL+C.
```

## Structure

Below, I provide an incomplete but representative overview of the bot structure

```bash
bot
├── cmd                 #  entry point of the application.                 
├── env                 #  Environment variables for the application.     
├── services            #  Services that are used in the application (Stooq api and RabbitMQ)    
├── Dockerfile          # Instructions to build a Docker image.    
```

# Frontend

## Usage

The first thing you should do is install the npm dependencies for Angular.

Then you have several ways to run it

### Makefile

The following commands should always be executed at the root of the repository

To install dependencies:

```bash
make install-app
```

To run:

```bash
make run-chat-app
```

The front end will be available at  http://localhost:4200

### Manual

The following commands should always be executed at the root of the repository

To install:

```bash
cd chat-app
npm run i
```

To run:

```bash
cd chat-app
npm run dev
```

The front end will be available at  http://localhost:4200

## Structure

Below, I provide an incomplete but representative overview of the front-end structure

```bash
chat-app
├── config                 #  Configuration files for the application.                  
├── src     
│   ├── app      
│        ├── core          #  Core components that are used in the application.
│        ├── modules       #  Modules that are used in the application.
│        ├── routes        #  Routing configuration for the applicationn.
│   ├── assets 
│        ├── fonts              
│        ├── images              
│   ├── envinronments  
├── Dockerfile             # Instructions to build a Docker image.    
├── tailwind.config.js     # Tailwind configuration file with it's styles     
```
