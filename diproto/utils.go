package diproto

import "errors"

// Support functions to convert different int subtypes to a sequence of bytes
// Implemented with bitwise ops
func int32ToBytes(i int32) []byte {
	ui := uint32(i)
	r := make([]byte, 4)
	r[0] = byte((ui & (0b11111111 << 24)) >> 24)
	r[1] = byte((ui & (0b11111111 << 16)) >> 16)
	r[2] = byte((ui & (0b11111111 << 8)) >> 8)
	r[3] = byte(ui & 0b11111111)
	return r
}

func bytesToInt32(bi []byte) (int32, error) {
	if len(bi) != 4 {
		return 0, errors.New("wrong []byte format to decode into int32: must be 4 bytes straight")
	}
	var ur uint32
	ur = ur + uint32(bi[0])<<24
	ur = ur + uint32(bi[1])<<16
	ur = ur + uint32(bi[2])<<8
	ur = ur + uint32(bi[3])
	return int32(ur), nil
}

func int16ToBytes(i int16) []byte {
	ui := uint16(i)
	r := make([]byte, 2)
	r[0] = byte((ui & (0b11111111 << 8)) >> 8)
	r[1] = byte(ui & 0b11111111)
	return r
}

func bytesToInt16(bi []byte) (int16, error) {
	if len(bi) != 2 {
		return 0, errors.New("wrong []byte format to decode into int16: must be 2 bytes straight")
	}
	var ur uint16
	ur = ur + uint16(bi[0])<<8
	ur = ur + uint16(bi[1])
	return int16(ur), nil
}
