package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	var j = []byte(`{"size": "nầm", "color": 444}`)
	// a map container to decode the JSON structure into
	c := make(map[string]json.RawMessage)

	// unmarschal JSON
	e := json.Unmarshal(j, &c)

	// panic on error
	if e != nil {
		panic(e)
	}

	// a string slice to hold the keys
	k := make([]string, len(c))

	// iteration counter
	i := 0

	// copy c's keys into k
	for s, v := range c {
		k[i] = s
		value := string(v)
		value = strings.ReplaceAll(value, "\"", "")
		fmt.Println(string(s))
		fmt.Println(value)
		i++
	}

	// output result to STDOUT
	fmt.Printf("%#v\n", k)
}
