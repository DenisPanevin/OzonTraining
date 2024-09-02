package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/*Вы живете в городе прямоугольной формы с длиной n и шириной m, в котором находится k различных видов природных ресурсов.
Вам известны месторождения каждого ресурса в виде списка координат
[x,y], где
x — позиция по длине города n, а
y — по ширине m.
Чтобы сократить объемы расходов на освоение ресурсов, найдите участок города:
с минимально возможной площадью;
прямоугольной формы со сторонами параллельными сторонам города;
с хотя бы одним месторождением каждого ресурса.
*/

type Res struct {
	X int
	Y int
}
type Testcase struct {
	Id        int
	CitySizeX int
	CitySizeY int
	Resources [][]Res
}

var permutations = [][]int{
	{1, 2, 3, 4},
	{1, 2, 4, 3},
	{1, 3, 2, 4},
	{1, 3, 4, 2},
	{1, 4, 2, 3},
	{1, 4, 3, 2},
	{2, 1, 3, 4},
	{2, 1, 4, 3},
	{2, 3, 1, 4},
	{2, 3, 4, 1},
	{2, 4, 1, 3},
	{2, 4, 3, 1},
	{3, 1, 2, 4},
	{3, 1, 4, 2},
	{3, 2, 1, 4},
	{3, 2, 4, 1},
	{3, 4, 1, 2},
	{3, 4, 2, 1},
	{4, 1, 2, 3},
	{4, 1, 3, 2},
	{4, 2, 1, 3},
	{4, 2, 3, 1},
	{4, 3, 1, 2},
	{4, 3, 2, 1},
}
var cs = [][]int{
	{1, 3, 4, 2},
}

func main() {
	file, err := os.Open("./resources/3")
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
	//var wg sync.WaitGroup
	//var mu sync.Mutex

	// Read number of test cases
	var numOfTests int
	fmt.Fscan(in, &numOfTests)
	testCases := make([]*Testcase, numOfTests)
	//answers := make([]int, numOfTests)
	//wg.Add(numOfTests)
	for i := 0; i < numOfTests; i++ {
		t := new(Testcase)
		t.Id = i
		fmt.Fscan(in, &t.CitySizeX)
		fmt.Fscan(in, &t.CitySizeY)
		numOfRes := 0
		fmt.Fscan(in, &numOfRes)
		t.Resources = make([][]Res, numOfRes)

		for n := range t.Resources {
			resAmt := 0
			fmt.Fscan(in, &resAmt)
			resource := make([]Res, resAmt)
			for r := range resource {
				rs := Res{}
				fmt.Fscan(in, &rs.X)
				fmt.Fscan(in, &rs.Y)
				resource[r] = rs
			}
			t.Resources[n] = resource
		}
		testCases[i] = t

		//go Boxes(t, &wg, &mu, &answers)
	}

	for _, v := range testCases {
		fmt.Fprint(out, fmt.Sprintf("%d\n", DoPattern(v, permutations)))
	}
	/*a := []int{1, 2, 3, 4, 5}
	a = a[:2]
	fmt.Printf("test %v\n", a)*/
	//a = a[:2] - {1,2}//exclude
	//a = a[2:]  - {3,4,5}//include
	//println(extractResRight(testCases[647], 3))
	//println(DoPattern(testCases[0], cs))

	//wg.Wait()
}

func DoPattern(t *Testcase, premutations [][]int) int {
	//println(t.CitySizeX)
	//println(t.CitySizeY)
	x1 := 100001
	x2 := 0
	y1 := 100001
	y2 := 0
	Area := 10000000000
	orig := make([][]Res, len(t.Resources))
	for i := range orig {
		orig[i] = append([]Res{}, t.Resources[i]...)
	}
	for i := range premutations {
		for s := range orig {
			t.Resources[s] = append([]Res{}, orig[s]...)
		}
		for k := range premutations[i] {
			switch premutations[i][k] {
			case 1:
				//println("RL")
				x1 = CleanR(t) - 1
			case 2:
				//println("TD")
				y2 = CleanT(t)
			case 3:
				//println("LR")
				x2 = CleanL(t)
			case 4:
				//println("BT")
				y1 = CleanB(t) - 1
			}
		}
		if x1 < 0 {
			x1 = 0
		}
		if x2 > t.CitySizeX {
			x2 = t.CitySizeX
		}
		if y1 < 0 {
			y1 = 0
		}
		if y2 > t.CitySizeY {
			y2 = t.CitySizeY
		}
		//fmt.Printf("x1=%d x2=%d y1=%d y2=%d\n", x1, x2, y1, y2)
		mArea := (y2 - y1) * (x2 - x1)
		if mArea != 0 && mArea < Area {
			Area = mArea
		}
	}

	return Area
}

