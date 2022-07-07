package TTYHandler

import (
	"log"
	"strings"

	"github.com/tarm/serial"
)

type PortWorker struct {
	serial.Config
	IsWorking      bool
	ReadPipe       chan string
	WritePipe      chan string
	stopReadPipe   chan int
	stopWritePipe  chan int
	stopLoggerPipe chan int
}

func NewPortWorker(port string, bauld int) PortWorker {
	return PortWorker{
		Config: serial.Config{
			Name: port,
			Baud: bauld,
		},
		IsWorking:      false,
		ReadPipe:       make(chan string),
		WritePipe:      make(chan string),
		stopReadPipe:   make(chan int),
		stopWritePipe:  make(chan int),
		stopLoggerPipe: make(chan int),
	}
}

func (p *PortWorker) Start() {
	s, err := serial.OpenPort(&p.Config)
	if err != nil {
		log.Fatal(err)
	}
	p.IsWorking = true
	go Reader(s, p.ReadPipe, p.stopReadPipe)
	go Writer(s, p.WritePipe, p.stopWritePipe)
	go Logger(p.ReadPipe, p.stopLoggerPipe)
}

func (p *PortWorker) Stop() {
	p.IsWorking = false
	p.stopReadPipe <- 1
	p.stopWritePipe <- 1
	p.stopLoggerPipe <- 1
}

func Reader(s *serial.Port, pipe chan string, stop chan int) {
	handler := NewReadHandler(pipe)
	for {
		select {
		case <-stop:
			return
		default:
			buf := make([]byte, 128)
			n, err := s.Read(buf)
			if err != nil {
				log.Fatal(err)
			} else if n > 0 {
				handler(string(buf[:n]))
			}
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

func Writer(s *serial.Port, pipe chan string, stop chan int) {
	for {
		select {
		case <-stop:
			return
		case msg := <-pipe:
			n, err := s.Write([]byte(msg))
			if err != nil {
				log.Fatal(err, n)
			}
		}
	}
}
