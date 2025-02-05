package main
import "fmt"
func main() {
	nums:=[]int{1,2,3} // Step 1: Create a slice

	nums=append(nums,4,5) // Step 2: Append new elements
	fmt.Println("After append:",nums) // Output: [1 2 3 4 5]

	// Step 3: Remove an element (e.g., remove element at index 2)
	index:=2 
	nums=append(nums[:index], nums[index+1:]...) 
	fmt.Println("After removal:", nums) // Output: [1 2 4 5]
}