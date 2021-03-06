package kafkaservice

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

//KafkaSvc will be used as reciver to associate kafka methods to it
type KafkaSvc struct {
	Writer *kafka.Writer
}

//Services is a collection of all the operation required in kafka service
type Services interface {
	WriteToKafka(ctx context.Context, message interface{}) error
}

//NewKafkaService creates new writer to kafka
func NewKafkaService() *KafkaSvc {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{os.Getenv("BROKER")},
		Topic:   os.Getenv("TOPIC"),
	})

	return &KafkaSvc{Writer: w}
}

//WriteToKafka writes message to kafka
func (k *KafkaSvc) WriteToKafka(ctx context.Context, message interface{}) error {

	log.Println("Saving to kafka")
	jsonString, err := json.Marshal(message)
	if err != nil {
		log.Println("Error while saving to kafka. ERROR:", err)
		return err
	}

	err = k.Writer.WriteMessages(ctx, kafka.Message{
		Value: []byte(jsonString),
	})

	if err != nil {
		log.Println("Could not write message. ERROR:", err)
		return err
	}
	return nil
}
