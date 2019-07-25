package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	url := "http://www.baidu.com"
	log, _ := zap.NewProduction()
	defer log.Sync() // flushes buffer, if any
	sugar := log.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
	log.Info("User does not exist", zap.Int("uid", 123))
	log.Error("User does not exist", zap.Int("uid", 123))


}
