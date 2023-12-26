package server

import (
	"bufio"
	"log"
	"net"
	"strings"
)

type Server struct {
	listener net.Listener
}

func New() *Server {
	return &Server{}
}

func (s *Server) Start(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s.listener = listener

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return err
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) Stop() error {
	return s.listener.Close()
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReaderSize(conn, 1024)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				log.Println("server: connection closed")
			}
			return
		}

		msg = strings.TrimSuffix(msg, "\n")

		log.Printf("server: received message: %s\n", msg)

		err = s.write(conn, "message received")
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (s *Server) write(conn net.Conn, msg string) error {
	_, err := conn.Write([]byte(msg + "\n"))
	return err
}
