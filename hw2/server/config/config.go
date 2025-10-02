package config

import (
	"time"
)

type Config struct {
	Addr            string
	IOTimeout       time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
	HardMinSleep    time.Duration
	HardMaxSleep    time.Duration
	HardErrPct      int
}

func CreateConfig() Config {
	config := Config{
		Addr:            ":8080",
		IOTimeout:       time.Second * 15,
		ShutdownTimeout: time.Second * 10,
		IdleTimeout:     time.Second * 60,
		HardMinSleep:    time.Second * 10,
		HardMaxSleep:    time.Second * 20,
		HardErrPct:      50,
	}
	return config
}
