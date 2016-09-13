# Day #2


## Goals

* [Scoping](2.1_scoping.md)
* Visibility Export / Unexport
* Collection Type Array/Slice/Map
* Controls if/for/switch/defer/panic/recover
* User-defined Type
* Method
* Interface Type
* Enum
* Error Type
* Interface
* Casting
* Type Assertion
* JSON
* Web Service

## Scoping (https://play.golang.org)

Scoping is straight forward in GO.

```go
var a = 10

func main() {
	fmt.Println("Global A", a)

	a := 20
	fmt.Println("Main A", a)
	if true {
		a := 30
		fmt.Println("If A", a)
	}
	fmt.Println("Main A", a)

    blee()
    fred(1000)
    foo(5)
}

func blee() {
    fmt.Println("Blee A", a)
}

func fred(a int) {
    fmt.Println("Fred A", a)
}

func foo(b int) (a int) {
    fmt.Println("Foo A", a)
    return a
}
```

## Exported vs Unxported Types

Until now we've been mainly staying in main and the playground.
Hence visibility did not matter much. 
Let's take a closer look by writing a library...
 
(Live Demo)[https://github.com/derailed/vmw_vis]

## Errors 

Idiom Errors are just values! Don't ignore them!
How do you handle errors?

Idiom: Just don't check errors, handle them gracefully!

Caveats:

Don't compare errors on string
May not be practical to expose error type as a lib vendor. 
Interface specification with an custom err type, will force every implementor to import that package (Dependency!/Coupling)

### Basic declaration

```go
import (
	"errors"
	"fmt"
)

func main() {
	// It works!!
	fred(1)
	fred(10)
}

func fred(i int) (int, error) {
	if i == 10 {
		return 0, fmt.Errorf("Found a really bad value %d", i)
	}

	return i, nil
}
```

### Let's check!

```go
import (
	"errors"
	"fmt"
)

func main() {
	i, err := fred(1)
	if err != nil {
		fmt.Printf("Fred capped out %#v\n", err)
	} else {
		fmt.Println("Fred Returned", i)
	}
}

func fred(i int) (int, error) {
	if i == 10 {
		return 0, fmt.Errorf("Found a really bad value %d", i)
	}

	return i, nil
}
```

### Sentinel Error

// Sentinel errors ex. io.EOF
// NOTE: Part of your public API!!
var (
	errCrapola = errors.New("Crapola!")
	errBoom   = errors.New("Boom!")
	errBadAss  = errors.New("Badass!")
)

func main() {
	err := blee(10)
	if err == errCrapola {
		fmt.Printf("Crapola happened -- %#v", err)
	} else if strings.Contains(err.Error(), "badAss") { // Code Smell! err.Error() is for humans!
		fmt.Println("Bad Ass happened! Get into the chopper!")
	} else if err == errBoom {
		fmt.Printf("Boom happened %#v", err)
	} else {
		fmt.Println("Some other bad stuff")
	}
}

func blee(i int) error {
	if i == 10 {
		return errCrapola
		// return errBadAss
	} else if i == 20 {
		return errBoom
	}

	return fmt.Errorf("Bad stuff here too!")
}
```

NOTE: Problem is want to provide more context with another error. Breaks the code ;-(

### Error Types

```go
// Pb: Must be made public + code coupling!
type userError struct {
	user, reason string
}

func (e userError) Error() string {
	return fmt.Sprintf("UserError: [%s] Fail! %s", e.user, e.reason)
}

func main() {
	err := blee()
	if err != nil {
		log.Println("Main", err)
	}
}

func blee() error {
	err := fred(20)
	switch err.(type) {
	case nil:
		// Success!
	case userError:
		log.Println(err)
		return nil
	default:
		fmt.Println(err)
	}
	return err
}

func fred(i int) error {
	if i == 10 {
		return userError{user: "Fernand", reason: "You suck!"}
	} else if i == 20 {
		return fmt.Errorf("BumbleBee Tuna!")
	}
	return nil
}
```

## Error behavior

Assert error on behavior!

```go
type Retryer interface {
	Retry() bool
}

type RetryError struct {
	reason string
}

func (r RetryError) Error() string { return fmt.Sprintf(r.reason) }
func (r RetryError) Retry() bool   { return true }

func isRetryable(err error) bool {
	e, ok := err.(Retryer)
	return ok && e.Retry()
}

func main() {
	for i, j := 1, 0; isRetryable(fred(i)); j++ {
		if j == 5 {
			fmt.Println("Giving up! Done retrying...")
			break
		}
		i++
	}

	fmt.Println("All done here. Good bye!")
}

func fred(i int) error {
	fmt.Println("Calling Fred!", i)
	if i < 5 {
		return RetryError{reason: "Play it again, Sam!"}
	}
	fmt.Println("Fred Succeeded!")
	return nil
}
```

### Recommendation (Dave Cheney)[github.com/pkg/errors]

(GopherCon 2016)[https://www.youtube.com/watch?v=lsBF58Q-DnY]

## Collection Types

### Array

Zero based indexing. Fast as static continuous (same size) blocks of mem are preallocated.

```go
var a [5]int
var b [10]bool
c := [5]float64{}
d := [5]float64{100, 200, 300, 400, 500}
e := [5]float64{0: 100, 4: 1000}
f := [...]float64{0: 100, 4: 1000}

fmt.Println(a, b, c, d, e, f)
fmt.Println(len(e))
fmt.Println(cap(e))

fmt.Println(f[0])
// fmt.Println(f[10])

f[1] = 20
fmt.Println(f)

for i := 0; i < len(f); i++ {
    fmt.Println(f[i])
}

g := [5]*int{}
fmt.Println(g)

h := [5]*int{0: new(int), 1: new(int)}
fmt.Println(h)

*h[0] = 10
*h[1] = 20
fmt.Println(h)

for i := 0; i < len(h); i++ {
    if h[i] != nil {
        fmt.Println(*h[i])
    } else {
        fmt.Println("Nil")
    }
} 
```

#### Assignments

```go
c := [5]float64{}
d := [5]float64{100, 200, 300, 400, 500}

fmt.Println(c)
fmt.Println(d)
c = d
fmt.Println(c)
fmt.Println(d)

c[1] = 1
d[1] = 10
fmt.Println(c)
fmt.Println(d)
```

#### Pointer Array Assignment

```go
f := [5]*int{}
g := [5]*int{0: new(int)}

f = g
fmt.Println(f)
fmt.Println(g)

*f[0] = 10
*g[0] = 20
fmt.Println(*f[0])
fmt.Println(*g[0])
```

#### Warning!! Arrays are passed by value

```go
a := [2e6]int{}
// Doh! 64-bit int = 8byte -> Alloc 2m*8 = 16Mb on heap
fred(a)
// Problem with passing a ptr is array can be mutated

// BAD!
func fred(a [2e6]int) {
}
// Better!
func fred(a *[2e6]int) {
}
```

NOTE: Array are not commonly used because of usability issues. Hence slice are preferred.

#### Mutli Dimensional arrays

```go
    a := [3][3]int{{1,1,1}, {2,2,2}, {3,3,3}}
    a[1][2] = 10
    fmt.Printf("%#v\n", a)
```

## Your turn...

### Lab2.1 ![alt text](https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon24.png "Lab2.1") 
---

### Mission
> Write a program (with tests!) to compute 3x3 matrix addition and multiplication.
> Extra credit for computing a nice output formatter ;-)

### Expectations

1 1 1     3 3 3    4 4 4
2 2 2  +  4 4 4 =  6 6 6
3 3 3     5 5 5    8 8 8

1 1 1     3 3 3     3  3  3
2 2 2  *  4 4 4 =   8  8  8
3 3 3     5 5 5    15 15 15
---

### Slice

Built on array, same operations. But, can grow dynamically and resolves passing as args.

```go
slice := make([]int, 3, 5)
slice[0], slice[1], slice[2] = 10, 20, 30
info(slice)

s := []string{"Hello", "World"} // len=2, cap=2
info(s)
s = append(s, "BumbleBee", "Tuna")
info(s)

// BOOM!
slice[3] = 40

slice = append(slice, 40)
info(slice)

slice = append(slice, 50)
info(slice)

slice = append(slice, 60)
info(slice)

slice1 := slice[1:3]
info(slice1)

slice1[1] = 100
info(slice1)
info(slice)

slice2 := slice[3:]
info(slice)
info(slice2)

slice3 := slice[:3]
info(slice3)

slice4 := slice[:3:3]
info(slice4)

slice4 = append(slice4, 200)
info(slice4)
info(slice)

func info(s []int) {
    fmt.Printf("%v -- %d|%d\n", s, len(s), cap(s))
}
``` 

#### Working with Slices

```go
// Iteration
s := []int{0:1, 3:3} // len 4 | cap 4

// This is ok
// NOTE: More control then range since always iterate from start of slice
for i := 0; i < len(s); i++ {
    fmt.Println(s[i])
}

// Better
// !!!! NOTE: Range makes a copy of each element in the slice!
for i, v := range(s) {
    fmt.Printf("%X\n", v)
    fmt.Println(i, v)
}

// Or
for _, v := range s {
    fmt.Println(v)
}

// Or
for i, _ := range s {
    fmt.Println(i)
}

// Slice as arguments
s := make([]int, 2e6)
proccess(s)

// Always passing a constant 24 bytes value 8 addr + 8 len + 8 cap, no matter what the size is!
func process(s []int) {
}
```

#### TODO LAB ON SLICE!!

### Map

UNORDERED Collection of key/value pairs. Use hash function to determine bucket. 
Key values can be anything that is comparable (==). Thus func, slice, channel are not a key option

```go
m := make(map[string]int)
fmt.Printf("%#v\n", m)
m1 := map[string]int{}
fmt.Printf("%#v\n", m1)
m2 := map[string]int{"fred": 1, "blee": 2, "duh": 3}
fmt.Printf("%#v\n", m2)

m2["yo"] = 10
fmt.Printf("%#v\n", m2)

fmt.Println(m2["yo"])

// Bad!
fmt.Println(m2["boom"])

// Better
v, ok := m2["boom"]
if ok {
    fmt.Println(v, ok)
}

// Best
if v, ok := m2["boom"]; ok {
    m2["boom"] = v + 1
}

if v, ok := m2["yo"]; ok {
    m2["yo"] = v + 1
}
fmt.Println(m2)

for k, v := range m2 {
    fmt.Println(k, v)
}

delete(m2, "yo")
fmt.Println(m2)
```

## Controls

### If Statement

```go
  	a := 1

	if a > 1 {
		fmt.Println("a>1", a)
	} else if a == 1 {
		fmt.Println("a==1", a)
	} else {
		fmt.Println("Anything else", a)
	}
	
	m := map[int]int{1: 10, 2: 20}
	
	if v, ok := m[3]; ok {
	   fmt.Println("Value of 3", v)
	} else { // NOTE: Scope!
	   fmt.Println("No value set", v, ok)
	}
```

### For Statement

```go
	s := []int{1, 2, 3, 4}

    // Traditional
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

    // While
	var bail bool
	for bail {
		fmt.Println("YO!!")
		time.After(1 * time.Second)
		bail = true
	}

    // While
	for {
		fmt.Println("HERE!")
		time.Sleep(1 * time.Second)
		break
	}

    // Omit increment	
	for bail := false; bail; {
		fmt.Println("YO!!")
		time.After(1 * time.Second)
		bail = true
	}

    // Multi inits
	for i, j := 0, 10; i < 10; i, j = i+1, j+10 {
		fmt.Println(i, j)
	}
```

### Case Statement

```go
	i := 15

	switch i {
	case 10:
		fmt.Println("It is 10")
	case 20:
		fmt.Println("It is 10")
	default:
		fmt.Println("It is, what it is", i)
	}

	switch {
	case i == 10:
		fmt.Println("It is 10")
	case i == 20:
		fmt.Println("It is 20")
	}
```

### Goto label / Continue label

Just FYI! It's there if you absolutely need it! If continue from a nested loop.

### Defer

Used for cleaning up at the end of a function or block. Useful and powerful idiom!

NOTE: Defer functions are stacks LSFO!!

```go
func main() {
    // Main use case
	defer fmt.Println("Yo Mama!")

    // LSFO
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Fred", fred())
}

// Can be an anonymous function. Here modifies return named parameters
func fred() (i int) {
	defer func() { i++ }()
	return 1
}
```

```go
func main() {
	i := 1
 
    // Note: binding of variables!
	defer func(j int) { fmt.Println("Defer", j, i + 1) }(i)

        i = 10
	fmt.Println(i)
}
```

#### Defer Poor man timer

Takes advantage of binding to time a function!!

```go
func count(min, max int) {
	defer func(t time.Time) {
		fmt.Printf("Elapsed %v\n", time.Since(t))
	}(time.Now())

	for i := min; i < max; i++ {
		fmt.Printf("% d", i)
		time.Sleep(1 * time.Microsecond)
	}
}
```

### Panic

Panic can triggered manually. Runtime errors will also trigger a panie ex. array out of bound. 
All functions in the stack will be unwinded when panic occurs, thus defer functions will be called!!

NOTE: Don't panic!
NOTE: Checkout json package for best examples!

```go
func main() {
	defer fmt.Println("YO!")
	panic("Something bad happened!")
}
```

### Recover

```go
func main() {
	caller()
	fmt.Println("Normal exit!")
}

func caller() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering!")
		}
	}()

	crap()
}

