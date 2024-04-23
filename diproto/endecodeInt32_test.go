package diproto

import "testing"

func TestEncodeInt32Positive(t *testing.T) {
	var ref int32 = 1387456
	res, i, err := decodeInt32(encodeInt32(ref))
	if res != ref {
		t.Errorf("decodeInt32(encodeInt32(%d)) result=%d, want=%d", ref, res, ref)
	}
	wantDelta := 5
	if i != wantDelta {
		t.Errorf("decodeInt32(encodeInt32(%d)) delta=%d, want=%d", ref, i, wantDelta)
	}
	if err != nil {
		t.Errorf("decodeInt32(encodeInt32(%d)) returned error: %v", ref, err)
	}
}

func TestEncodeInt32Negative(t *testing.T) {
	var ref int32 = -1387456
	res, i, err := decodeInt32(encodeInt32(ref))
	if res != ref {
		t.Errorf("decodeInt32(encodeInt32(%d)) result=%d, want=%d", ref, res, ref)
	}
	wantDelta := 5
	if i != wantDelta {
		t.Errorf("decodeInt32(encodeInt32(%d)) delta=%d, want=%d", ref, i, wantDelta)
	}
	if err != nil {
		t.Errorf("decodeInt32(encodeInt32(%d)) returned error: %v", ref, err)
	}
}

func TestEncodeInt32Zero(t *testing.T) {
	var ref int32 = 0
	res, i, err := decodeInt32(encodeInt32(ref))
	if res != ref {
		t.Errorf("decodeInt32(encodeInt32(%d)) result=%d, want=%d", ref, res, ref)
	}
	wantDelta := 5
	if i != wantDelta {
		t.Errorf("decodeInt32(encodeInt32(%d)) delta=%d, want=%d", ref, i, wantDelta)
	}
	if err != nil {
		t.Errorf("decodeInt32(encodeInt32(%d)) returned error: %v", ref, err)
	}
}

func TestDecodeInt32WrongLength(t *testing.T) {
	_, _, err := decodeStr([]byte{})
	if err == nil {
		t.Errorf("decodeStr() with 0 byte input returned nil, want: error")
	}
	_, _, err = decodeStr([]byte{di_type_int32, 0, 1, 2})
	if err == nil {
		t.Errorf("decodeStr() with 4 byte input returned nil, want: error")
	}
	_, _, err = decodeStr([]byte{di_type_int32, 0, 1, 2, 3, 4})
	if err == nil {
		t.Errorf("decodeStr() with 6 byte input returned nil, want: error")
	}
}

func TestDecodeInt32WrongCode(t *testing.T) {
	_, _, err := decodeStr([]byte{23, 0, 0, 0, 2, 128, 129})
	if err == nil {
		t.Errorf("decodeStr() with wrong code input returned nil, want: error")
	}
}
