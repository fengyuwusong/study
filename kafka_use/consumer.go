package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func InitKafkaConsumer() (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V2_7_0_0
	return sarama.NewConsumer([]string{"localhost:9092"}, config)
}

//ConsumerMsg 调用该函数消费消息
func ConsumerMsg() {
	cons, err := InitKafkaConsumer()
	if err != nil {
		panic(err)
	}
	defer cons.Close()
	//获取分区数
	res, _ := cons.Partitions("log")
	fmt.Println(res)

	// sarama.OffsetOldest: 从最初数据获取 sarama.OffsetNewest: 从最新投递数据开始获取
	pc, err := cons.ConsumePartition("log", 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	msgchan := pc.Messages()
	errchan := pc.Errors()
	for {
		select {
		case msg := <-msgchan:
			fmt.Println("msg:", msg.Partition, string(msg.Key), string(msg.Value), msg.Timestamp.String())
		case err = <-errchan:
			fmt.Println(err)
		}
	}
}
