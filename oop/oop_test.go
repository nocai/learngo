package oop

import (
	"image/color"
	"testing"
	"fmt"
	"unicode"
	"strings"
	"io"
)

type ColoredPoint struct {
	color.Color // 匿名字段（嵌入）
	x, y int    // 具名字段（聚合）
}

func TestOop1(t *testing.T) {
	cp := new(ColoredPoint)
	fmt.Println(cp.Color)
}

type Integer int

func (i Integer) Double() Integer {
	return i * 2
}
func TestOop2(t *testing.T) {
	var i Integer = 2
	fmt.Println(i.Double())
	fmt.Println(i + 1)
}

type Count int
type StringMap map[string]string
type FloatChan chan float64

func TestOop3(t *testing.T) {
	var i Count = 7
	i++
	fmt.Println(i)
	sm := make(StringMap)
	sm["key1"] = "value1"
	sm["key2"] = "value2"
	fmt.Println(sm)

	fc := make(FloatChan, 1)
	fc <- 2.29558714939
	fmt.Println(<-fc)
}

type RuneForRuneFunc func(rune) rune

func TestOop4(t *testing.T) {
	var removePunctuation RuneForRuneFunc
	phrases := []string{"Day; dusk, and night", "All day long"}
	removePunctuation = func(char rune) rune {
		if unicode.Is(unicode.Terminal_Punctuation, char) {
			return -1
		}
		return char
	}
	processPhrases(phrases, removePunctuation)
}

func processPhrases(phrases []string, function RuneForRuneFunc) {
	for _, phrase := range phrases {
		fmt.Println(strings.Map(function, phrase))
	}
}

func (count *Count) Increment() {
	*count++
}

func (count *Count) Decrement() {
	*count--
}

func (count Count) IsZero() bool {
	return count == 0
}

func TestOop5(t *testing.T) {
	var count Count
	i := int(count)
	count.Increment()
	j := int(count)
	count.Decrement()
	k := int(count)
	fmt.Println(count, i, j, k, count.IsZero())
}

type Part struct {
	Id   int
	Name string
}

func (part *Part) LowerCase() {
	part.Name = strings.ToLower(part.Name)
}
func (part *Part) UpperCase() {
	part.Name = strings.ToUpper(part.Name)
}
func (part Part) String() string {
	return fmt.Sprintf("<<%d %q>>", part.Id, part.Name)
}
func (part Part) HasPrefix(prefix string) bool {
	return strings.HasPrefix(part.Name, prefix)
}

func TestOop6(t *testing.T) {
	//part := Part{5, "Wrench"}
	part := new(Part)
	part.Id = 5
	part.Name = "Wrench"
	part.UpperCase()
	part.Id += 11
	fmt.Println(part, part.HasPrefix("W"))
}

type Item struct {
	id       string
	price    float64
	quantity int
}

func (item *Item) Cost() float64 {
	return item.price * float64(item.quantity)
}

type SpecialItem struct {
	Item
	catalogId int
}

func TestOop7(t *testing.T) {
	special := SpecialItem{Item{"Green", 3, 5}, 2007}
	fmt.Println(special.id, special.price, special.quantity, special.catalogId)
	fmt.Println(special.Cost())
	fmt.Println(special.Item.Cost())
}

type LuxuryItem struct {
	Item
	markup float64
}

// method1
//func (item *LuxuryItem) Cost() float64 {
//	return item.Item.price * float64(item.Item.quantity) * item.markup
//}

// method2
//func (item *LuxuryItem) Cost() float64 {
//	return item.price * float64(item.quantity) * item.markup
//}

func (item *LuxuryItem) Cost() float64 {
	return item.Item.Cost() * item.markup
}

func TestOop8(t *testing.T) {
	part := Part{5, "Wrench"}
	asStringV := Part.String
	sv := asStringV(part)
	hasPrefix := Part.HasPrefix
	asStringP := (*Part).String
	sp := asStringP(&part)
	lower := (*Part).LowerCase
	lower(&part)
	fmt.Println(sv, sp, hasPrefix(part, "w"), part)

}

