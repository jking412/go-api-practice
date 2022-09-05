package cmd

import (
	"github.com/spf13/cobra"
	"go-api-practice/pkg/console"
	"go-api-practice/pkg/redis"
	"time"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

func runPlay(cmd *cobra.Command, args []string) {
	redis.Redis.Set("hello", "hi from redis", 10*time.Second)
	console.Success(redis.Redis.Get("hello"))
}
