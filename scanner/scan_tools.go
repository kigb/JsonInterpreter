package scanner
/* tools for scanning */
import (
	"strconv"
)

func get_forward() string {
	current_ = current_ + 1
	return string(Jsource[current_])
}

func is_eof() bool {
	return current_ >= J_length
}

func get_string() string {
	/* get the string from the source */
	var got_str string = ""
	current_ = current_ + 1
	for string(Jsource[current_]) != "\"" {
		got_str = got_str + string(Jsource[current_])
		current_ = current_ + 1
	}
	return got_str
}

func get_num() float64 {
	/* get the number from the source */
	var got_num string = ""
	for string(Jsource[current_]) != "," &&
		string(Jsource[current_]) != "}" &&
		string(Jsource[current_]) != "]" &&
		string(Jsource[current_]) != " " &&
		string(Jsource[current_]) != "\n" &&
		string(Jsource[current_]) != "\t" &&
		string(Jsource[current_]) != "\r" {
		got_num = got_num + string(Jsource[current_])
		current_ = current_ + 1
	}
	f, err := strconv.ParseFloat(got_num, 64)
	if err != nil {
		panic("Bad number format")
	}
	current_ = current_ - 1
	return f
}

func get_bool() bool {
	/* get the bool from the source */
	var got_bool string = ""
	for string(Jsource[current_]) != "," &&
		string(Jsource[current_]) != "}" &&
		string(Jsource[current_]) != "]" &&
		string(Jsource[current_]) != " " &&
		string(Jsource[current_]) != "\n" &&
		string(Jsource[current_]) != "\t" &&
		string(Jsource[current_]) != "\r" {
		got_bool = got_bool + string(Jsource[current_])
		current_ = current_ + 1
	}
	current_ = current_ - 1
	if got_bool == "true" {
		return true
	} else if got_bool == "false" {
		return false
	} else {
		panic("Bad bool format")
	}
}

func get_null() bool {
	/* get the null from the source */
	var got_null string = ""
	for string(Jsource[current_]) != "," &&
		string(Jsource[current_]) != "}" &&
		string(Jsource[current_]) != "]" &&
		string(Jsource[current_]) != " " &&
		string(Jsource[current_]) != "\n" &&
		string(Jsource[current_]) != "\t" &&
		string(Jsource[current_]) != "\r" {
		got_null = got_null + string(Jsource[current_])
		current_ = current_ + 1
	}
	current_ = current_ - 1
	if got_null == "null" || got_null == "NULL" {
		return true
	} else {
		panic("Bad null format")
	}
}

func get_array() []Element{
	/* get the array from the source */
	ans := make([]Element, 0)
	for !is_eof(){
		switch get_forward() {
			case "]":
				return ans
			case "[":
				ans = append(ans, Array_token{BEGIN_ARRAY, get_array()})//recursive call
			case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
				ans = append(ans, Num_token{NUMBER, get_num()})
			case "t", "f":
				ans = append(ans, Bool_token{BOOLEAN, get_bool()})
			case "n", "N":
				get_null()
				ans = append(ans, Null_token{NULL, 0})
			case "\"":
				ans = append(ans, Str_token{STRING, get_string()})	
		}
	}
	
	return ans
}