# Go Expirables
Expirables variables in Golang.

## Installation
To install this as a dependency, simply call :
`go get github.com/steve-nzr/expirables`

## Usage
How to use of this package :
```go
package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/steve-nzr/expirables"
)

var token = expirables.NewExpirable(func() interface{} {
	return uuid.New().String()
}, time.Second*4)

func main() {
	val := token.Get().(string)

	if uuid.MustParse(token.Get().(string)).String() == val {
		fmt.Printf("'%s' is an UUID !\n", val)
	}

	time.Sleep(time.Second * 4)

	val = token.Get().(string)

	if uuid.MustParse(token.Get().(string)).String() == val {
		fmt.Printf("'%s' stills an UUID (but is another) !\n", val)
	}
}
```

Which outputs (UUID are examples) :
```bash
user@deb:~$ go run main.go 
'c1b4b6da-fdcf-4644-a207-8a95e751b9d8' is an UUID !
'd1c208a2-c5e4-4c64-b553-1cd5afafd00c' stills an UUID !
```
