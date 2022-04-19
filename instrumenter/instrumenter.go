package instrumenter

type Instrumenter interface {
	Instrumenter(string) ([]byte, error)
}
