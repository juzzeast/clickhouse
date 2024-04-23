package main

import (
	"fmt"

	"github.com/juzzeast/clickhouse/diproto"
)

func main() {
	di := diproto.DataInput{
		int32(3498579),
		"laijkfhvoiovjqpiowerfnb p9w oiasusfgipouip",
		diproto.DataInput{
			int32(12), int32(13), int32(14), "fifteen"}}
	fmt.Printf("Data Input: %v \n", di)
	encoded, _ := diproto.Encode(di)
	fmt.Printf("Encoded Data Input: %v \n", encoded)

	decoded, err := diproto.Decode(encoded)
	if err != nil {
		fmt.Printf("Not able to decode DataInput: %v \n", err.Error())
	}
	fmt.Printf("Decoded Data Input: %v \n", decoded)
}
