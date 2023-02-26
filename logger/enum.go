package logger

type RunStatus int

const (
	OK RunStatus = iota
	Fail
	Skip
	Yet
)

var runStatusStrings = [4]string{"✅", "❌", "⏸", "⛔"}

func (s RunStatus) String() string {
	return runStatusStrings[s]
}
