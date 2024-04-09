package main

import (
	"fmt"
	"github.com/ediallocyf/math-skills/app"
	"github.com/ediallocyf/math-skills/util"
	"os"
)

func main() {

	args := os.Args[1]
	data, err := util.ReadFileToStr(args)
	if err != nil {
		fmt.Println(err)
	}

	average := app.Average(data)
	fmt.Println("Average: ", average)
	//fmt.Printf("average is type of: %T\n", average)

	median := app.Median(data)
	fmt.Println("Median: ", median)
	//fmt.Printf("Median is type of: %T\n", median)

	variance := app.Variance(data)
	fmt.Println("Variance: ", variance)
	//fmt.Printf("Variance is type of: %T\n", variance)

	StandardDeviation := app.StandardDeviation(data)
	fmt.Println("Standard_Deviation:", StandardDeviation)
	//fmt.Printf("StandardDeviation is type of: %T\n", StandardDeviation)
}
