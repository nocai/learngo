package ref

import (
	"testing"
	"fmt"
	"reflect"
	"io"
	"os"
	"github.com/astaxie/beego/logs"
)

type MyInt int

func TestRef(t *testing.T) {
	var i int
	var j MyInt
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(j))
	fmt.Println(reflect.ValueOf(j))
}

func TestRef2(t *testing.T) {
	var r io.Reader
	tty, err := os.OpenFile("D://t.txt", os.O_RDWR, 0)
	if err != nil {
		logs.Error(err)
		return
	}
	r = tty

	bs := make([]byte, 1024)
	for {
		n, err := r.Read(bs)
		if err == io.EOF {
			logs.Error(err)
			break
		}
		//if n == 0 {
		//	logs.Error(err)
		//	break
		//}
		fmt.Println(string(bs[:n]))
	}

	fmt.Println(reflect.TypeOf(r))
	fmt.Println(reflect.ValueOf(r))

	var w io.Writer
	w = r.(io.Writer)

	fmt.Println(reflect.TypeOf(w))
	fmt.Println(reflect.ValueOf(w))

}

func TestRef3(t *testing.T) {
	var x float64 = 3.4
	fmt.Println("value:", reflect.ValueOf(x))
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("type:", reflect.ValueOf(x).Type())
	fmt.Println("kind:", reflect.TypeOf(x).Kind())
	fmt.Println("kind:", reflect.ValueOf(x).Kind())
	fmt.Println(fmt.Sprintf("Interface:%T", reflect.ValueOf(x).Interface()))
}

func TestRef4(t *testing.T) {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println(v.CanSet())
	v.SetFloat(7.1)
}

func TestRef5(t *testing.T) {
	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
	v := p.Elem()
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)

}

type T struct {
	A int
	B string
}

func TestRef6(tt *testing.T) {
	var t = T{A:23, B:"skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now:", t)
}