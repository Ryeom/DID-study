package issuer

import (
	"context"
	"errors"
	"fmt"
	"log"
	"ssikr/core"
	"ssikr/protos"
)

type Server struct {
	protos.UnimplementedSimpleIssuerServer

	Issuer *Issuer
}

type Issuer struct {
	kms         *core.ECDSAManager
	did         *core.DID
	didDocument *core.DIDDocument
}

func (server *Server) IssueSimpleVC(_ context.Context, msg *protos.MsgIssueVC) (*protos.MsgIssueVCResponse, error) {
	log.Printf("IssueSimpleVC MSG: %+v \n", msg)

	response := new(protos.MsgIssueVCResponse)

	vcToken, err := server.Issuer.GenerateSampleVC()
	if err != nil {

	}
	response.Vc = vcToken

	return response, nil
}

func (issuer *Issuer) GenerateDID() {
	// 키생성(ECDSA) - 향후 KMS로 대체.
	issuer.kms = core.NewEcdsa()

	// DID 생성.
	issuerDid, _ := core.NewDID("comnic", issuer.kms.PublicKeyBase58())

	issuer.did = issuerDid

	// DID Document 생성.
	verificationId := fmt.Sprintf("%s#keys-1", issuerDid)
	verificationMethod := []core.VerificationMethod{
		{
			Id:                 verificationId,
			Type:               core.VERIFICATION_KEY_TYPE_SECP256K1,
			Controller:         issuerDid.String(),
			PublicKeyMultibase: issuer.kms.PublicKeyMultibase(),
		},
	}
	didDocument := core.NewDIDDocument(issuerDid.String(), verificationMethod)
	issuer.didDocument = didDocument

	fmt.Printf("검증용 issuer pbKey: %s\n", issuer.didDocument.VerificationMethod[0].PublicKeyMultibase)

	RegisterDid(issuerDid.String(), didDocument)
}

func (issuer *Issuer) GenerateSampleVC() (string, error) {
	// 영지식증명 :거래 상대방에게 어떠한 정보도 제공하지 않은 채, 자신이 해당 정보를 가지고 있다는 사실을 증명하는 것
	// (조금의 이슈있음, 하지만 vc를 따로 빼면 보완가능)
	// 여기는 하나만 만들게 되어있으나 서비스에서는 아토믹 vc를해서 다수의 vc를 넘겨야 한다. 왜냐면
	// 각 요소[이름, 생일, 입학일, 졸업일]를 따로따로 jwt로 만들어서 저장해야 한다.
	// 각요소를 따로따로 vp로 넘겨야 하기때문 (만약에 그렇게 하지않아도 되는 정보일 경우 한데 묶어도 되긴함)
	// VC 생성.
	vc, err := core.NewVC(
		"1234567890",
		[]string{"VerifiableCredential", "AlumniCredential"},
		issuer.did.String(),
		map[string]interface{}{
			"id": "1234567890",
			"name": map[string]interface{}{
				"id": "1234567",
				"name": []map[string]string{
					{
						"value": "Example University",
						"lang":  "en",
					}, {
						"value": "Exemple d'Université",
						"lang":  "fr",
					},
				},
			},
		},
	)

	if err != nil {
		return "", errors.New("Failed creation VC.")
	}

	// VC에 Issuer의 private key로 서명한다.(JWT 사용)
	token := vc.GenerateJWT(issuer.didDocument.VerificationMethod[0].Id, issuer.kms.PrivateKey)

	return token, nil
}

func RegisterDid(did string, document *core.DIDDocument) error {
	err := core.RegisterDid(did, document.String())
	if err != nil {
		return err
	}
	return nil

}

// 이슈어는 자기의 디아이디를 이미,미리 공개하고 있음
