package processor

import (
	"testing"
)

func TestNewProcessor(t *testing.T) {
	processor, _ := NewExampleProcessor("topic", 123)

	if(processor.topic != "topic") {
		t.Error("Expected \"topic\", got: ", processor.topic)
	}

	if(processor.offset != 123) {
		t.Error("Expected 123, got: ", processor.offset)
	}

	_, err := NewExampleProcessor("", 123)
	if(err.Error() != "topic can not be empty") {
		t.Error("Expected \"topic can not be empty\", got: ", err.Error())
	}

}

func TestGetTopic(t *testing.T) {
	processor, _ := NewExampleProcessor("topic", 123)
	if(processor.GetTopic() != "topic") {
		t.Error("Expected \"topic\", got: ", processor.GetTopic())
	}
}

func TestGetOffset(t *testing.T) {
	processor, _ := NewExampleProcessor("topic", 123)
	if(processor.GetOffset() != 123) {
		t.Error("Expected 123, got: ", processor.GetOffset())
	}
}

func TestProcessMessage(t *testing.T) {
	processor, _ := NewExampleProcessor("topic", 123)
	success, _ := processor.ProcessMessage("key", "value")
	if(success != 1) {
		t.Error("Expected 1, got: ", success)
	}
}