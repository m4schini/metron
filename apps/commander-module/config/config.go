package config

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"time"
)

func readConfig() (*yaml.File, error) {
	file, err := yaml.ReadFile("config/metron.yaml")
	if err != nil {
		return nil, err
	}

	return file, nil
}

func GetConfig() *yaml.File {
	config, err := readConfig()

	fmt.Println(config)

	val, err := config.Get("targets[1].account")
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

	return config
}

type Target struct {
	Account  string
	Interval time.Duration
}

func GetTargets() []*Target {
	targets := make([]*Target, 0)

	config, e := readConfig()
	if e != nil {
		return targets
	}

	formatGetAccount := `targets[%d].account`
	formatGetInterval := `targets[%d].interval`

	var err error
	for i := 0; err == nil; i++ {
		var name, interval string
		name, err = config.Get(fmt.Sprintf(formatGetAccount, i))
		if err != nil {
			break
		}

		interval, err = config.Get(fmt.Sprintf(formatGetInterval, i))
		var duration time.Duration
		if err != nil {
			duration = 15 * time.Minute
		} else {
			duration, err = time.ParseDuration(interval)
			if err != nil {
				duration = 15 * time.Minute
			}
		}

		targets = append(targets, &Target{
			Account:  name,
			Interval: duration,
		})
	}

	return targets
}
