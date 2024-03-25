package sortPkg

import "log"

func BubbleSort(arr []int) {
	for i := range arr {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				swap(&arr, j, j+1)
			}
		}
	}
	log.Println(arr)
}

func swap(arr *[]int, x int, y int) {
	aux := (*arr)[x]
	(*arr)[x] = (*arr)[y]
	(*arr)[y] = aux
}
