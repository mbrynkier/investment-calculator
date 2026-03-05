package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	//investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	//futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("The future value of our investment is: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("The future real value of our investment is: %.2f\n", futureRealValue)

	//fmt.Println(""The future value of our investment is:", futurelValue)
	//fmt.Println("The future real value of our investment is: ", futureRealValue)

	//fmt.Printf("The future value of our investment is: %.2f\n The future real value of our investment is: %.2f", futureValue, futureRealValue)
	fmt.Print(formattedFV, formattedFRV)

}

func ouputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
