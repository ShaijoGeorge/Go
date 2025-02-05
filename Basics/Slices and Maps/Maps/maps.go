// Maps: Key-value pairs
package main
import "fmt"
func main() {
	user:= map[string]string{
		"name":"Shaijo",
		"email":"abc@gmail.com",
	}
	fmt.Println(user["name"])
}