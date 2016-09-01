package processor

type Processor interface {
	GetTopic() string
	GetOffset() int64
	ProcessMessage(key string, value string) (int, error)
}
