package gobserve

//EventInterface This interface is the "event enveloppe" for dispatching elements through all subscribes
type EventInterface interface {
	//Name the name of the event
	Name() string
}
