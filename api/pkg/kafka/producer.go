package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

var topic = "test-topic"
var producer sarama.AsyncProducer

func InitProducer() {
	config := sarama.NewConfig()
	client, err := sarama.NewClient([]string{"localhost:9092"}, config)
	if err != nil {
		hlog.Fatalf("new kafka client failed: %s", err.Error())
	}
	producer, err = sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		hlog.Fatalf("new kafka producer failed: %s", err.Error())
	}
}

func Send(data []byte) {
	encoder := sarama.ByteEncoder(data)
	producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: encoder}

}
