package scanner

// use package scanner to define the token types and the implementation
var Jsource string = "" //source string, to be scanned
var current_ = -1       //current position of the scanner, lowercase as private
var J_length = 0        //length of the source string, needed to be initialized in main.go
var is_inobject = false //flag to indicate if the scanner is in an object, if no, then panic
var JObject = Object{}  //object to store the result of the scanner
// var key = true          //to judge whether is the key, before : is key, after : is value
var current_key = "" //to store the current key

func Scan() {
	if is_eof() {
		return
	}

	cur_str := get_forward()

	/* judge whether in object */
	if !is_inobject && cur_str != "{" {
		panic("Not in object, JSON should start with {")
	}

	switch cur_str {
	case "{":
		is_inobject = true
		JObject.Object_ = make(map[string]Element) //initialize the map

	case "}":
		is_inobject = false
		return

	case "\"":
		/* string */
		str := get_string()
		if current_key == "" {
			current_key = str
			break
		} // if it is key, then store the key and wait for the value
		tmp := Str_token{STRING, str}
		var tmp_ele Element = tmp
		JObject.Object_[current_key] = tmp_ele
		// println(current_key, ":", JObject.Object_[current_key].GetStringValue())
		current_key = "" //reset the key

	case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
		/* number */
		if current_key == "" {
			panic("No key is available, bad json format")
		}
		num := get_num()
		tmp := Num_token{NUMBER, num}
		var tmp_ele Element = tmp
		JObject.Object_[current_key] = tmp_ele
		// println(current_key, ":", JObject.Object_[current_key].GetNumValue())
		current_key = "" //reset the key

	case "t", "f":
		/* boolean */
		if current_key == "" {
			panic("No key is available, bad json format")
		}
		bool_cur := get_bool()
		tmp := Bool_token{BOOLEAN, bool_cur}
		var tmp_ele Element = tmp
		JObject.Object_[current_key] = tmp_ele
		// println(current_key, ":", JObject.Object_[current_key].GetBoolValue())
		current_key = "" //reset the key

	case "n", "N":
		/* null */
		if current_key == "" {
			panic("No key is available, bad json format")
		}
		get_null() //if it is null, then no need to store the value,if not, panic
		tmp := Null_token{NULL, 0}
		var tmp_ele Element = tmp
		JObject.Object_[current_key] = tmp_ele
		// println(current_key, ":", JObject.Object_[current_key].GetNullValue())
		current_key = "" //reset the key

	case "[":
		/* array */
		if current_key == "" {
			panic("No key is available, bad json format")
		}
		ary := get_array()
		tmp := Array_token{BEGIN_ARRAY, ary}
		var tmp_ele Element = tmp
		JObject.Object_[current_key] = tmp_ele
		// println(current_key, ":", JObject.Object_[current_key].GetElementValue())
		current_key = "" //reset the key

	}

	Scan()
}
