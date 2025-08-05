package types
type MQTTProtocolVersion byte

const (
	MQTT_3_1   MQTTProtocolVersion = 3
	MQTT_3_1_1 MQTTProtocolVersion = 4 
	MQTT_5_0   MQTTProtocolVersion = 5
)

func (v MQTTProtocolVersion) String() string {
	switch v {
	case MQTT_3_1:
		return "3.1"
	case MQTT_3_1_1:
		return "3.1.1"
	case MQTT_5_0:
		return "5.0"
	}
	return ""
}