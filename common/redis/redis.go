package common

import (
	"context"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type RedisCli struct {
	cli redis.Conn
}

func NewRedisCli() *RedisCli {
	return &RedisCli{}
}

func (cli *RedisCli) Conn() error {
	var err error
	cli.cli, err = redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Printf("redis.Dial() error:%v", err)
		return err
	}
	return nil
}

func (cli *RedisCli) Close() {
	cli.cli.Close()
}

func (cli *RedisCli) SetData(ctx context.Context, key string, data string) error {
	replySet, err := cli.cli.Do("SET", key, data)
	if err != nil {
		fmt.Printf("SET error: %v\n", err)
		return err
	}
	fmt.Println(replySet)
	return nil
}

func (cli *RedisCli) GetData(ctx context.Context, key string) ([]byte, error) {
	reply, err := cli.cli.Do("GET", key)
	if err != nil {
		fmt.Printf("SET error: %v\n", err)
		return []byte{}, err
	}

	ret, _ := reply.([]byte)
	return ret, nil
}
