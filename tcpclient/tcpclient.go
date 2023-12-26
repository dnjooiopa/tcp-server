package tcpclient

import (
	"bufio"
	"log"
	"net"
	"strings"
)

type Client struct {
	conn net.Conn
}

func New() *Client {
	return &Client{}
}

func (c *Client) Connect(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) SendMsg(msg string) error {
	_, err := c.conn.Write([]byte(msg + "\n"))
	return err
}

func (c *Client) HandleIncomingMsg(msg chan<- string) {
	reader := bufio.NewReader(c.conn)
	srvMsg, err := reader.ReadString('\n')
	if err != nil {
		if err.Error() == "EOF" {
			log.Println("tcpclient: server closed")
			c.Close()
		}
		return
	}

	msg <- strings.TrimSuffix(srvMsg, "\n")
}
