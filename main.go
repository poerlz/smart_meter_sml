package main


import (
	"fmt"
	"github.com/tarm/serial"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	fmt.Println("Hello Ostfriesland")
	tty := &serial.Config{
		Name: "/dev/ttyAMA0",
		Baud: 9600,
		ReadTimeout: 1,
		Size: 8,
	}


	spew.Dump(tty)
}