func crap() {
    defer fmt.Println("Crap is Done!" )
	panic("blee")
}
```

## Error Type

Erro
## User-defined Type (https://play.golang.org)

Go is a static language. Thus the compiler must know the kind/size of a given 
type to keep us from bad memory access and also can perform optimization for us. 
We've already seen some of the primitive types, and now we can built our own types.

```go
type car struct {
    year  int
    make, model string
    color string
}

var myCar car
fmt.Printf("%v\n", myCar)

// Idiomatic!! use var when init to zero values

domsCar := car{1975, "Chevrolet", "Impala", "Bronco Orange"}
fmt.Printf("%v\n", domsCar)

// Use short initializer
fredsCar := car{
    year:  1971, 
    make:  "Ford", 
    model: "Pinto", 
    color: "Puke Green",
}
fmt.Printf("%v\n", fredsCar)

// NOTE: Scoping
func paint(c car) {}

// Using user-defined types
type owner struct {
    name string
    owns car
}

fred := owner{
    name: "Fred",
    owns: car{
        year:  1971,
        make:  "Ford",
        model: "Pinto",
        color: "Puke Green",
    },
}
fmt.Printf("%#v\n", fred)

// Base Types
type policeCar car

doylesCar := policeCar{
    year:  1971,
    make:  "Ford",
    model: "Pinto",
    color: "Puke Green",
}
fmt.Printf("%#v\n", doylesCar)

