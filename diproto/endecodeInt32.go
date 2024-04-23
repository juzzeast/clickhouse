package diproto

import "errors"

// Model for type int32 record encoding: [TypeCode - 1 byte][Data - 4 bytes]
func encodeInt32(i int32) []byte {
	r := make([]byte, 1)
	r[0] = di_type_int32
	bs := int32ToBytes(i)
	r = append(r, bs...)
	return r
}

// Assuming to have a slice of []byte starting with encoded int32 record
// Returns decoded int32 AND number of slice bytes processed
func decodeInt32(bs []byte) (int32, int, error) {
	var rslt int32
	i := 0
	if len(bs) < 5 {
		return rslt, i, errors.New("wrong format of record to decode: must be exactly 5 bytes for int32")
	}
	if bs[i] != di_type_int32 {
		return rslt, i, errors.New("wrong type of record to decode: must be of di_type_int32")
	} else {
		i++
		var err error
		rslt, err = bytesToInt32(bs[i : i+4])
		if err != nil {
			return rslt, i, errors.New("wrong format of record to decode: can't retrieve int32 value")
		}
		i += 4
	}
	return rslt, i, nil
}
