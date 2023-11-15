package redis_message

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"time"
	"tom_example/global"
)

// 要运行多个
var ListenerCmd = &cobra.Command{
	Use:   "redis_listener",
	Short: "redis 监听",
	Run: func(cmd *cobra.Command, args []string) {
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

			fmt.Printf("%s on key:%s \n", msg[1], msg[0])
		}
	},
}