type Object interface {

}

func TestOop9(t *testing.T) {
	var o Object = []string{"A", "B"}
	fmt.Println(o)
}

type Exchanger interface {
	Exchange()
}

type StringPair struct {
	first, second string
}

func (pair *StringPair) Exchange() {
	pair.first, pair.second = pair.second, pair.first
}

func (pair StringPair) String() string {
	return fmt.Sprintf("%q + %q", pair.first, pair.second)
}

type Point [2]int

func (point *Point) Exchange() {
	point[0], point[1] = point[1], point[0]
}

func TestOop10(t *testing.T) {
	jekyll := StringPair{"Henry", "Jekyll"}
	hyde := StringPair{"Edward", "Hyde"}
	point := Point{5, -3}
	fmt.Println("Before:", jekyll, hyde, point)

	jekyll.Exchange()
	hyde.Exchange()
	point.Exchange()
	fmt.Println("After #1:", jekyll, hyde, point)

	exchangeThses(&jekyll, &hyde, &point)
	fmt.Println("After #2:", jekyll, hyde, point)
}

func exchangeThses(exchangers ...Exchanger) {
	for _, exchanger := range exchangers {
		exchanger.Exchange()
	}
}

func (pair *StringPair) Read(data []byte) (n int, err error) {
	if pair.first == "" && pair.second == "" {
		return 0, io.EOF
	}
	if pair.first != "" {
		n = copy(data, pair.first)
		pair.first = pair.first[n:]
	}
	if n < len(data) && pair.second != "" {
		m := copy(data[n:], pair.second)
		pair.second = pair.second[m:]
		n += m
	}
	return n, nil
}

func TestOop11(t *testing.T) {
	const size = 16
	robert := &StringPair{"Robert L.", "Stevenson"}
	david := StringPair{"David", "Balfour"}
	for _, reader := range []io.Reader{robert, &david} {
		raw, err := ToBytes(reader, size)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%q\n", raw)
	}
}

func ToBytes(reader io.Reader, size int) ([]byte, error) {
	data := make([]byte, size)
	n, err := reader.Read(data)
	if err != nil {
		return data, err
	}
	return data[:n], nil
	//_ = n
	//return data, nil
}

type LowerCaser interface {
	LowerCase()
}
type UpperCaser interface {
	UpperCase()
}
type LowerUpperCaser interface {
	LowerCaser
	UpperCaser
}
type FixCaser interface {
	FixCase()
}
type ChangeCaser interface {
	LowerUpperCaser
	FixCaser
}

func fixcase(s string) string {
	var chars []rune
	upper := true
	for _, char := range s {
		if upper {
			char = unicode.ToUpper(char)
		} else {
			char = unicode.ToLower(char)
		}
		chars = append(chars, char)
		upper = unicode.IsSpace(char) || unicode.Is(unicode.Hyphen, char)
	}
	return string(chars)
}

func (part *Part) FixCase() {
	part.Name = fixcase(part.Name)
}

func (pair *StringPair) UpperCase() {
	pair.first = strings.ToUpper(pair.first)
	pair.second = strings.ToUpper(pair.second)
}

func (pair *StringPair) FixCase() {
	pair.first = fixcase(pair.first)
	pair.second = fixcase(pair.second)
}

func (pair *StringPair) LowerCase() {
	pair.first = strings.ToLower(pair.first)
	pair.second = strings.ToLower(pair.second)
}

func TestOop12(t *testing.T) {
	toaskRack := Part{8427, "TOAST RACK"}
	toaskRack.LowerCase()
	lobelia := StringPair{"LOBELIA", "SACKVILLE-BAGGINS"}
	lobelia.FixCase()
	fmt.Println(toaskRack, lobelia)
}

type IsValider interface {
	IsValid() bool
}

func TestOop13(t *testing.T) {
	points := [][2]int{{4, 6}, {}, {-7, 11}, {15, 17}, {14, -8}}
	for _, point := range points {
		fmt.Printf("(%d %d)", point[0], point[1])
	}
}
func TestOop14(t *testing.T) {
	points := []struct{ x, y int }{{4, 6}, {}, {-7, 11}, {15, 17}, {14, -8}}
	for _, point := range points {
		fmt.Printf("(%d %d)", point.x, point.y)
	}
}

