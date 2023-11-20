package redis_message

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"runtime"
	"time"
	"tom_example/global"
)

//  go run main.go
// 要运行多个
var ListenerCmd = &cobra.Command{
	Use:   "redis_listener",
	Short: "redis 监听",
	Run: func(cmd *cobra.Command, args []string) {
		//node := "default"
		//if len(args) > 0 {
		//	node = args[0]
		//}
		ch := make(chan bool, 5)
		client := global.RedisClient
		for {
			msg, err := client.BRPop(context.Background(), 2*time.Second, "test_msg").Result()
			if err == redis.Nil {
				fmt.Println("got no message, till next loop.")
				continue
			}
			if err != nil {
				continue
			}
			ch <- true
			fmt.Println("NumGoroutine num", runtime.NumGoroutine())
			go processMessage(msg[1], msg[0], ch)
		}
	},
}

func processMessage(msg, channel string, ch chan bool) {
	time.Sleep(time.Second)
	fmt.Printf("%s on key:%s \n", msg, channel)
	<-ch
}
