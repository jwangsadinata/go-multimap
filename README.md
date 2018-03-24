[![GoDoc](https://godoc.org/github.com/jwangsadinata/go-multimap?status.svg)](https://godoc.org/github.com/jwangsadinata/go-multimap) [![Build Status](https://travis-ci.org/jwangsadinata/go-multimap.svg)](https://travis-ci.org/jwangsadinata/go-multimap) [![Go Report Card](https://goreportcard.com/badge/github.com/jwangsadinata/go-multimap)](https://goreportcard.com/report/github.com/jwangsadinata/go-multimap) [![Coverage Status](https://coveralls.io/repos/github/jwangsadinata/go-multimap/badge.svg?branch=master&service=github)](https://coveralls.io/github/jwangsadinata/go-multimap?branch=master&service=github) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/jwangsadinata/go-multimap/blob/master/LICENSE)

# Go-Multimap

Implementation of the `multimap` data structure in [Go](https://www.golang.org/project/).

A multimap (sometimes also multihash or multidict) is a generalization of a map 
or associative array abstract data type in which more than one value may be 
associated with and returned for a given key. 

Current implementation only supports slice-based multimap, but an implementation 
for set-based multimap will also be implemented soon.

This package was heavily inspired by the Google Guava interface of MultiMap and 
written in the style of the [container](https://golang.org/pkg/container/) package.


References: 
[Wikipedia](https://en.wikipedia.org/wiki/Multimap), 
[Guava](https://google.github.io/guava/releases/19.0/api/docs/com/google/common/collect/Multimap.html)

## Installation ##

Install the package via the following:

    go get -u github.com/jwangsadinata/go-multimap

## Usage ##

The go-multimap package can be used similarly to the following:
```go
// example/example.go
package main

import (
	"fmt"

	"github.com/jwangsadinata/go-multimap/multimap"
)

func main() {
	m := multimap.New()
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a")

	fmt.Printf("All Entries: %v\n", m.Entries())
	fmt.Printf("All Keys: %v\n", m.Keys())
	fmt.Printf("Distinct Keys: %v\n", m.KeySet())
	fmt.Printf("All Values: %v\n\n", m.Values())

	value, _ := m.Get(1)
	fmt.Printf("The values with key 1 is: %v\n\n", value)

	m.Remove(4, "d")
	m.RemoveAll(1)

	fmt.Printf("Current size of multimap after deletion: %v\n\n", m.Size())

	fmt.Printf("Assert that (2, \"b\") is in the map: %v\n", m.Contains(2, "b"))
	fmt.Printf("Assert that there is a key 4 in the map: %v\n", m.ContainsKey(4))
	fmt.Printf("Assert that the value \"c\" is in the map: %v\n", m.ContainsValue("c"))

	m.Clear()
	fmt.Printf("Assert that multimap is empty: %v\n", m.Empty())
}
```

Example output:
```sh
$ go run example.go
All Entries: [{3 c} {4 d} {1 x} {1 a} {2 b}]
All Keys: [1 1 2 3 4]
Distinct Keys: [2 3 4 1]
All Values: [c d x a b]

The values with key 1 is: [x a]

Current size of multimap after deletion: 2

Assert that (2, "b") is in the map: true
Assert that there is a key 4 in the map: false
Assert that the value "c" is in the map: true
Assert that multimap is empty: true
```

Please see [the GoDoc API page](http://godoc.org/github.com/jwangsadinata/go-multimap) for a
full API listing.
