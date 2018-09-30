package node

import (
	"fmt"
	"io"
	"net"
	"strings"
	"time"

	"./message"
)

//定义一个实例
var NetworkInstant *Network

type Node struct {
	Net        *net.TCPConn //连接实例
	Port       string       //端口
	ConnetTime int          // 连接时间
}

type Nodes map[string]*Node     //存放所有的节点
type NodeChan chan *Node        //接收*Node的chan
type ConnectionChan chan string //接收地址：端口
//核心结构
type Network struct {
	Nodes          Nodes                //保存所有节点
	ConnectionChan ConnectionChan       //新的连接进入，通过该Chan触发处理机制
	Address        string               //保存当前节点的地址
	Port           string               // 端口
	NodeCallback   NodeChan             //新的节点连接完成，调用该chan
	BroadCastChan  chan message.Message //所有Message丢入该chan，该chan会广播
	IncomingChan   chan message.Message //如果接收到外部消息，丢入到该chan
}

// 创建新的网络
// 需要传入本机地址及端口
func CreateNetwork(address, port string) *Network {
	n := &Network{}

	n.Address = address
	n.Port = port
	n.Nodes = Nodes{}
	n.BroadCastChan, n.IncomingChan = make(chan message.Message), make(chan message.Message)
	n.ConnectionChan, n.NodeCallback = CreateConnectHandlerForReceive()

	return n
}

// 处理连接函数
func CreateConnectHandlerForReceive() (ConnectionChan, NodeChan) {
	incomingAddr := make(ConnectionChan)
	connCb := make(NodeChan)

	// 相当于启动一个线程
	go func() {
		for {
			// 读取一个地址，当没有地址传输过来时，channel挂起
			address := <-incomingAddr
			fmt.Println("start connect to node: " + address)
			// 将地址分为IP和port
			s := strings.Split(address, ":")
			if len(s) != 2 {
				fmt.Println("incoming address error: " + address)
				continue
			}
			localAddress := NetworkInstant.Address + ":" + NetworkInstant.Port
			// 当接收到的address和本地ip地址不同，且address不在已知的地址map中
			if localAddress != address && NetworkInstant.Nodes[address] == nil {
				// 连接此实例
				// todo：限制最大连接上限
				go ConnectNode(s[0], s[1], 10*time.Second, false, connCb)
			}
		}
	}()
	return incomingAddr, connCb
}

//Timeout 定时任务，时间到触发channel
func Timeout(t time.Duration) chan bool {
	i := make(chan bool)
	// 设置休眠时间，时间未到，i不写入，chan挂起
	go func() {
		time.Sleep(t)
		i <- true
	}()
	return i
}

// 连接到某个节点
func ConnectNode(nodehost string, port string, timeout time.Duration, retry bool, cb NodeChan) {
	host := nodehost + ":" + port
	fmt.Println("try to connect to node " + host)
	// 将字符串转为地址
	addr, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		fmt.Println(err.Error())
	}
	tryConnecting := true
	for tryConnecting {
		go func() {
			conn, err := net.DialTCP("tcp", nil, addr)
			if err != nil {
				fmt.Println("DialTCP " + host + " err: " + err.Error())
			}
			if conn != nil {
				cb <- &Node{conn, port, int(time.Now().Unix())}
				// 发送自己端口
				portMsg, _ := message.NewMessage(message.MessageTypePort)
				portMsg.Data = []byte(NetworkInstant.Port)
				NetworkInstant.BroadCastChan <- *portMsg

				tryConnecting = false
			}
		}()
		select {
		case <-Timeout(timeout):
			if tryConnecting && !retry {
				tryConnecting = false
			}
		}
	}
}
func Run(n *Network, port string) {
	fmt.Println("net start run: " + NetworkInstant.Address)
	listencb := StartListening(NetworkInstant.Address, port)
	for {
		select {
		case node := <-listencb:
			NetworkInstant.Nodes.AddNode(node)
		case node := <-n.NodeCallback:
			NetworkInstant.Nodes.AddNode(node)
		case m := <-n.BroadCastChan:
			go n.BroadCastMessage(&m)
		}
	}
}

// 开始接受其他节点发来的信息
func StartListening(address, port string) NodeChan {
	cb := make(NodeChan)
	addr, err := net.ResolveTCPAddr("tcp4", address+":"+port)
	if err != nil {
		fmt.Println("resolve tcp address error: " + err.Error())
	}
	listening, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		fmt.Println("listen tcp address error: " + err.Error())
	}
	go func(listener *net.TCPListener) {
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept tcp address error: " + err.Error())
			}
			cb <- &Node{conn, "0", int(time.Now().Unix())}
		}
	}(listening)
	return cb
}

//BroadCastMessage 广播所有信息
func (n *Network) BroadCastMessage(m *message.Message) {
	b := m.MarshalBinary()
	for _, node := range n.Nodes {
		fmt.Println("broadcast message to " + node.Net.RemoteAddr().String())
		go func() {
			_, err := node.Net.Write(b)
			if err != nil {
				fmt.Println("err when broadcast message to " + node.Net.RemoteAddr().String())
			}
		}()
	}
}

//AddNode 有新的
func (n Nodes) AddNode(node *Node) bool {
	addr := node.Net.RemoteAddr().String()
	if addr != NetworkInstant.Address && NetworkInstant.Nodes[addr] == nil {
		fmt.Println("node connect from: " + addr)
		// n[addr] = node	// 这里不加入，等到对方传来port的时候，再加入
		if node.Port != "0" {
			n[addr+":"+node.Port] = node
		}
		go HandleNode(n, node)
		return true
	}
	return false
}

//HandleNode 接收Node的消息，并且如果有返回则传递回去
func HandleNode(nodes Nodes, node *Node) {
	for {
		data := make([]byte, 1024*1000)
		l, err := node.Net.Read(data[0:])
		if err != nil {
			fmt.Println("err when read byte from node, err: " + err.Error())
		}
		if err == io.EOF {
			fmt.Println("EOF")
			node.Net.Close()
			NetworkInstant.Nodes[node.Net.RemoteAddr().String()] = nil
			break
		}

		m := message.Message{}
		m.UnmarshalBinary(data[0:l])
		if err != nil {
			fmt.Println("err when convert bytes to Message, err: " + err.Error())
			continue
		}

		// 如果Type只是传递port，则不需要对传递给其他方法处理了
		if m.Type == message.MessageTypePort {
			node.Port = string(m.Data)
			nodeAddr := node.Net.RemoteAddr().String()
			s := strings.Split(nodeAddr, ":")
			if len(s) > 1 {
				nodeAddr = s[0]
			}
			fmt.Println("add new node to Nodes: " + nodeAddr + ":" + node.Port)
			nodes[nodeAddr+":"+node.Port] = node
			continue
		}

		m.Reply = make(chan message.Message)
		//TODO: 如果出错导致返回消息没有传递过来，需要有个timeout机制
		go func(cb chan message.Message) {
			for {
				reply, ok := <-cb
				if !ok {
					close(cb)
					break
				}

				replyBytes := reply.MarshalBinary()
				i := 0
				for i < 1 {
					writeResult, _ := node.Net.Write(replyBytes[i:])
					i += writeResult
				}
			}
		}(m.Reply)

		NetworkInstant.IncomingChan <- m
	}
}
