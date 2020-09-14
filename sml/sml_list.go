package sml

type SMLGetListReq struct{}

type SMLGetListRes struct {
	ClientID        OctetString // Octet string OPTIONAL
	ServerID        OctetString // Octet string
	ListName        OctetString // Octet string OPTIONAL
	ACTSensorTime   SMLTime     // OPTIONAL
	ValList         SMLList
	ListSignature   SMLSignature // OPTIONAL
	ACTGeatewayTime SMLTime      // OPTIONAL
}

type SMLList struct {
	ValListEntry []SMLListEntry
}

type SMLListEntry struct {
	ObjName        OctetString // Octet string
	Status         SMLStatus   // OPTIONAL
	ValTime        SMLTime     // OPTIONAL
	Unit           SMLUnit     // OPTIONAL
	Scaler         int8        // OPTIONAL
	Value          SMLValue
	ValueSignature SMLSignature // OPTIONAL
}

type SMLListType struct {
	SmlTime SMLTime
}
