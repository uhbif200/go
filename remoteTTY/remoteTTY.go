package main

import (
	"log"
	"remoteTTY/TCPServer"
)

func main() {
	// pw := TTYHandler.NewPortWorker("COM7", 115200)
	// pw.Start()

	// stopLoggerPipe := make(chan int)

	// go Logger(pw.ReadPipe, stopLoggerPipe)

	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	if pw.IsWorking {
	// 		pw.WritePipe <- scanner.Text()
	// 	}
	// }

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