func CleanR(t *Testcase) int {
	for i := range t.Resources {
		sort.Slice(t.Resources[i], func(k, j int) bool {
			return t.Resources[i][k].X < t.Resources[i][j].X
		})
	}
	//fmt.Printf("resources sorted by X=%v\n", t.Resources)
	fairest := []Res{}
	smallest := 100001
	for i := range t.Resources {
		fairest = append(fairest, t.Resources[i][(len(t.Resources[i])-1)])
	}
	for i := range fairest {
		if fairest[i].X < smallest {
			smallest = fairest[i].X
		}
	}
	//fmt.Printf("Smallest X1=%d\n", smallest)

	for i := 0; i < len(t.Resources); i++ {

		if len(t.Resources[i]) > 1 {

			for r := 0; r < len(t.Resources[i]); r++ {
				if t.Resources[i][r].X > smallest {
					//fmt.Printf("%v\n", t.Resources[i][r])
					t.Resources[i] = t.Resources[i][r:]
					break
				}
			}
		}

	}
	return smallest
}
func CleanT(t *Testcase) int {
	for i := range t.Resources {
		sort.Slice(t.Resources[i], func(k, j int) bool {
			return t.Resources[i][k].Y < t.Resources[i][j].Y
		})
	}
	//fmt.Printf("resources sorted by Y=%v\n", t.Resources)
	fairest := []Res{}
	smallest := 0
	for i := range t.Resources {
		fairest = append(fairest, t.Resources[i][0])
	}
	for i := range fairest {
		if fairest[i].Y > smallest {
			smallest = fairest[i].Y
		}
	}
	//fmt.Printf("Smallest Y2=%d\n", smallest)

	for i := 0; i < len(t.Resources); i++ {
		if len(t.Resources[i]) > 1 {
			for r := 0; r < len(t.Resources[i]); r++ {
				if t.Resources[i][r].Y > smallest {
					//fmt.Printf("%v\n", t.Resources[i][r])
					t.Resources[i] = t.Resources[i][:r]
					break
				}
			}
		}

	}
	//fmt.Printf("SAfter trim=%v\n", t.Resources)
	return smallest
}
func CleanL(t *Testcase) int {
	for i := range t.Resources {
		sort.Slice(t.Resources[i], func(k, j int) bool {
			return t.Resources[i][k].X < t.Resources[i][j].X
		})
	}
	//fmt.Printf("resources sorted by X=%v\n", t.Resources)
	fairest := []Res{}
	smallest := 0
	for i := range t.Resources {
		fairest = append(fairest, t.Resources[i][0])
	}
	for i := range fairest {
		if fairest[i].X > smallest {
			smallest = fairest[i].X
		}
	}
	//fmt.Printf("Smallest X2=%d\n", smallest)

	for i := 0; i < len(t.Resources); i++ {
		if len(t.Resources[i]) > 1 {
			for r := 0; r < len(t.Resources[i]); r++ {
				if t.Resources[i][r].X > smallest {
					//fmt.Printf("%v\n", t.Resources[i][r])
					t.Resources[i] = t.Resources[i][:r]
					break
				}
			}
		}

	}
	return smallest
}
func CleanB(t *Testcase) int {
	for i := range t.Resources {
		sort.Slice(t.Resources[i], func(k, j int) bool {
			return t.Resources[i][k].Y < t.Resources[i][j].Y
		})
	}
	//fmt.Printf("resources sorted by Y1=%v\n", t.Resources)
	fairest := []Res{}
	smallest := 100001
	for i := range t.Resources {
		fairest = append(fairest, t.Resources[i][(len(t.Resources[i])-1)])
	}
	for i := range fairest {
		if fairest[i].Y < smallest {
			smallest = fairest[i].Y
		}
	}
	//fmt.Printf("Smallest Y1=%d\n", smallest)
	//fmt.Printf("Fairest are=%v\n", fairest)

	for i := 0; i < len(t.Resources); i++ {
		if len(t.Resources[i]) > 1 {
			for r := 0; r < len(t.Resources[i]); r++ {
				if t.Resources[i][r].Y > smallest {
					//fmt.Printf("%v\n", t.Resources[i][r])
					t.Resources[i] = t.Resources[i][r:]
					break
				}
			}
		}

	}
	//fmt.Printf("Array after trim=%v\n", t.Resources)
	return smallest
}
