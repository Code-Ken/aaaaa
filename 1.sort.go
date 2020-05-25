package main

import "fmt"

func main(){
	arr := []int{1,2,3,4,5,1,2,3,2}
	quickSort(arr)
	fmt.Println(arr)
}

//1.1 快排
//=============================================

func quickSort(arr []int){
	l := len(arr)-1
	recursionQuickSort(arr,0,l)
}

func recursionQuickSort(arr []int,begin int,end int){
	if end < begin{
		return
	}
	p := partition(arr,begin,end)
	recursionQuickSort(arr,begin,p-1)
	recursionQuickSort(arr,p+1,end)
}

func partition(arr []int,begin int,end int) int {
	p := end
	counter := begin

	for i:=begin;i<p;i++{
		if arr[i] < arr[p]{
			arr[i],arr[counter] = arr[counter],arr[i]
			counter++
		}
	}
	arr[p],arr[counter] = arr[counter],arr[p]
	return counter
}

//1.2 归并
//=============================================

func mergeSort(arr []int){
	recursionMergeSort(arr,0,len(arr)-1)
}

func recursionMergeSort(arr []int,left int,right int){
	if right <= left{
		return
	}
	mid := (left+right)>>1
	recursionMergeSort(arr,left,mid)
	recursionMergeSort(arr,mid+1,right)
	merge(arr,left,mid,right)
}

func merge(arr []int,left int,mid int,right int){
	tmpArr := make([]int,right-left+1)
	p:=0
	i:= left
	j := mid+1
	for i<=mid &&j<=right{
		if arr[i] < arr[j]{
			tmpArr[p] = arr[i]
			i++
		}else{
			tmpArr[p] = arr[j]
			j++
		}
		p++
	}

	for i<=mid{
		tmpArr[p] = arr[i]
		i++
		p++
	}

	for j<=right{
		tmpArr[p] = arr[j]
		j++
		p++
	}
	for k:=0;k<len(tmpArr);k++{
		arr[left+k] = tmpArr[k]
	}
}



//1.3 堆排序
//完全二叉树
//=============================================
//堆排序从h-1层构建
// parent = (i-1)>>1
// children1 = 2*i+1
// children2 = 2*i+2


func heapSort(tree []int){
	n := len(tree)
	buildHeap(tree,n)
	for i:=n-1;i>=0;i--{
		tree[0],tree[i] = tree[i],tree[0]
		heapify(tree,i,0)
	}
}

//n总节点
//i对哪一个节点做heapify
func heapify(tree []int,n int,i int){
	if i >= n{
		return
	}
	c1 := 2*i+1
	c2 := 2*i+2
	max := i

	if c1 < n && tree[max] < tree[c1]{
		max = c1
	}
	if c2 < n && tree[max] < tree[c2]{
		max = c2
	}
	if max != i{
		tree[max],tree[i] = tree[i],tree[max]
		heapify(tree,n,max)
	}
}

func buildHeap(tree []int,n int){
	lastNode := n - 1
	parent := (lastNode - 1) >> 1

	for i:=parent;i>=0;i--{
		heapify(tree,n,i)
	}
}




