// Goroutines: Lightweight threads.
// Channels: Communicate between goroutines.

package main
import (
	"fmt"
	"time"
)

func printNumbers(){
	for i:=1;i<=5;i++{
		fmt.Println(i)
		time.Sleep(500*time.Millisecond)
	}
}

func main(){
	go printNumbers() // Run in a goroutine
	go printNumbers() // Run another goroutine
	time.Sleep(3*time.Second) // Wait for goroutines to finish
}