package znet

import (
	"fmt"
	"net"
)

// Server IServer接口的实现
type Server struct {
	// 服务器的名称
	Name string
	// 服务器绑定的ip版本
	IPVersion string
	// 服务器监听的ip
	IP string
	// 服务器监听的端口
	Port int
}

func NewServer(name string) *Server {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
}

func (s *Server) Start() {
	fmt.Println("[Start] Server at IP: ", s.IP, " Port: ", s.Port, " is starting...")

	go func() {
		// 1、获取一个tcp地址
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err)
			return
		}
		// 2、监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, " error: ", err)
			return
		}

		fmt.Println("Start Zinx server successfully ", s.Name, "Listening...")
		// 阻塞等待客户连接，处理客户端连接业务
		for {
			// 如果有连接过来， 阻塞返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error", err)
				continue
			}

			// 处理业务
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("rec buf error: ", err)
						continue
					}

					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write buf error: ", err)
						continue
					}
				}
			}()
		}
	}()
}

func (s *Server) Serve() {
	s.Start()

	// TODO 做一些启动之后的额外业务

	// 阻塞等待
	select {}
}

func (s *Server) Stop() {

}
