package sml

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"unsafe"

	"github.com/davecgh/go-spew/spew"
	"github.com/reiver/go-cast"

	crc "gopkg.in/snksoft/crc.v1"
)

var (
	Version = []byte("\x01\x01\x01\x01")     // start sequenz and version number
	End     = []byte("\x1b\x1b\x1b\x1b\x1a") // escape sequenz
	Search  = []byte("\x1b\x1b\x1b\x1b\x01\x01\x01\x01")
	// Begin = []byte{01, 01, 01, 01} // start sequenz and version number
	// End   = []byte{'1b', '1b', '1b', '1b'}     // escape sequenz
)

type Stream []byte
type SMLTest struct {
	Sequence SMLStart
	CRC16    uint16
	Text     OctetString
}

type SML struct {
	Stream   []byte
	Sequence SMLStart
	Version  SMLVersion
	Messages []SMLMessage
	CRC16    uint16
}

func New(raw []byte) (sml SML) {

	trim(&raw)
	// tmp := SML{}
	sml.Parse(&raw)

	// sml.Stream = raw
	// sml.Cut()
	// checksum := bytes.NewBuffer(sml.Stream[len(sml.Stream)-2:])
	// if err := binary.Read(checksum, binary.LittleEndian, &sml.CRC16); err != nil {
	// 	panic(err)
	// }
	// spew.Dump(SMLParse(sml.Stream, tmp))
	// SMLParse2(sml.Stream, tmp)
	// spew.Dump(tmp)
	return sml
}

func read(b *[]byte, data interface{}) {
	r := bytes.NewBuffer(*b)
	if err := binary.Read(r, binary.LittleEndian, data); err != nil {
		panic(err)
	}
}

func Cut(raw *[]byte, len int) {
	tmp := *raw
	*raw = tmp[len:]
}

func (s *SML) Parse(raw *[]byte) {

	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// spew.Dump(&raw)
	switch v.Kind() {
	case reflect.Slice:
		fmt.Println("v => slice")
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			t := v.Field(i).Interface()
			switch t.(type) {
			case SMLStart:
				s.Sequence.Parse(raw)
			case SMLVersion:
				s.Version.Parse(raw)
			case []SMLMessage:
				fmt.Println("test SMLMessage")
				buf := bytes.NewBuffer(*raw)
				if bytes.Equal(buf.Next(1), []byte("\x76")) {
					fmt.Println("Message found")
					Cut(raw, 1)
					tmp := SMLMessage{}
					tmp.Parse(raw)
					spew.Dump(tmp.GroupNo)
				}
				// s.Messages
			case uint16:
				fmt.Println("test uint16")
			case []byte:
				fmt.Println("test byte")
				s.Stream = *raw
			}
		}
	}
}

func (s *SMLTest) Parse(raw *[]byte) {

	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Ptr:
		v = v.Elem()
	case reflect.Slice:
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			t := v.Field(i).Interface()
			switch t.(type) {
			case uint16:
				read(raw, &s.CRC16)
				Cut(raw, 2)
			case SMLStart:
				s.Sequence.Parse(raw)
			case OctetString:
				s.Text.Parse(raw)
			}
		}
	default:
		fmt.Println("v => other")
	}

}

func (s *SML) SMLCheckCRC16() bool {
	checksum := crc.CalculateCRC(crc.X25, s.Stream[0:len(s.Stream)-2])
	if smlSum, err := cast.Uint64(s.CRC16); err == nil {
		if checksum == smlSum {
			return true
		}
	}
	return false
}

func trim(raw *[]byte) {
	Stream := *raw
	start := bytes.Index(Stream, Search)
	stop := bytes.Index(Stream[start+len(Search):], Search)
	if start != 0 && stop != -1 {
		*raw = Stream[start : stop+len(Search)+start]
	}
}

func (s *SML) Cut() {
	start := bytes.Index(s.Stream, Search)
	stop := bytes.Index(s.Stream[start+len(Search):], Search)
	if start != 0 && stop != -1 {
		s.Stream = s.Stream[start : stop+len(Search)+start]
	}
}

func (s *SML) Total() {
	// fmt.Printf("%q\n", *s)
	start := []byte("\x07\x01\x00\x01\x08\x00\xff")
	beginTotal := bytes.Index(s.Stream, start)
	fmt.Printf("found sequence: %d\n", beginTotal)
	totalData := s.Stream[beginTotal : beginTotal+30]
	fmt.Printf("%x\n", totalData)
	fmt.Printf("data: %d\n", bytes.Index(totalData, []byte("\x56"))+1)
	totalData = totalData[bytes.Index(totalData, []byte("\x56"))+2:]
	fmt.Printf("%x\n", totalData)
	// totalData = totalData[:bytes.Index(totalData, []byte("\x01"))]
	totalData = totalData[:5]
	fmt.Printf("%x\n", totalData)
	fmt.Printf("%q\n", totalData)
	// fmt.Println(ByteArrayToInt(totalData))

	var nowTotal uint32
	nowBuffer := bytes.NewBuffer(totalData)
	binary.Read(nowBuffer, binary.BigEndian, &nowTotal)
	fmt.Println(nowTotal)
	// fmt.Printf("%q\n", totalData)
	// fmt.Printf("%d\n", totalData[1:])
	// fmt.Printf("%s", hex.Dump(totalData[1:]))
	// fmt.Printf("%q\n", string(totalData[1:]))
	// fmt.Printf("%x\n", strconv.Atoi(string(totalData[1:])))

}

func ByteArrayToInt(arr []byte) int64 {
	val := int64(0)
	size := len(arr)
	for i := 0; i < size; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}
	return val
}
