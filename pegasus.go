package pegasus

import (
	"flag"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"strings"
	"github.com/truppert/pegasus/processor"
)

type Pegasus struct {
	brokerList []string
}

// TODO : use a broker list
func NewPegasus(brokers string) *Pegasus {

	pegasus := new(Pegasus)

	if brokers == "" {
		log.Println("brokers can not be empty")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Split our brokers into an array
	brokerList := strings.Split(brokers, ",")

	pegasus.brokerList = brokerList

	return pegasus
}

func (pegasus Pegasus) StartProcessor(processor processor.Processor) {
	// Start our consumer
	consumer, err := startConsumer(pegasus.brokerList, processor)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Consume messages forever
	consumeMessages(consumer, processor)

}

func startConsumer(brokerList []string, processor processor.Processor) (sarama.PartitionConsumer, error) {
	master, err := sarama.NewConsumer(brokerList, nil)
	if err != nil {
		return nil, err
	}

	consumer, err := master.ConsumePartition(processor.GetTopic(), 0, processor.GetOffset())
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func consumeMessages(consumer sarama.PartitionConsumer, processor processor.Processor) error {
	log.Printf("Consuming topic: %s\n", processor.GetTopic())
	for {
		select {
		case message := <-consumer.Messages():
			_, err := processor.ProcessMessage(string(message.Key), string(message.Value))
			if err != nil {
				return err
			}
		case err := <-consumer.Errors():
			return err
		}
	}
}
