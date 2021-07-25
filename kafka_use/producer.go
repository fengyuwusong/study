package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/Shopify/sarama"
)

// 异步生产者
func InitKafkaProducer() (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // 等待服务器所有副本都保存成功
	//config.Producer.Partitioner = sarama.NewRandomPartitioner // 随机分区类型
	config.Producer.Partitioner = sarama.NewManualPartitioner
	// 是否等待成功和失败后的响应 只有RequireAcks设置不是NoReponse才生效
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V2_7_0_0 //设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	return sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
}

//ProducerMsg 调用该函数生产消息
func ProducerMsg() {
	prod, err := InitKafkaProducer()
	if err != nil {
		panic(err)
	}
	defer prod.Close()

	msg := &sarama.ProducerMessage{
		Topic: "log",
		//Key:       sarama.StringEncoder("test"),
		Partition: 0, // config.Producer.Partitioner 为manual时生效
	}
	msgchan := prod.Input()

	for i := 0; i < 100; i++ {
		msg.Value = sarama.StringEncoder("msg id is :" + strconv.Itoa(rand.Intn(100)))
		msgchan <- msg
		select {
		case suc := <-prod.Successes():
			fmt.Println(suc)
		case err := <-prod.Errors():
			fmt.Println(err)
		}
	}
}