// Compiler won't implicitly convert. You need to cast it!
johnsonCar := car{1975, "Chevrolet", "Impala", "Bronco Orange"}
fmt.Printf("%#v\n", policeCar(johnsonCar))
```

### Type Embedding

```go
type (
painter interface {
    paint(color string)
}

printer interface {
    print()
}

car struct {
    brand, color string
}

truck struct {
    car     // Embedded Type
    payload int
}
)

func (c *car) paint(color string) {
c.color = color
}

func (c car) print() {
fmt.Printf("%#v\n", c)
}

func (t truck) print() {
fmt.Printf("%#v\n", t)
}

func main() {
c := car{brand: "Ford", color: "Puke Green"}

paintIt(&c, "Cool Blue")
printIt(c)

var c1 = new(car)
paintIt(c1, "Blue")
printIt(c1)

c2 := car{}
//paintIt(c2, "Red")
printIt(c2)

t1 := truck{car: car{brand: "Chevy", color: "Red"}, payload: 10000}
paintIt(&t1, "Yellow")
printIt(t1)
}

func printIt(p printer) {
p.print()
}

func paintIt(p painter, color string) {
p.paint(color)
}
```

## Function (one more thing...)

### Var args
Variadic functions are used thru out GO packages. Means a given function can take 0 or more arguments of a given type.

```go
func variadic(args... int) {
  fmt.Println(args)
}

