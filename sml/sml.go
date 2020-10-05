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

type SMLTest struct {
	Sequence SMLStart
	CRC16    uint16
	Text     OctetString
}

type SML struct {
	Sequence SMLStart
	Version  SMLVersion
	Stream   []byte
	Messages []SMLMessage
	CRC16    uint16
}

func New(raw []byte) (sml SML) {

	sml.Stream = raw
	sml.Cut()
	checksum := bytes.NewBuffer(sml.Stream[len(sml.Stream)-2:])
	if err := binary.Read(checksum, binary.LittleEndian, &sml.CRC16); err != nil {
		panic(err)
	}
	// tmp := SML{}
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

func (s *SMLStart) Parse(raw *[]byte) {
	read(raw, s)
}

func (s *OctetString) Parse(raw *[]byte) {
	len := make([]byte, 1)
	read(raw, len)
	text := make([]byte, len[0])
	read(raw, text)
	*s = text
}

func (s *SMLTest) Parse(raw *[]byte) {

	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		fmt.Printf("---------------- [refect ptr] -----------------\n")
		spew.Dump(v.Kind())
		fmt.Printf("---------------- [/refect ptr] -----------------\n")
	}
	switch v.Kind() {
	case reflect.Ptr:
		v = v.Elem()
	case reflect.Slice:
		fmt.Println("v => slice")
	case reflect.Struct:
		fmt.Println("v => struct")
		for i := 0; i < v.NumField(); i++ {
			t := v.Field(i)
			// spew.Dump(t.Type().String())
			spew.Dump(t.Interface())
			// spew.Dump(f.Kind())

			switch t.Type().String() {
			case "uint16":
				fmt.Println("test uint16")
				read(raw, &s.CRC16)
			case "sml.SMLStart":
				fmt.Println("test smlStart")

				s.Sequence.Parse(raw)
			case "sml.OctetString":
				fmt.Println("test octetstring")
				s.Text.Parse(raw)
			}
		}
	default:
		fmt.Println("v => other")
	}

}

func SMLParse2(raw []byte, data interface{}) {

	fmt.Printf("---------------- [parse2] -----------------\n")

	switch data := data.(type) {
	case *uint32:
		fmt.Println("test uint32")
	case *uint16:
		fmt.Println("test uint16")

	case *SMLMessage:
		fmt.Println("test SMLMessage")
		spew.Dump(data)
	case *SML:
		fmt.Println("test SML")
		spew.Dump(data)
	case *SMLVersion:
		fmt.Println("test SMLVersion")
		read(&raw, data)
	case *SMLStart:
		fmt.Println("test SMLStart")
		br := make([]byte, 10)
		// read(&raw, br)
		spew.Dump(br)

	case *SMLTime:
		fmt.Println("test SMLTime")
		// read(&raw, data)
	case *OctetString:
		fmt.Println("test OctetString")
		// read(&raw, &data)
	default:
		fmt.Println("test default")
		// read(&raw, data)
		// spew.Dump(data)
	}
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		fmt.Printf("---------------- [refect ptr] -----------------\n")
		spew.Dump(v.Kind())
		fmt.Printf("---------------- [/refect ptr] -----------------\n")
	}
	switch v.Kind() {
	case reflect.Ptr:
		v = v.Elem()
	case reflect.Slice:
		fmt.Println("v => slice")
	case reflect.Struct:
		fmt.Println("v => struct")
		for i := 0; i < v.NumField(); i++ {
			t := v.Field(i)
			spew.Dump(t.Type().String())
			// tmp := SMLType(t)
			// f := reflect.ValueOf(&tmp)
			// spew.Dump(f.Kind())

			switch t.Type().String() {
			case "sml.SMLStart":
				tmp := SMLStart{}

				SMLParse2(raw, &tmp)
				spew.Dump(tmp)

			case "sml.OctetString":
				tmp := OctetString{}
				SMLParse2(raw, &tmp)
				spew.Dump(tmp)
			}
			// spew.Dump(v.Field(i).Type())

		}
	default:
		fmt.Println("v => other")
	}

	// spew.Dump(v.NumField())
	// spew.Dump(v.Field(1))
	// spew.Dump(v.Field(1).Type().Name())
	// bs := make([]byte, 4)
	// read(&raw, &bs)
	// data = bs
	// fmt.Println(bs)

	// spew.Dump(tmp)
	fmt.Printf("\n---------------- [/parse2] -----------------\n")
	// spew.Dump(data)
	// spew.Dump(st.Elem())

	fmt.Printf("\n---------------- [//parse2] -----------------\n")
}

