package sml

import (
	"encoding/binary"
)

type SMLTime struct {
	SecIndex       uint32
	Timestamp      SMLTimestamp
	TimestampLocal SMLTimestampLocal
}

type SMLTimestampLocal struct {
	Timestamp        SMLTimestamp
	LocalOffset      int16
	SeasonTimeOffset int16
}

type SMLStatus struct {
	Status8  uint8
	Status16 uint16
	Status32 uint32
	Status64 uint64
}

type SMLTimestamp uint32
type SMLUnit uint8
type OctetString []byte
type SMLSignature OctetString
type SMLStart [4]byte
type SMLVersion [4]byte

type SMLUint16 uint16

type SMLValue struct {
	BooleanValue bool
	Integer8     int8
	Integer16    int16
	Integer32    int32
	Integer64    int64
	ByteList     OctetString
	Unsigned8    uint8
	Unsigned16   uint16
	Unsigned32   uint32
	Unsigned64   uint64
	SMLList      SMLListType
}

type SMLTreePath struct {
	PathEntry OctetString
}

type SMLTree struct {
	ParameterName  OctetString
	ParameterValue SMLProcParValue
	ChildList      ListOfSMLTree
}

type SMLProcParValue struct {
	SmlValue       SMLValue
	SmlPeriodEntry SMLPeriodEntry // [0x02]
	SmlTupelEntry  SMLTupelEntry  // [0x03]
	SmlTime        SMLTime        // [0x04]
	SmlListEntry   SMLListEntry   // [0x05]
}

type ListOfSMLTree struct {
	treeEntry []SMLTree
}

type SMLTupelEntry struct {
	ServerID        OctetString
	SecIndex        SMLTime
	Status          uint64
	UunitPA         SMLUnit
	ScalerPA        int8
	ValuePA         int64
	UnitR1          SMLUnit
	ScalerR1        int8
	ValueR1         int64
	UnitR4          SMLUnit
	ScalerR4        int8
	ValueR4         int64
	SignaturePAR1R4 OctetString
	UnitmA          SMLUnit
	ScalermA        int8
	ValuemA         int64
	UnitR2          SMLUnit
	ScalerR2        int8
	ValueR2         int64
	UnitR3          SMLUnit
	ScalerR3        int8
	ValueR3         int64
	SignaturemAR2R3 OctetString
}

func (s *SMLStart) Parse(raw *[]byte) {
	read(raw, s)
	Cut(raw, len(s))
}

func (s *SMLVersion) Parse(raw *[]byte) {
	read(raw, s)
	Cut(raw, len(s))
}

func (s *OctetString) Parse(raw *[]byte) {
	len := make([]byte, 1)
	read(raw, len)
	text := make([]byte, len[0])
	read(raw, text)
	*s = text
	Cut(raw, binary.Size(text))
}
