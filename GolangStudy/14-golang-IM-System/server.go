package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"io"
	"time"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户的列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的channel
	Message chan string
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

// 监听Message广播消息channel的goroutine，一旦有消息，就发送给全部的在线User
func (this *Server) ListenMessager() {
	for {
		// 不断尝试从Message channel中读消息
		msg := <-this.Message

		// 将读到的msg发送给全部的在线User
		this.mapLock.Lock()
		for _, user := range this.OnlineMap {
			user.C <- msg
		}
		this.mapLock.Unlock()
	}
}


// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ": " + msg

	this.Message <- sendMsg
}

// 处理当前连接的业务的方法
func (this *Server) Handler(conn net.Conn) {
	// ...处理当前连接的业务
	//fmt.Println("连接建立成功")

	// 用户上线业务
	user := NewUser(conn, this)
	user.Online()

	// 监听用户是否活跃的channel
	isLive := make(chan bool)

	// 创建一个goroutine，接收客户端发送的消息，读到buf中
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf) // 返回成功读取的字节数n
			if n == 0 {              // 用户下线
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}

			// 提取用户的消息（去除'\n'）
			msg := string(buf[:n-1]) // 提取buf中前n-1个字节，转换为string

			// 用户针对msg进行消息处理
			user.DoMessage(msg)

			// 用户发送任意消息，代表该用户是活跃的
			isLive <- true
		}
	}()

	// 当前Handler阻塞
	for {
		select {
		case <-isLive:
			// 如果isLive中有数据可读，代表当前用户是活跃的，应该重置定时器
			// 不用做任何事情，select会进入这个case，不进入下面的case，但下面的case语句也会执行

		case <-time.After(time.Second * 300): // time.After()本质是一个channel，当它可读时定时器触发，执行它的时候就是重置定时器
			// 已经超时，将当前的User强制关闭
			user.SendMsg("长时间无活动, 您已退出连接")

			// 销毁当前User的资源
			close(user.C)

			// 关闭连接
			conn.Close()

			// 退出当前Handler
			return // 或在runtime.Goexit()
		}
	}

}

// 启动服务器的接口，绑定到Server类
func (this *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port)) // Sprintf返回格式化的字符串
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}
	// close listen socket
	defer listener.Close()

	// 启动监听Message的goroutine
	go this.ListenMessager()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err: ", err)
			continue
		}

		// do handler
		go this.Handler(conn) // 创建一个sub go去处理业务，main go继续循环监听

	}

}
