package main

import (
	"fmt"

	"github.com/valyala/fastjson"
)

func main() {
	var p fastjson.Parser
	jsonData := `{ "name": "John", "age": 30, "car": null, "cars": ["Ford", "BMW", "Fiat"]}`

	v, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("name: %s\n", v.GetStringBytes("name"))
	fmt.Printf("age: %d\n", v.GetInt("age"))
	fmt.Printf("car: %s\n", v.GetStringBytes("car"))

	cars := v.GetArray("cars")
	for i, v := range cars {
		fmt.Printf("cars[%d]: %s\n", i, v.GetStringBytes())
	}
}
