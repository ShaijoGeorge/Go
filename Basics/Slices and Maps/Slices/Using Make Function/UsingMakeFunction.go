package main

import "fmt"

func main(){
	nums := make([]int, 3) // Creates a slice with length 3 and capacity 3
    nums[0]=3
    nums[1]=4
    nums[2]=5
    fmt.Println(nums)
}