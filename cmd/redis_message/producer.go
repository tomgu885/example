package redis_message

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"time"
	"tom_example/global"
)

var ProducerCmd = &cobra.Command{
	Use:   "redis_produce",
	Short: "producer message",
	Run: func(cmd *cobra.Command, args []string) {
		client := global.RedisClient
		fmt.Println("starting sending")
		for i := 0; i < 1_000; i++ {
			t := time.Now().UnixMicro() / 1000
			client.LPush(context.Background(), "test_msg", fmt.Sprintf("ms_%d_%d", i, t))

			if i%100 == 0 {
				fmt.Printf("sent %d\n", i)
			}

			//u := rand.Int63n(100) + 50
			//time.Sleep(time.Duration(u) * time.Millisecond)

		}

		fmt.Println("sent all")
	},
}
