package main

import (
	"JsonInterpreter/scanner"
	"fmt"
	"log"
	"os"
	"JsonInterpreter/judger"
)

func print_value_of_object(obj scanner.Json_object, key string) {
	/* get the value of the key in the object */
	if obj.Get_element(key) == nil {
		fmt.Println("No such key")
		return
	}
	print_element(obj.Get_element(key))
}

func print_element(ele scanner.Element) {
	/* print the value of the element */
	switch ele.GetType() {
	case scanner.NUMBER:
		fmt.Print(ele.GetNumValue())
	case scanner.STRING:
		fmt.Print(ele.GetStringValue())
	case scanner.BOOLEAN:
		fmt.Print(ele.GetBoolValue())
	case scanner.NULL:
		fmt.Println("$null")
	case scanner.BEGIN_ARRAY: //print the array
		tmp := ele.GetElementValue()
		fmt.Print("[")
		for i := 0; i < len(tmp); i++ {
			print_element(tmp[i])
			if i != len(tmp)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Print("]")
	}
}

func main() {
	content, err := os.ReadFile("json.txt")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(content))

	scanner.Jsource = string(content)
	if !judger.JsonValid(scanner.Jsource){
		panic ("Bad json format")
	}//optional, judge the json format
	scanner.J_length = len(scanner.Jsource)
	scanner.Scan()
	var JsonObject scanner.Json_object = scanner.JObject
	print_value_of_object(JsonObject, "aaa")
}
