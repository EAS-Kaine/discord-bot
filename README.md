# discord-bot

Tasked with producing microservices, which utilise a database for storage, this repository holds a Discord bot to interface between the Discord server and game APIs.

## Usage

```Golang
    go run main.go
```

## NOTES

- roles, username sent in every request header

## MVP

- [x] MySQL DB (Containerised?)

- [ ] User permissions

- [ ] standard syn/ack stuff

- [ ] adding/listing/removing actions from the db with user group permissions

## Diagram

![Task diagram](img/disc-bot-diagram.png "Task diagram")