func main() {
	fmt.Println("Hello, playground")
	
	variadic(1)
	variadic(2, 3, 4)
	
	s := []int{1,2,3}
	variadic(s...)
}
```

## Methods

```go
type car struct {
	year        int
	make, model string
	color       string
}

func (c car) isNew() bool {
	return c.year == 2016
}

// Immutable
func (c car) paint(color string) {
    c.color = color
}

func main() {
	fredsCar := car{
		year:  1971,
		make:  "Ford",
		model: "Pinto",
		color: "Puke Green",
	}
	fmt.Println("New?", fredsCar.isNew())

    // Wat?
    fredsCar.paint("Cool Blue")
    fmt.Printf("%#v\n", fredsCar)

    // NOTE: Slides
    // NOTE: Same call either value or pointer receiver. Compiler adjust with & or * depending on type.

    domsCar := &car{
		year:  1971,
		make:  "Ford",
		model: "Pinto",
		color: "Puke Green",        
    }

    // (*domsCar).isNew()
    domsCar.isNew()
    // domsCar.paint()
    domsCar.paint("Bronco Orange")

}

func paint(c car, color string) {
    c.color = color
}
```

They're are expections as slice, function, map, channel, interface are Reference types ie they are intrisic pointers

```go
func 
```

> Run thru installing GO on OSX. Ensure GOROOT, GOPATH are set!


## Your turn...

### Lab1.1 ![alt text](https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon24.png "Lab1.1") 
---

### Mission
> Grab the default GO program from the playground and run it locally to make
sure you have a valid installation.

### Directions

```shell
brew update && brew install go git mercurial
export GOPATH=$HOME/your_go_workspace
export PATH=$PATH:$GOPATH/bin 
mkdir -p $GOPATH/src/github.com/your_git_username/go_training/labs/lab1.1
cd $GOPATH/src/github.com/your_git_username/go_training/labs/lab1.1
touch main.go
# Using your fav editor, copy/paste playground sample app
```

### Expectations

```shell
go run main.go # Outputs Hello, playground
go build -o lab1.1 main.go # Builds a go exec named lab1.1
lab1.1 # Run your executable
```

---

### Interface and Polymorphism

GO is very cool when it comes to defining interfaces ie You really don't have to statically declare that
a given user-defined type implements an interface. Very much inline with DuckTyping here with a bit of a twist ;-(
GO stdlib is filled with interface. Checkout the cool io.Reader/io.Writer

```go
type (
	painter interface {
		paint(color string)
	}

	printer interface {
		print()
	}

	car struct {
		brand, color string
    }

    truck struct {
        brand, color string
    }
)

