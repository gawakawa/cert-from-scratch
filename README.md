# cert-from-scratch

X.509 certificate implementation from scratch in Go.

## Usage

### Dummy Certificate Generator

Generate a dummy certificate:

```sh
go run cmd/main.go basecert test/basecert
```

Verify the generated certificate with OpenSSL:

```sh
openssl x509 -in test/basecert.pem -inform PEM -noout -text
```
