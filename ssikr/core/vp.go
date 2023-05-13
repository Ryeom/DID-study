package core

import (
	"crypto/ecdsa"
	"crypto/x509"
	"fmt"
	"github.com/getlantern/deepcopy"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/multiformats/go-multibase"
	"log"
	"strings"
	"time"
)

type VP struct {
	Context      []string `json:"@context"`
	Id           string   `json:"id,omitempty"`
	Type         []string `json:"type,omitempty"`
	Issuer       string   `json:"issuer,omitempty"`
	IssuanceDate string   `json:"issuanceDate,omitempty"`

	// jwt의 token형식으로 저장한다.
	VerifiableCredential []string `json:"verifiableCredential"`
	// VP의 Proof 사용.
	Proof                *Proof   `json:"proof,omitempty"`
}

// JWT를 위한 claim
type JwtClaimsForVP struct {
	jwt.StandardClaims

	Nonce string
	Vp    VP `json:"vp,omitempty"`
}
func NewVP(id string, typ []string, issuer string, vcTokens []string) (*VP, error) {
	newVP := &VP{
		Context: []string{
			"https://www.w3.org/2018/credentials/v1",
			"https://www.w3.org/2018/credentials/v2",
		},
		Id:                   id,
		Type:                 typ,
		Issuer:               issuer,
		VerifiableCredential: vcTokens,
	}
	return newVP, nil
}
func (vp *VP) GenerateJWT(verificationId string, pvKey *ecdsa.PrivateKey) string {
	aud := ""
	exp := time.Now().Add(time.Minute * 5).Unix()       //만료 시간. 현재 + 5분
	jti := uuid.NewString()                             // JWT ID
	t, err := time.Parse(time.RFC3339, vp.IssuanceDate) //unixtime으로 바꾸기 위해.
	iat := t.Unix()
	nbf := iat
	iss := vp.Issuer
	sub := "Verifiable Presentation"

	// Proof를 제거하고 JWT를 만들기 위해 복제한다.
	vpTmp := new(VP)
	deepcopy.Copy(vpTmp, vp)
	vpTmp.Proof = nil

	jwtClaims := JwtClaimsForVP{
		jwt.StandardClaims{
			Audience:  aud,
			ExpiresAt: exp,
			Id:        jti,
			IssuedAt:  iat,
			Issuer:    iss,
			NotBefore: nbf,
			Subject:   sub,
		},
		"qwasd!234",
		*vpTmp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwtClaims)
	token.Header["kid"] = verificationId

	tokenString, err := token.SignedString(pvKey)

	if err != nil {

	}

	return tokenString
}
func (vp *VP) VerifyJwt(tokenString string, pbKey *ecdsa.PublicKey) (bool, error) {

	parts := strings.Split(tokenString, ".")
	err := jwt.SigningMethodES256.Verify(strings.Join(parts[0:2], "."), parts[2], pbKey)
	if err != nil {
		return false, nil
	}

	//개별적으로 내부 VC들을 다시 다 검증해야 한다.
	parseToken, err := jwt.ParseWithClaims(tokenString, &JwtClaimsForVP{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			log.Fatalln("unexpected signing method.")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		//did := token.Header["kid"].(string)
		//pbKeyBase58 := did // getPbKey(did, "") //DID를 통해 DID-Document의 pbKey를 구한다.
		//pbKey, _ := x509.ParsePKIXPublicKey(base58.Decode(pbKeyBase58))

		return pbKey, nil
	})

	fmt.Println("parseToken: ", parseToken)
	claims, ok := parseToken.Claims.(*JwtClaimsForVP)
	fmt.Println("claims: ", claims)

	if ok && parseToken.Valid {

		if &claims.Vp != nil {
			vpMapClaims := claims.Vp
			vcList := vpMapClaims.VerifiableCredential

			for _, vcToken := range vcList {
				fmt.Println("VC: ", vcToken)
				//verify jwt
				//VC.VerifyJwt(vcToken)

			}
		} else {
			return false, fmt.Errorf("VC is not exist.")
		}

	} else {
		return false, fmt.Errorf("VP is not valid.")
	}

	return true, nil
}

func ParseAndVerifyJwtForVP(tokenString string) (bool, *JwtClaimsForVP, error) {
	//개별적으로 내부 VC들을 다시 다 검증해야 한다.
	parseToken, _ := jwt.ParseWithClaims(tokenString, &JwtClaimsForVP{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			log.Fatalln("unexpected signing method.")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		//DID를 추출한다.
		parts := strings.Split(token.Header["kid"].(string), "#")
		did := parts[0]

		//Resolve한다.
		didDocumentStr, err := ResolveDid(did)
		if err != nil {
			log.Printf("Failed to Resolve DID.\nError: %x\n", err)
		}

		//Json string을 DID Document 객체로 생성한다.
		didDocument, err := NewDIDDocumentForString(didDocumentStr)
		if err != nil {
			log.Printf("Failed generate DID Document from string.\nError: %x\n", err)
		}
		// 첫 번째를 사용한다고 가정한다.
		// TODO: 키 ID(위의 kid)에 해당하는 키 값 구하기.
		pbKeyBaseMultibase := didDocument.VerificationMethod[0].PublicKeyMultibase
		_, bytePubKey, err := multibase.Decode(pbKeyBaseMultibase)
		pbKey, err := x509.ParsePKIXPublicKey(bytePubKey)

		return pbKey, nil
	})

	fmt.Println("parseToken: ", parseToken)
	claims, ok := parseToken.Claims.(*JwtClaimsForVP)
	fmt.Println("claims: ", claims)

	if ok && parseToken.Valid {
		if &claims.Vp != nil {
			vpMapClaims := claims.Vp
			vcList := vpMapClaims.VerifiableCredential

			for idx, vcToken := range vcList {
				fmt.Printf("VC[%d]: %s", idx, vcToken)
				verify, _, err := ParseAndVerifyJwtForVC(vcToken)
				if !verify || err != nil {
					log.Printf("Failed to verify VC[%d] in VP.", idx)
					// VC가 한 건이라도 오류라면 바로 리턴해 버린다.
					return false, nil, err
				}
				fmt.Printf("	==> VC[%d] is Verified.\n", idx)
			}
		} else {
			return false, nil, fmt.Errorf("VC is not exist.")
		}

	} else {
		return false, nil, fmt.Errorf("VP is not valid.")
	}

	return true, claims, nil
}
