package config

import (
	"time"

	"github.com/adjust/rmq/v5"
	"github.com/opa-oz/simple-queue/pkg"
	"github.com/opa-oz/simple-queue/pkg/utils"
	"github.com/spf13/viper"
)

type configFile struct {
	Targets pkg.Targets `yaml:"targets"`
}

func GetTargets(cfg *Environment) (*pkg.Targets, error) {
	viper.SetConfigName(".config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(cfg.ConfigFile)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config configFile
	viper.Unmarshal(&config)

	return &config.Targets, nil
}

func PrepareQueues(connection *rmq.Connection, targets *pkg.Targets, startImmediately bool) (*pkg.RMQueues, error) {
	queues := make(pkg.RMQueues)
	for target := range *targets {
		qName := utils.GetQ(target)
		q, err := (*connection).OpenQueue(qName)
		if err != nil {
			return nil, err
		}

		queues[qName] = &q
		if startImmediately {
			q.StartConsuming(10, time.Second)
		}
	}

	return &queues, nil
}
