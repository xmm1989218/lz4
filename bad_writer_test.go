package lz4_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/xmm1989218/lz4"
)

func TestBadWriter(t *testing.T) {
	var buf bytes.Buffer
	writer := lz4.NewBadWriter(&buf)
	if _, err := writer.Write([]byte("1234567890")); err != nil {
		t.Fatalf("write fail: %v", err)
		return
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("close fail: %v", err)
		return
	}

	fmt.Printf("%s\n", hex.EncodeToString(buf.Bytes()))
}
