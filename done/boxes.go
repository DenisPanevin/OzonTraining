package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"sync"
)

type testCase struct {
	Id        int
	NumOfCars int
	Capacity  int
	Boxes     []int
}

func main() {
	// Open input file
	file, err := os.Open("./casesBoxes/2")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Initialize bufio reader and writer
	in := bufio.NewReader(file)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// Initialize WaitGroup and Mutex
	var wg sync.WaitGroup
	var mu sync.Mutex

	// Read number of test cases
	var numOfTests int
	fmt.Fscan(in, &numOfTests)
	answers := make([]int, numOfTests)
	wg.Add(numOfTests)

	for i := 0; i < numOfTests; i++ {
		t := new(testCase)
		t.Id = i
		fmt.Fscan(in, &t.NumOfCars)
		fmt.Fscan(in, &t.Capacity)
		var numOfBoxes int
		fmt.Fscan(in, &numOfBoxes)
		t.Boxes = make([]int, numOfBoxes)

		for j := 0; j < numOfBoxes; j++ {
			var exp int
			fmt.Fscan(in, &exp)
			t.Boxes[j] = exp
		}

		go Boxes(t, &wg, &mu, &answers)
	}

	wg.Wait()

	for num := range answers {
		out.WriteString(strconv.Itoa(answers[num]) + "\n")
		out.Flush()
	}
}

func Boxes(t *testCase, wg *sync.WaitGroup, mu *sync.Mutex, answers *[]int) {
	defer wg.Done()

	for i := range t.Boxes {

		t.Boxes[i] = 1 << t.Boxes[i]
	}

	sort.Ints(t.Boxes)

	currLoad := 0
	pickupCounter := 0
	carsCounter := 1
	trips := 0
	for pickupCounter < len(t.Boxes) {

		for i := len(t.Boxes) - 1; i > -1; i-- {
			if t.Boxes[i] > 0 && t.Boxes[i]+currLoad <= t.Capacity {
				currLoad += t.Boxes[i]

				t.Boxes[i] = 0
				pickupCounter++

			}
		}

		carsCounter++
		if pickupCounter == len(t.Boxes) {

			trips++
			break
		}

		currLoad = 0
		if carsCounter > t.NumOfCars {
			carsCounter = 1

			trips++

		}

	}

	mu.Lock()
	(*answers)[t.Id] = trips
	mu.Unlock()

}
