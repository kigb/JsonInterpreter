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