package etcd

import (
	"go.etcd.io/etcd/clientv3"
	"time"
	"fmt"
)

func Conn() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer cli.Close()
}
