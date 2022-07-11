package TCPServer

import (
	"log"
	"net"
	"regexp"
	"remoteTTY/TTYHandler"
)

type Server struct {
	host     string
	port     string
	protocol string
}

func NewServer(host, port, protocol string) Server {
	return Server{
		host:     host,
		port:     port,
		protocol: protocol,
	}
}

func (s *Server) Start() {
	listen, err := net.Listen(s.protocol, s.host+":"+s.port)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		} else {
			go s.handleIncomingRequest(conn)
		}
	}
}

func (*Server) Stop() {

}

func (*Server) handleIncomingRequest(conn net.Conn) {
	buffer := make([]byte, 1024)

	var port_worker TTYHandler.PortWorker

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Print(err)
			conn.Close()
		}
		msg := string(buffer[:n])

		if !port_worker.IsWorking {
			port := parcePort(msg)
			if port != "" {
				port_worker = TTYHandler.NewPortWorker(port, 115200)
				port_worker.Start()
				go func() {
					for {
						msg := <-port_worker.ReadPipe
						conn.Write([]byte(msg))
					}
				}()
			}
		} else {
			port_worker.WritePipe <- msg
		}
	}
}

func parcePort(port string) string {
	reg, _ := regexp.Compile(`^[a-zA-Z0-9\/]*`)
	str := reg.FindString(port)
	return str
}
