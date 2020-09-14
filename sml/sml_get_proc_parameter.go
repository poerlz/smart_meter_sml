package sml

type SMLGetProcParameterReq struct {
	ServerID          OctetString
	Username          OctetString
	Password          OctetString
	ParameterTreePath SMLTreePath
	Attribute         OctetString
}
type SMLGetProcParameterRes struct {
	ServerID          OctetString
	ParameterTreePath SMLTreePath
	ParameterTree     SMLTree
}
type SMLSetProcParameterReq struct {
	ServerID          OctetString
	Username          OctetString
	Password          OctetString
	ParameterTreePath SMLTreePath
	ParameterTree     SMLTree
}
type SMLSetProcParameterRes struct {
}
