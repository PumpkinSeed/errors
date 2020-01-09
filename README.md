# Errors

Simple error package for Go with minimal allocation and high performance. Optimized to keep error on function stack.

### Motivation

I found out that the available error packages too complex for the simple purpose or not achieve the simple goal.

### Purpose & Goal

Maintain an error chain (some kind of list of errors), and let the system to check whether the error's type caused the actual error or not.

### Usage

```go
package main

import (
    stdlib "errors"

    "github.com/PumpkinSeed/errors"
)

var ErrGlobal = errors.New("global err")
var ErrGlobal2 = errors.New("global err 2")
var ErrNotUsed = errors.New("not used err")

func main() {
    err := f3()
    
    stdlib.Is(err, ErrGlobal) // true
    stdlib.Is(err, ErrGlobal2) // true
    stdlib.Is(err, ErrNotUsed) // false

    println(err.Error()) // "global err 2: global err: string1"
}

func f1() error {
    return errors.New("string1")
}

func f2() error {
    return errors.Wrap(f1(), ErrGlobal)
}

func f3() error {
    return errors.Wrap(f2(), ErrGlobal2)
}
```  