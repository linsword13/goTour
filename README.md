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
* [Toolings](#user-content-toolings)
* [Mathematica](#user-content-mathematica)

## Basic structure
Hello world!

`./src/main.go`

---
```.go
package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)

    http.ListenAndServe(":3001", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, 你好!"))
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

[benchmarkLink]: http://benchmarksgame.alioth.debian.org/u64q/which-programs-are-fastest.html

[lineageImg]: images/lineage.png