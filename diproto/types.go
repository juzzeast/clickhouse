package diproto

type DataInput []interface{}

// Internal protocol codes for supported data types
const (
	di_type_string byte = byte(1)
	di_type_int32  byte = byte(2)
	di_type_di     byte = byte(255)
)

/* Encoding protocol:

<encoded-datainput> ::= <type-code> <payload>

<type-code> ::= <string-type-code> | <int32-type-code> | <datainput-type-code>
<string-type-code> ::= 1
<int32-type-code> ::= 2
<datainput-type-code> ::= 255

<payload> ::= <string-payload> | <int32-payload> | <datainput-payload>
<string-payload> ::= <string-length> <string-bytes>
<int32-payload> ::= <int32-bytes>
<datainput-payload> ::= <datainput-elements-count> (<string-payload> | <int32-payload> | <datainput-payload>)

<string-length> ::= <4-bytes>
<int32-bytes> ::= <4-bytes>
<datainput-elements-count> ::= <2 bytes>

*/
