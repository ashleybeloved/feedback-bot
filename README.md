# Feedback Bot — simple bot for collecting messages
A lightweight Telegram bot for receiving messages via the bot to the administrator. It's built in Go using the telego library. English localization only.

## Commands

Any non-command user message replies to admin. 
  
`/start` - An introductory command for the user about what the bot does.  
  
Admin Routes:

`/ban <ID>` and `/unban <ID>` - ban and unban users in bot.  

And admin have inline buttons under any message "Reply" and "Ban".


## Configuration

.env file:

```
# Bot
BOT_TOKEN=123456789:yOuR_ToKeN

# Admin
TELEGRAM_USERNAME=your_username
TELEGRAM_ID=your_id
```

## Quick Start

### Prerequisites

- Go 1.20+  
- Telegram Bot Token (@BotFather)  
- `go get github.com/mymmrac/telego`  
- `go get github.com/joho/godotenv`

### Installation

- Clone the repository  
`git clone https://github.com/ashleybeloved/feedback-bot.git`  
`cd feedback-bot`

- Create environment file  
`nano .env`

- Run the bot  
`go run main.go`

##### If you have any issues with bot you can message me in telegram: [bot](https://t.me/hback_bot) or [@fuckcensor](https://t.me/fuckcensor) 
