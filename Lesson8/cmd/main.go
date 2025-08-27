package main

import (
	"flag"
	"time"
)

func main() {
	var (
		port        = flag.Int("port", 8080, "gRPC server port")
		maxChats    = flag.Int("max-chats", 100, "max number of chats")
		maxMessages = flag.Int("max-msgs", 1000, "max messages per chat")
		cleanup     = flag.Duration("cleanup", time.Minute, "cleanup interval")
	)
	flag.Parse()
}
