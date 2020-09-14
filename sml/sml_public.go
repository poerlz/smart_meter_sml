package sml

type SMLPublicOpenReq struct {
	Codepage   OctetString
	ClientID   OctetString
	ReqFileID  OctetString
	ServerID   OctetString
	Username   OctetString
	Password   OctetString
	SMLVersion uint8
}
type SMLPublicOpenRes struct {
	Codepage   OctetString
	ClientID   OctetString
	ReqField   OctetString
	ServerID   OctetString
	RefTime    SMLTime
	SMLVersion uint8
}
type SMLPublicCloseReq struct {
	GlobalSignature SMLSignature
}
type SMLPublicCloseRes struct {
	GlobalSignature SMLSignature
}
