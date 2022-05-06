package main

import (
	"fmt"
	"net"
	"flag" // 命令行解析库
	"io"
	"os"
)

type Client struct {
	serverIp   string
	serverPort int
	Name       string
	conn       net.Conn
	flag       int // 当前客户端选择的模式
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		serverIp:   serverIp,
		serverPort: serverPort,
		flag:       -1,
	}

	// 连接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error: ", err)
		return nil
	}
	client.conn = conn

	// 返回对象
	return client

}

// 处理server回应的消息，直接显示到标准输出即可
func (this *Client) DealResponse() {
	// 一旦this.conn有消息，就直接copy到Stdout标准输出上，永久阻塞监听
	io.Copy(os.Stdout, this.conn)
}

// 菜单
func (this *Client) menu() bool {
	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		this.flag = flag
		return true
	} else {
		fmt.Println(">>>>>请输入合法范围内的数字<<<<<")
		return false
	}
}

// 查询在线用户
func (this *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := this.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write error: ", err)
		return
	}
}

// 私聊模式的方法
func (this *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	this.SelectUsers()
	fmt.Println(">>>>>清输入聊天对象的用户名,输入exit退出私聊模式.")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>>>请输入聊天内容，输入exit退出当前私聊")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			// 消息不为空则发送给服务器
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				_,err := this.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn.Write error: ", err)
					break
				}
			}

			// 发送完后，chatMsg置空，准备接收下一条消息
			chatMsg = ""
			fmt.Println(">>>>>请输入聊天内容，输入exit退出当前私聊")
			fmt.Scanln(&chatMsg)
		}

		// 退出当前私聊后，可以选择其他用户私聊
		this.SelectUsers()
		fmt.Println(">>>>>清输入聊天对象的用户名,输入exit退出私聊模式.")
		fmt.Scanln(&remoteName)

	}
}

// 公聊模式的方法
func (this *Client) PublicChat() {
	var chatMsg string
	fmt.Println(">>>>>请输入聊天内容,输入exit退出公聊模式.")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		// 消息不为空则发送给服务器
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_,err := this.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn.Write error: ", err)
				break
			}
		}

		// 发送完后，chatMsg置空，准备接收下一条消息
		chatMsg = ""
		fmt.Println(">>>>>请输入聊天内容，输入exit退出公聊模式.")
		fmt.Scanln(&chatMsg)
	}
}

// 更新用户名的方法
func (this *Client) UpdateName() bool {
	fmt.Println(">>>>>请输入用户名: ")
	fmt.Scanln(&this.Name)

	sendMsg := "rename|" + this.Name + "\n"
	_, err := this.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write error: ", err)
		return false
	}

	return true
}


// 客户端主业务
func (this *Client) Run() {
	// 只要flag != 0就一直循环，flag = 0时退出
	for this.flag != 0 {
		// 只要menu()不返回true就一直循环调用menu()，返回true时退出
		for this.menu() != true {

		}

		// 根据不同的模式处理不同的业务
		switch this.flag {
		case 1:
			// 公聊模式
			this.PublicChat()
			break
		case 2:
			// 私聊模式
			this.PrivateChat()
			break
		case 3:
			// 更新用户名
			this.UpdateName()
			break
		}
	}
}

// 全局变量
var serverIp string
var serverPort int

// init函数在main函数之前执行
func init() {
	// 1、解析得到的变量地址；2、命令行参数；3、参数默认值；4、输入-h时的提示信息
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口")
}

func main() {
	// 解析命令行 格式：./client -ip 127.0.0.1 -port 8888
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>服务器连接失败...")
		return
	}

	// 单独开启一个goroutine去处理sever回应的消息
	go client.DealResponse()

	fmt.Println(">>>>>服务器连接成功...")

	// 启动客户端的业务
	client.Run()
}