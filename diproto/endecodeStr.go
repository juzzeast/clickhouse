package diproto

import "errors"

// Model for type string record encoding: [TypeCode - 1 byte][Length - 4 bytes][Data <Length> bytes]
// By spec, string could be up to 1,000,000 chars, hence it requires 4 byte integer to accomodate this length
func encodeStr(s string) []byte {
	r := make([]byte, 1)
	r[0] = di_type_string
	bytesOfString := []byte(s)
	bytesOfLen := int32ToBytes(int32(len(bytesOfString)))
	r = append(r, bytesOfLen...)
	r = append(r, bytesOfString...)
	return r
}

// Assuming to have a slice of []byte starting with encoded string record
// Returns decided string AND number of slice bytes processed
func decodeStr(bs []byte) (string, int, error) {
	i := 0
	var rslt string
	if len(bs) < 5 {
		return rslt, i, errors.New("wrong format of record to decode: not enough bytes for string")
	}
	if bs[i] != di_type_string {
		return rslt, i, errors.New("wrong type of record to decode: must be of di_type_string")
	} else {
		i++
		rsltLen, err := bytesToInt32(bs[i : i+4])
		if err != nil {
			return rslt, i, errors.New("wrong format of record to decode: can't retrieve string length")
		}
		i += 4
		if int(rsltLen) > len(bs[i:]) {
			return rslt, i, errors.New("wrong format of record to decode: stored string length mismatch actual data")
		}
		rslt = string(bs[i : i+int(rsltLen)])
		i += int(rsltLen)
	}
	return rslt, i, nil
}
