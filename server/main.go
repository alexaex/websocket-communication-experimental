package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		for {
			// 从WebSocket连接读取消息
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}

			// 打印收到的JSON消息
			fmt.Printf("Received: %s\n", msg)

			// 对msg进行处理 比如json反序列化 json->Object

			// 发送消息到client
			// responseMsg := []byte("Hello from server")
			// err = conn.WriteMessage(websocket.TextMessage, responseMsg)
			// if err != nil {
			//     fmt.Println(err)
			//     return
			// }
		}
	})

	r.Run(":8080")
}
