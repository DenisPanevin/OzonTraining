package main

/*
Маркетплейс получает комиссию p% с каждой продажи товаров.
Если был продан товар, который стоил ai рублей, комиссия должна составить ai * p / 100 рублей.
По техническому заданию комиссия должна округляться в меньшую сторону до второго знака после запятой,
то есть до целого числа копеек. Но, из-за допущенной программистом ошибки, комиссия всегда округлялась
в меньшую сторону до целого числа рублей, то есть копейки отбрасывались.

Найдите сумму, упущенную из-за ошибки, после продажи маркетплейсом n товаров со стоимостью ai рублей.
*/
import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type TestCase struct {
	Fee         int64
	WrongPrices []int64
}

func main() {
	file, err := os.Open("./rounds/2")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	in := bufio.NewReader(file)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	numOfTest := 0
	fmt.Fscan(in, &numOfTest)
	TestCases := make([]*TestCase, numOfTest)
	for i := 0; i < numOfTest; i++ {
		t := TestCase{}
		wpl := 0
		fmt.Fscan(in, &wpl)
		t.WrongPrices = make([]int64, wpl)
		fmt.Fscan(in, &t.Fee)
		for k := range t.WrongPrices {
			fmt.Fscan(in, &t.WrongPrices[k])
		}
		TestCases[i] = &t

	}
	for i := range TestCases {

		fmt.Fprint(out, fmt.Sprintf(" %.2f\n", Calculate(TestCases[i])))
	}
	//fmt.Printf("%d\n", len((*TestCases[0]).WrongPrices))
}
func Calculate(t *TestCase) float64 {
	//correctPrices := []float32{}
	var f int64
	var ans float64

	for i := range t.WrongPrices {
		f = t.Fee * t.WrongPrices[i]
		cnts := (f % 100)
		ans = float64(cnts)/100.0 + ans
	}
	ans = roundToTwoDecimalPlaces(ans)
	return ans
	//fmt.Printf("the fee  is %.2f\n", ans)

}

func roundToTwoDecimalPlaces(num float64) float64 {
	return math.Round(num*100) / 100
}
