package fibonacci

import (
	"testing"
	"fmt"
)

func TestFinonacci(t *testing.T) {
	fmt.Println("Fibonacci(45) = ", Fibonacci(45).(int))
}
