package types


type ConnectReturnCode byte


const (
	ConnectionAccepted            ConnectReturnCode = 0x00
	UnacceptableProtocolVersion   ConnectReturnCode = 0x01
	IdentifierRejected            ConnectReturnCode = 0x02
	ServerUnavailable             ConnectReturnCode = 0x03
	BadUsernameOrPassword         ConnectReturnCode = 0x04
	NotAuthorized                 ConnectReturnCode = 0x05
)