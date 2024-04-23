package diproto

import (
	"reflect"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	ref := DataInput{
		int32(3498579),
		"laijkfhvoiovjqpiowerfnb p9w oiasusfgipouip",
		DataInput{int32(12), int32(13), int32(14), "fifteen"}}
	refEncoded, _ := Encode(ref)
	res, err := Decode(refEncoded)
	if !reflect.DeepEqual(res, ref) {
		t.Errorf("Decode(Encode(ref)) not equal to ref, want: equal")
	}
	if err != nil {
		t.Errorf("decodeStr(encodeStr(ref)) returned error: %v", err)
	}
}

func TestEncodeDecodeEmpty(t *testing.T) {
	ref := DataInput{}
	refEncoded, _ := Encode(ref)
	res, err := Decode(refEncoded)
	if reflect.DeepEqual(res, ref) {
		t.Errorf("Decode(Encode(ref)) not equal to ref, want: equal")
	}
	if err != nil {
		t.Errorf("decodeStr(encodeStr(ref)) returned error: %v", err)
	}
}

func TestEncodeDecodeEmptyNested(t *testing.T) {
	ref := DataInput{DataInput{}}
	refEncoded, _ := Encode(ref)
	res, err := Decode(refEncoded)
	if reflect.DeepEqual(res, ref) {
		t.Errorf("Decode(Encode(ref)) not equal to ref, want: equal")
	}
	if err != nil {
		t.Errorf("decodeStr(encodeStr(ref)) returned error: %v", err)
	}
}

func TestEncodeWrongType(t *testing.T) {
	ref := DataInput{
		int32(3498579),
		"laijkfhvoiovjqpiowerfnb p9w oiasusfgipouip",
		DataInput{int16(12), int32(13), int32(14), "fifteen"}}
	_, err := Encode(ref)
	if err == nil {
		t.Errorf("decodeStr(encodeStr(ref)) with wrong type input returned nil, want: error")
	}
}

func TestDecodeWrongFormat1(t *testing.T) {
	_, err := Decode([]byte{254, 0, 1, 34})
	if err == nil {
		t.Errorf("decodeStr(encodeStr(ref)) with wrong type input returned nil, want: error")
	}
}

func TestDecodeWrongFormat2(t *testing.T) {
	_, err := Decode([]byte{255, 1})
	if err == nil {
		t.Errorf("decodeStr(encodeStr(ref)) with wrong type input returned nil, want: error")
	}
}

func TestDecodeWrongType(t *testing.T) {
	_, err := Decode([]byte{255, 0, 1, 5, 0, 0, 0, 1})
	if err == nil {
		t.Errorf("decodeStr(encodeStr(ref)) with wrong type input returned nil, want: error")
	}
}

func TestDecodeMalformed(t *testing.T) {
	_, err := Decode([]byte{255, 0, 1, 1, 0, 0, 0, 25, 127, 128, 129})
	if err == nil {
		t.Errorf("decodeStr(encodeStr(ref)) with wrong type input returned nil, want: error")
	}
}
