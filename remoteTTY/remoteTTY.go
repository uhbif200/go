package main

import (
	"bufio"
	"os"
	"remoteTTY/TTYHandler"
)

func main() {
	pw := TTYHandler.NewPortWorker("COM7", 115200)
	pw.Start()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if pw.IsWorking {
			pw.WritePipe <- scanner.Text()
		}
	}
}
