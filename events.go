package gobserve

type EventInterface interface {
	Name() string
	Source() string
}
