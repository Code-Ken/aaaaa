package main

import "fmt"

func main(){
	arr := []int{1,2,3,4,5,1,2,3,2}
	quickSort(arr)
	fmt.Println(arr)
}
