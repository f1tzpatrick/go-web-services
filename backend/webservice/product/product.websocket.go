package product

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

type message struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

func productSocket(ws *websocket.Conn) {

	done := make(chan struct{})
	fmt.Println("New websocket connection")

	go func(c *websocket.Conn) {
		for {
			var msg message
			if err := websocket.JSON.Receive(ws, &msg); err != nil {
				log.Println(err)
				break
			}
			fmt.Printf("Received Message: %s\n", msg.Data)
		}
		close(done)
	}(ws)
loop:
	for {
		select {
		case <-done:
			fmt.Println("Connection was Closed")
			break loop
		default:
			products, err := GetTopTenProducts()
			if err != nil {
				log.Println(err)
				break
			}
			if err := websocket.JSON.Send(ws, products); err != nil {
				log.Println(err)
				break
			}
			time.Sleep(10 * time.Second)
		}
	}
	fmt.Println("closing the connection")
	defer ws.Close()
}
