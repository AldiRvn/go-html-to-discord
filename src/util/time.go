package util

import (
	"math/rand"
	"time"
)

func SleepRandom(min, max int) {
	d := time.Duration(rand.Intn(max-min)+min) * time.Second
	time.Sleep(d)
}
