package t

import (
	"testing"
	"fmt"
	"strings"
	"strconv"
	"unicode/utf8"
	"unicode"
	"image/color"
	"reflect"
	"sort"
	"math/rand"
	"os"
	"log"
	"math"
)

func Test1(t *testing.T) {
	str := "ABC爱比次D"
	fmt.Println(str)
	fmt.Println(len(str))

	rs := []rune(str)
	fmt.Println(rs)
	fmt.Println(string(rs[3]))

}

func Test2(t *testing.T) {
	names := "刘军*刘军*刘军**刘军"
	fmt.Println(names)

	fmt.Print("|")
	for _, name := range strings.SplitAfter(names, "*") {
		fmt.Printf("%s|", name)
	}
	fmt.Println()
}

func Test3(t *testing.T) {
	for _, record := range []string{"刘军*刘军", "刘军|刘军", "刘军$刘军"} {
		rs := strings.FieldsFunc(record, func(char rune) bool {
			switch char {
			case '*', '|', '$':return true
			}
			return false
		})
		fmt.Println(rs)
	}
}

func Test4(t *testing.T) {
	str := "ABC"
	strs := strings.Fields(str)
	fmt.Println(strs)
	fmt.Println(len(strs))
	fmt.Println(strings.Contains(str, "AC"))
	fmt.Println(strings.Count(str, "A"))
	fmt.Println(strings.EqualFold(str, "abc"))

	fmt.Println(len("我是中国人"))
	fmt.Println(strings.IndexAny("我是中国人", "中")) // 在存在返回 6
	fmt.Println(strings.IndexAny("我是中国人", "和")) // 在存在返回 -1

}

func Test5(t *testing.T) {
	for _, truth := range []string{"1", "t", "TRUE", "FALSE", "0", "A"} {
		if b, err := strconv.ParseBool(truth); err != nil {
			fmt.Printf("\n{%v}", err)
		} else {
			fmt.Println(b, "")
		}
	}
	fmt.Println()
}

func Test6(t *testing.T) {
	x, err := strconv.ParseFloat("-99.7", 64)
	fmt.Printf("%8T %6v %v\n", x, x, err)
	y, err := strconv.ParseInt("71309", 10, 0)
	fmt.Printf("%8T %6v %v\n", y, y, err)
	z, err := strconv.Atoi("71309")
	fmt.Printf("%8T %6v %v\n", z, z, err)
}

func Test7(t *testing.T) {
	str := "我们都是中国人"
	char, size := utf8.DecodeLastRune([]byte(str))
	fmt.Println(string(char), size)

	char, size = utf8.DecodeLastRuneInString(str)
	fmt.Println(string(char), size)

	fmt.Println(utf8.RuneCountInString(str))
	fmt.Println(utf8.RuneLen(char))

	fmt.Println(utf8.RuneStart([]byte(str)[0]))
	fmt.Println(utf8.Valid([]byte("ABC")))
}

func Test8(t *testing.T) {
	fmt.Println(unicode.IsNumber('1'))
}

func Test9(t *testing.T) {
	grades := []int{87, 55, 43, 71, 60, 43, 32, 19, 63}
	inflate(grades, 3)
	fmt.Println(grades)
}

func inflate(numbers []int, factor int) {
	for i := range numbers {
		numbers[i] *= factor
	}
}

type rectangle struct {
	x0, y0, x1, y1 int
	fill           color.RGBA
}

func Test10(t *testing.T) {
	rect := rectangle{4, 8, 20, 10, color.RGBA{0xFF, 0, 0, 0xFF}}
	fmt.Println(rect)
	resizeRect(&rect, 5, 5)
	fmt.Println(rect)

}

func resizeRect(rect *rectangle, width, height int) {
	(*rect).x1 += width
	rect.y1 += height
}

