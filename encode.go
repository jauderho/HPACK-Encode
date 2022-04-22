package main

import (
	"fmt"
	"golang.org/x/net/http2/hpack"
	"os"
	"strconv"
)

func main() {
	fmt.Println(Encode(os.Args[1]))
}

func Encode(s string) string {
	var result string
	var count int

	hd := hpack.AppendHuffmanString(nil, s)
	hl := hpack.HuffmanEncodeLength(s) | 0x80

	result += RenderByte(byte(hl))

	for _, b := range hd {
		result += RenderByte(b)
		count += 1
	}

	return "static const u_char nginx[" + strconv.Itoa(count+1) + "] = \"" + string(result) + "\";"
}

func RenderByte(b byte) string {
	return fmt.Sprintf("\\x%x", b)
}
