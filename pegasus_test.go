package main

import (
	"testing"
	"github.com/truppert/pegasus/processor"
	"github.com/Shopify/sarama"
)

func TestPegasus(t *testing.T) {
	t.Skip("This function would block")
	pegasus := NewPegasus("kafka:9092")
	exampleProcessor, _ := processor.NewExampleProcessor("example", sarama.OffsetNewest)
	pegasus.StartProcessor(exampleProcessor)
}

func TestNewPegasus(t *testing.T) {
	pegasus := NewPegasus("kafka:9092")
	if(len(pegasus.brokerList) != 1) {
		t.Error("Expected array of length 1, got: ", len(pegasus.brokerList))
	}
	if(pegasus.brokerList[0] != "kafka:9092") {
		t.Error("Expected [\"kafka:9092\"], got: ", pegasus.brokerList)
	}
}