func TestArray(t *testing.T) {
	var buffer [20]byte
	var grid1 [3][3]int
	grid1[1][0], grid1[1][1], grid1[1][2] = 8, 6, 2
	grid2 := [3][3]int{{4, 3}, {8, 6, 2}}
	cities := [...]string{"ShangHai", "Mumbai", "Istanbul", "Beijin"}
	cities[len(cities) - 1] = "Karachi"
	fmt.Println("Type Len Contents")
	fmt.Printf("%-8T %2d %v\n", buffer, len(buffer), buffer)
	fmt.Printf("%-8T %2d %q\n", cities, len(cities), cities)
	fmt.Printf("%-8T %2d %v\n", grid1, len(grid1), grid1)
	fmt.Printf("%-8T %2d %v\n", grid2, len(grid2), grid2)

	var array [2] int
	fmt.Println(len(array))
	fmt.Println(cap(array))

	array[0] = 1
	array[1] = 2
	fmt.Println(array)
	fmt.Println(reflect.TypeOf(array))
	fmt.Println(reflect.TypeOf(array[:]))
}

func TestSlice(t *testing.T) {
	s := make([]int, 2)
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	ss := make([]int, 2, 4)
	fmt.Println(ss)
	fmt.Printf("%T\n", ss)
}

func TestSlice2(t *testing.T) {
	amounts := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sum := 0
	for _, amount := range amounts {
		sum += amount
	}
	fmt.Println(sum)

	sum2 := 0.0
	for i := range amounts {
		temp := float64(amounts[i]) * 1.05
		sum2 += temp
	}
	fmt.Printf("%.2f", sum2)
}

type Product struct {
	name  string
	price float64
}

func (pro Product) String() string {
	return fmt.Sprintf("%s (%.2f)", pro.name, pro.price)
}

func TestSlice3(t *testing.T) {
	pros := []*Product{
		{"A", 1.1},
		{"B", 2.2},
		{"C", 3.3},
	}
	fmt.Println(pros)

	for _, pro := range pros {
		pro.price += 0.5
	}
	fmt.Println(pros)
}

func TestSlice4(t *testing.T) {
	s := []string{"A", "B", "C", "D", "E", "F", "G"}
	a := []string{"K", "L", "M", "N"}
	u := []string{"m", "n", "o", "p", "q", "r"}

	s = append(s, "h", "i", "j")
	s = append(s, a...)
	s = append(s, u[2:5]...)

	letters := "WXY我"
	b := []byte{'U', 'V'}
	b = append(b, letters...)
	fmt.Printf("%v\n%s\n", s, b)
}

func TestSort(t *testing.T) {
	files := []string{"Test.conf", "util.go", "Makefile", "misc.go", "main.go"}
	fmt.Printf("Unsorted:            %q\n", files)
	sort.Strings(files)
	fmt.Printf("Underlying bytes:    %q\n", files)
	SortFoldeStrings(files)
	fmt.Printf("Case insensitive:    %q\n", files)

	target := "Makefile"
	for i, file := range files {
		if file == target {
			fmt.Printf("found \"%s\" at files[%d]\n", target, i)
			break
		}
	}

	i := sort.Search(len(files), func(i int) bool {
		return strings.ToLower(files[i]) >= strings.ToLower(target)
	})
	fmt.Println(files[i])
}

func SortFoldeStrings(s []string) {
	sort.Sort(FoldeStrings(s))
}

type FoldeStrings []string

func (slice FoldeStrings) Len() int {
	return len(slice)
}

func (slice FoldeStrings) Less(i, j int) bool {
	return strings.ToLower(slice[i]) < strings.ToLower(slice[j])
}

