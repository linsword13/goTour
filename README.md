# A quick tour of golang

> Complexity is multiplicative.
>
> -- <cite>Rob Pike</cite>

The **Go Programming Language** (golang) was conceived in 2007 by Robert Griesemer, Rob Pike and Ken Thompson, and officially announced in 2009. "Radical Simplicity" is the agenda of the design of the language. 

Golang is
* a [compiled][benchmarkLink] language,
* typed (somewhat strongly),
* high-level,
    - package system,
    - garbage collection,
    - first-class functions,
    - portable,
* open-sourced,
* "batteries included" with a standard library and good tooling support.

**Why are we even talking about another language here?!!**

## Outline
* [Basic structure](#user-content-basic-structure)
* [Language features](#user-content-language-features)
    - [The new C?](#user-content-the-new-c)
    - [Structural typing](#user-content-structural-typing)
    - [Concurrency](#user-content-concurrency)
* [Toolings](#user-content-toolings)
* [Mathematica](#user-content-mathematica)

## Basic structure
Hello world!

`./src/main.go`

---
```.go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)

    http.ListenAndServe(":3001", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, 你好!")
}
```
---

Time to build it:

---
```.bash
# read-compile-run
$ go run main.go

# or just read-compile
$ go build main.go
$ ./main
$ go clean

# or put the executable to a system location
$ go install main.go
$ main
$ go clean -i main.go
```
---

## Language features
The benefit of hindsight

![golang-family-tree][lineageImg]

### The new C?
Must be, since it is from Ken Thompson, right?

Short answer: no. The good news though, they did learn.

<h4>Better type declaration syntax</h4>
Remember this C mess?

---
```.c
#include <signal.h>

void (*signal(int signo, void (*func)(int)))(int);
```
---

With a bit help from `typedef`, better.

---
```.c
typedef void Sigfunc(int);

Sigfunc *signal(int, Sigfunc *);
```
---

Why can't we just read from left to right? Exactly.

---
```.go
// basic declaration
var i int
var arr [5]string
var hash map[string]float64
func(str string) int {
    tmp := str // type inference inside function block
    //
}
var foo func(string) int

// if we were to declare signal in golang
type Sigfunc func() Sigfunc

func signal(signo int, foo Sigfunc) Sigfunc {
    //
}
```
---
BTW, no forward function declaration is necessary in golang.

<h4>Function is first-class</h4>
Together with lexical scoping, first-class function has many interesting use cases. Closure is one example.

`./src/examples/func/func.go`

---
```.go
func age(initAge int) func() {
    i := initAge
    return func() {
        i++
        fmt.Printf("%v years old\n", i)
    }
}

func main() {
    celBirthday := age(18)
    for i := 0; i < 5; i++ {
        celBirthday()
    }
}
```
---

Being a functional programming language, Mathematica can easily do this as well.

---
```.Mathematica
makeAge[age_Integer] := Module[{year = age},
    Print[++year]&
]

celBDay = makeAge[18];

Do[celBDay[], 5]
```
---

<h4>Pointers, but safer</h4>
Pointer exists, but no pointer arithmetic. Golang is "call-by-value".

---
```.go
func add1(nPtr *int) int {
    *nPtr++ // interpreted as (*nPtr)++
    return *nPtr
}

var n int = 1
add1(&n)
```
---

Just a side note: in golang it's usually safe to return pointers from function calls.

---
```.c
int *foo(void) {
    int i = 1;
    return &i;
}

int *p = foo();
*p; /* compile error, if you are lucky */
```
---

---
```.go
func foo() *int {
    i := 1
    return &i
}

*foo() 
```
---

<h4>Multiple return values</h4>
---
```.go
func foo() (int, bool) {
    //
    return 1, true
}

a, ok := foo()
```
---

Let's bash C one more time...

---
```.c
#include <string.h>

/* int strcmp(const char *str1, const char *str2) */

int main(void) {
    /* define str1, str2 here */
    if (!strcmp(str1, str2)) {
        println("They are equal!");
    }
}
```
---

Having multiple return values also gives golang a cleaner way for error handling.

---
```.go
func foo() (int, error) {
    //
    return 1, true
}

res, err := foo()
if err != nil {
    // continue, if no error
}
```
--- 

<h4>Defer calls</h4>
Say we have the following code.

---
```.Mathematica
Catch@Module[{conn},
    conn = OpenSQLConnection["demo"];
    Check[op1[], Throw["err1"]];
    Check[op2[], Throw["err2"]];
    (* ... *)
    CloseSQLConnection[conn];
]
```
---

In golang, `defer` can help with this.

---
```.go
func foo() {
    conn := openConn()
    defer conn.Close()
    //...
}
```
---

<h4>Consistent Unicode encoding support</h4>
Golang uses UTF-8 behind the scene. 

Mathematica for instance, doesn't fully support the whole range of UTF-8 encoding.

---
```.Mathematica
FromCharacterCode[FromDigits["FFFF", 16], "UTF-8"] (* fine *)

FromCharacterCode[FromDigits["10000", 16], "UTF-8"] (* out-of-range *)
```
---

### Structural typing
Golang uses `interface` to generalize the behaviour of types. A concrete type satisfies an `interface` implicitly.

>In other words, don't check whether it IS-a duck: check whether it QUACKS-like-a duck, WALKS-like-a duck, etc, etc, depending on exactly what subset of duck-like behaviour you need to play your language-games with.
>
>-- <cite>Alex Martelli</cite>

Other object-oriented languages have similar concepts. In Java for example:

---
```.java
interface Thief {
    void steal(void);
}

// Person1 is a thief
class Person1 implements Thief {
    void steal(void) {
        //
    }
}

// Person2 is not
class Person2 {
    void steal(void) {
        //
    }
}
```
--- 

The "hello world" example is an example of using the interface.

---
```.go
// fmt.Fprint looks like this
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
    //
}

// io.Writer is an interface defined as
type Writer interface {
    Write(p []byte) (n int, err error)
}

// say we want to use fmt.Fprint for our purpose
type MyWriter struct{
    filePath string
}

func (w *MyWriter) Write(p []byte) (n int, err error) {
    f, err := os.Create(w.filePath)
    if err != nil {
        return 0, err
    }
    return f.Write(p)
}

// what's the meaning of interface{}
```
---

Interface embedding can be used to extend functionality.

---
```.go
type WLExpert interface {
    knowWL()
}

type CloudExpert interface {
    knowCloud()
}

type SW interface {
    WLExpert
    CloudExpert
    knowNKS()
    //...
}
```
---

###Concurrency

This is where golang really shines, I think.

<h4>Goroutine</h4>
`Goroutine` is what golang uses for running concurrent executions. One can regard it as similar to a native OS thread, only more lightweight.

A sequential program runs in a single goroutine. Additional ones are launched by calling the `go` statements.

`./src/examples/goroutines/goroutines.go`

---
```.go
func main() {
    i := []int{1, 2}
    go func() { i[0] = 3 }()
    time.Sleep(1 * time.Second)
    fmt.Printf("%v\n", i)
}
```
---

<h4>Channels</h4>
Channel is the mechanism for inter-communication between goroutines.

The previous example can be safer using channel.

`./src/examples/channels/channels.go`

---
```.go
func main() {
    i := []int{1, 2}
    done := make(chan struct{})
    go func() {
        i[0] = 3
        done <- struct{}{}
    }()
    <-done
    fmt.Printf("%v\n", i)
}
```
---

The repo contains a more complete example (`./src/examples/crawler/crawler.go`), a simple web link crawler, to illustrate how goroutines and channels fit together.  

[benchmarkLink]: http://benchmarksgame.alioth.debian.org/u64q/which-programs-are-fastest.html

[lineageImg]: images/lineage.png