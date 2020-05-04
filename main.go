package main

import (
  "os"
  "fmt"
  "log"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/line/line-bot-sdk-go/linebot"
)
func main() {
  bot, err := linebot.New(
    os.Getenv("LINEBOT_CHANNEL_SECRET"),
    os.Getenv("LINEBOT_CHANNEL_TOKEN"),
  )
  if err != nil {
    log.Fatal(err)
  }
  router := gin.Default()

  router.GET("/hello", func(c *gin.Context) {
    c.String(http.StatusOK, "Hello World!!")
  })

  router.POST("/callback", func(c *gin.Context) {
    events, err := bot.ParseRequest(c.Request)
    if err != nil {
			log.Print("ここまではきてる")
			log.Print(err)
      if err == linebot.ErrInvalidSignature {
        c.Writer.WriteHeader(400)
      } else {
        c.Writer.WriteHeader(500)
      }
      return
    }
    for _, event := range events {
      if event.Type == linebot.EventTypeMessage {
        switch message := event.Message.(type) {
          case *linebot.TextMessage:
            if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
              log.Print(err)
            }
          case *linebot.StickerMessage:
            replyMessage := fmt.Sprintf(
              "sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
            if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
              log.Print(err)
            }
        }
      }
    }
  })
  router.Run(":" + os.Getenv("PORT"))
}