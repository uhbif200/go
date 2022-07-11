package main

import (
	"log"
	"remoteTTY/TCPServer"
)

func main() {
	s := TCPServer.NewServer("localhost", "63668", "tcp")
	s.Start()
}

func Logger(pipe chan string, stop chan int) {
	for {
		select {
		case <-stop:
			return
		case msg := <-pipe:
			log.Print(msg)
		}
	}
}
