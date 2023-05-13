package main

import (
	"fmt"
	"log"
	"ssikr/core"
)

// examples/vdr/ex_vdr.go
func main() {
	var method = "ssikr"

	kms := new(core.ECDSAManager)
	kms.Generate()
	//키쌍 만들고


	// 만든 키쌍의 퍼블릭키를 사용해서 did를 만둔다
	did, err := core.NewDID(method, kms.PublicKeyMultibase())

	if err != nil {
		log.Printf("Failed to generate DID, error: %v\n", err)
	}

	// DID Document 생성.
	verificationId := fmt.Sprintf("%s#keys-1", did)
	verificationMethod := []core.VerificationMethod{
		{
			Id:                 verificationId,
			Type:               core.VERIFICATION_KEY_TYPE_SECP256K1,
			Controller:         did.String(),
			PublicKeyMultibase: kms.PublicKeyMultibase(), // 인코딩형식을 멀티베이스로 해따
		},
	}
	didDocument := core.NewDIDDocument(did.String(), verificationMethod) // 디아듸 도큐먼트를 만들어따

	fmt.Println("### New DID ###")
	fmt.Printf("did => %s\n", did)
	fmt.Printf("did document => %+v\n", didDocument)

	RegisterDid(did.String(), didDocument) // 저장해라 grpc 클라이언트 만들어서

	//Resolve한다.
	didDocumentStr, err := core.ResolveDid(did.String()) // 이걸로 리졸버 서버에 접속해서 did document를 가져와라
	if err != nil {
		log.Printf("Failed to Resolve DID.\nError: %x\n", err)
	}

	fmt.Printf("did document =3 %+v\n", didDocumentStr)
}

func RegisterDid(did string, document *core.DIDDocument) error {
	err := core.RegisterDid(did, document.String())
	if err != nil {
		return err
	}
	return nil
}

/*
### New DID ###
did => did:ssikr:EWmur3VuL8tNPXxgYAJAR7YNL5Ri1kAHNx6vYpf4G7Cy
did document => {"@context":["https://www.w3.org/ns/did/v1"],"id":"did:ssikr:EWmur3VuL8tNPXxgYAJAR7YNL5Ri1kAHNx6vYpf4G7Cy","verificationMethod":[{"id":"did:ssikr:EWmur3VuL8tNPXxgYAJAR7YNL5Ri1kAHNx6vYpf4G7Cy#keys-1","type":"EcdsaSecp256k1VerificationKey2019","controller":"did:ssikr:EWmur3VuL8tNPXxgYAJAR7YNL5Ri1kAHNx6vYpf4G7Cy","PublicKeyMultibase":"zaSq9DsNNvGhYxYyqA9wd2eduEAZ5AXWgJTbTK9seHcD86CZR78bRcL3YKmcMSQduxZBAbgsknXvc1zAxWLbQ357T7EsmhL5dqePeSizdnex7B9cLPfnd3echgvRy"}]}
Registrar Response: result:"OK"
Result: didDocument:"{\"@context\":[\"https://www.w3.org/ns/did/v1\"],\"id\":\"did:ssikr:EWmur3VuL8tNPXxgYAJAR7YNL5Ri1kAHNx6vYpf4G7Cy\",\"verificationMethod\":[{\"id\":\"did:ssikr:EWmur3VuL8tNPXxgYAJAR7YNL5Ri1kAHNx6vYpf4G7Cy#keys-1\",\"type\":\"EcdsaSecp256k1VerificationKey2019\",\"controller\":\"did:ssikr:EWmur3VuL8tNPXxgYAJAR7YNL5Ri1kAHNx6vYpf4G7Cy\",\"PublicKeyMultibase\":\"zaSq9DsNNvGhYxYyqA9wd2eduEAZ5AXWgJTbTK9seHcD86CZR78bRcL3YKmcMSQduxZBAbgsknXvc1zAxWLbQ357T7EsmhL5dqePeSizdnex7B9cLPfnd3echgvRy\"}]}"
did document =3 {"@context":["https://www.w3.org/ns/did/v1"],"id":"did:ssikr:EWmur3VuL8tNPXxgYAJAR7YNL5Ri1kAHNx6vYpf4G7Cy","verificationMethod":[{"id":"did:ssikr:EWmur3VuL8tNPXxgYAJAR7YNL5Ri1kAHNx6vYpf4G7Cy#keys-1","type":"EcdsaSecp256k1VerificationKey2019","controller":"did:ssikr:EWmur3VuL8tNPXxgYAJAR7YNL5Ri1kAHNx6vYpf4G7Cy","PublicKeyMultibase":"zaSq9DsNNvGhYxYyqA9wd2eduEAZ5AXWgJTbTK9seHcD86CZR78bRcL3YKmcMSQduxZBAbgsknXvc1zAxWLbQ357T7EsmhL5dqePeSizdnex7B9cLPfnd3echgvRy"}]}


*/


//jwt [header, payload, signature ]
//ex] xxxxxxxx.yyyyyyyyy.zzzzzzzzz
// .이게 한개면 서명이 없는거
// .. 두개면 서명까지 다있는겅
//시그니처가 검증이 안되면 앞의 데이터는 쓰면 안된다.
// h
// p  vc를 담기도하고 claim을 담기도 함
// s
