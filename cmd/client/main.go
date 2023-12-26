package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dnjooiopa/tcp-server/tcpclient"
)

func main() {
	client := tcpclient.New()

	if err := client.Connect("localhost:8080"); err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	err := client.SendMsg("hello from client")
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter text: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSuffix(text, "\n")

			err := client.SendMsg(text)
			if err != nil {
				log.Fatalln(err)
			}

			msg := make(chan string)
			go client.HandleIncomingMsg(msg)
			m := <-msg
			fmt.Println("received message:", m)
		}
	}()

	select {}
}
