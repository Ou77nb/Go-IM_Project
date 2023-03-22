package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// Message 消息
type Message struct {
	gorm.Model
	UserId     int64  //发送者
	TargetId   int64  //接受者
	Type       int    //发送类型  1私聊  2群聊  3心跳
	Media      int    //消息类型  1文字 2表情包 3语音 4图片 /表情包
	Content    string //消息内容
	CreateTime uint64 //创建时间
	ReadTime   uint64 //读取时间
	Pic        string
	Url        string
	Desc       string
	Amount     int //其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}

type Node struct {
	Conn      *websocket.Conn //连接
	Addr      string          //客户端地址
	DataQueue chan []byte     //消息
	GroupSets set.Interface   //好友 / 群
}

//映射关系
var clientMap = make(map[int64]*Node, 0)

//读写锁
var rwLocker sync.RWMutex

//	需要:发送者ID 、接受者ID 、消息类型、发送的内容、发送类型
func Chat(writer http.ResponseWriter, request *http.Request) {
	//获取请求体
	query := request.URL.Query()
	//获取发送者ID
	Id := query.Get("userId")
	userId, _ := strconv.ParseInt(Id, 10, 64)
	//设置token为true
	isvalida := true
	conn, err := (&websocket.Upgrader{
		//token 校验
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println("message.go 65:", err)
		return
	}
	//获取conn
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	//用户关系
	//userid 跟 node绑定 并加锁 标明信息的发送者
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()
	//完成发送逻辑
	go sendProc(node)
	//完成接受逻辑
	go receiveProc(node)
	sendMsg(userId, []byte("欢迎进入聊天系统"))

}

//发送消息
func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue: //取出信息主体内容
			fmt.Println("[ws]sendProc >>>> msg :", string(data))
			//将信息发送到局域网
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println("message.go 95:", err)
				return
			}
		}
	}
}

//接收消息
func receiveProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println("message.go 106:", err)
			return
		}
		msg := Message{}
		err = json.Unmarshal(data, &msg)
		if err != nil {
			fmt.Println("message.go 113:", err)
		}
		dispatch(data)
		//把消息保存到切片
		broadMsg(data)
		fmt.Println("[ws] recvProc <<<<< ", string(data))
	}
}

var udpSendChan = make(chan []byte, 1024)

func broadMsg(data []byte) {
	udpSendChan <- data
}

//init 函数会在程序执行开始的时候被调用。
func init() {
	go udpSendProc()
	go udpReceiveProc()
}

//完成udp数据发送协程
func udpSendProc() {
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 0, 255),
		Port: 3000,
	})
	defer con.Close()
	if err != nil {
		fmt.Println("message.go 142:", err)
	}

	fmt.Println("udpSendProc  data :", udpSendChan)
	for {
		select {
		case data := <-udpSendChan:

			_, err := con.Write(data)
			if err != nil {
				fmt.Println("message.go 150:", err)
				return
			}
		}
	}

}

//完成udp数据接收协程
func udpReceiveProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		fmt.Println("message.go 165:", err)
	}
	defer con.Close()
	for {
		var buf [512]byte
		n, err := con.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("udpReceiveProc  data :", string(buf[0:n]))
		dispatch(buf[0:n])
	}
}

//后端调度逻辑处理
func dispatch(data []byte) {
	msg := Message{}
	msg.CreateTime = uint64(time.Now().Unix())
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("message.go 195:", msg.Type)
	switch msg.Type {
	case 1: //私信
		fmt.Println("dispatch  data :", string(data))
		sendMsg(msg.TargetId, data)
	}
}

func sendMsg(userId int64, msg []byte) {
	//获取信息的发送者ID
	rwLocker.RLock()
	node, ok := clientMap[userId]
	rwLocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
