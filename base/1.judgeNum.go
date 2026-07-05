package main
import (
	"fmt"
)
func main() {
	var num int
	fmt.Scanln(&num)
	if (num % 2 != 0) {
		fmt.Println("奇数")
	} else {
		fmt.Println("偶数")
	}
}