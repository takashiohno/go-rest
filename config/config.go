package config

import (
    "log"
)

type Config struct {
    LogLevel int
    LogFlag int
}

func DefaultConf() Config {
    return Config {
        LogLevel: 4,
        LogFlag: log.Ldate|log.Ltime|log.Lshortfile,
    }
}
