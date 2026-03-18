package privkey

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
)

type RSAPrivateKey struct {
	key *rsa.PrivateKey
}

// 秘密鍵を生成する関数
func New(bits int) (*RSAPrivateKey, error) {
	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	return &RSAPrivateKey{key: key}, nil
}

// 秘密鍵を DER 形式でエンコードする関数
func (r *RSAPrivateKey) Marshal() ([]byte, error) {
	return x509.MarshalPKCS8PrivateKey(r.key)
}

// 秘密鍵に対応する公開鍵を取得する関数
func (r *RSAPrivateKey) Public() *rsa.PublicKey {
	return &r.key.PublicKey
}

// 秘密鍵を用いてハッシュ済データの署名を生成する関数
func (r *RSAPrivateKey) Sign(hashed []byte) ([]byte, error) {
	return rsa.SignPKCS1v15(nil, r.key, crypto.SHA256, hashed)
}
