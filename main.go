package main

import (
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"log"
	"os"
)

func main() {
	portpath := "/dev/ttyUSB0"
	if len(os.Args) > 1 {
		portpath = string(os.Args[1])
	}

	options := serial.OpenOptions{
		PortName:        portpath,
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	defer port.Close()

	chargeBytes := []byte{0x55, 0x55, 0x55, 0x55, 0x4, 0x1}
	n, err := port.Write(chargeBytes)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

	fmt.Println("Wrote", n, "bytes.")
}
