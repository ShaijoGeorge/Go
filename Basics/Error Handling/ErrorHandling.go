package main
import(
	"fmt"
	"errors"
)

func divide(a,b float64)(float64, error){
	if b==0{
		return 0, errors.New("Division by zero")
	}
	return a/b, nil
}
func main() {
	result, err:=divide(10,0)
	if err != nil{
		fmt.Println("Error: ", err)
	} else{
		fmt.Println("Result:", result)
	}
}