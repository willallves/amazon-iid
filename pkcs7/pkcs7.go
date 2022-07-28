package pkcs7

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/fullsailor/pkcs7"
	"github.com/willallves/amazon-iid/common"
)

// AWSPKCS7IIDCert is the PKCS7 public certificate
const AWSPKCS7IIDCert = `-----BEGIN CERTIFICATE-----
MIIC7TCCAq0CCQCWukjZ5V4aZzAJBgcqhkjOOAQDMFwxCzAJBgNVBAYTAlVTMRkw
FwYDVQQIExBXYXNoaW5ndG9uIFN0YXRlMRAwDgYDVQQHEwdTZWF0dGxlMSAwHgYD
VQQKExdBbWF6b24gV2ViIFNlcnZpY2VzIExMQzAeFw0xMjAxMDUxMjU2MTJaFw0z
ODAxMDUxMjU2MTJaMFwxCzAJBgNVBAYTAlVTMRkwFwYDVQQIExBXYXNoaW5ndG9u
IFN0YXRlMRAwDgYDVQQHEwdTZWF0dGxlMSAwHgYDVQQKExdBbWF6b24gV2ViIFNl
cnZpY2VzIExMQzCCAbcwggEsBgcqhkjOOAQBMIIBHwKBgQCjkvcS2bb1VQ4yt/5e
ih5OO6kK/n1Lzllr7D8ZwtQP8fOEpp5E2ng+D6Ud1Z1gYipr58Kj3nssSNpI6bX3
VyIQzK7wLclnd/YozqNNmgIyZecN7EglK9ITHJLP+x8FtUpt3QbyYXJdmVMegN6P
hviYt5JH/nYl4hh3Pa1HJdskgQIVALVJ3ER11+Ko4tP6nwvHwh6+ERYRAoGBAI1j
k+tkqMVHuAFcvAGKocTgsjJem6/5qomzJuKDmbJNu9Qxw3rAotXau8Qe+MBcJl/U
hhy1KHVpCGl9fueQ2s6IL0CaO/buycU1CiYQk40KNHCcHfNiZbdlx1E9rpUp7bnF
lRa2v1ntMX3caRVDdbtPEWmdxSCYsYFDk4mZrOLBA4GEAAKBgEbmeve5f8LIE/Gf
MNmP9CM5eovQOGx5ho8WqD+aTebs+k2tn92BBPqeZqpWRa5P/+jrdKml1qx4llHW
MXrs3IgIb6+hUIB+S8dz8/mmO0bpr76RoZVCXYab2CZedFut7qc3WUH9+EUAH5mw
vSeDCOUMYQR7R9LINYwouHIziqQYMAkGByqGSM44BAMDLwAwLAIUWXBlk40xTwSw
7HX32MxXYruse9ACFBNGmdX2ZBrVNGrN9N2f6ROk0k9K
-----END CERTIFICATE-----`

var (
	PKCS7Cert       *x509.Certificate
	PKCS7CertPEM, _ = pem.Decode([]byte(AWSPKCS7IIDCert))
)

func VerifyPKCS7Cert() {

	var err error

	if PKCS7Cert, err = x509.ParseCertificate(PKCS7CertPEM.Bytes); err != nil {
		panic(err)
	}

	// The /pkcs7 endpoint contains the document and the signature

	fmt.Println("Fetching PKCS7 Signature")
	PKCS7Sig, err := common.FetchDatas("http://169.254.169.254/latest/dynamic/instance-identity/pkcs7")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(PKCS7Sig))

	fmt.Println("Checking against PKCS7 Certificate")
	PKCS7SigNew := fmt.Sprintf("-----BEGIN PKCS7-----\n%s\n-----END PKCS7-----", string(PKCS7Sig))

	PKCS7SigBER, PKCS7SigRest := pem.Decode([]byte(PKCS7SigNew))
	if len(PKCS7SigRest) != 0 {
		panic("Failed to decode the PEM encoded PKCS7 signature")
	}

	PKCS7Data, err := pkcs7.Parse(PKCS7SigBER.Bytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(PKCS7Data.Content))

	PKCS7Data.Certificates = []*x509.Certificate{PKCS7Cert}

	err = PKCS7Data.Verify()
	if err != nil {
		fmt.Println("Unable to verify: " + err.Error())
	} else {
		fmt.Println("Verified OK")
	}

}
