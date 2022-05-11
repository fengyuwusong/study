package main

import (
	"context"
	"fmt"
	"strconv"
)

type ClientV2 struct {
}

type ConfigV2 struct{}

// Get 获取key对应的value.
func (c *ClientV2) Get(ctx context.Context, key string) (string, error)

func NewClientV2(serviceName string, config *ConfigV2) (*ClientV2, error)

type DemotionClient struct {
	*ClientV2
}

func NewDemotionClient(serviceName string, config *ConfigV2) (*DemotionClient, error) {
	clientV2, err := NewClientV2(serviceName, config)
	if err != nil {
		return nil, err
	}
	client := &DemotionClient{clientV2}
	return client, nil
}

// GetInt parse value to int
func (d *DemotionClient) GetInt(ctx context.Context, key string) (int, error) {
	value, err := d.Get(ctx, key)
	if err != nil {
		return 0, err
	}
	ret, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("GetInt Error: Key = %s; value = %s is not int", key, value)
	}
	return ret, nil
}

// GetBool parse value to bool:
//     if value=="0" return false;
//     if value=="1" return true;
//     if value!="0" && value!="1" return error;
func (d *DemotionClient) GetBool(ctx context.Context, key string) (bool, error) {
	// 类似GetInt方法
	return false, nil
}
