package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"io"
	"log"
	"net"
	"os"
)

var (
	privPEM = []byte(`-----BEGIN PRIVATE KEY-----
MIIEpAIBAAKCAQEA3yiEanpybB6EULjTePPZqqvf6g6o6OgbLIirix1Giu9hKpid
CQsd0js0G4SSGrjd5dz23YdK1pAydpfDRGM7aK0VxKJPQoVCY5nHyr7ZeqDDQw6l
FRWjZxHZ+jZvd+zz/pRy/yaql/c6I/CPh0jyobH7mmF1gwkP2smuhzSCXRg2QcKE
2o0PWjLS2PxxJGSHBpmBMCXe+dxlUjpu3SEBBjePCnXXEi4G9u0VIQxnp/Yl4ZJO
xZ80fy6GAa/tgPiNdpxrnOzNcTojhIO5gPHAlbvj5mw5tmABJ7Ye1blCrB9PnA04
CEqfCfy8fIYeBEKe5RQr843Mh5PPYPkDsbuLAwIDAQABAoIBAHNfxE40rlG92VTO
qe7fzAqUP/kHyUZZMittIZuT8DPGnWrjalURnUJi/4a2nL8aEwdTnYJc/1E9TgnI
XtlNoJp22klGTUosEW3jRMtKrq/ay/kwaiMjg863CQZ2/Fx5cpCNeYL3H73fwPxx
1BLyoRb+KQHhH1s25S6NlSEsAbZU5hPHXP6u5eub3gM3iqbu6urwUOktwH0tt+HZ
VolOdnjm78ypMQ9HdHOqa0WL86BUqI0e80cFPwhvTTNsDSS3ZeTeJLzf5Oy7tp1E
koxeRahNB3BRGHut+B62qZob4EqYmAvHqkD+1AKXhFraJtHH8vwNopFHT9bjZ65S
mEOdPHkCgYEA66E7vIsn0kELqtIJfR4143QucJfBMXebHsOKR9ujg/VfCPst9hEd
dtp73WkivxujehthnSkdSJQbLKbaSfPR5Ow8VJrvlt34H9a/uXIUmZgq9yaBi5zG
fjCMgxRaoCG0yP2gWHz8cmiGoMwWevEoXZhI55gKv5ahL5fdDD/IGS0CgYEA8nNG
GEDri/FUegCyT2Ndxz38F5KfnMv4uHVeRX6diF+djGL43dW7A9n+TtDIkHtli5js
BQfnWkCZrfCSxOr2Y8jGdBzrXHvJM2I/dQBMY/NmHijRzyYpL9IRLeyPcomfqICF
7Ty0r6VbINEIqrLTYSk/qTceUDh7BIgQCDtOcu8CgYArXvce4kJHKh/apmSGuivT
HQx7PwOZdlmAFR/70ArN/Dks7wbrtwTEXrzT6UkzAgRaMnYKNookVNaXpnKhkBZ/
W5hTPl2BWIQWYDHUEZKHHwIxkc3gg8/pZEhIzFNODEY5hK2h/Ad/i4vURxyQLplu
eNxafJrl3vT2TK6sVYUgVQKBgQDZo1zSvMQ09UfN1P47gYlXJ286geZBzF/SxZkA
bS5gkuRMdzPLfubICFHe0lCYUgzPiVClG3k0bTUHDSPTHNBctohfBu7IMF5mf9VS
5BeuyXlrrVzZxPnH8Zx2SxphyOTHT8fpNEtGOPtatApBoRFa8Loy3kWZ4Xmckb9C
hGphhQKBgQDrMwTxmkuyIDoEQC4pmhwQYkJfDkkOLJlwOo5pM6np3NaDCK3Qo9Ct
W56qEwfi9FWBJD21SzRyehqJXtzgxJ11P3TxHxRVaDXjrdrvs81Wnx203lbmMHze
zfl1i78vRRARvUvPgVPUq9m8jbCkIcrsvCW9rHuaH5K8NnjLOXNuFA==
-----END PRIVATE KEY-----`)

	certPEM = []byte(`-----BEGIN CERTIFICATE-----
MIIDUTCCAjmgAwIBAgIIHk5UJxhdzucwDQYJKoZIhvcNAQELBQAwPDELMAkGA1UE
BhMCQ0ExGTAXBgNVBAoTEEFkcmlhbm9TZWxhIEluYy4xEjAQBgNVBAMTCWxvY2Fs
aG9zdDAeFw0yMzEwMDgxNjI5NDBaFw0zMzEwMDUxNjM0NDBaMDwxCzAJBgNVBAYT
AkNBMRkwFwYDVQQKExBBZHJpYW5vU2VsYSBJbmMuMRIwEAYDVQQDEwlsb2NhbGhv
c3QwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDfKIRqenJsHoRQuNN4
89mqq9/qDqjo6BssiKuLHUaK72EqmJ0JCx3SOzQbhJIauN3l3Pbdh0rWkDJ2l8NE
YztorRXEok9ChUJjmcfKvtl6oMNDDqUVFaNnEdn6Nm937PP+lHL/JqqX9zoj8I+H
SPKhsfuaYXWDCQ/aya6HNIJdGDZBwoTajQ9aMtLY/HEkZIcGmYEwJd753GVSOm7d
IQEGN48KddcSLgb27RUhDGen9iXhkk7FnzR/LoYBr+2A+I12nGuc7M1xOiOEg7mA
8cCVu+PmbDm2YAEnth7VuUKsH0+cDTgISp8J/Lx8hh4EQp7lFCvzjcyHk89g+QOx
u4sDAgMBAAGjVzBVMA4GA1UdDwEB/wQEAwICpDATBgNVHSUEDDAKBggrBgEFBQcD
ATAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBSXOUWHZ8i3Dhcl8vsM77QUpWWJ
vzANBgkqhkiG9w0BAQsFAAOCAQEAywpCM+t8gmH4Dcwws1DAiKS3I/9YigOZRV5l
9GMTH3rYAYHJHOIubOxikvX/BftDKoH0jABoSVbswcIr1NVv+YmedFHZoLCAk2mb
IE+PrfnIcymOXqvxJhviV0hZs0NbNqOVIu/BikNAGNOZcf37JwUrvnr7fQaNgV92
W63bnFZbBgaMH38emLvjDlb0VXWrm98BLWahyNIGsuPKN+KPhbJ7pmvvSLPU33x6
m/LkxjKMuEIjrKRLlpVxmQP5X+YLhhXVVsGPnvnpqykbcRsUqoobKSzf6HVwL2zm
qZbJV1TRCICEFPS5d5c/L7BuD05bCejpkrxbzjvFhwKOnwq2jQ==
-----END CERTIFICATE-----`)
)

func main() {
	trustedCertificate := os.Args[1]
	pemBytes, err := os.ReadFile(trustedCertificate)
	if err != nil {
		log.Fatalf("failed to read trusted cert file: %v", err)
	}
	der, _ := pem.Decode(pemBytes)
	cert, err := x509.ParseCertificate(der.Bytes)
	if err != nil {
		log.Fatalf("failed to parse certificate: %v", err)
	}
	cp := x509.NewCertPool()
	cp.AddCert(cert)

	serverCert, err := tls.X509KeyPair(certPEM, privPEM)
	if err != nil {
		log.Fatalf("failed to load server tls cert: %v", err)
	}

	listener, err := tls.Listen("tcp", ":2424", &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		MinVersion:   tls.VersionTLS12,
		ClientCAs:    cp,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	})
	if err != nil {
		log.Fatalf("failed to start TLS listener: %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("failed to accept connection: %v", err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			log.Fatalf("error reading from connection: %v", err)
		}
	}()

	_, err := io.Copy(conn, os.Stdin)
	if err != nil {
		log.Fatalf("error writing to connection: %v", err)
	}
}
