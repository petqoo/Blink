package types


type FixedHeader struct {
	PacketType PacketType
	Flags      ControlFlags
	RemainingLength int
}
func (h FixedHeader) EncodeControlByte() byte {
	return byte(h.PacketType)<<4 | byte(h.Flags)
}

func DecodeControlByte(b byte) (PacketType, ControlFlags) {
	packetType := PacketType(b >> 4)
	flags := ControlFlags(b & 0x0F)
	return packetType, flags
}


type ConnectPacket struct {
	FixedHeader FixedHeader
	ProtocolName    string
	ProtocolVersion MQTTProtocolVersion
	ConnectFlags    ConnecFlag 
	KeepAlive       uint16
	ClientID        string
	WillTopic       string 
	WillMessage     string 
	Username        string 
	Password        string 
}
type CONNACKPacket struct {
	FixedHeader FixedHeader
	ConnectReturnCode ConnectReturnCode
	SessionPresent  SessionPresentFlag	
}