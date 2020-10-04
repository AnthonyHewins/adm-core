package util

import (
	"os"
	"strconv"
)

func EnvInt(name string, def int64) int64 {
	val, wasSet := os.LookupEnv(name)

	if !wasSet {
		return def
	}

	if integer, err := strconv.ParseInt(val, 10, 10); err != nil {
		panic(err)
	} else {
		return integer
	}
}
