package main

import (
	"client/handle"
	"client/model"
	"time"
)

func main() {
	Alice := model.Student{Id: "202315155533", Name: "Alice", Gender: "Female", Age: 20, Credit: 107}
	Handle := &handle.StudentHandle{}
	for i := 0; i < 10; i++ {
		err := Handle.SendMessageToServer(Alice, "ws://localhost:8080/ws")
		if err != nil {
			time.Sleep(time.Second * 30)
		} else {
			time.Sleep(time.Second * 60)
		}
	}
}
