package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//cors ...
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, HEAD, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(cors())

	port := GetEnv("PORT", "8080")

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello chatbot")
	})

	// get message from AT
	r.POST("/new-sms", func(ctx *gin.Context) {
		data := IncomingSMS{}
		if err := ctx.Bind(&data); err != nil {
			log.Println("error", err.Error())
			return
		}

		prompt, err := GetPrompt(ctx, data.From)
		if err != nil {
			log.Println("error", err.Error())
		}

		ans, err := QueryOpenAI(ctx, fmt.Sprintf(prompt, data.Text))
		if err != nil {
			log.Println("error", err.Error())
		}

		ansText := ans.Choices[0].Text

		res, err := SendSMS(ctx, SMS{
			Mobile:    data.From,
			Text:      ansText,
			ShortCode: "28009",
		})

		if err != nil {
			log.Println("error", err.Error())
			return
		}

		log.Println(&res)

		ctx.String(http.StatusOK, ansText)
	})

	// get api callback from chat AI
	r.POST("/ai-response", func(ctx *gin.Context) {

	})

	r.Run(":" + port)
}
