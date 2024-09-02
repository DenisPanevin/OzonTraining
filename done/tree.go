package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
вершины дерева закодированы так что первое число кода вершины это ее номер, второе это кол-во сыновей,
затем номера сыновей в произвольном порядке
требуется восстановить дерево по его коду
*/
type Branch struct {
	Name     int
	Children []Branch
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer

	file, err := os.Open("./tree/1")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)

	defer out.Flush()
	var testcases [][]Branch

	numOfTests := 0
	fmt.Fscan(in, &numOfTests)
	for i := 0; i < numOfTests; i++ {
		branchCodeLen := 0
		fmt.Fscan(in, &branchCodeLen)
		branchCode := []int{}
		for l := 0; l < branchCodeLen; l++ {
			element := 0
			fmt.Fscan(in, &element)
			branchCode = append(branchCode, element)
		}

		testcases = append(testcases, constructbranches(branchCode))

	}

	for _, v := range testcases {
		fmt.Fprint(out, fmt.Sprintf("%d\n", FindRoot(v)))
	}

}

func constructbranches(code []int) []Branch {
	//fmt.Printf("input code is %v\n", code)
	branches := []Branch{}
	excludeIndexes := []int{}
	cleanedCode := []int{}
	if len(code) > 2 {
		for i := 0; i < len(code); i++ {
			if code[i] == 0 {
				excludeIndexes = append(excludeIndexes, i)
				excludeIndexes = append(excludeIndexes, i-1)
			}
		}
		//fmt.Printf("indexes are %v\n", excludeIndexes)
		for v, s := range code {
			if !contains(excludeIndexes, v) {
				cleanedCode = append(cleanedCode, s)
			}
		}
		//fmt.Printf("clean code is %v\n", cleanedCode)
		for i := 0; i < len(cleanedCode); i++ {
			b := Branch{Name: cleanedCode[i]}
			i++
			cheldren := make([]int, cleanedCode[i])
			for v := 0; v < len(cheldren); v++ {
				i++
				child := Branch{Name: cleanedCode[i], Children: []Branch{}}
				b.Children = append(b.Children, child)
			}
			branches = append(branches, b)
		}
		//fmt.Printf("branches is %v\n", branches)
		return branches
	} else {
		return []Branch{
			{
				Name:     code[0],
				Children: []Branch{},
			},
		}
	}

}
func contains(slice []int, element int) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func FindRoot(branches []Branch) int {
	//fmt.Printf("branches is %v len=%d\n", branches, len(branches))
	leaves := []*Branch{}

	for len(branches) != 1 {
		element := branches[0]
		branches = branches[1:]
		for v := range branches {
			FindLeaf(&branches[v], &leaves)
		}
		//printLeaves(leaves)
		if !compareBranches(element, &leaves) {
			branches = append(branches, element)
		} else {
			//println("found parent")
		}
		//fmt.Printf("branches now  %v len=%d\n", branches, len(branches))
	}
	return branches[0].Name
	//println(branches[0].Name)
}

func FindLeaf(branch *Branch, leaves *[]*Branch) {
	for i := range branch.Children {
		if len(branch.Children[i].Children) != 0 {
			FindLeaf(&branch.Children[i], leaves)
		} else {
			*leaves = append(*leaves, &branch.Children[i])
		}
	}
}

func compareBranches(element Branch, leaves *[]*Branch) bool {
	for _, v := range *leaves {
		if element.Name == v.Name {
			*v = element
			//fmt.Printf("\033[33m%s,%d\033[0m\n", "found", v.Name)
			return true
		}
	}
	return false
}
func printLeaves(leaves []*Branch) {
	fmt.Printf("\033[32m%s\033[0m\n", "leaves are")
	for v := range leaves {
		fmt.Printf("\033[32m%v\033[0m\n", leaves[v].Name)
	}
}
