package main

import (
	"bytes"
	"fmt"
	"log"

	"code.mukd.de/smart_meter_sml/sml"

	"github.com/davecgh/go-spew/spew"
	"github.com/tarm/serial"
)

func main() {
	tty := &serial.Config{
		Name:        "/dev/ttyUSB0",
		Baud:        9600,
		ReadTimeout: 0,
		Size:        8,
	}
	spew.Dump(tty)

	ser, err := serial.OpenPort(tty)

	if err != nil {
		fmt.Println()
		log.Fatal(err)
	}
	defer ser.Close()

	buf := make([]byte, 8)
	// var stream string
	var big []byte

	streaming := false
	// for i := 0; i < 100; i++ {
	// 	_, err := ser.Read(buf)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	s := [][]byte{big, buf}
	// 	big = bytes.Join(s, []byte(""))
	// }
	for {
		_, err := ser.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		s := [][]byte{big, buf}
		big = bytes.Join(s, []byte(""))
		// fmt.Printf("%x", big)
		if streaming {
			// stream += string(buf)
			if bytes.Index(big, sml.End) >= 0 {
				fmt.Printf("found end: %d\n", bytes.Index(big, sml.End))
				fmt.Printf("found end: %d\n", bytes.Index(big[len(sml.Search):], sml.Search))
				break
			}
		} else if bytes.Index(big, sml.Search) >= 0 {
			// fmt.Printf("%x", big)
			big = big[bytes.Index(big, sml.Search):]
			// fmt.Printf("found start: %d\n", bytes.Index(buf, sml.Begin))
			// stream += string(buf)
			streaming = true
		}
	}
	fmt.Printf("\n---------------------------\n")
	// test := []byte(stream)
	protocol := sml.SML{Stream: big}

	// fmt.Printf("%x\n", protocol.Stream)
	// fmt.Printf("length: %d\n", len(protocol.Stream))
	fmt.Printf("% x\n", protocol.Stream)
	fmt.Printf("%x\n", protocol.Stream)
	fmt.Printf("%q\n", protocol.Stream)
	// protocol.Cut()
	// fmt.Printf("%x\n", protocol.Stream)
	// fmt.Printf("length: %d\n", len(protocol.Stream))
	fmt.Printf("\n---------------------------\n")
	// fmt.Printf("%q\n", protocol.Stream)
	protocol.Total()
	// all := bytes.IndexRune (protocol.Stream, sml.Begin)
	// stop := bytes.Index(protocol.Stream[start:], sml.End)
	// test := bytes.Index(protocol.Stream, sml.Search)
	// test2 := bytes.Index(protocol.Stream[test+len(sml.Search):], sml.Search)
	// fmt.Printf("start: %d\n", start)
	// fmt.Printf("stop: %d\n", stop)
	// fmt.Printf("test1: %d\n", test)
	// fmt.Printf("test2: %d\n", test2)
	// fmt.Printf("%q\n", test[start:stop])

	fmt.Printf("\n")

}
