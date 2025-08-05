package types


const (
	// CONNECT byte = 0x10
	// CONNACK byte = 0x20
	// PUBLISH byte = 0x30
	// PUBACK byte = 0x40
	// SUBSCRIBE byte = 0x50
	// SUBACK byte = 0x60
	// PINGREQ byte = 0xC0
	// DISCONNECT byte = 0xE0

	QoS0 byte = 0x00
	QoS1 byte = 0x01
	QoS2 byte = 0x02

	DUPFLAG byte = 0x80
	QOSFLAG byte = 0x06
    RETAINFLAG byte = 0x01
)