type Person struct {
	Title     string
	Forenames []string
	Surname   string
}

type Author1 struct {
	Names    Person
	Title    []string
	YearBorn int
}

func TestOop15(t *testing.T) {
	author1 := Author1{Person{"Mr", []string{"Robert", "Louis", "Balfour"}, "Stevenson"}, []string{"Kidnapped", "Treasure Island"}, 1850}
	fmt.Printf("%#v\n", author1)
	author1.Names.Title = ""
	author1.Names.Forenames = []string{"Oscar", "Fingal", "Fingal", "O'Flahertie", "Wills"}
	author1.Names.Surname = "Wilde"
	author1.Title = []string{"The Picture of Dorian Gray"}
	author1.YearBorn += 4
	fmt.Printf("%#v", author1)
}

type Author2 struct {
	Person
	Title    []string
	YearBorn int
}

func TestOop16(t *testing.T) {
	author2 := Author2{Person{"Mr", []string{"Robert", "Louis", "Balfour"}, "Stevenson"}, []string{"Kidnappen", "Treasure Island"}, 1850}
	fmt.Println(author2)

	author2.Title = []string{"The Picture of Dorian Gray"}
	author2.Person.Title = ""
	author2.Forenames = []string{"Oscar", "Fingal", "O'Flahertie", "Wills"}
	author2.Surname = "Wilde"
	author2.YearBorn += 4
	fmt.Println(author2)
}

type Tasks struct {
	slice []string
	Count
}

func (tasks *Tasks) Add(task string) {
	tasks.slice = append(tasks.slice, task)
	tasks.Increment()
}

func (tasks *Tasks) Tally() int {
	return int(tasks.Count)
}

func TestOop17(t *testing.T) {
	tasks := Tasks{}
	fmt.Println(tasks.IsZero(), tasks.Tally(), tasks)
	tasks.Add("One")
	tasks.Add("Two")
	fmt.Println(tasks.IsZero(), tasks.Tally(), tasks)
}

type Optioner interface {
	Name() string
	IsValid() bool
}

type OptionCommon struct {
	ShortName string "short option name"
	LongName string "long option name"
}

type IntOption struct {
	OptionCommon
	Value, Min, Max int
}

func (option IntOption) Name() string {
	return name(option.ShortName, option.LongName)
}

func (option IntOption) IsValid() bool {
	return option.Min <= option.Value && option.Value <= option.Max
}
func name(shortName, longName string) string {
	if longName == ""{
		return shortName
	}
	return longName
}

type StringOption struct {
	OptionCommon
	Value string
}
func (so StringOption) Name() string {
	return name(so.ShortName, so.LongName)
}
func (so StringOption) IsValid() bool {
	if so.Value != "" {
		return true
	}
	return false
}
type FloatOption struct {
	Optioner
	Value float64
}
type GenericOption struct {
	OptionCommon
}

func (option GenericOption) Name() string {
	return name(option.ShortName, option.LongName)
}
func (option GenericOption) IsValid() bool {
	return true
}
func TestOop18(t *testing.T) {
	fmt.Println(new(OptionCommon))

	fileOption := StringOption{OptionCommon{"f", "file"}, "index.html"}
	topOption := IntOption{OptionCommon:OptionCommon{"t", "top"}, Max:100}
	sizeOption := FloatOption{GenericOption{OptionCommon{"s","size"}}, 19.5}
	for _, option := range []Optioner {topOption, fileOption, sizeOption} {
		fmt.Print("name = ", option.Name(), "*valid = ", option.IsValid())
		fmt.Print("*value = ")
		switch option := option.(type) {
		case IntOption:
			fmt.Print(option.Value, "*min = ", option.Min, " *max = ", option.Max, "\n")
		case StringOption:
			fmt.Println(option.Value)
		case FloatOption:
			fmt.Println(option.Value)

		}
	}

}