package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string // user channel，传输的数据类型为string
	conn net.Conn    // 用于和客户端通信的连接

	server *Server
}

// 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String() // 客户端的地址

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server: server,
	}

	// 启动监听当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}

// 用户的上线业务
func (this *User) Online() {
	// 用户上线，将用户加入到OnlineMap
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	// 广播当前用户上线的消息
	this.server.BroadCast(this, "已上线")
}

// 用户的下线业务
func (this *User) Offline() {
	// 用户下线，将用户从OnlineMap中删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	// 广播当前用户上线的消息
	this.server.BroadCast(this, "下线")
}

// 给当前User对应的客户端发送消息
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 用户的处理消息业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		// 如果收到客户端发来的消息为“who”，表示查询当前在线用户
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ": " + "在线...\n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 消息格式为“rename|xxx”时，表示修改当前用户名
		newName := strings.Split(msg, "|")[1] // 根据“|”拆分字符串，放入数组中，“xxx”索引为1

		// 判断newName是否已经存在
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("该用户名已被占用\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name) // 删除旧用户
			this.server.OnlineMap[newName] = this    // 添加新用户
			this.server.mapLock.Unlock()

			this.Name = newName // 更新用户名
			this.SendMsg("成功修改用户名: " + this.Name + "\n")
		}

	} else if len(msg) > 3 && msg[:3] == "to|" {
		// 消息格式为“to|xxx|消息内容”时，表示给xxx用户发送消息

		// 1 获取对方用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMsg("消息格式错误，请使用\"to|xxx|消息内容\"格式。\n")
			return
		}

		// 2 根据对方用户名得到对方User对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("该用户不存在\n")
			return
		}

		// 3 获取消息内容，通过对方User对象将消息内容发送给过去
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMsg("无消息内容，清重发\n")
			return
		}
		remoteUser.SendMsg(this.Name + "对您说: " + content + "\n")
	} else {
		// 否则，消息广播
		this.server.BroadCast(this, msg)
	}


}


// 监听当前user channel的方法，一旦有消息，就直接发送给对应客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}
