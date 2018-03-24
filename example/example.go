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
