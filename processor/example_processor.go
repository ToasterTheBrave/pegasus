package processor

import (
	"log"
	"fmt"
)

type ExampleProcessor struct{
	topic string
	offset int64
}

func (processor ExampleProcessor) GetTopic() string {
	return processor.topic
}

func (processor ExampleProcessor) GetOffset() int64 {
	return processor.offset
}

func NewExampleProcessor(topic string, offset int64) (*ExampleProcessor, error) {
	c := new(ExampleProcessor)
	if topic == "" {
		return nil, fmt.Errorf("topic can not be empty")
	}
	c.topic = topic

	c.offset = offset

	return c, nil
}

func (processor ExampleProcessor) ProcessMessage(key string, value string) (int, error) {
	log.Printf("Message received: ['%s':'%s']\n", key, value)
	return 1, nil
}
