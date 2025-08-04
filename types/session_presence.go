package types

type SessionPresentFlag byte

const (
	SessionNotPresent SessionPresentFlag = 0
	SessionWasPresent  SessionPresentFlag = 1
)

func (f SessionPresentFlag) Bool() bool {
	return f == SessionWasPresent
}

func (f SessionPresentFlag) String() string {
	if f == SessionWasPresent {
		return "SessionWasPresent"
	}
	return "SessionNotPresent"
}
