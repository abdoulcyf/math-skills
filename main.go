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
	average := api.Average("data.text")
	fmt.Println("average", average)
	median := api.Median("data.txt")
	fmt.Println("median", median)
	variance := api.Variance("data.txt")
	fmt.Println("variance", variance)
	StandardDeviation := api.StandardDeviation("data.txt")
	fmt.Println("Standard_Deviation", StandardDeviation)
}
