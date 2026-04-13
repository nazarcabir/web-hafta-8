package handlers

import (
    "log"
    "net/http"
    "sync"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

type Client struct {
    Conn     *websocket.Conn
    CourseID string
}

var clients = make(map[*Client]bool)
var mu sync.Mutex

// @Summary Sınıf Sohbetine/Canlı Yayına Katıl
// @Tags WebSocket
// @Security BearerAuth
// @Param courseId path string true "Kurs ID"
// @Router /ws/classroom/{courseId} [get]
func ClassroomWS(c *gin.Context) {
    courseID := c.Param("courseId")
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Println("Upgrade error:", err)
        return
    }
    defer conn.Close()

    client := &Client{Conn: conn, CourseID: courseID}
    mu.Lock()
    clients[client] = true
    mu.Unlock()

    defer func() {
        mu.Lock()
        delete(clients, client)
        mu.Unlock()
    }()

    for {
        var msg map[string]interface{}
        if err := conn.ReadJSON(&msg); err != nil {
            break
        }
        mu.Lock()
        for c := range clients {
            if c.CourseID == courseID {
                c.Conn.WriteJSON(msg)
            }
        }
        mu.Unlock()
    }
}
