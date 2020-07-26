# geoshp
[![GoDoc](https://godoc.org/github.com/ShreyasSubhedar/geoshp?status.svg)](https://godoc.org/github.com/ShreyasSubhedar/geoshp)
[![Go Report Card](https://goreportcard.com/badge/github.com/ShreyasSubhedar/geoshp)](https://goreportcard.com/report/github.com/ShreyasSubhedar/geoshp)
[![Build Status](https://travis-ci.org/ShreyasSubhedar/geoshp.svg?branch=master)](https://travis-ci.org/ShreyasSubhedar/geoshp)
- Go package for computing volumes and surface area of  high dimensional shapes.

## Usage
#### Installation
go get github.com/ShreyasSubhedar/geoshp
#### Importing
import "github.com/ShreyasSubhedar/geoshp"

## Example: method invoke
```go
import (
	"fmt"

	"github.com/ShreyasSubhedar/geoshp"
)

func main() {
	// Declaring some variables
	var radius float64
	var height float64

	// Calling Scan() function for
	// scanning and reading the input
	// texts given in standard input
	fmt.Scan(&radius)
	fmt.Scan(&height)

	// Printing the given texts
	vol := geoshp.CyclinderVol(radius, height)
	fmt.Println("The volume of given cylinder : ", vol)
}

```
