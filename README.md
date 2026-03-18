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

### Self-Signed Certificate

Generate certificate and key:

```sh
go run cmd/main.go selfsigned test/selfsigned
```

Rename for nginx:

```sh
mv test/selfsigned-cert.pem test/cert.pem
mv test/selfsigned-key.pem test/key.pem
```

Start HTTPS server and verify:

```sh
sudo nginx -c $(pwd)/nginx.conf
curl -4 --cacert test/cert.pem https://localhost:8443
```
