package scanner

// use package scanner to define the token types and the implementation
var Jsource string = "" //source string, to be scanned
var current_ = -1       //current position of the scanner, lowercase as private
var J_length = 0        //length of the source string, needed to be initialized in main.go
// var is_inobject = false //flag to indicate if the scanner is in an object, if no, then panic
var JObject []Object = make([]Object, 0) //object to store the result of the scanner
var current_object_index = -1            //current index of the object
// var key = true          //to judge whether is the key, before : is key, after : is value
var current_key = "" //to store the current key
var isFatherObject = true
var JArray []Element = make([]Element, 0)

func Scan_object() Object {
	if is_eof() {
		return Object{}
	}

	cur_str := get_forward()

	if cur_str == "{" && isFatherObject {
		isFatherObject = false
		JObject = append(JObject, Object{make(map[string]Element), BEGIN_OBJECT})
		current_object_index = current_object_index + 1
		JObject[current_object_index].Object_ = make(map[string]Element) //initialize the map
		Scan_object()
		isFatherObject = true //restart
		return JObject[current_object_index]
	}
	/* judge whether in object */
	// if !is_inobject && cur_str != "{" {
	// 	panic("Not in object, JSON should start with {")
	// }

	switch cur_str {
	case "{":
		// is_inobject = true
		JObject = append(JObject, Object{make(map[string]Element), BEGIN_OBJECT}) //append a new object
		pre_index := current_object_index
		// print("pre_index:",pre_index)
		current_object_index = current_object_index + 1
		JObject[current_object_index].Object_ = make(map[string]Element) //initialize the map
		if !isFatherObject && current_key == "" {
			panic("No key is available, bad json format")
		}
		save_cur_key := current_key
		current_key = ""
		Scan_object() //recursive call,get the obj
		JObject[pre_index].Object_[save_cur_key] = JObject[pre_index+1]
		current_key = ""
		current_object_index = pre_index

	case "}":
		// is_inobject = false
		return JObject[current_object_index]

	case "\"":
		/* string */
		str := get_string()
		if current_key == "" {
			current_key = str
			break
		} // if it is key, then store the key and wait for the value
		tmp := Str_token{STRING, str}
		var tmp_ele Element = tmp
		JObject[current_object_index].Object_[current_key] = tmp_ele
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
		JObject[current_object_index].Object_[current_key] = tmp_ele
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
		JObject[current_object_index].Object_[current_key] = tmp_ele
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
		JObject[current_object_index].Object_[current_key] = tmp_ele
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
		JObject[current_object_index].Object_[current_key] = tmp_ele
		// println(current_key, ":", JObject.Object_[current_key].GetElementValue())
		current_key = "" //reset the key

	}
	Scan_object()
	return JObject[current_object_index]
}

func Scan_array() {
	if is_eof() {
		return
	}

	cur_str := get_forward()

	if cur_str == "[" {
		JArray = get_array()
	}

}

func Scan(){
	if(J_length == 0){
		panic("No source string")
	}

	if(string(Jsource[0]) == "["){
		Scan_array()
	}
	if(string(Jsource[0]) == "{"){
		Scan_object()
	}
}
