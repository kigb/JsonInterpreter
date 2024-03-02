package scanner
/* element interface, like union in C/C++ */
/* to represent elements in Objects, like string, boolean and etc. */
type Element interface{
	GetType() int
	GetStringValue() string
	GetNumValue() float64
	GetBoolValue() bool
	GetNullValue() int
	GetElementValue() []Element
}

/* STR */
type Str_token struct {
	Token_type int
	Value string
}

func (s Str_token)GetType() int{
	return s.Token_type
}

func (s Str_token)GetStringValue() string{// return the string value
	return s.Value
}

func (s Str_token)GetNumValue() float64{//unused, return 0
	return 0
}

func (s Str_token)GetBoolValue() bool{//unused, return false
	return false
}

func (s Str_token)GetNullValue() int{//unused, return 0
	return 0
}

func (s Str_token)GetElementValue() []Element{//unused, return nil
	return nil
}

/* NUM */
type Num_token struct {
	Token_type int
	Value float64
}

func (n Num_token)GetType() int{
	return n.Token_type
}

func (n Num_token)GetNumValue() float64{// return the number value
	return n.Value
}

func (n Num_token)GetStringValue() string{//unused, return ""
	return ""
}

func (n Num_token)GetBoolValue() bool{//unused, return false
	return false
}

func (n Num_token)GetNullValue() int{//unused, return 0
	return 0
}

func (n Num_token)GetElementValue() []Element{//unused, return nil
	return nil
}
/* BOOL */
type Bool_token struct {
	Token_type int
	Value bool
}

func (b Bool_token)GetType() int{
	return b.Token_type
}

func (b Bool_token)GetBoolValue() bool{// return the bool value
	return b.Value
}

func (b Bool_token)GetStringValue() string{//unused, return ""
	return ""
}

func (b Bool_token)GetNumValue() float64{//unused, return 0
	return 0
}

func (b Bool_token)GetNullValue() int{//unused, return 0
	return 0
}

func (b Bool_token)GetElementValue() []Element{//unused, return nil
	return nil
}

/* NULL */
type Null_token struct {
	Token_type int
	Value int	// 0 for null
}

func (n Null_token)GetType() int{
	return n.Token_type
}

func (n Null_token)GetStringValue() string{//unused, return ""
	return ""
}

func (n Null_token)GetNumValue() float64{//unused, return 0	
	return 0
}

func (n Null_token)GetBoolValue() bool{//unused, return false
	return false
}

func (n Null_token)GetNullValue() int{// return the null value
	return 0
}

func (n Null_token)GetElementValue() []Element{//unused, return nil
	return nil
}

/* ARRAY */
type Array_token struct {
	Token_type int
	Value []Element
}

func (a Array_token)GetType() int{
	return a.Token_type
}

func (a Array_token)GetElementValue() []Element{// return the array value
	return a.Value
}

func (a Array_token)GetStringValue() string{//unused, return ""
	return ""
}

func (a Array_token)GetNumValue() float64{//unused, return 0
	return 0
}

func (a Array_token)GetBoolValue() bool{//unused, return false
	return false
}

func (a Array_token)GetNullValue() int{//unused, return 0
	return 0
}

func (a Array_token)GetArrayValue() []Element{// return the array value
	return a.Value
}