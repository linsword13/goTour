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
* [Basic structure](#basic-structure)
* [Language features](#language-features)
    - [The new C?](#the-new-c?)
* [Toolings](#toolings)
* [Mathematica](#mathematica)

## Basic structure
Hello world!

```{.go}
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

## Language features


[benchmarkLink]: http://benchmarksgame.alioth.debian.org/u64q/which-programs-are-fastest.html