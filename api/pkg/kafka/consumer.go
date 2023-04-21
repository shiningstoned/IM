package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type ConsumerCallback func(data []byte)

var consumer sarama.Consumer

func InitConsumer() {
	config := sarama.NewConfig()
	client, err := sarama.NewClient([]string{"localhost:9092"}, config)
	if err != nil {
		hlog.Fatalf("new kafka client failed: %s", err.Error())
	}
	consumer, err = sarama.NewConsumerFromClient(client)
	if err != nil {
		hlog.Fatalf("new kafka consumer failed: %s", err.Error())
	}
}

func ConsumeMessage(callBack ConsumerCallback) {
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		hlog.Fatalf("consume partition message failed: %s", err.Error())
	}
	defer partitionConsumer.Close()
	for {
		msg := <-partitionConsumer.Messages()
		callBack(msg.Value)
	}
}
