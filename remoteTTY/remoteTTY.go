package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{
		Name:        "COM7",
		Baud:        115200,
		ReadTimeout: time.Millisecond * 100,
		Size:        0,
		Parity:      0,
		StopBits:    0,
	}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	readPipe := make(chan string)
	writePipe := make(chan string)

	go Reader(s, readPipe)
	go Logger(readPipe)
	go Writer(s, writePipe)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		writePipe <- scanner.Text() + string('\n')
	}
}

func Reader(s *serial.Port, pipe chan string) {

	handler := NewReadHandler(pipe)
	for {
		buf := make([]byte, 128)
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		} else if n > 0 {
			handler(string(buf[:n]))
		}
	}
}

func NewReadHandler(pipe chan string) func(data string) {
	buff := ""
	return func(data string) {
		buff += data
		for endlPos := strings.IndexByte(buff, '\n'); endlPos >= 0; {
			msg := buff[0:endlPos]
			buff = buff[endlPos+1:]
			pipe <- msg
			endlPos = strings.IndexByte(buff, '\n')
		}
	}
}

func Logger(pipe chan string) {
	for {
		log.Print(<-pipe)
	}
}

func Writer(s *serial.Port, pipe chan string) {
	for {
		n, err := s.Write([]byte(<-pipe))
		if err != nil {
			log.Fatal(err, n)
		}
	}
}
