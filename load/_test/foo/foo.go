package foo

type Example struct {
	a int
	b int
}

func NewExample() *Example {
	return &Example{1, 2}
}

func (e *Example) Sum() int {
	return e.a + e.b
}
