package types


type FixedHeader struct {
    ControlBytes byte
	RemainingLength int
}

type ConnectPacket struct {
	FixedHeader FixedHeader
	ProtocolName    string
	ProtocolVersion byte
	ConnectFlags    byte //bch mnnsach kyn bites hdo  CleanSession, WillFlag, WillQoS, WillRetain, PasswordFlag, UserNameFlag
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
	

}