package types

type PacketType byte

const (
	CONNECT     PacketType = 1
	CONNACK     PacketType = 2
	PUBLISH     PacketType = 3
	PUBACK      PacketType = 4
	SUBSCRIBE   PacketType = 8
	SUBACK      PacketType = 9
	UNSUBSCRIBE PacketType = 10
	PINGREQ     PacketType = 12
	PINGRESP    PacketType = 13
	DISCONNECT  PacketType = 14
)

type ControlFlags byte

const (
	DUPFlag    ControlFlags = 1 << 3 // 00001000
	QoS1Flag   ControlFlags = 1 << 1 // 00000010
	QoS2Flag   ControlFlags = 1 << 2 // 00000100
	RetainFlag ControlFlags = 1 << 0 // 00000001
)
