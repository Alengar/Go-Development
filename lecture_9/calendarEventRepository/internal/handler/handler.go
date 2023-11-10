package handler

import (
	"calendarEventManagement/internal/app/event"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

type EventHandler struct {
	eventUseCase  event.UseCase
	htmlTemplates map[string]*template.Template
}

func NewEventHandler(eventUC event.UseCase) *EventHandler {
	htmlTemplates := make(map[string]*template.Template)
	htmlTemplates["layout"] = template.Must(template.ParseFiles("html/layout.html"))
	htmlTemplates["event_list"] = template.Must(template.ParseFiles("html/event_list.html"))
	htmlTemplates["create_event"] = template.Must(template.ParseFiles("html/create_event.html"))

	return &EventHandler{
		eventUseCase:  eventUC,
		htmlTemplates: htmlTemplates,
	}
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	// Handle event creation, similar to previous examples
}

func (h *EventHandler) ListEvents(c *gin.Context) {
	events, err := h.eventUseCase.GetEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Render the "event_list" template with the events data
	c.HTML(http.StatusOK, "layout", gin.H{
		"Title":   "Event List",
		"Content": "event_list",
		"Events":  events,
	})
}
