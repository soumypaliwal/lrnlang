package main

import "fmt"

func main() {
	list := []int{45, 67, 23, 42, 6, 8}
	fmt.Println(list)
	//fmt.Println("Selection Sort :o(n^2): ", selectionSort(list))
	//fmt.Println("Bubble Sort :o(n^2) : ", bubbleSort(list))
	fmt.Println("Insertion sort :O(n^2)", insertionSort(list))
}

// take the element and move it to right direction
func insertionSort(list []int) []int {
	for i := 0; i < len(list); i++ {
		fmt.Println("i=> ", list)
		for j := i; j > 0 && list[j-1] > list[j]; {
			swap(j-1, j, list)
			fmt.Println(list)
			j--
		}
	}

	return list
}

// comppare with adjuscent value and swap
func bubbleSort(list []int) []int {
	isSwaped := false
	for i := len(list) - 1; i >= 1; i-- {
		fmt.Println("i=> ", list)
		for j := 0; j < i; j++ {
			if list[j] > list[j+1] {
				swap(j, j+1, list)
				fmt.Println(list)
				isSwaped = true
			}
		}

		if !isSwaped {
			break
		}
	}
	return list
}

//find thin minimum value index and swap
func selectionSort(list []int) []int {
	for i := 0; i < len(list)-1; i++ {
		indexOfmin := i
		for j := i; j < len(list); j++ {
			if list[indexOfmin] > list[j] {
				indexOfmin = j
			}
		}
		swap(i, indexOfmin, list)
	}
	return list
}
func swap(x int, y int, list []int) []int {
	temp := list[x]
	list[x] = list[y]
	list[y] = temp
	return list
}
