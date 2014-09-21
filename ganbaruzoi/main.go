package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	dec := [][]string{{"今日", "ぞい"}, {"も"}, {"1", "ぞい"}, {"日", "ぞい"}, {"がん", "ぞい"}, {"ばる", "ぞい"}, {"ぞい！"}}
	rand.Seed(time.Now().Unix())
	for i := 0; ; i++ {
		str := ""
		for _, arr := range dec {
			str += arr[rand.Intn(len(arr))]
		}
		fmt.Println(str)
		if str == "今日も1日がんばるぞい！" {
			fmt.Printf("がんばるまで%dzoiでした\n", i)
			return
		}
	}

}
