module github.com/eas-kaine/discord-bot

go 1.18

require (
	github.com/bwmarrin/discordgo v0.25.0
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.5
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b // indirect
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68 // indirect
)

replace github.com/bwmarrin/discordgo => github.com/Enterprise-Automation/discordgo v0.25.0
