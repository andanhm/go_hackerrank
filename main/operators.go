package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"log"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var mealCostInput string
	var tipPercentInput string
	var taxPercentInput string
	scanner.Scan()
	mealCostInput = scanner.Text()
	scanner.Scan()
	tipPercentInput = scanner.Text()
	scanner.Scan()
	taxPercentInput = scanner.Text()

	mealCost, err := strconv.ParseFloat(mealCostInput,64); if err != nil {
		log.Fatal(err)
	}
	tipPercent,err:=strconv.ParseFloat(tipPercentInput,64); if err != nil {
		log.Fatal(err)
	}
	taxPercent,err:=strconv.ParseFloat(taxPercentInput,64); if err != nil {
		log.Fatal(err)
	}

	var tip float64 = mealCost * (tipPercent/100)
	var tax float64 = mealCost * (taxPercent/100)
	var totalCost = float64(mealCost + tip +tax)
	fmt.Printf("The total meal cost is %d dollars.",Round(totalCost))
}

func Round(val float64) int {
	if val < 0 {
		return int(val-0.5)
	}
	return int(val+0.5)
}