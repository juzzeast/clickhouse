package diproto

import "testing"

func TestEncodeStrAscii(t *testing.T) {
	ref := "lorem ipsum. Lo-rem IP-SUM!"
	res, i, err := decodeStr(encodeStr(ref))
	if res != ref {
		t.Errorf("decodeStr(encodeStr(%s)) result=%s, want=%s", ref, res, ref)
	}
	wantDelta := len([]byte(ref)) + 5
	if i != wantDelta {
		t.Errorf("decodeStr(encodeStr(%s)) delta=%d, want=%d", ref, i, wantDelta)
	}
	if err != nil {
		t.Errorf("decodeStr(encodeStr(%s)) returned error: %v", ref, err)
	}
}

func TestEncodeStrUTF8(t *testing.T) {
	ref := "lorem ipsum. and then \xbd\xb2\x3d\xbc\x20\xe2\x8c\x98, Дякую"
	res, i, err := decodeStr(encodeStr(ref))
	if res != ref {
		t.Errorf("decodeStr(encodeStr(%s)) result=%s, want=%s", ref, res, ref)
	}
	wantDelta := len([]byte(ref)) + 5
	if i != wantDelta {
		t.Errorf("decodeStr(encodeStr(%s)) delta=%d, want=%d", ref, i, wantDelta)
	}
	if err != nil {
		t.Errorf("decodeStr(encodeStr(%s)) returned error: %v", ref, err)
	}
}

func TestEncodeStrEmpty(t *testing.T) {
	ref := ""
	res, i, err := decodeStr(encodeStr(ref))
	if res != ref {
		t.Errorf("decodeStr(encodeStr(%s)) result=%s, want=%s", ref, res, ref)
	}
	wantDelta := len([]byte(ref)) + 5
	if i != wantDelta {
		t.Errorf("decodeStr(encodeStr(%s)) delta=%d, want=%d", ref, i, wantDelta)
	}
	if err != nil {
		t.Errorf("decodeStr(encodeStr(%s)) returned error: %v", ref, err)
	}
}

func TestDecodeStrWrongLength(t *testing.T) {
	_, _, err := decodeStr([]byte{})
	if err == nil {
		t.Errorf("decodeStr() with 0 byte input returned nil, want: error")
	}
	_, _, err = decodeStr([]byte{di_type_string, 0, 1, 2})
	if err == nil {
		t.Errorf("decodeStr() with 4 byte input returned nil, want: error")
	}
}

func TestDecodeStrWrongCode(t *testing.T) {
	_, _, err := decodeStr([]byte{22, 0, 0, 0, 2, 128, 129})
	if err == nil {
		t.Errorf("decodeStr() with wrong code input returned nil, want: error")
	}
}

func TestDecodeStrMalformed(t *testing.T) {
	_, _, err := decodeStr([]byte{22, 0, 0, 0, 5, 128, 129, 130})
	if err == nil {
		t.Errorf("decodeStr() with malformed input returned nil, want: error")
	}
}
