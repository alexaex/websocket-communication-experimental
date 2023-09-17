package handle

import (
	"client/model"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

type StudentHandle struct{}

func (handle *StudentHandle) SendMessageToServer(student model.Student, link string) error {
	// 解析url
	Url, err := url.Parse(link)
	if err != nil {
		log.Println("解析url链接失败.")
		return err
	}
	log.Println("url链接解析成功.")

	// 建立websocket链接
	conn, _, err := websocket.DefaultDialer.Dial(Url.String(), nil)
	if err != nil {
		log.Println("建立websocket连接失败.")
		return err
	}
	log.Println("建立websocket连接成功.")

	// 数据序列化
	StudentInfo, err := json.Marshal(student)
	if err != nil {
		log.Println("数据序列化失败.")
		return err
	}
	log.Println("数据序列化成功.")

	// 延时关闭websocket连接
	defer conn.Close()

	//传递websocket数据
	err = conn.WriteJSON(string(StudentInfo))
	if err != nil {
		log.Println("数据传输失败.")
	}
	time.Sleep(10 * time.Second)
	log.Println("数据传输成功.\n")

	// 读取来自服务器的数据
	// conn.ReadMessage()...
	return nil
}
