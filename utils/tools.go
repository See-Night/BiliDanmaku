package utils

import (
	"bytes"
	"io"

	"github.com/andybalholm/brotli"
)

func Uint32ToBytes(v uint32) []byte {
	return []byte{
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	}
}

func BytesToUint32(b []byte) uint32 {
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

func Uint16ToBytes(v uint16) []byte {
	return []byte{
		byte(v >> 8),
		byte(v),
	}
}

func BrotliDecompress(data []byte) ([]byte, error) {
	var decompressed []byte
	reader := brotli.NewReader(bytes.NewReader(data))
	decompressed, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return decompressed, nil
}
