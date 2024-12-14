package api

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type ConnectionManager struct {
	connections map[*websocket.Conn]bool
	mutex       sync.RWMutex
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[*websocket.Conn]bool),
	}
}

func (m *ConnectionManager) AddConnection(conn *websocket.Conn) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.connections[conn] = true
}

func (m *ConnectionManager) RemoveConnection(conn *websocket.Conn) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.connections, conn)
}

func (m *ConnectionManager) Broadcast(message []byte) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for conn := range m.connections {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("Error broadcasting to connection: %v", err)
			conn.Close()
			delete(m.connections, conn)
		}
	}
}
