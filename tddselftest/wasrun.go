package tddselftest

type WasRun struct {
	wasRun bool
	name   string
}

func NewWasRun(name string) *WasRun {
	return &WasRun{false, name}
}

func (w *WasRun) run() {
	w.testMethod()
}

func (w *WasRun) testMethod() {
	w.wasRun = true
}
