package types


type ConnecFlag byte


const (
	CleanSessionFlag ConnecFlag = 1 << 1 // 00000010
	WillFlag         ConnecFlag = 1 << 2 // 00000100
	WillQoS1         ConnecFlag = 1 << 3 // 00001000
	WillQoS2         ConnecFlag = 1 << 4 // 00010000
	WillRetainFlag   ConnecFlag = 1 << 5 // 00100000
	PasswordFlag     ConnecFlag = 1 << 6 // 01000000
	UsernameFlag     ConnecFlag = 1 << 7 // 10000000
)

func (f ConnecFlag) IsSet(flag ConnecFlag) bool {
	return f&flag != 0	
}

func (f *ConnecFlag) Set(flag ConnecFlag){
	*f|= flag
}
func (f *ConnecFlag) Clear(flag ConnecFlag){
	*f&^= flag
}
func (f *ConnecFlag) SetWillQoS(qos byte) {
	*f &^= (WillQoS1 | WillQoS2)              
	*f |= ConnecFlag((qos & 0x03) << 3)     
}


func (f *ConnecFlag) WillQoS() byte {
	return byte((*f >> 3) & 0x03)
}

func (f ConnecFlag) Parse() []ConnecFlag{
	var setflags []ConnecFlag
	AllFlag := []ConnecFlag{CleanSessionFlag, WillFlag, WillQoS1, WillQoS2, WillRetainFlag, PasswordFlag, UsernameFlag}
	for _, flag := range AllFlag {
		if f.IsSet(flag) {
			setflags = append(setflags, flag)
		}
	}
	return setflags
}