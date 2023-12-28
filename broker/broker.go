package broker

type Producer interface {
	Produce(message []byte) error
}
