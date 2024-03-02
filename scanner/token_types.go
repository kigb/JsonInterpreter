package scanner
/* token types */
const (
	/* Non-value tokens */

	/* Begin non-value tokens */
	BEGIN_OBJECT    = iota // {
	END_OBJECT             // }
	BEGIN_ARRAY            // [
	END_ARRAY              // ]
	NAME_SEPARATOR         // :
	VALUE_SEPARATOR        // ,
	/* End non-value tokens */

	/* Value tokens */

	/* Begin value tokens */
	/* String */
	STRING // "string"
	/* Number */
	NUMBER // 123 or 1e10
	/* Boolean */
	BOOLEAN // true or false
	/* Null */
	NULL // null
	/* End value tokens */

	/* End of file */
	EOF // EOF
)