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
    -  [Package system](#user-content-package-system)
    -  [Cross compilation](#user-content-cross-compilation)
    -  [Other tools](#user-content-other-tools)
* [Mathematica](#user-content-mathematica)
* [Summary](#user-content-summary)

## Basic structure
Hello world!

[`./src/main.go`](./src/main.go)

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

Time to build it:

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

[Back to TOC][toc]

## Language features
The benefit of hindsight

![golang-family-tree][lineageImg]

### The new C?
Must be, since it is from Ken Thompson, right?

Short answer: no. The good news though, they did learn.

<h4>Better type declaration syntax</h4>
Remember this C mess?

```.c
#include <signal.h>

void (*signal(int signo, void (*func)(int)))(int);
```

With a bit help from `typedef`, better.

```.c
typedef void Sigfunc(int);

Sigfunc *signal(int, Sigfunc *);
```

Why can't we just read from left to right? Exactly.

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

Being a functional programming language, Mathematica can easily do this as well.

```.Mathematica
makeAge[age_Integer] := Module[{year = age},
    Print[++year]&
]

celBDay = makeAge[18];

Do[celBDay[], 5]
```

<h4>Pointers, but safer</h4>
Pointer exists, but no pointer arithmetic. Golang is "call-by-value".

```.go
func add1(nPtr *int) int {
    *nPtr++ // interpreted as (*nPtr)++
    return *nPtr
}

var n int = 1
add1(&n)
```

Just a side note: in golang it's usually safe to return pointers from function calls.

```.c
int *foo(void) {
    int i = 1;
    return &i;
}

int *p = foo();
*p; /* compile error, if you are lucky */
```

```.go
func foo() *int {
    i := 1
    return &i
}

*foo() 
```

<h4>Multiple return values</h4>
```.go
func foo() (int, bool) {
    //
    return 1, true
}

a, ok := foo()
```

Let's bash C one more time...

```.c
/* confusing return value */
#include <string.h>

/* int strcmp(const char *str1, const char *str2) */

int main(void) {
    /* define str1, str2 here */
    if (!strcmp(str1, str2)) {
        println("They are equal!");
    }
}

/* another example: errno in UNIX programming */
#include <errno.h>

int foo(void) {
    /* ... */
    if (wrong()) {
        errno = ENOENT;
        return(-1);
    }
    return(0);
}

int main(void) {
    int res = foo();
    if (res < 0) {
        log(errno);
    }
    /* another call may overwrite the errno */
    foo2();
    return(0);
}

```

Having multiple return values also gives golang a cleaner way for error handling.

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

<h4>Defer calls</h4>
Say we have the following code.

```.Mathematica
Catch@Module[{conn},
    conn = OpenSQLConnection["demo"];
    Check[op1[], Throw["err1"]];
    Check[op2[], Throw["err2"]];
    (* ... *)
    CloseSQLConnection[conn];
]
```

In golang, `defer` can help with this.

```.go
func foo() {
    conn := openConn()
    defer conn.Close()
    //...
}
```

<h4>Consistent Unicode encoding support</h4>
Golang uses UTF-8 behind the scene. 

Mathematica for instance, doesn't fully support the whole range of UTF-8 encoding.

```.Mathematica
FromCharacterCode[FromDigits["FFFF", 16], "UTF-8"] (* fine *)

FromCharacterCode[FromDigits["10000", 16], "UTF-8"] (* out-of-range *)
```

### Structural typing
Golang uses `interface` to generalize the behaviour of types. A concrete type satisfies an `interface` implicitly.

>In other words, don't check whether it IS-a duck: check whether it QUACKS-like-a duck, WALKS-like-a duck, etc, etc, depending on exactly what subset of duck-like behaviour you need to play your language-games with.
>
>-- <cite>Alex Martelli</cite>

Other object-oriented languages have similar concepts. In Java for example:

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

The "hello world" example is an example of using the interface.

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

Interface embedding can be used to extend functionality.

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

###Concurrency

This is where golang really shines, I think.

<h4>Goroutine</h4>
`Goroutine` is what golang uses for running concurrent executions. One can regard it as similar to a native OS thread, only more lightweight.

A sequential program runs in a single goroutine. Additional ones are launched by calling the `go` statements.

`./src/examples/goroutines/goroutines.go`

```.go
func main() {
    i := []int{1, 2}
    go func() { i[0] = 3 }()
    time.Sleep(1 * time.Second)
    fmt.Printf("%v\n", i)
}
```

<h4>Channels</h4>
Channel is the mechanism for inter-communication between goroutines.

The previous example can get rid of the race condition using channel.

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

The repo contains a more complete example (`./src/examples/crawler/crawler.go`), a simple web link crawler, to illustrate how goroutines and channels fit together.

[Back to TOC][toc]

##Toolings

###Package system
Golang is opinionated. One example is its package system.

`./src/examples/testing/main.go`, `./src/examples/testing/helper/helper.go`

---
```.go

// helper package
package helper

import "fmt"

var ExportedVar = 1
var secretVar = 2

func ExportedFunc() int {
    fmt.Println("Exported")
    return secretVar
}

func secretFunc() {
    fmt.Println("You can't see me!")
}


// main package
package main

import (
    "./helper"
    "fmt"
)

func main() {
    int := helper.ExportedFunc()
    fmt.Println("we get", int, helper.ExportedVar)
}

```

Golang also makes it easy to work with version-controlled packages. For instance, the tool `go get` can grab packages from common host sites such as GitHub, GitLab.

```.bash
$ go get github.com/linsword13/goTour
```

###Cross compilation
`go build` is all we need for generating binaries that work with other platforms.

```.bash
# This gives all the supported OSs and architectures
$ go tool dist list

# This compiles the source to Windows x86-64
GOOS=windows GOARCH=amd64 go build source.go
```

A sample makefile can be found in `./src/examples/crawler/Makefile`.

###Other tools
| Tool | Description |
| :---: | :---: |
| `go build` | compile packages and dependencies |
| `go clean` | remove object files |
| `go doc` | show documentation for package or symbol |
| `go env` | print Go environment information |
| `go fix` | run go tool fix on packages |
| `go fmt` | run gofmt on package sources |
| `go generate` | generate Go files by processing source |
| `go get` | download and install packages and dependencies |
| `go install` | compile and install packages and dependencies |
| `go list` | list packages |
| `go run` | compile and run Go program |
| `go test` | test packages |
| `go tool` | run specified go tool |
| `go version` | print Go version |
| `go vet` | run go tool vet on packages |

`go test` and `go doc` are commonly used.

`./src/examples/testing/helper/helper_test.go`

---
```.go
package helper

import "testing"

func TestExportedFunc(t *testing.T) {
    res := ExportedFunc()
    if res != 2 {
        t.Errorf("expect %v, instead got %v", 2, res)
    }
}
```

```.bash
# run the test
$ go test ./src/examples/testing/helper

# ask for function documentation
go doc fmt.Sprintf
```

And there are more. Profiling, benchmarking, documentation generation, text editor support..., all these come with the standard golang installation.

[Back][toc]

##Mathematica
Golang can generate C-style shared libraries. `LibraryLink` can then be used to access these golang-generated libraries from Mathematica.

Here is the complete workflow of a trivial example.

1. Write up the implementation in golang.

    `./src/examples/mathematica/goSquare/main.go`

    ---
    ```.go
    package main

    import "C"

    //export GoSquare
    func GoSquare(n int) int {
        return n * n
    }

    func main() {}
    ```

2. Build the shared library.

    ```.bash
    # use .dylib on Mac
    go build -o ./goSquare.so -buildmode=c-shared ./goSquare
    ```
The actual Makefile can be found at `./src/examples/mathematica/Makefile`.

3. Include the shared library in a C file, using LibraryLink's interface.

    `./src/examples/mathematica/goSquare.c`

    ---
    ```.c
    #include "WolframLibrary.h"
    #include "libgoSquare.h"

    DLLEXPORT int callGoFunc(WolframLibraryData libData, mint Argc, MArgument *Args, MArgument Res) {
        mint in;
        in = MArgument_getInteger(Args[0]);
        int res = GoSquare((int)in);
        MArgument_setInteger(Res, res);
        return LIBRARY_NO_ERROR;
    }
    ```

4. Final step can be done in Mathematica, using the CCompilerDriver package.

    `./src/examples/mathematica/example.nb`

    ---
    ```.Mathematica
    SetDirectory[NotebookDirectory[]];
    Needs["CCompilerDriver`"]

    lib = CreateLibrary[
        {"callGoFunc.c"}, 
        "callGoFunc", 
        "Debug" -> True,
        "IncludeDirectories" -> NotebookDirectory[],
        "Libraries" -> "goSquare"
    ];

    callGo = LibraryFunctionLoad[lib, "callGoFunc", {Integer}, Integer]

    callGo[3]
    9                 
    ```

Note: the above works fine on Linux, but I haven't got it to work on Mac yet. Golang's support for creating C-style dynamic library on Windows is being [worked on](https://github.com/golang/go/issues/11058).

[Back to TOC][toc]

##Summary
There are some obvious lackings of golang, but overall it is a language with simplicity at its core and an ever-increasing community.

In case you want to read more:

* [golang.org](https://golang.org)
* [The Go Programming Language](https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440)
* [The source](https://github.com/golang/go)

[Back to top][top]

[benchmarkLink]: http://benchmarksgame.alioth.debian.org/u64q/which-programs-are-fastest.html

[lineageImg]: images/lineage.png

[toc]: #user-content-outline

[top]: #user-content-a-quick-tour-of-golang