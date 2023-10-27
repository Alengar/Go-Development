package main

import (
	"github.com/gin-gonic/gin"
	"goDeveloping/lecture_9/calendarEventRepository/internal/app/event"
	"goDeveloping/lecture_9/calendarEventRepository/internal/handler"
	"net/http"
)

func main() {
	r := gin.Default()

	eventRepo := event.NewRepository()
	eventUC := event.NewEventUseCase(eventRepo)

	eventHandler := handler.NewEventHandler(eventUC)

	r.LoadHTMLGlob("html/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "layout", gin.H{
			"Title":   "Welcome to Event Calendar",
			"Content": "Welcome to the Event Calendar!",
		})
	})

	r.GET("/events", eventHandler.ListEvents)
	r.POST("/events", eventHandler.CreateEvent)

	r.Run(":8080")
}