func (s FoldeStrings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func TestMap1(t *testing.T) {
	massForPlanet := make(map[string]float64)
	massForPlanet["Mercury"] = 0.06
	massForPlanet["Venus"] = 0.82
	massForPlanet["Earth"] = 1.00
	massForPlanet["Mars"] = 0.11
	fmt.Println(massForPlanet)
}

type Point struct {
	x, y, z int
}

func (point Point) String() string {
	return fmt.Sprintf("(%d, %d, %d)", point.x, point.y, point.z)
}

func TestMap2(t *testing.T) {
	triangle := make(map[*Point]string, 3)
	triangle[&Point{89, 47, 27}] = "a"
	triangle[&Point{86, 65, 86}] = "b"
	triangle[&Point{7, 44, 45}] = "c"
	fmt.Println(triangle)

	var pa = Point{1, 2, 3}
	var pb = Point{1, 2, 3}
	fmt.Println(pa == pb)

	nameForPoint := make(map[Point]string)
	nameForPoint[pa] = "a"
	nameForPoint[pb] = "b"
	fmt.Println(nameForPoint)

	populationForCity := map[string]int{"A":1, "B":2, "C":3}
	for key, value := range populationForCity {
		fmt.Printf("%-10s %8d\n", key, value)
	}
	fmt.Println(populationForCity["D"])

	key := "B"
	if population, found := populationForCity[key]; found {
		fmt.Println(population)
	} else {
		fmt.Println("not found!")
	}
	fmt.Println(populationForCity)
	delete(populationForCity, "B")
	fmt.Println(populationForCity)
}

func TestOrder(t *testing.T) {
	a, b, c := 2, 3, 5
	for a := 7; a < 8; a ++ {
		b := 11
		c = 13
		fmt.Printf("Inner: a->%d b->%d c->%d\n", a, b, c)
	}
	fmt.Printf("outer: a->%d b->%d c->%d\n", a, b, c)
}
//
//func shadow() (err error) {
//	//x, err := check1()
//	//if err != nil {
//	//	return
//	//}
//	//if y, err := check2(x); err != nil {
//	//	return
//	//} else {
//	//	fmt.Println(y)
//	//}
//	//return
//}

type StringSlice []string

func (ss StringSlice) String() string {
	result := "StringSlice {"
	for _, str := range ss {
		result += "\"" + str + "\"" + ", "
	}
	result = result[:len(result) - 1]
	return result + "}"
}
func TestOrder2(t *testing.T) {
	fancy := StringSlice{"Lithiun", "Soodium", "Potammsium", "Rubidium"}
	fmt.Println(fancy)
	plain := []string(fancy)
	fmt.Println(plain)
}

func TestOrder3(t *testing.T) {
	var i interface{} = 99
	var s interface{} = []string{"Left", "Right"}
	fmt.Println(s)

	j := i.(int)
	fmt.Printf("%T -> %d\n", j, j)

	if i, ok := i.(int); ok {
		fmt.Printf("%T -> %d\n", i, j)
	}
	if s, ok := s.([]string); ok {
		fmt.Printf("%T ->%q\n", s, s)
	}
}

func createCounter(start int) (chan int) {
	next := make(chan int)
	go func(i int) {
		for {
			next <- i
			i ++
		}
	}(start)
	return next
}

func TestOrder4(t *testing.T) {
	counterA := createCounter(2)
	counterB := createCounter(102)
	for i := 0; i < 5; i ++ {
		a := <-counterA
		fmt.Printf("(A->%d, B->%d)", a, <-counterB)
	}
	fmt.Println()
}

func TestOrder5(t *testing.T) {
	channels := make([]chan bool, 6)
	for i := range channels {
		channels[i] = make(chan bool)
	}

	go func() {
		for {
			channels[rand.Intn(6)] <- true
		}
	}()

	counter := make(map[int]int)

	for i := 0; i < 36; i++ {
		var x int
		select {
		case <-channels[0]:
			x = 1
		case <-channels[1]:
			x = 2
		case <-channels[2]:
			x = 3
		case <-channels[3]:
			x = 4
		case <-channels[4]:
			x = 5
		case <-channels[5]:
			x = 6
		}
		fmt.Printf("%d", x)
		count(x, counter)

	}
	fmt.Println("\n", counter)
}

func count(i int, m map[int]int) map[int]int {
	if v, ok := m[i]; ok {
		m[i] = v + 1
	} else {
		m[i] = 1
	}
	return m
}

func TestOrder6(t *testing.T) {
	var file *os.File
	var err error
	if file, err = os.Open("D:\\DTLFolder\\DriversBackup\\backuplist.db"); err != nil {
		log.Println("failed to open the file :", err)
		return
	}
	defer file.Close()
}

func Fibonacci(n int) int {
	if n < 2 {
		return n;
	}
	return Fibonacci(n - 1) + Fibonacci(n - 2)
}

func TestOrder7(t *testing.T) {
	for n := 0; n < 20; n ++ {
		fmt.Print(Fibonacci(n), " ")
	}
	fmt.Println()
}

func HofstadterFemale(n int) int {
	if n <= 0 {
		return 1
	}
	return n - HofstadterMale(HofstadterFemale(n - 1))
}

func HofstadterMale(n int) int {
	if n <= 0 {
		return 0
	}
	return n - HofstadterFemale(HofstadterMale(n - 1))
}

func TestOrder8(t *testing.T) {
	females := make([]int, 20)
	males := make([]int, len(females))
	for n := range females {
		females[n] = HofstadterFemale(n)
		males[n] = HofstadterMale(n)
	}
	fmt.Println("F", females)
	fmt.Println("M", males)
}

func IsPalindrome(word string) bool {
	if utf8.RuneCountInString(word) <= 1 {
		return true
	}
	first, sizeOfFirst := utf8.DecodeRuneInString(word)
	last, sizeOfLast := utf8.DecodeLastRuneInString(word)
	if first != last {
		return false
	}
	return IsPalindrome(word[sizeOfFirst:len(word) - sizeOfLast])
}

func TestOrder9(t *testing.T) {
	str := "AbbAc"
	fmt.Println(IsPalindrome(str))
}
//
//var FunctionForSuffix = map[string]func(string)([]string, error) {
//	".gz":GzipFileList,
//	".tar":TarFileList,
//	".tar":TarFileList,
//	".tgz":TarFileList,
//	".zip":ZipFileList
//}
//
//func ArchiveFileListMap(file string) ([]string, error) {
//	if function, ok := FunctionForSuffix[Suffix(file)]; ok {
//		return function(file)
//	}
//	return nil, errors.New("unrecognized archive")
//}

func Minimum(first interface{}, rest ...interface{}) interface{} {
	minimum := first
	for _, x := range rest {
		switch x := x.(type) {
		case int :
			if x < minimum.(int) {
				minimum = x
			}
		case float64:
			if x < minimum.(float64) {
				minimum = x
			}
		case string :
			if x < minimum.(string) {
				minimum = x
			}
		}
	}
	return minimum
}

func TestOrder10(t *testing.T) {
	i := Minimum(4, 3, 8, 2, 9).(int)
	fmt.Printf("%T %v\n", i, i)
}

func Index(xs interface{}, x interface{}) int {
	switch slice := xs.(type) {
	case []int :
		for i, y := range slice {
			if y == x.(int) {
				return i
			}
		}
	case []string :
		for i, y := range slice {
			if y == x.(string) {
				return i
			}
		}
	}
	return -1
}

func IndexReflectX(xs interface{}, x interface{}) int {
	if slice := reflect.ValueOf(xs); slice.Kind() == reflect.Slice {
		for i := 0; i < slice.Len(); i++ {
			switch y := slice.Index(i).Interface().(type) {
			case int :
				if y == x.(int) {
					return i
				}
			case string :
				if y == x.(string) {
					return i
				}
			}
		}
	}
	return -1
}

func IndexReflect(xs interface{}, x interface{}) int {
	if slice := reflect.ValueOf(xs); slice.Kind() == reflect.Slice {
		for i := 0; i < slice.Len(); i ++ {
			if reflect.DeepEqual(x, slice.Index(i)) {
				return i
			}
		}
	}
	return -1
}

func TestOrder11(t *testing.T) {
	xs := []int{2, 4, 6, 8}
	fmt.Println("5 @", Index(xs, 5), " 6 @", Index(xs, 6))
	fmt.Println("5 @", IndexReflectX(xs, 5), " 6 @", IndexReflectX(xs, 6))

	ys := []string{"C", "B", "K", "A"}
	fmt.Println("Z @", Index(ys, "Z"), " A @", Index(ys, "A"))
	fmt.Println("Z @", IndexReflect(ys, "Z"), " A @", IndexReflect(ys, "A"))
}

func IntSliceIndex(xs []int, x int) int {
	for i, y := range xs {
		if x == y {
			return i
		}
	}
	return -1
}

type Slicer interface {
	EqualTo(i int, x interface{}) bool
	Len() int
}

type IntSlice []int

func (slice IntSlice) EqualTo(i int, x interface{}) bool {
	return slice[i] == x.(int)
}
func (slice IntSlice) Len() int {
	return len(slice)
}

func IntIndexSlicer(ints []int, x int) int {
	return IndexSlicer(IntSlice(ints), x)
}
func IndexSlicer(slice Slicer, x interface{}) int {
	for i := 0; i < slice.Len(); i ++ {
		if slice.EqualTo(i, x) {
			return i
		}
	}
	return -1
}

type StringSlice2 []string

func (slice StringSlice2) EqualTo(i int, x interface{}) bool {
	return slice[i] == x.(string)
}
func (slice StringSlice2) Len() int {
	return len(slice)
}
func StringIndexSlicer(xs []string, x string) int {
	return IndexSlicer(StringSlice2(xs), x)
}
func TestOrder12(t *testing.T) {
	xs := []int{2, 4, 6, 8}
	fmt.Println("5 @", IntIndexSlicer(xs, 5), " 6 @", IntIndexSlicer(xs, 6))

	ys := []string{"C", "B", "K", "A"}
	fmt.Println("Z @", StringIndexSlicer(ys, "Z"), " A @", StringIndexSlicer(ys, "A"))
}

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i ++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func TestOrder13(t *testing.T) {
	xs := []int{2, 4, 6, 8}
	ys := []string{"C", "B", "K", "A"}
	fmt.Println(SliceIndex(len(xs), func(i int) bool {
		return xs[i] == 5
	}),
		SliceIndex(len(xs), func(i int) bool {
			return xs[i] == 6
		}),
		SliceIndex(len(ys), func(i int) bool {
			return ys[i] == "Z"
		}),
		SliceIndex(len(ys), func(i int) bool {
			return ys[i] == "A"
		}))
}

func TestOrder14(t *testing.T) {
	i := SliceIndex(math.MaxInt32, func(i int) bool {
		return i > 0 && i % 27 == 0 && i % 51 == 0
	})
	fmt.Println(i)
}

func IntFilter(slice []int, predicate func(int) bool) []int {
	filtered := make([]int, 0, len(slice))
	for i := 0; i < len(slice); i ++ {
		if predicate(slice[i]) {
			filtered = append(filtered, slice[i])
		}
	}
	return filtered
}

func TestOrder15(t *testing.T) {
	readings := []int{4, -3, 2, -7, 8, 19, -11, 7, 18, -6}
	even := IntFilter(readings, func(i int) bool {
		return i % 2 == 0
	})
	fmt.Println(even)
}

func Filter(limit int, predicate func(int) bool, appender func(int)) {
	for i := 0; i < limit; i ++ {
		if predicate(i) {
			appender(i)
		}
	}
}

func TestOrder16(t *testing.T) {
	readings := []int{4, -3, 2, -7, 8, 19, -11, 7, 18, -6}
	even := make([]int, 0, len(readings))
	Filter(len(readings), func(i int) bool {
		return readings[i] % 2 == 0
	}, func(i int) {
		even = append(even, readings[i])
	})
	fmt.Println(even)
}

func TestOrder17(t *testing.T) {
	parts := []string{"X15", "X23", "A41", "L19", "X57", "A63"}
	var Xparts []string
	Filter(len(parts), func(i int) bool {
		return parts[i][0] == 'X'
	}, func(i int) {
		Xparts = append(Xparts, parts[i])
	})
	fmt.Println(Xparts)
}

func TestOrder18(t *testing.T) {
	var product int64 = 1
	Filter(26, func(i int) bool {
		return i % 2 != 0
	}, func(i int) {
		product *= int64(i)
	})
	fmt.Println(product)
}

func TestAdmin(t *testing.T) {
	fmt.Println(new(Admin))

}

type Admin struct {
	Id        int64 `json:"id"`
	AdminName string `json:"admin_name"`
	UserId    int64 `orm:"unique" json:"user_id"`

	Invalid   interface{} `json:"invalid"`
}