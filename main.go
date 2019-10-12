package main

import (
	"backend/websocket"
	"fmt"
	"net/http"
	"strings"
)

// 定义 WebSocket 服务处理函数
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	// 将连接更新为 WebSocket 连接
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	id := strings.Split(r.RequestURI, "=")[1]
	client := &websocket.Client{
		ID: id,
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()

}


func setupRoutes()  {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})

	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		serveWs(pool, writer, request)
	})
}

func main()  {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
