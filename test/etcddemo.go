package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func connect() (client *clientv3.Client, err error) {
	fmt.Println("出发??2")
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 2 * time.Second,
	})
	fmt.Println("出发?3?")
	if err != nil {
		fmt.Println("connect err:", err)
		return nil, err
	}
	return client, err
}

func main() {

	//连接
	cli, err := connect()
	defer cli.Close()
	fmt.Println("出发??")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cli)
	// 获取 etcd读写对象
	kv := clientv3.NewKV(cli)

	//添加键值对
	r1, err := kv.Put(context.TODO(), "/lesson/math", "100")
	if err != nil {
		fmt.Printf("put key1 err:", err)
		return
	}
	// 继续添加键值对
	r2, err := kv.Put(context.TODO(), "/lesson/music", "50")
	if err != nil {
		fmt.Println("put key2 err:", err)
		return
	}
	fmt.Println("添加结果r1: ", r1)
	fmt.Println("添加结果r1: ", r2)

	// 获取整个 /lesson目录下的数据
	getAll, err := kv.Get(context.TODO(), "/lesson", clientv3.WithPrefix())
	if err != nil {
		fmt.Println("select all err: ", err)
		return
	}
	fmt.Println("查询所有：", getAll.Kvs)


}
