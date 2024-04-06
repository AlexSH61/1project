package main

import (
	"fmt"
	"os"

	logger "github.com/AlexSH61/payment/foundation"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

func main() {
	log, err := logger.NewLogger()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()
	if err != nil {
		log.Errorw("startup", "ERROR", err)
		os.Exit(1)
	}
}
func Run(log *zap.SugaredLogger) error {
	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return err
	}
	log.Infow("startup", "STATUS", "OK!")
	log.Infow("cfg", "ENV", &cfg.Env)
	return nil
}
