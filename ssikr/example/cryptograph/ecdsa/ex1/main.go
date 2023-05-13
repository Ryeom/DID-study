package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

func main() {
	pvKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	// elliptic.Curve 를 뭘로 쓸꺼냥 elliptic.p224, p384 p521 ...
	// 보통 privatekey를 알고있으면 public key를 생성할수이씀 (그 속에 포함)
	if err != nil {
		fmt.Println("generage key err")
	}

	pbKey := &pvKey.PublicKey
fmt.Println("######################################")
	fmt.Printf("[ privatekey D ] %x ",pvKey.D )
	fmt.Printf("[ public key X ] %s ",pbKey.X )
	fmt.Printf("[ public key Y ] %s "  ,pbKey.Y )
	fmt.Printf("[ privatekey byte ] %x ",pbKey.X.Bytes() )
	fmt.Printf("[ public key byte ] %x  ",pbKey.Y.Bytes() )
}
