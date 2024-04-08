package main

import (
	"fmt"
	"github.com/ediallocyf/math-skills/api"
	"github.com/ediallocyf/math-skills/util"
	"os"
)

func main() {
	//div := api.DoMath(88, 2)
	//fmt.Println(div)

	args := os.Args[1]

	data, err := util.ReadFileToStr(args)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
	average := api.Average()
	fmt.Println("average", average)
}
