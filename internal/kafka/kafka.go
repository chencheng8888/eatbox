package kafka

import (
	"eat_box/pkg/errcode"
	"fmt"

	"github.com/IBM/sarama"
)

type Producer struct {
	Pd sarama.SyncProducer
}

func NewProducer(addrs []string, config *sarama.Config) Producer {
	producer, err := sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println(err)
		return Producer{}
	}
	p := Producer{
		Pd: producer,
	}
	return p
}
func (p Producer) SendMsg(msg *sarama.ProducerMessage) *errcode.Error {
	_, _, err := p.Pd.SendMessage(msg)
	if err != nil {
		return errcode.ErrKafkaSend
	}
	return errcode.Success
}
