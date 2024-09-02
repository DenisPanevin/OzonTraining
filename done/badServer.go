


func main() {
	/*var in *bufio.Reader
	var out *bufio.Writer

	file, err := os.Open("./cases/15")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	in = bufio.NewReader(file)
	out = bufio.NewWriter(os.Stdout)

	defer out.Flush()
	var testcases [][]int

	input := 0
	fmt.Fscan(in, &input)
	k := input

	for i := 0; i < k; i++ {
		var testcase []int
		fmt.Fscan(in, &input)
		s := input
		for i := 0; i < s; i++ {
			a := 0
			fmt.Fscan(in, &a)
			testcase = append(testcase, a)

		}

		testcases = append(testcases, testcase)

	}

	for _, v := range testcases {
		fmt.Fprint(out, fmt.Sprintf("%d\n", BadServer(v)))
	}
	*/
	//test2 := []int{7, 1, 4, 1, 9, 1, 1, 9, 1, 7, 9}
	//test2 := []int{1, 2, 3, 4, 5}
	//test2 := []int{1, 1, 2}
	//test2 := []int{3, 5, 5, 4, 4}
	//test2 := []int{3, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 4, 4, 4}

	//test2 := []int{7, 7, 17, 7, 12, 17, 17, 17, 7, 17}
	//println(BadServer(test2))

}

func BadServer(reqArr []int) int {
	//fmt.Printf("input=%v\n", reqArr)
	var users []int
	currLen := 0
	maxLen := 0
	lasti := 0

	for i := 0; i < len(reqArr); i++ {
		if isUnique(reqArr[i], users) {
			if len(users) < 2 {
				users = append(users, reqArr[i])
				currLen++
			} else {
				if maxLen < currLen {
					maxLen = currLen
				}
				//fmt.Printf("uniq=%vvalue%d\n", users, reqArr[i])
				//fmt.Printf("maxlen=%d\n", maxLen)
				lasti = reqArr[i-1]

				for k := i; k > 0; k-- {
					//	fmt.Printf("going back=%d_%d\n", lasti, reqArr[k])
					if reqArr[k-1] != lasti {
						users = []int{reqArr[k]}
						currLen = 1
						i = k
						break
					}
				}

			}
		} else {
			currLen++
			//fmt.Printf("not Unique=%v value%d\n", users, reqArr[i])
			//fmt.Printf("currLen=%d\n", currLen)

		}
	}
	if maxLen < currLen {
		maxLen = currLen
	}
	return maxLen
}

func isUnique(x int, a []int) bool {

	for _, v := range a {
		if v == x {
			return false
		}
	}

	return true
}

/*for i := 0; i < len(reqArr); i++ {
	if len(users) < 2 {
		if isUnique(reqArr[i], users) {
			users = append(users, reqArr[i])
		}
		currLen++
	} else {

		if isUnique(reqArr[i], users) {
			if maxLen < currLen {
				maxLen = currLen
			}
			fmt.Printf("uniq=%vvalue%d\n", users, reqArr[i])
			fmt.Printf("maxlen=%d\n", maxLen)
			if reqArr[i-1] != reqArr[i] {
				users = []int{reqArr[i-1], reqArr[i]}
			} else {
				users = []int{reqArr[i]}
			}
			currLen = 2

		} else {
			currLen++
			fmt.Printf("not Unique=%v value%d\n", users, reqArr[i])
			fmt.Printf("cl=%d\n", currLen)

		}
	}

}
if maxLen < currLen {
	maxLen = currLen
}
return maxLen*/
