// util/hash.go

package util

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/btcsuite/btcutil/base58"
)

func MakeHash(plain string) []byte {
	digest := sha256.Sum256([]byte(plain))
	return digest[:]
}

func MakeHashBase58(plain string) string {
	return base58.Encode(MakeHash(plain))
}

func MakeHashHex(plain string) string {
	return hex.EncodeToString(MakeHash(plain))
}
// RSA : 보안적위험성이 있어서 ECC를 사용할거당!

// ECDSA : RSA와 동일한 수준의 보안을 제공하지만 길이가 더짧다.
