package selfsigned

import (
	"crypto/sha256"
	"encoding/asn1"
	"math/big"
	"time"

	"github.com/gawakawa/cert-from-scratch/privkey"
)

var (
	oidRSAEncryption = asn1.ObjectIdentifier{
		1,
		2,
		840,
		113549,
		1,
		1,
		1,
	}
	oidSHA256WithRSAEnc = asn1.ObjectIdentifier{
		1,
		2,
		840,
		113549,
		1,
		1,
		11,
	}
	oidCommonName = asn1.ObjectIdentifier{2, 5, 4, 3}
)

type Certificate struct {
	TBSCertificate     TBSCertificate
	SignatureAlgorithm AlgorithmIdentifier
	SignatureValue     asn1.BitString
}

type TBSCertificate struct {
	Version      int `asn1:"tag:0,explicit"`
	SerialNumber int
	Signature    AlgorithmIdentifier
	Issuer       Name
	Validity     Validity
	Subject      Name
	PublicKey    SubjectPublicKeyInfo
}

type AlgorithmIdentifier struct {
	Algorithm asn1.ObjectIdentifier
}

type Name struct {
	RDNSequence []AttributeTypeAndValue `asn1:"set"`
}

type AttributeTypeAndValue struct {
	Type  asn1.ObjectIdentifier
	Value string
}

type Validity struct {
	NotBefore time.Time `asn1:"utc"`
	NotAfter  time.Time `asn1:"utc"`
}

type SubjectPublicKeyInfo struct {
	Algorithm AlgorithmIdentifier
	PublicKey asn1.BitString
}

type RSAPublicKey struct {
	N *big.Int
	E int
}

func New(key *privkey.RSAPrivateKey) *Certificate {
	signatureAlgorithm := AlgorithmIdentifier{
		Algorithm: oidSHA256WithRSAEnc,
	}

	// CommonName
	name := Name{
		RDNSequence: []AttributeTypeAndValue{
			{
				Type:  oidCommonName,
				Value: "localhost",
			},
		},
	}

	publicKey := RSAPublicKey{
		N: key.Public().N,
		E: key.Public().E,
	}
	encodePublicKey, err := asn1.Marshal(publicKey)
	if err != nil {
		panic("failed to marshal public key: " + err.Error())
	}
	publicKeyBitString := asn1.BitString{
		Bytes:     encodePublicKey,
		BitLength: len(encodePublicKey) * 8,
	}
	subjectPublicKeyInfo := SubjectPublicKeyInfo{
		Algorithm: AlgorithmIdentifier{
			Algorithm: oidRSAEncryption,
		},
		PublicKey: publicKeyBitString,
	}

	tbsCertificate := TBSCertificate{
		Version:      2, // X.509 Ver.3
		SerialNumber: 1, // この値は同一の CA が発行する証明書で一意になっていればよい
		Signature:    signatureAlgorithm,
		Issuer:       name,
		Validity: Validity{
			NotBefore: time.Now().UTC(),
			NotAfter:  time.Now().AddDate(1, 0, 0).UTC(),
		},
		Subject:   name, // Issuer と Subject が同じなのでいわゆるオレオレ証明書
		PublicKey: subjectPublicKeyInfo,
	}

	// 署名
	encodedTBS, err := asn1.Marshal(tbsCertificate)
	if err != nil {
		panic("failed to marshal TBS certificate: " + err.Error())
	}
	hashed := sha256.Sum256(encodedTBS)
	signature, err := key.Sign(hashed[:])
	if err != nil {
		panic("failed to sign TBS certificate: " + err.Error())
	}

	return &Certificate{
		TBSCertificate:     tbsCertificate,
		SignatureAlgorithm: signatureAlgorithm,
		SignatureValue: asn1.BitString{
			Bytes:     signature,
			BitLength: len(signature) * 8,
		},
	}
}

func (b *Certificate) Marshal() ([]byte, error) {
	return asn1.Marshal(*b)
}
