package diproto

import (
	"math"
	"testing"
)

func TestInt32ToByteMax(t *testing.T) {
	var ref int32 = math.MaxInt32
	res, err := bytesToInt32(int32ToBytes(ref))
	if res != ref {
		t.Errorf("bytesToInt32(int32ToBytes(%d) result: %d, want: %d", ref, res, ref)
	}
	if err != nil {
		t.Errorf("bytesToInt32(int32ToBytes(%d) returned error: %v", ref, err)
	}
}

func TestInt32ToByteMin(t *testing.T) {
	var ref int32 = math.MinInt32
	res, err := bytesToInt32(int32ToBytes(ref))
	if res != ref {
		t.Errorf("bytesToInt32(int32ToBytes(%d) result: %d, want: %d", ref, res, ref)
	}
	if err != nil {
		t.Errorf("bytesToInt32(int32ToBytes(%d) returned error: %v", ref, err)
	}
}

func TestInt32ToBytePositive(t *testing.T) {
	var ref int32 = 204573
	res, err := bytesToInt32(int32ToBytes(ref))
	if res != ref {
		t.Errorf("bytesToInt32(int32ToBytes(%d) result: %d, want: %d", ref, res, ref)
	}
	if err != nil {
		t.Errorf("bytesToInt32(int32ToBytes(%d) returned error: %v", ref, err)
	}
}

func TestInt32ToByteNegative(t *testing.T) {
	var ref int32 = -1287
	res, err := bytesToInt32(int32ToBytes(ref))
	if res != ref {
		t.Errorf("bytesToInt32(int32ToBytes(%d) result: %d, want: %d", ref, res, ref)
	}
	if err != nil {
		t.Errorf("bytesToInt32(int32ToBytes(%d) returned error: %v", ref, err)
	}
}

func TestInt32ToByteZero(t *testing.T) {
	var ref int32 = 0
	res, err := bytesToInt32(int32ToBytes(ref))
	if res != ref {
		t.Errorf("bytesToInt32(int32ToBytes(%d) result: %d, want: %d", ref, res, ref)
	}
	if err != nil {
		t.Errorf("bytesToInt32(int32ToBytes(%d) returned error: %v", ref, err)
	}
}

func TestInt32ToByteDecodeWrongInput(t *testing.T) {
	_, err := bytesToInt32([]byte{})
	if err == nil {
		t.Errorf("bytesToInt32() with 0 byte input returned: nil, want: error")
	}
	_, err = bytesToInt32([]byte{0, 2, 3})
	if err == nil {
		t.Errorf("bytesToInt32() with 3 byte input returned: nil, want: error")
	}
	_, err = bytesToInt16([]byte{0, 1, 2, 3, 5})
	if err == nil {
		t.Errorf("bytesToInt32() with 5 byte returned: nil, want: error")
	}
}

func TestInt16ToByteMax(t *testing.T) {
	var ref int16 = math.MaxInt16
	res, err := bytesToInt16(int16ToBytes(ref))
	if res != ref {
		t.Errorf("bytesToInt32(int32ToBytes(%d) result: %d, want: %d", ref, res, ref)
	}
	if err != nil {
		t.Errorf("bytesToInt32(int32ToBytes(%d) returned error: %v", ref, err)
	}
}

func TestInt16ToByteMin(t *testing.T) {
	var ref int16 = math.MinInt16
	res, err := bytesToInt16(int16ToBytes(ref))
	if res != ref {
		t.Errorf("bytesToInt32(int32ToBytes(%d) result: %d, want: %d", ref, res, ref)
	}
	if err != nil {
		t.Errorf("bytesToInt32(int32ToBytes(%d) returned error: %v", ref, err)
	}
}

func TestInt16ToBytePositive(t *testing.T) {
	var ref int16 = 1001
	res, err := bytesToInt16(int16ToBytes(ref))
	if res != ref {
		t.Errorf("bytesToInt32(int32ToBytes(%d) result: %d, want: %d", ref, res, ref)
	}
	if err != nil {
		t.Errorf("bytesToInt32(int32ToBytes(%d) returned error: %v", ref, err)
	}
}

func TestInt16ToByteNegative(t *testing.T) {
	var ref int16 = -13985
	res, err := bytesToInt16(int16ToBytes(ref))
	if res != ref {
		t.Errorf("bytesToInt32(int32ToBytes(%d) result: %d, want: %d", ref, res, ref)
	}
	if err != nil {
		t.Errorf("bytesToInt32(int32ToBytes(%d) returned error: %v", ref, err)
	}
}

func TestInt16ToByteZero(t *testing.T) {
	var ref int16 = -13985
	res, err := bytesToInt16(int16ToBytes(ref))
	if res != ref {
		t.Errorf("bytesToInt32(int32ToBytes(%d) result: %d, want: %d", ref, res, ref)
	}
	if err != nil {
		t.Errorf("bytesToInt32(int32ToBytes(%d) returned error: %v", ref, err)
	}
}

func TestInt16ToByteDecodeWrongInput(t *testing.T) {
	_, err := bytesToInt16([]byte{})
	if err == nil {
		t.Errorf("bytesToInt16() with 0 byte input returned: nil, want: error")
	}
	_, err = bytesToInt16([]byte{0})
	if err == nil {
		t.Errorf("bytesToInt16() with 1 byte input returned: nil, want: error")
	}
	_, err = bytesToInt16([]byte{0, 1, 2})
	if err == nil {
		t.Errorf("bytesToInt16() with 3 byte returned: nil, want: error")
	}
}
