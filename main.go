package main

import (
	"JsonInterpreter/judger"
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
		// print(len(tmp))
		fmt.Print("[")
		for i := 0; i < len(tmp); i++ {
			// print(tmp[i].GetType())
			print_element(tmp[i])
			if i != len(tmp)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Print("]")
	case scanner.BEGIN_OBJECT: //print the object
		tmp := ele.GetObjectValue()
		fmt.Print("{")
		i := 0
		length := len(tmp.Object_)
		for k, v := range tmp.Object_ {
			i = i + 1
			fmt.Print(k, ": ")
			print_element(v)
			if condition := i == length; condition {
				break
			}
			fmt.Print(", ")
		}
		fmt.Print("}")
	}

}

func print_all_elements(o scanner.Object){
	/* print all the elements in the object */
	fmt.Println("{")
	for k, v := range o.Object_ {
		fmt.Print("\t")
		fmt.Print(k, ": ")
		print_element(v)
		fmt.Print(",")
		fmt.Println("")
	}
	fmt.Print("}")

}

func main() {
	content, err := os.ReadFile("json.txt")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(content))

	scanner.Jsource = string(content)
	if !judger.JsonValid(scanner.Jsource) {
		panic("Bad json format")
	} //optional, judge the json format
	scanner.J_length = len(scanner.Jsource)
	scanner.Scan()
	// var JsonObject scanner.Json_object = scanner.JObject[0]
	// print_value_of_object(JsonObject, "aaa")
	print_all_elements(scanner.JObject[0])
}
