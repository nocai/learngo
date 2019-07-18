package effective

import (
	"fmt"
	"io"
	"os"
	"testing"
)

// Contents returns the file's contents as a string
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var result []byte
	buf := make([]byte, 100)

	for {
		n, err := f.Read(buf)
		result = append(result, buf[0:n]...)

		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}
	return string(result), nil
}

func Test(t *testing.T) {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func trace(s string) string {
	fmt.Println("enter:", s)
	return s
}

func untrace(s string) {
	fmt.Println("leaving:", s)
}

func TestA(t *testing.T) {
	defer untrace("a")
	trace("a")
}

func un(s string) {
	fmt.Println("leaving:",s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func TestB(t *testing.T) {
	b()
}
