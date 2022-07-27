package main

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
)

// AWSRSAIIDCert is the RSA public certificate
const AWSRSAIIDCert = `-----BEGIN CERTIFICATE-----
MIIDIjCCAougAwIBAgIJAKnL4UEDMN/FMA0GCSqGSIb3DQEBBQUAMGoxCzAJBgNV
BAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdTZWF0dGxlMRgw
FgYDVQQKEw9BbWF6b24uY29tIEluYy4xGjAYBgNVBAMTEWVjMi5hbWF6b25hd3Mu
Y29tMB4XDTE0MDYwNTE0MjgwMloXDTI0MDYwNTE0MjgwMlowajELMAkGA1UEBhMC
VVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1NlYXR0bGUxGDAWBgNV
BAoTD0FtYXpvbi5jb20gSW5jLjEaMBgGA1UEAxMRZWMyLmFtYXpvbmF3cy5jb20w
gZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBAIe9GN//SRK2knbjySG0ho3yqQM3
e2TDhWO8D2e8+XZqck754gFSo99AbT2RmXClambI7xsYHZFapbELC4H91ycihvrD
jbST1ZjkLQgga0NE1q43eS68ZeTDccScXQSNivSlzJZS8HJZjgqzBlXjZftjtdJL
XeE4hwvo0sD4f3j9AgMBAAGjgc8wgcwwHQYDVR0OBBYEFCXWzAgVyrbwnFncFFIs
77VBdlE4MIGcBgNVHSMEgZQwgZGAFCXWzAgVyrbwnFncFFIs77VBdlE4oW6kbDBq
MQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHU2Vh
dHRsZTEYMBYGA1UEChMPQW1hem9uLmNvbSBJbmMuMRowGAYDVQQDExFlYzIuYW1h
em9uYXdzLmNvbYIJAKnL4UEDMN/FMAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQEF
BQADgYEAFYcz1OgEhQBXIwIdsgCOS8vEtiJYF+j9uO6jz7VOmJqO+pRlAbRlvY8T
C1haGgSI/A1uZUKs/Zfnph0oEI0/hu1IIJ/SKBDtN5lvmZ/IzbOPIJWirlsllQIQ
7zvWbGd9c9+Rm3p04oTvhup99la7kZqevJK0QRdD/6NpCKsqP/0=
-----END CERTIFICATE-----`

var (
	RSACert       *x509.Certificate
	RSACertPEM, _ = pem.Decode([]byte(AWSRSAIIDCert))
)

func main() {

	var err error

	if RSACert, err = x509.ParseCertificate(RSACertPEM.Bytes); err != nil {
		panic(err)
	}

	fmt.Println("buscando certificado RSA")
	RSASig, err := http.Get("http://169.254.169.254/latest/dynamic/instance-identity/signature")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer RSASig.Body.Close()
	bod1, err := ioutil.ReadAll(RSASig.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(bod1))

	DecodedRSASig, err := base64.StdEncoding.DecodeString(string(bod1))
	if err != nil {
		fmt.Println("Failed to Decode Signature: " + err.Error())
	}

	fmt.Println("Buscando iid")
	document, err := http.Get("http://169.254.169.254/latest/dynamic/instance-identity/document")
	if err != nil {
		panic(err)
	}

	defer document.Body.Close()
	bod2, err := ioutil.ReadAll(document.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(bod2))

	fmt.Printf("Checking RSA Certificate SHA256WithRSA\n")
	err = RSACert.CheckSignature(x509.SHA256WithRSA, bod2, DecodedRSASig)
	if err != nil {
		fmt.Println("Unable to verify: " + err.Error())
	} else {
		fmt.Println("Verified OK")
	}
}
