package testframe

type WasRun struct {
	wasRun bool
	method func(*WasRun)
}

func NewWasRun(method func(*WasRun)) *WasRun {
	return &WasRun{false, method}
}

func (w *WasRun) run() {
	w.method(w)
}

func (w *WasRun) testMethod() {
	w.wasRun = true
}
