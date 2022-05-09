package main

import (
	"github.com/jprieto92/golang-training/learnWithTests/mocking/countdown/countdown"
	"os"
	"time"
)

func main() {
	sleeper := &countdown.ConfigurableSleeper{1 * time.Second, time.Sleep}
	countdown.Countdown(os.Stdout, sleeper)
}
