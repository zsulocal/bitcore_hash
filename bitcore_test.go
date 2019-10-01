package bitcore_hash

import (
	"encoding/hex"
	"testing"
)

func TestBitcore(t *testing.T) {
	hash := "00000020c8e194c5bea7ee02d91f7862cd03aa7859364463919b88b0ee07d65aaa4919e9ee2eb5d5c11ca1f5c9ec0be6e2a3051191a49afa64df1dbcaa7a0cb567050013fc39ae5997942e1b61a8b39b"
	hash_bin, _ := hex.DecodeString(hash)
	powhash := GetPowHash(hash_bin, 1504590332)
	powhash_hex := hex.EncodeToString(powhash)
	if powhash_hex != "264240b50bf86271babf615cffb644fb3487ef2b4525336ef7be1f0000000000" {
		t.Error("Test bitcore powhash failed")
		return
	}
}
