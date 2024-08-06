package handler

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/iterator"
	"net/http"
)

type MessageHandler struct {
	client *firestore.Client
}

type MessageHandlerInterface interface {
	SendMessage(c echo.Context)
	GetMessages(c echo.Context)
}

type MessageRequest struct {
	User      string `json:"user"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
}

func NewMessage(client *firestore.Client) *MessageHandler {
	return &MessageHandler{
		client: client,
	}
}

func (m *MessageHandler) SendMessage(c echo.Context) {
	var msg MessageRequest
	if err := c.Bind(&msg); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, _, err := m.client.Collection("messages").Add(context.Background(), msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "status : MessageRequest sent")
}

func (m *MessageHandler) GetMessages(c echo.Context) {
	iter := m.client.Collection("messages").OrderBy("timestamp", firestore.Asc).Documents(context.Background())
	var messages []MessageRequest
	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		var msg MessageRequest
		doc.DataTo(&msg)
		messages = append(messages, msg)
	}
	c.JSON(http.StatusOK, messages)
}
