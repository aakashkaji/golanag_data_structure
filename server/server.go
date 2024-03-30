package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type Server struct {
	listenAddr string
	ln         net.Listener
	clients    map[net.Conn]struct{}
	mu         sync.Mutex
	quitchan   chan struct{}
}

func NewServer(listenAddr string) *Server {
	return &Server{listenAddr: listenAddr,
		clients:  make(map[net.Conn]struct{}),
		quitchan: make(chan struct{})}

}

func (s *Server) StartServer() error {

	ln, err := net.Listen("tcp", s.listenAddr)

	if err != nil {
		log.Println("Error occured: ", err)
	}
	defer ln.Close()

	s.ln = ln
	go s.AcceptLoop()

	<-s.quitchan
	return nil

}

func (s *Server) AcceptLoop() {

	for {
		conn, err := s.ln.Accept()

		if err != nil {
			log.Println("Error on AcceptLoop: ", err)

		}
		fmt.Println("Accept New connection:", conn.RemoteAddr())

		go s.Readloop(conn)
	}

}

func (s *Server) Readloop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	s.mu.Lock()
	s.clients[conn] = struct{}{}
	s.mu.Unlock()

	for {
		n, err := conn.Read(buf)

		if err != nil {
			log.Println("ReadLoop err:", err)
			continue
		}
		msg := buf[:n]
		fmt.Println(msg)
		fmt.Println(string(msg))
		_, err = conn.Write(msg)
		s.BrodcastMessage(conn, msg)
		if err != nil {
			log.Println("ReadLoop err:", err)

		}

	}

}

func (s *Server) BrodcastMessage(sender net.Conn, msg []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for con := range s.clients {

		fmt.Println(con)

		if con != sender {
			con.Write(msg)

		}

	}

}

func mainss() {

	fmt.Println("Tcp server connections")

	newServer := NewServer(":3000")
	newServer.StartServer()
}
