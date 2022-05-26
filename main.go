package main

import (
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

		res, err := SendSMS(ctx, SMS{
			Mobile:    data.From,
			Text:      "Hi from API",
			ShortCode: "28009",
		})

		if err != nil {
			log.Println("error", err.Error())
			return
		}

		log.Println(&res)

		prompt := "Jackson is the the child of Jane Juma. Jackson has left school at 1:20am. He is in bus with number plate KLC2393 with driver Wycliffe. Jackson is using route North. Jackson is arriving at home at 2:30 am.\n\nQuestion: When is her coming?\nChat bot:"

		ans, err := QueryOpenAI(ctx, prompt)
		if err != nil {
			log.Println("error", err.Error())
		}

		ansText := ans.Choices[0].Text

		ctx.String(http.StatusOK, ansText)
	})

	// get api callback from chat AI
	r.POST("/ai-response", func(ctx *gin.Context) {

	})

	r.Run(":" + port)
}
