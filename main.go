package main

import (
	"github.com/willallves/amazon-iid/pkcs7"
	"github.com/willallves/amazon-iid/rsa"
)

func main() {
	rsa.VerifyRSACert()

	pkcs7.VerifyPKCS7Cert()
}
