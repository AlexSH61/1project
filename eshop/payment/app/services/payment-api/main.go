package main

import (
	"fmt"
	"os"

	"github.com/AlexSH61/payment/foundation/logger"
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
	if err := Run(log); err != nil {
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
	fmt.Printf("%+v\n", cfg)

	log.Infow("startup", "STATUS", "OK!")
	log.Infow("cfg", "ENV", cfg.ENV)
	return nil
}
