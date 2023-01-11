#
#
#
##  Make movies crud using go

### Before starting the web server run the following command:
```sh
go mod init github.com/username/slack-age-bot
go mod tidy
go get github.com/shomali11/slacker
```

### To start web server, run :
```sh
go build
go run main.go 
```

### You can call the following routes in postman and see results 

```go
	    os.Setenv("SLACK_BOT_TOKEN", "xoxb-*******...")
	    os.Setenv("SLACK_APP_TOKEN", "xapp-1-****...")

```
#### add bot in message like this:
```sh
@age-bot my yob is 1996
```
Enjoy it :)
