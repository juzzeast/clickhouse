package diproto

import (
	"errors"
	"reflect"
)

// Model for type DataInput encoding: [TypeCode - 1 byte][Data Records Count - 2 bytes][Data]
// Unlike other supported types, DataInput encodes number of elements instead of length
// Supports values of string, int32 and (nested) DataInput types, any other would result in an error
func Encode(di DataInput) ([]byte, error) {
	var r []byte
	r = append(r, di_type_di)
	diLen := int16(len(di))
	bdiLen := int16ToBytes(diLen)
	r = append(r, bdiLen...)
	for _, v := range di {
		it := reflect.TypeOf(v)
		switch it.String() { // TODO: think of switching to type comparison through types.Identical
		case "int32":
			iv := v.(int32)
			bi32 := encodeInt32(iv)
			r = append(r, bi32...)
		case "string":
			str := v.(string)
			bstr := encodeStr(str)
			r = append(r, bstr...)
		case "diproto.DataInput":
			di := v.(DataInput)
			bdi, err := Encode(di)
			if err != nil {
				return r, err
			}
			r = append(r, bdi...)
		default:
			return r, errors.New("can't encode DataInput: record of unsupported type")
		}
	}
	return r, nil
}

// Assuming to have a slice of []byte starting with encoded DataInput record
func Decode(bDI []byte) (DataInput, error) {
	i := 0
	rsltDI := DataInput{}
	if len(bDI) < 3 {
		return rsltDI, errors.New("wrong format of record to decode: not enough bytes for DataInput")
	}
	if bDI[i] != di_type_di {
		return rsltDI, errors.New("wrong type of record to decode: must be of di_type_di")
	} else {
		i++
		subDiLen, err := bytesToInt16(bDI[i : i+2])
		if err != nil {
			return rsltDI, errors.New("wrong format of record to decode: can't retrieve DataInput length")
		}
		i += 2
		rsltDI, _, err = decodeR(bDI[i:], int(subDiLen))
		if err != nil {
			return rsltDI, err
		}
	}
	return rsltDI, nil
}

func decodeR(bDI []byte, count int) (DataInput, int, error) {
	var rsltDI DataInput
	counter := 0
	i := 0
	for counter < count {
		if i >= len(bDI) {
			return rsltDI, i, errors.New("wrong format of record to decode: reached the end of the encoded DataInput unexpectedly")
		}
		switch bDI[i] {
		case di_type_string:
			str, delta, err := decodeStr(bDI[i:])
			if err != nil {
				return rsltDI, i, err
			}
			// fmt.Printf("decoded string=%s; delta=%d; counter=%d\n", str, delta, counter)
			rsltDI = append(rsltDI, str)
			i += delta
			counter++
		case di_type_int32:
			i32, delta, err := decodeInt32(bDI[i:])
			if err != nil {
				return rsltDI, i, err
			}
			//fmt.Printf("decoded int32=%d; delta=%d; counter=%d\n", i32, delta, counter)
			rsltDI = append(rsltDI, i32)
			i += delta
			counter++
		case di_type_di:
			i++
			subDILen, err := bytesToInt16(bDI[i : i+2])
			if err != nil {
				return rsltDI, i, errors.New("wrong format of record to decode: can't retrieve nested DataInput length")
			}
			//fmt.Printf("decoding sub slice: counter=%d; numberOfItems=%d\n", counter, subL)
			i += 2
			subDI, delta, err := decodeR(bDI[i:], int(subDILen))
			if err != nil {
				return rsltDI, i, err
			}
			rsltDI = append(rsltDI, subDI)
			i += delta
			counter++
		default:
			return rsltDI, i, errors.New("can't decode DataInput: record of unsupported type")
		}
	}
	return rsltDI, i, nil
}
