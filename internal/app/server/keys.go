package server

import (
	"bytes"
)

// Keys This is only for POC. In a production setting, this should be processed by a KDF before being used in a higher level protocol
type Keys struct {
	KeyA []byte
	KeyB []byte
}

func (k *Keys) keysAreEqual() bool {
	return bytes.Equal(k.KeyA, k.KeyB)
}
