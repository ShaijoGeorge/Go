package main

import "fmt"

func main(){
	arr:= [5]int{1,2,3,4,5} //Create an array
	nums:=arr[1:4] //Slice from index 1 to 3 (inclusive start, exclusive end)
	fmt.Println(nums)
}