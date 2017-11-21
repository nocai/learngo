package fibonacci

import (
	"fmt"
	"testing"
)

func TestFinonacci(t *testing.T) {
	fmt.Println("Fibonacci(45) = ", Fibonacci(45).(int))
}
