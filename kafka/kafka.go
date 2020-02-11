package kafka

import (
	log "github.com/sirupsen/logrus"

	"github.com/Shopify/sarama"
)

func SendMessage(brokers []string, topic string, payload string) {
	log.Info("Initializing Kafka")

	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5

	// brokers := []string{"192.168.59.103:9092"}
	//brokers := []string{"localhost:9092"}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	//topic := "test"
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(payload),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
}
