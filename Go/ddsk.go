package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const ddskTemp string = "ドドスコスコスコドドスコスコスコドドスコスコスコ"
	ddsk := [2]string{"ドド", "スコ"}
	text := ""
	flag := false
	n := 0

	for flag == false {
		rand.Seed(time.Now().UnixNano())
		t := rand.Intn(2)
		fmt.Println(ddsk[t])
		text += ddsk[t]
		if len(text) >= 3*24 {
			if text[2*n:] == ddskTemp {
				fmt.Println("ラブ注入♡")
				flag = true
			}
			n += 3
		}
	}
}
