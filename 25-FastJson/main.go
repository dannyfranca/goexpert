package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var p fastjson.Parser
	jsonData := `{ "name": "John Jr", "age": 5, "cat": null, "cars": ["Ford", "BMW", "Fiat"], "father": { "name": "John", "age": 30 }}`

	value, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("name: %s\n", value.GetStringBytes("name"))
	fmt.Printf("age: %d\n", value.GetInt("age"))
	fmt.Printf("cat: %s\n", value.GetStringBytes("car"))
	fmt.Printf("cars: %s\n", value.GetStringBytes("cars"))
	fmt.Printf("father: %s\n", value.GetStringBytes("father"))

	cars := value.GetArray("cars")
	for i, v := range cars {
		fmt.Printf("cars[%d]: %s\n", i, v.GetStringBytes())
	}

	father := value.GetObject("father")
	fmt.Printf("father.name: %s\n", father.Get("name"))
	fmt.Printf("father.age: %s\n", father.Get("age"))

	userJSON := father.String()
	var user User
	if err := json.Unmarshal([]byte(userJSON), &user); err != nil {
		panic(err)
	}
	fmt.Println(user.Name, user.Age)

	var sc fastjson.Scanner

	sc.Init(`   { "name": "John Jr", "age": 5, "cat": null, "cars": ["Ford", "BMW", "Fiat"], "father": { "name": "John", "age": 30 }} null "3"`)
	for sc.Next() {
		fmt.Printf("key: %s, value: %s\n", sc.Value(), sc.Value())
	}
	if sc.Error() != nil {
		panic(sc.Error())
	}
}
