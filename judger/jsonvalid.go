package judger
import(
	"encoding/json"
)
/* check if the json string is valid */
/* actually should be judged by DFA, but anyway */
func JsonValid(jsonStr string) bool {
	isJSON := json.Valid([]byte(jsonStr))
	return isJSON
}