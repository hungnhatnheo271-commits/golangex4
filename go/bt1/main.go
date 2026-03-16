package main

import "fmt"

func main() {


	tree1Data := []int{10, 5, 20, 3, 7}
	tree2Data := []int{50, 30, 70, 20, 40}
	tree3Data := []int{8, 4, 12, 2, 6}
	tree4Data := []int{15, 10, 25, 5, 12}


	tree1 := buildTree(tree1Data)
	tree2 := buildTree(tree2Data)
	tree3 := buildTree(tree3Data)
	tree4 := buildTree(tree4Data)

	fmt.Println("Tree 1 DESC:")
	printDESC(tree1)

	fmt.Println("\nTree 2 DESC:")
	printDESC(tree2)

	fmt.Println("\nTree 3 DESC:")
	printDESC(tree3)

	fmt.Println("\nTree 4 DESC:")
	printDESC(tree4)
}