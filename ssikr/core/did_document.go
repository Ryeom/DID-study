package core

// DID Document : 마찬가지로 표준화 되어있지않음 표준화 과정에 있다.
// Decentralized Identifiers (DIDs) v1.0 https://www.w3.org/TR/did-core/
// Simple Example: https://www.w3.org/TR/did-core/#a-simple-example

// verification method
// id 키에대한 아이디 유일값, 타입은 키타입, 콘트롤러는 키를 콘트롤하는거, 퍼블릭키멀티베이스(비트코인형식)스트링으로 되어있음 에검증할때 그대로 퍼블릭키로 복원한다음 사용한다ㅣ
// context는 여러개가들어간다아
// github.com/comnics/did-example/core/did_document.go

// producer : jwt나 json 등 다른 타입으로 변경해서 다른곳으로 보내줄때 이렇게 애기한다
// vdr : public key를 추출하기 위한 저장소
// 저장하는 서비스를 레지스트라 라고 부르고이따

import "encoding/json"

const (
	VERIFICATION_KEY_TYPE_SECP256K1 = "EcdsaSecp256k1VerificationKey2019"
	VERIFICATION_KEY_TYPE_JWS       = "JsonWebKey2020"
	VERIFICATION_KEY_TYPE_X25519    = "X25519KeyAgreementKey2019"
	VERIFICATION_KEY_TYPE_ED25519   = "Ed25519VerificationKey2018"

	PROOF_TYPE_ED25519 = "Ed25519Signature2018"
	PROOF_TYPE_JWS     = "JsonWebSignature2020"
)

type didDocumentInterface interface {
	produce(doc DIDDocument) string
	consume(str string) DIDDocument
}

// DID Document
// Decentralized Identifiers (DIDs) v1.0 https://www.w3.org/TR/did-core/
// Simple Example: https://www.w3.org/TR/did-core/#a-simple-example
type DIDDocument struct {
	// @context
	// Mendatory
	Context []string `json:"@context"`

	Id                   string               `json:"id"`
	AlsoKnownAs          []string             `json:"alsoKnownAs,omitempty"`
	Controller           string               `json:"controller,omitempty"`
	VerificationMethod   []VerificationMethod `json:"verificationMethod,omitempty"`
	Authentication       []Authentication     `json:"authentication,omitempty"`
	AssertionMethod      string               `json:"assertionMethod,omitempty"`
	KeyAgreement         string               `json:"keyAgreement,omitempty"`
	CapabilityInvocation string               `json:"capabilityInvocation,omitempty"`
	CapabilityDelegation string               `json:"capabilityDelegation,omitempty"`
	Service              []Service            `json:"service,omitempty"`
}

type VerificationMethod struct {
	Id                 string `json:"id"`
	Type               string `json:"type"`
	Controller         string `json:"controller"`
	PublicKeyMultibase string `json:"PublicKeyMultibase,omitempty"`
}

type Authentication struct {
	Id                 string `json:"id"`
	Type               string `json:"type"`
	Controller         string `json:"controller"`
	PublicKeyBase58    string `json:"publicKeyBase58,omitempty"`
	PublicKeyMultibase string `json:"PublicKeyMultibase,omitempty"`
}

type Service struct {
	Id              string `json:"id"`
	Type            string `json:"type"`
	ServiceEndpoint string `json:"serviceEndpoint"`
}

func NewDIDDocument(did string, verificationMethod []VerificationMethod) (doc *DIDDocument) {
	var docTmp = new(DIDDocument)
	docTmp.Context = []string{"https://www.w3.org/ns/did/v1"}
	docTmp.Id = did
	docTmp.VerificationMethod = verificationMethod

	return docTmp
}

func NewDIDDocumentForString(docStr string) (didDoc *DIDDocument, err error) {
	didDoc = new(DIDDocument)
	e := json.Unmarshal([]byte(docStr), didDoc)
	if e != nil {
		return nil, e
	}

	return didDoc, nil
}

////////////////////////////////////////////////
// methods define
////////////////////////////////////////////////

func (doc *DIDDocument) Produce() string {
	str, err := json.Marshal(doc)
	if err != nil {
		// fmt.Println("Fail to marshal: ", err)
		return ""
	}
	return string(str)
}

func (doc *DIDDocument) Consume(str string) {
	err := json.Unmarshal([]byte(str), doc)
	if err != nil {

	}
}

func (doc *DIDDocument) GetVerificationMethod() []VerificationMethod {
	return doc.VerificationMethod
}

func (doc *DIDDocument) AddVerificationMethod(id string, typ string, controller string, publicKeyMultibase string) {
	newVm := VerificationMethod{
		Id:                 id,
		Type:               typ,
		Controller:         controller,
		PublicKeyMultibase: publicKeyMultibase,
	}

	doc.VerificationMethod = append(doc.VerificationMethod, newVm)
}

func (doc *DIDDocument) String() string {
	return doc.Produce()
}
