//structs are used to define custom types.
package main

import "fmt"

type Person struct{
	Name string
	Age int
}

func main(){
	p:=Person{Name:"Shaijo", Age:20}
	fmt.Println(p.Name,p.Age)
}