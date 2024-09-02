package main

import (
	"bufio"
	"fmt"
	"os"
)

type Testcase struct {
	Emp []string
	New []string
}

func main() {
	file, err := os.Open("./logins/1")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Initialize bufio reader and writer
	in := bufio.NewReader(file)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	/*var numOfTests int
	fmt.Fscan(in, &numOfTests)*/

	el := 0
	fmt.Fscan(in, &el)
	emp := make([]string, el)
	for l := 0; l < len(emp); l++ {
		fmt.Fscan(in, &emp[l])
		//println(emp[l])
	}

	fmt.Fscan(in, &el)
	nw := make([]string, el)
	for l := 0; l < len(nw); l++ {
		fmt.Fscan(in, &nw[l])
		//println(emp[l])
	}

	//answ := Compare(nw, emp)
	for _, v := range Compare(nw, emp) {
		fmt.Fprint(out, fmt.Sprintf("%d\n", v))
	}
}

func Compare(newbies, employees []string) []int {
	answers := make([]int, len(newbies))
	for l, v := range newbies {

		for _, k := range employees {
			//fmt.Printf("newbie is %s employee is %s\n", v, k)
			for i := 1; i < len(v); i++ {
				if comparePair(v, k, i-1, i) {
					answers[l] = 1

				}
			}
		}
	}
	return answers
}

func comparePair(n, e string, i, j int) bool {
	if i < 0 || i >= len(e) || j < 0 || j >= len(e) {
		return false
	}
	//fmt.Printf("compare pair of %s and %s\n", string(n[i]), string(n[j]))
	if n == e {
		return true
	}
	runes := []rune(n)

	// Swap the runes at positions i and j
	runes[i], runes[j] = runes[j], runes[i]

	// Convert the slice of runes back to a string
	swapped := string(runes)
	if swapped == e {
		return true
	}

	return false
}
