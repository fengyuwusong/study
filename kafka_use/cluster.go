package main

import "github.com/Shopify/sarama"

func ClusterMge() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0 //版本要高于2.4
	admin, err := sarama.NewClusterAdmin([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}
	//这里添加两个分区，也可以做其他操作
	err = admin.CreatePartitions("log", 3, [][]int32{{0}, {0}}, false)
	if err != nil {
		panic(err)
	}
	admin.Close()
}
