# DIProto

## DataInput Encoding/Decoding module

Package **github.com/juzzeast/clickhouse/diproto** exports two functions:  
Encode(diproto.DataInput) ([]byte, error)  
Decode([]byte) (diproto.DataInput, error)  

and a **DataInput** type (effectively, alias to **[]interface{}**).

To be successfully encoded, **DataInput** var should ONLY contain elements of types string, int32 and nested **DataInput** itself.  
If any element of different type is present, Encode will return error.

## DataInput Encoding/Decoding protocol

For space optimization and encoding/decoding performance, the **DataInput** instances are being incoded into binary format.

Every record (item) is preceeding with a type code (1 byte). **DataInput** itself has a code 255 (math.MaxUint8).  
For fixed-length type (*int32*) it follows with value.  
For variable length (*string*) it follows with length coded in bytes, then byte representation of data.  
For nested self (*DataInput*), it follows with elements count, then encoded **DataInput**.  

The encoding protocol specification (also included into types.go):
```
[encoded-datainput] ::= [type-code] [payload]

[type-code} ::= [string-type-code} | [int32-type-code} | [datainput-type-code}
[string-type-code} ::= 1
[int32-type-code} ::= 2
[datainput-type-code} ::= 255

[payload} ::= [string-payload} | [int32-payload} | [datainput-payload}
[string-payload} ::= [string-length} [string-bytes}
[int32-payload} ::= [int32-bytes}
[datainput-payload} ::= [datainput-elements-count} ([string-payload} | [int32-payload} | [datainput-payload})

[string-length} ::= [4-bytes}
[int32-bytes} ::= [4-bytes}
[datainput-elements-count} ::= [2 bytes}
```
