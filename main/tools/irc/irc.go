package irc

import (
	"fmt"
	"log"
	"net"

	"gopkg.in/irc.v2"
)

type ircaddr struct {
	ip   int //ip
	port int //端口
}

func ThreadIRCSeed() {
	conn, err := net.Dial("tcp", "chat.freenode.net:6667")
	if err != nil {
		panic(err)
	}
	config := irc.ClientConfig{
		Nick: "i_have_a_nick",
		Pass: "password",
		User: "username",
		Name: "Full Name",
		Handler: irc.HandlerFunc(func(c *irc.Client, m *irc.Message) {
			if m.Command == "001" {
				// 001表示连接服务器成功
				c.Write("JOIN #bitcoin_ls\r")

			} else if m.Command == "PRIVMSG" {
				// PRIVMSG表示接收到消息

				fmt.Printf("%s\n", m.Params[1])

				c.WriteMessage(&irc.Message{
					Command: "PRIVMSG",
					Params: []string{
						m.Params[0],
						m.Trailing(),
					},
				})

			} else if m.Command == "JOIN" {
				// JOIN表示加入一个频道（channel）
				// who命令
				c.Write("WHO #bitcoin_ls\r")

			} else if m.Command == "352" {
				// who命令的返回消息
				fmt.Printf("%s\n", m.Params[3])
				fmt.Printf("%s\n", m.Params[5])
			}

		}),
	}

	// Create the client
	client := irc.NewClient(conn, config)
	err = client.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
