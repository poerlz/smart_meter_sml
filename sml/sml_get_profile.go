package sml

type SMLGetProfilePackReq struct {
	ServerID          OctetString
	Username          OctetString
	Password          OctetString
	WithRawData       bool
	BeginTime         SMLTime
	EndTime           SMLTime
	ParameterTreePath SMLTreePath
	ObjectList        ListOfSMLObjReqEntry
	DasDetatils       SMLTree
}
type SMLGetProfilePackRes struct {
	ServerID          OctetString
	ActTime           SMLTime
	RefPeriod         uint32
	ParameterTreePath SMLTreePath
	HeaderList        ListOfSMLProfObjHeaderEntry
	PeriodList        ListOfSMLProfObjPeriodEntry
	RawData           OctetString
	ProfileSignature  SMLSignature
}
type SMLGetProfileListReq struct {
	ServerID          OctetString
	Username          OctetString
	Password          OctetString
	WithRawData       bool
	BeginTime         SMLTime
	EndTime           SMLTime
	ParameterTreePath SMLTreePath
	ObjectList        ListOfSMLObjReqEntry
	DasDetails        SMLTree
}
type SMLGetProfileListRes struct {
	ServerID          OctetString
	ActTime           SMLTime
	RegPeriod         uint32
	ParameterTreePath SMLTreePath
	valTime           SMLTime
	Status            uint64
	PeriodList        ListOfPeriodEntry
	RawData           OctetString
	PeriodSignature   SMLSignature
}

type ListOfSMLObjReqEntry struct {
	ObjectListEntry []SMLObjReqEntry
}

type SMLObjReqEntry OctetString

type ListOfSMLProfObjHeaderEntry struct {
	HeaderListEntry []SMLProfObjHeaderEntry
}

type SMLProfObjHeaderEntry struct {
	ObjName OctetString
	Unit    SMLUnit
	Scaler  int8
}

type ListOfSMLProfObjPeriodEntry struct {
	PeriodListEntry []SMLProfObjPeriodEntry
}

type SMLProfObjPeriodEntry struct {
	ValTime         SMLTime
	Status          uint64
	ValueList       ListOfSMLValueEntry
	PeriodSignature SMLSignature
}

type ListOfSMLValueEntry struct {
	ValueListEntry []SMLValueEntry
}

type SMLValueEntry struct {
	Value          SMLValue
	valueSignature SMLSignature
}

type ListOfPeriodEntry struct {
	PeriodListEntry []SMLPeriodEntry
}

type SMLPeriodEntry struct {
	ObjName        OctetString
	Unit           SMLUnit
	Scaler         int8
	Value          SMLValue
	valueSignature SMLSignature
}
