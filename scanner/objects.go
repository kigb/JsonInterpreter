package scanner
/* object definition, a map of string and elements */
type Json_object interface{
	Get_element(string) Element
}

type Object struct{
	Object_ map[string]Element
	Token_type int
}

func (o Object)Get_element(str string) Element{
	if _, ok := o.Object_[str]; !ok{
		return nil
	}
	return o.Object_[str]
}