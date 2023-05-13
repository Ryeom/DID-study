package main

import (
	"fmt"
	"ssikr/core"
)

func main() {
	method := "ssikr"
	// 선행조건 : 키를 만든다.
	kms := new(core.ECDSAManager)
	kms.Generate()

	did, err := core.NewDID(method, kms.PublicKeyMultibase())
	if err != nil {

	}
	// did document 생성
	// kms의 pb기준으로 did를 만든다
	verificationId := fmt.Sprintf("%s#keys-1", did)
	verificationMethod := []core.VerificationMethod{
		{Id: verificationId,
			Type:               "EcdsaSecp256k1VerificationKey2019",
			Controller:         did.String(),
			PublicKeyMultibase: kms.PublicKeyMultibase(),
	// 키를 베리피케이션 메소드나 어솔리티에서 참조를 한다.
		},
	}

	didDocument := core.NewDidDocument(did.String(),verificationMethod)

	fmt.Printf("[DID]%s\n",did)
	fmt.Printf("[DID Document]%+v\n",didDocument)

	/*
이 예제는 키를 계속 생성하게 된다.
	[DID]did:ssikr:EZXW78SFYtTcFDJRvE2VJssvnRWg9PhwpjgUBbPficer
	[DID Document]&{Context:[https://www.w3.org/ns/did/v1] Id:did:ssikr:EZXW78SFYtTcFDJRvE2VJssvnRWg9PhwpjgUBbPficer AlsoKnownAs:[] Controller: VerificationMethod:[{Id:did:ssikr:EZXW78SFYtTcFDJRvE2VJssvnRWg9PhwpjgUBbPficer#keys-1 Type:EcdsaSecp256k1VerificationKey2019 Controller:did:ssikr:EZXW78SFYtTcFDJRvE2VJssvnRWg9PhwpjgUBbPficer PublicKeyMultibase:zaSq9DsNNvGhYxYyqA9wd2eduEAZ5AXWgJTbTGKcwJJdMa93assvD5iJnVTfAv9kVGKeTcniTswHBfispgSmnrv9BGTCMJenmLX45bED6bGXpg4Nk8JRTBQxoz7bD}] Authentication:[] AssertionMethod: KeyAgreement: CapabilityInvocation: CapabilityDelegation: Service:[]}

	*/




	//
}
