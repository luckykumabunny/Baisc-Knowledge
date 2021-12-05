package main

import "fmt"

func main() {
   fmt.Println("Hello, World!")
	var array =[]int{56,99,-35,0,8}
	quick_sort(array, 0, len(array)-1)
	fmt.Println(array)
}

func quick_sort(array []int, left int, right int){
	if left > right{
		return
	}
	
	low := left
	high := right
	povit := array[left]
	
	for high > low{
		for high > low && array[high] > povit{
			high--
		}
		array[low] = array[high]
		
		for high > low && array[low] <= povit{
			low++
		}
		array[high] = array[low]
	}
	array[low] = povit
	
	quick_sort(array, left, low-1)
	quick_sort(array, low+1, right)
}