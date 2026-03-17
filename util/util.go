package util

import (
	"encoding/base64"
	"os"
)

type Marshaler interface {
	Marshal() ([]byte, error)
}

func SaveAsPEM(
	filename string,
	marshaler Marshaler,
	pemType string,
	perm os.FileMode,
) error {
	data, err := marshaler.Marshal()
	if err != nil {
		return err
	}
	// ファイルを新規作成 ( すでにある内容は空にして切り詰め ) 、書き込み専用で開く
	f, err := os.OpenFile(
		filename,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		perm,
	)
	if err != nil {
		return err
	}
	defer f.Close()

	header := "-----BEGIN " + pemType + "-----\n"
	footer := "-----END " + pemType + "-----\n"
	if _, err := f.WriteString(header); err != nil {
		return err
	}

	enc := base64.StdEncoding.EncodeToString(data)
	// 64 文字ずつファイルに書き込む
	for i := 0; i < len(enc); i += 64 {
		end := min(i+64, len(enc))
		if _, err := f.WriteString(enc[i:end] + "\n"); err != nil {
			return err
		}
	}
	if _, err := f.WriteString(footer); err != nil {
		return err
	}
	return nil
}

// 証明書を PEM 形式で保存する関数
func MarshalAndSaveCert(filename string, marshaler Marshaler) error {
	return SaveAsPEM(filename+".pem", marshaler, "CERTIFICATE", 0644)
}

// 秘密鍵を PEM 形式で保存する関数
func MarshalAndSaveKey(filename string, marshaler Marshaler) error {
	return SaveAsPEM(filename+".pem", marshaler, "PRIVATE KEY", 0600)
}
