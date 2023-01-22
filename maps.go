// Maps are go's built-in associative data type (hashes or dicts in other languages)
package main

import "fmt"

func main() {
	// 1. empty map: make(map[key-type]val-type)
	m := make(map[string]int)
	// 2. set key pairs
	m["K1"] = 100
	m["K2"] = 20
	// 3. get value with key
	fmt.Println("map:", m)
	fmt.Println("map:", m["K1"])
	// 4. delete key/value pairs from a map
	delete(m, "K2")
	fmt.Println("map:", m)
	//5. The second return value when getting a value from a map indicates if the key was present in the map.
	// This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
	// ?????????????????????????????????????????????????/
	_, prs := m["k2"]
    fmt.Println("prs:", prs)
}
