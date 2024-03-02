package main

import (
	"JsonInterpreter/scanner"
	"fmt"
	"log"
	"os"
)

func print_value_of_object(obj scanner.Json_object, key string) {
	/* get the value of the key in the object */
	if obj.Get_element(key) == nil {
		fmt.Println("No such key")
		return
	}
	switch obj.Get_element(key).GetType() {
	case scanner.NUMBER:
		fmt.Println(obj.Get_element(key).GetNumValue())
	case scanner.STRING:
		fmt.Println(obj.Get_element(key).GetStringValue())
	case scanner.BOOLEAN:
		fmt.Println(obj.Get_element(key).GetBoolValue())
	case scanner.NULL:
		fmt.Println("null")
	}
}

func main() {
	content, err := os.ReadFile("json.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(content))
	scanner.Jsource = string(content)
	scanner.J_length = len(scanner.Jsource)
	scanner.Scan()
	var JsonObject scanner.Json_object = scanner.JObject
	print_value_of_object(JsonObject, "123sd4")
}
