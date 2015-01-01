package main

import (
	"math/rand"
	"time"
)

var omikuji = []string{
	"大吉",
	"中吉",
	"小吉",
	"吉",
	"半吉",
	"末吉",
	"末小吉",
	"凶",
	"小凶",
	"半凶",
	"末凶",
	"大凶",
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	result := omikuji[rand.Intn(len(omikuji))]
	println(result)
}
