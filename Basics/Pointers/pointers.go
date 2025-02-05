//Go uses pointers, but no pointer arithmetic.
package main
import "fmt"
func main() {
	x:=10
	p:=&x // Pointer to x
	*p=20 // Dereference and update value
	fmt.Println(x)
}

/*
x := 10 → A variable x is declared and initialized with the value 10.
p := &x → A pointer p is created, storing the memory address of x.
*p = 20 → The value at the memory address p (which is x) is updated to 20.
fmt.Println(x) → Since x was modified through the pointer, printing x outputs 20.
*/