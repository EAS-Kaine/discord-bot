# discord-bot

Tasked with producing microservices, which utilise a database for storage, this repository holds a Discord bot to interface between the Discord server and game APIs.

## Usage

```Golang
    go run main.go
```

## Expected API response format

```
// Expected validation response on /api/validation/appname,command
{"status":"success","status_message":"valid_command"}

// Expected successful command response on /api/command/appname,command
{"status": "success", "status_message": "insert status message here", "discord_message": "message for discord here"}
```

## NOTES

## Features

- [x] Concurrently handle connections

- [x] Dockerfile

- [x] Handle syscall.ECONNREFUSED

- [x] Roles included in request headers

- [x] Callbacks

- [x] MySQL DB (Containerised?)

- [ ] User permissions

- [x] adding/listing/removing actions from the db with user group permissions

## Diagram

![Task diagram](img/disc-bot-diagram.png "Task diagram")
