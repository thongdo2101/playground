package main

import "fmt"

func main() {
	randomArr := []int{4, 7, 2, 6, 1, 3, 5}
	fmt.Println("origin arr: ", randomArr)
	// fmt.Println("bubble sort - sorted arr: ", BubbleSort(randomArr))
	// fmt.Println("selection sort - sorted arr: ", SelectionSort(randomArr))
	// fmt.Println("insertion sort - sorted arr: ", InsertionSort(randomArr))
	// fmt.Println("merge sort - sorted arr: ", MergeSort(randomArr))
	// fmt.Println("quick sort - sorted arr: ", QuickSort(randomArr, 0, len(randomArr)-1))
	// fmt.Println("simple quick sort - sorted arr: ", SimpleQuickSort(randomArr))
	fmt.Println("quick sort - sorted arr: ", quickSort(randomArr, 0, len(randomArr)-1))
}

func BubbleSort(originArr []int) []int {
	needSwap := true
	for needSwap {
		needSwap = false
		for i := 0; i < len(originArr)-1; i++ {
			if originArr[i] > originArr[i+1] {
				originArr[i], originArr[i+1] = originArr[i+1], originArr[i]
				needSwap = true
			}
		}
	}
	return originArr
}

func SelectionSort(originArr []int) []int {
	for i := 0; i < len(originArr); i++ {
		minIndex := i
		for j := i + 1; j < len(originArr); j++ {
			if originArr[j] < originArr[minIndex] {
				minIndex = j
			}
		}
		originArr[i], originArr[minIndex] = originArr[minIndex], originArr[i]
	}
	return originArr
}

func InsertionSort(originArr []int) []int {
	for i := 1; i < len(originArr); i++ {
		for j := i; j > 0 && originArr[j] < originArr[j-1]; j-- {
			originArr[j], originArr[j-1] = originArr[j-1], originArr[j]
		}
	}
	return originArr
}

func MergeSort(originArr []int) []int {
	if len(originArr) <= 1 {
		return originArr
	}
	mid := len(originArr) / 2
	leftArr := MergeSort(originArr[:mid])
	rightArr := MergeSort(originArr[mid:])
	return Merge(leftArr, rightArr)
}
func Merge(leftArr, rightArr []int) []int {
	result := make([]int, 0)
	for len(leftArr) > 0 && len(rightArr) > 0 {
		if leftArr[0] < rightArr[0] {
			result = append(result, leftArr[0])
			leftArr = leftArr[1:]
		} else {
			result = append(result, rightArr[0])
			rightArr = rightArr[1:]
		}
	}
	result = append(result, leftArr...)
	result = append(result, rightArr...)
	return result
}

func quickSort(arr []int, low, high int) []int {
	pivot := arr[(low+high)/2]
	i, j := low, high
	for i <= j {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	if low < j {
		quickSort(arr, low, j)
	}
	if i < high {
		quickSort(arr, i, high)
	}
	return arr
}
