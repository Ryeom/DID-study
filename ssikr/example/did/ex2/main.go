package main

import (
	"errors"
	"fmt"
	"log"
	"ssikr/core"
	"ssikr/util"
)

func NewDID(method string, pbKey string) (string, error) {
	if method == "" || pbKey == "" {
		return "", errors.New("왜 없냐 값이")
	}
	// 요즘 트렌드가 변수이름을 예쁘고 좋게 지어서 주석을 최소화하자~ 라서
	specificIdentifier := util.MakeHashBase58(pbKey)
	// public key를 기반으로 specificIdentifier를 hash로 만들어서 쓰기도함
	// 앞에 메소드가 없어도 유일한 값이다.

	did := fmt.Sprintf("did:%s:%s", method, specificIdentifier)

	return did, nil
}

func main() {
	//kms : key management system : 원래는 이거를 직접 만들어 줘야함.

	var method = "ssikr"

	kms := new(core.ECDSAManager)
	kms.Generate()

	did, err := NewDID(method, kms.PublicKeyMultibase())

	if err != nil {
		log.Printf("Failed to generate DID, error: %v\n", err)
	}

	fmt.Println("### New DID ###")
	fmt.Printf("did => %s\n", did)

}