func (c *car) paint(color string) {
	c.color = color
}

func (c car) print() {
	fmt.Printf("%#v\n", c)
}

func (t *truck) paint(color string) {
    c.color = color
}

func (t truck) print() {
    fmt.Printf("%#v\n", t)
}

func main() {
	c := car{brand: "Ford", color: "Puke Green"}

	paintIt(&c, "Cool Blue")
	printIt(c)

	var c1 = new(car)
	paintIt(c1, "Blue")
	printIt(c1)

	c2 := car{}
	//paintIt(c2, "Red")
	printIt(c2)

    t1 := truck{brand: "Chevy", color: "Red"}
    paintIt(&t1, "Yellow")
}

func printIt(p printer) {
	p.print()
}

func paintIt(p painter, color string) {
	p.paint(color)
}
```

## Dealing with JSON

### Interface Approach

```go
	b := []byte(`{"name":"Fernand","metrics":["6.3","180"]}`)

	m := map[string]interface{}{}

	fmt.Println("Decoder #1")
	// Decoding #1 Preferred!
	err := json.NewDecoder(bytes.NewReader(b)).Decode(&m)
	if err != nil {
		panic(err)
	}

	fmt.Println(m["name"])
	fmt.Println(m["metrics"].([]interface{})[0])
	fmt.Println(m["metrics"].([]interface{})[1])

	fmt.Println("Decoder #2")
	err = json.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}
	
	fmt.Println(m["name"])
	fmt.Println(m["metrics"].([]interface{})[0])
	fmt.Println(m["metrics"].([]interface{})[1])

	// Encoding JSON
	buff := new(bytes.Buffer)
	err = json.NewEncoder(buff).Encode(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(buff)
```

### JSON Types

```go
// NOTE: !!!Start LOWERCASE - BUG loading failed cozed non exported!!
type person struct {
	Name         string    `json:"name"`
	Measurements []float64 `json:"metrics"`
}

func main() {
	b := []byte(`{"name":"Fernand","metrics":[6.3,180]}`)

	p := new(person)

	fmt.Println("Decoder #1")
	// Decoding #1 Preferred!
	err := json.NewDecoder(bytes.NewReader(b)).Decode(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(p.Name)
	fmt.Println(p.Measurements)

	// Encoding JSON
	buff := new(bytes.Buffer)
	err = json.NewEncoder(buff).Encode(&p)
	if err != nil {
		panic(err)
	}
	fmt.Println(buff)
}
```

## Building a Web Service

GOTO Live example web_service

## Your turn...

### Lab2.2 ![alt text](https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon24.png "Lab2.2") 
---

### Mission
> Write a Web Service to convert roman to arabic and vice versa. The service output should be a JSON document with the following fields:
> o status : http response code
> o arabic_number: the arabic number
> o roman_numeral: the roman glyph
> o url: the url to run the inverse converter ie if roman was requested formulate the arabic conversion url or vice versa.

### Expectations

```shell
curl -XGET http://localhost:3000/roman?n=10 => {"status":200,"arabic_number":10,"roman_numeral":"X","url":"http://localhost:8080/arabic?g=X"}
curl -XGET http://localhost:3000/arabic?g=X => {"status":200,"arabic_number":10,"roman_numeral":"X","url":"http://localhost:8080/roman?n=10"}
```

## Web Service API

### TODO write a wrapper API that make http calls to basic web service

## Your turn...

### Lab2.3 ![alt text](https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon24.png "Lab2.3") 
---

### Mission
> Write an API client backed by the WebService you've just wrote. That is reuse the exact same calls as we did in 
> the roman to arabic converter but now by calling your DialARoman Web service.
> ExtraCredit - call your classmate DialARoman Web Service ;-)

### Directions

```shell
go run main.go # Starts your webserver on port 3000
```

### Expectations

```go
ToRoman(10) -> X # => by calling http://localhost:3000/roman?n=10
ToArabic("X") -> 10 # => by calling http://localhost:3000/arabic?g=X
```