func SMLParse(raw []byte, data interface{}) (s SML) {

	// raw := s.Stream
	st := reflect.TypeOf(data)
	for i := 0; i < st.NumField(); i++ {
		fmt.Println(st.Field(i).Type)
		// fmt.Println(st.Field(i).Type.Kind())
		// fmt.Println(st.Field(i).Type.Size())
		// fmt.Printf("%x \n", st.Field(i).Type.Size())
		switch st.Field(i).Type.String() {
		case "sml.SMLStart":
			fmt.Println("Kind Start")
			read(&raw, &s.Sequence)
			raw = raw[len(s.Sequence):]
		case "sml.SMLVersion":
			fmt.Println("Kind Version")
			read(&raw, &s.Version)
			raw = raw[len(s.Version):]
		case "[]sml.SMLMessage":
			fmt.Println("Kind SMLMessage")
			// s.Messages = append(s.Messages, SMLParse(raw, SMLMessage{}))
		case "uint16":
			fmt.Println("Kind uint16")
		case "sml.OctetString":
			fmt.Println("Kind Octetstring")
		default:
			fmt.Println("Kind other")
		}
		// switch st.Field(i).Type.Kind() {
		// case reflect.Uint8:
		// 	fmt.Println("Kind is 1")
		// case reflect.Uint16:
		// 	fmt.Println("size is 2")
		// case reflect.Array:
		// 	fmt.Println("kind is a array")
		// 	r := bytes.NewBuffer(raw)
		// 	bs := make([]byte, st.Field(i).Type.Size())
		// 	if err := binary.Read(r, binary.LittleEndian, bs); err != nil {
		// 		panic(err)
		// 	}
		// 	fmt.Printf("% x\n", bs)
		// case reflect.Slice:
		// 	fmt.Println("Kind is a Slice")
		// 	// s := sizeof(v.Type).ELem()
		// 	// fmt.Fprintln(s)
		// }
		// spew.Dump(st.Field(i))
		fmt.Println("---------------------")
	}
	// var data SML
	// buf := bytes.NewBuffer(s.Stream)
	// if err := binary.Read(buf, binary.LittleEndian, &data); err != nil {
	// 	panic(err)
	// }
	// VersionStart := bytes.Index(s.Stream, Version)
	// version := s.Stream[VersionStart:]
	// for _, v := range version {
	// 	if string(v) == "\x76" {
	// 		fmt.Printf("% x\n", v)
	// 	}
	// }

	// switch data.(type) {
	// case *uint16:
	// 	fmt.Println("test 1")
	// case SML:
	// 	fmt.Println("test SML")
	// case SMLMessage:
	// 	fmt.Println("test SMLMessage")
	// default:
	// 	fmt.Printf("test default")
	// }
	return s
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

func (s *SML) Cut() {
	start := bytes.Index(s.Stream, Search)
	stop := bytes.Index(s.Stream[start+len(Search):], Search)
	if start != 0 && stop != -1 {
		s.Stream = s.Stream[start : stop+len(Search)+start]
		// var checksum uint16

		// if err := binary.Read(s.Stream[len(s.Stream)-2:], binary.LittleEndian, &checksum); err != nil {
		// 	panic(err)
		// }
		// s.CRC16 = checksum
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
