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

Start HTTPS server:

```sh
nginx -c $(pwd)/nginx.conf -p $(pwd)/ -e /dev/stderr
```

Verify:

```sh
curl --cacert test/cert.pem https://localhost:8443
```

Stop server:

```sh
nginx -c $(pwd)/nginx.conf -p $(pwd)/ -e /dev/stderr -s stop
```

### CA-Signed Certificate

Generate CA certificate and server certificate:

```sh
go run cmd/main.go certified test/certified
```

Rename for nginx:

```sh
mv test/certified-cert.pem test/cert.pem
mv test/certified-key.pem test/key.pem
```

Start HTTPS server:

```sh
nginx -c $(pwd)/nginx.conf -p $(pwd)/ -e /dev/stderr
```

Verify:

```sh
curl --cacert test/certified-cacert.pem https://localhost:8443
```

Stop server:

```sh
nginx -c $(pwd)/nginx.conf -p $(pwd)/ -e /dev/stderr -s stop
```

## References

- [作って理解する HTTPS 証明書](https://techbookfest.org/product/vu1AyJ6UMHfVzAF0x0aSn0)
