# amazon-iid
To verify the instance identity document

??? Agents attested by the aws_iid attestor will be issued a SPIFFE ID like 
``` spiffe://example.org/spire/agent/aws_iid/ACCOUNT_ID/REGION/INSTANCE_ID.```

### Documentos de identidade da instância
| Dados	| Descrição |
| ----- | --------- |
| devpayProductCodes | Suspenso. |
| marketplaceProductCodes | O código do produto AWS Marketplace da AMI usada para iniciar a instância. |
| availabilityZone | A zona de disponibilidade na qual a instância está em execução. |
| privateIp | O endereço IPv4 privado da instância. |
| version | A versão do formato do documento de identidade da instância. |
| instanceId | O ID da instância. |
| billingProducts | Os produtos de faturamento da instância. |
| instanceType | O tipo de instância da instância. |
| accountId | O ID da conta da AWS que iniciou a instância. |
| imageId | A ID do AMI usado para executar a instância. |
| pendingTime | A data e a hora em que a instância foi iniciada. |
| architecture | A arquitetura da AMI usada para iniciar a instância (i386 | x86_64 | arm64). |
| kernelId | O ID do kernel associado à instância, se aplicável. |
| ramdiskId | O ID do disco de RAM associado a essa instância, se aplicável. |
| region | A região em que a instância está em execução. |


### Usar a assinatura PKCS7 para verificar o documento de identidade da instância	

	
```
fetch http://169.254.169.254/latest/dynamic/instance-identity/pkcs7
```

### Usar a assinatura RSA-2048 para verificar o documento de identidade da instância
	
```
fetch http://169.254.169.254/latest/dynamic/instance-identity/rsa2048
```


### type SignatureAlgorithm int
```
const (
	UnknownSignatureAlgorithm SignatureAlgorithm = iota

	MD2WithRSA  // Unsupported.
	MD5WithRSA  // Only supported for signing, not verification.
	SHA1WithRSA // Only supported for signing, and verification of CRLs, CSRs, and OCSP responses.
	SHA256WithRSA
	SHA384WithRSA
	SHA512WithRSA
	DSAWithSHA1   // Unsupported.
	DSAWithSHA256 // Unsupported.
	ECDSAWithSHA1 // Only supported for signing, and verification of CRLs, CSRs, and OCSP responses.
	ECDSAWithSHA256
	ECDSAWithSHA384
	ECDSAWithSHA512
	SHA256WithRSAPSS
	SHA384WithRSAPSS
	SHA512WithRSAPSS
	PureEd25519
)
	


	


	


	


