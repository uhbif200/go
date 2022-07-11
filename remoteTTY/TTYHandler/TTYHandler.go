package TTYHandler

import (
	"log"

	"github.com/tarm/serial"
)

type PortWorker struct {
	serial.Config
	IsWorking     bool
	ReadPipe      chan string
	WritePipe     chan string
	stopReadPipe  chan int
	stopWritePipe chan int
}

func NewPortWorker(port string, bauld int) PortWorker {
	return PortWorker{
		Config: serial.Config{
			Name: port,
			Baud: bauld,
		},
		IsWorking:     false,
		ReadPipe:      make(chan string),
		WritePipe:     make(chan string),
		stopReadPipe:  make(chan int),
		stopWritePipe: make(chan int),
	}
}

func (p *PortWorker) Start() {
	s, err := serial.OpenPort(&p.Config)
	if err != nil {
		log.Print(err)
		return
	}
	p.IsWorking = true
	go Reader(s, p.ReadPipe, p.stopReadPipe)
	go Writer(s, p.WritePipe, p.stopWritePipe)
}

func (p *PortWorker) Stop() {
	p.IsWorking = false
	p.stopReadPipe <- 1
	p.stopWritePipe <- 1
}

func Reader(s *serial.Port, pipe chan string, stop chan int) {
	//handler := NewReadHandler(pipe)
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
				pipe <- string(buf[:n])
				//handler(string(buf[:n]))
			}
		}
	}
}

// func NewReadHandler(pipe chan string) func(data string) { // handler removes all \r and \n, split lines by \n and throw it to pipe
// 	buff := ""
// 	return func(data string) {
// 		data = strings.Replace(data, string('\r'), "", -1)
// 		buff += data
// 		for endlPos := strings.IndexByte(buff, '\n'); endlPos >= 0; {
// 			msg := buff[0:endlPos]
// 			buff = buff[endlPos+1:]
// 			endlPos = strings.IndexByte(buff, '\n')
// 			if len(msg) > 0 {
// 				pipe <- msg
// 			}
// 		}
// 	}
// }

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
