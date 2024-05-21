package kafka

import (
	"eat_box/global"
	"eat_box/internal/dao"
	"eat_box/internal/model/swagger"
	"encoding/json"
	"fmt"
	"time"

	"github.com/IBM/sarama"
)

func ListenScore() {
	consumer, err := sarama.NewConsumer([]string{global.KafkaSetting.Addr}, global.KafkaConfig)
	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}
	datas := make([]swagger.ScoreData, 0)
	cp, err := consumer.ConsumePartition("score", 0, sarama.OffsetNewest)
	d := dao.NewDao(global.DBEngine)
	for {
		select {
		case msg := <-cp.Messages():
			var data swagger.ScoreData
			err := json.Unmarshal(msg.Value, &data)
			if err != nil {
				//处理错误
				fmt.Println(err)
			}
			datas = append(datas, data)
			if len(datas) >= 100 {
				//处理data
				err = d.UpdateBusinessesScore(datas)
				if err != nil {
					//处理错误
					fmt.Println(err)
				}
				//清空切片
				datas = datas[:0]
			}
		case <-time.After(time.Second * 1):
			//处理data
			if len(datas) > 0 {
				err = d.UpdateBusinessesScore(datas)
				if err != nil {
					//处理错误
					fmt.Println(err)
				}
				//清空切片
				datas = datas[:0]
			}
		}
	}
}
