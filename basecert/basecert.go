package basecert

import (
	"encoding/asn1"
	"time"
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
