package main

// import (
//   "github.com/mikepb/go-serial"
//   "log"
// )
//
// func main() {
//   options := serial.RawOptions
//   options.BitRate = 9600
//   p, err := options.Open("/dev/tty.usbserial")
//   if err != nil {
//     log.Panic(err)
//   }
//
//   defer p.Close()
//
//   chargeBytes := []byte{0x55, 0x55, 0x55, 0x55, 0x4, 0x1}
//
//   buf := make([]byte, 1)
//   if c, err := p.Read(buf); err != nil {
//     log.Panic(err)
//   } else {
//     log.Println(buf)
//     log.Println(c)
//   }
//   if c, err := p.Write(chargeBytes); err != nil {
//     log.Panic(err)
//   } else {
//     log.Println(c)
//   }
//
// }

import (
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"log"
	"os"
)

func main() {
	// Set up options.
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

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Make sure to close it later.
	defer port.Close()

	// Write 4 bytes to the port.
	chargeBytes := []byte{0x55, 0x55, 0x55, 0x55, 0x4, 0x1}
	n, err := port.Write(chargeBytes)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

	fmt.Println("Wrote", n, "bytes.")
}
