package sml

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
)

type SMLMessage struct {
	TransactionID OctetString
	GroupNo       uint8
	AbortOnError  uint8
	MessageBody   SMLMessageBody
	CRC16         uint16
	EndOfSmlMsg   EndOfSMLMsg
}

type SMLMessageBody struct {
	OpenRequest              SMLPublicOpenReq       // [0x00000100]
	OpenResponse             SMLPublicOpenRes       // [0x00000101]
	CloseRequest             SMLPublicCloseReq      // [0x00000200]
	CloseResponse            SMLPublicCloseRes      // [0x00000201]
	GetProfilePackRequest    SMLGetProfilePackReq   // [0x00000300]
	GetProfilePackResponse   SMLGetProfilePackRes   // [0x00000301]
	GetProfileListRequest    SMLGetProfileListReq   // [0x00000400]
	GetProfileListResponse   SMLGetProfileListRes   // [0x00000401]
	GetProcParameterRequest  SMLGetProcParameterReq // [0x00000500]
	GetProcParameterResponse SMLGetProcParameterRes // [0x00000501]
	SetProcParameterRequest  SMLSetProcParameterReq // [0x00000600]
	SetProcParameterResponse SMLSetProcParameterRes // [0x00000601]
	GetListRequest           SMLGetListReq          // [0x00000700]
	GetListResponse          SMLGetListRes          // [0x00000701]
	GetCosemRequest          SMLGetCosemReq         // [0x00000800]
	GetCosemResponse         SMLGetCosemRes         // [0x00000801]
	SetCosemRequest          SMLSetCosemReq         // [0x00000900]
	SetCosemResponse         SMLSetCosemRes         // [0x00000901]
	ActionCosemRequest       SMLActionCosemReq      // [0x00000A00]
	ActionCosemResponse      SMLActionCosemRes      // [0x00000A01]
	AttentionResponse        SMLAttentionRes        // [0x0000FF01]
}

type EndOfSMLMsg uint8

func ParseSMLMessage(raw []byte) (data []SMLMessage) {
	return data
}

func (s *SMLMessage) Parse(raw *[]byte) {

	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	spew.Dump(v)
	switch v.Kind() {
	case reflect.Slice:
		fmt.Println("v => slice")
	case reflect.Struct:
		fmt.Println("v => struct")
		for i := 0; i < v.NumField(); i++ {
			t := v.Field(i).Interface()
			switch t.(type) {
			case OctetString:
				s.TransactionID.Parse(raw)
			case uint8:
				fmt.Printf("---------------- [message] -----------------\n")
				buf := bytes.NewBuffer(*raw)
				spew.Dump(buf.Next(2))
				read(raw, &s.GroupNo)
				Cut(raw, 1)
				fmt.Printf("---------------- [/message] -----------------\n")
			case SMLMessageBody:
			case uint16:
				read(raw, &s.CRC16)
				Cut(raw, 2)
			case EndOfSMLMsg:
			}
		}
	}
}
