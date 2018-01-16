package trade

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
)

func GetPrivatekey() (result []byte, err error) {
	result = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXwIBAAKBgQDQLTlt25Vyhl/RkjL4dJ77lcJ0gfgsw04VFnjN66knK3ytBKEJ
6Pia9UjfzfcuYE7wxhjpI+lW/Ksf50oTSDhUH2It4YkFx1HVMSdZ7qPaGgvhcLqV
KSRBE07Kcb6P8Qiqt4B4WWdslx965p3whY3PRiYqaWI9OpwPOQODWSfU3QIDAQAB
AoGBAMbDx2ebFzBIGMjSnJQZVYrFTOtNBRZITA9aa3HBprpdjSbtmo0JwgTCWhhG
YdIH5peBrVs9DJgfm4xUm6eZdx0wR0ZTt2/gWyITs+HmtJ8KxT2SvJzxAalYtHkt
7vlX0IjW6ZsH9c3gKvUwLJ5Ixe+HZCYjTkxoMic5C8vEiZmBAkEA/NoXw9yQIucS
R2mLdz61rzoHwBgfms6edeQUu7ppyFPsc1rWwHRxBA/V5NaVFbIlkRsfmGZ9pyln
jdBCNxHorQJBANLEvUjgkwiyAMSV8aPGGHLzqIOuIv1mFIkmZWqTHQY57I5I/Lvc
EqpBKqLGUwcsMz8WluQ+JiOkeKPdXIKCMvECQQDi5ueqqLRjzc5WbT1tTcYGr+Gi
nUNHTaFfk8STTk59Keqm/d53KEb+6SL9zx5MMOiLVba9sUOTDZHS7g9tkdGlAkEA
gxQj/AzepIu/eoMeMpJiZisu5CYKULmJj/o3HF69sD+Z5KtzsomdehDpKS5aOJ2+
iT/NO8mDAquo85AZlnjOoQJBAKOPpWNL+Ps2X+of5tLDdvitkrarIxaQE8dsJI2g
UgiKZNAFz5vyN6XHJxkwvKsHyzOCk8cjpVp5Q/seh3fMFo0=
-----END RSA PRIVATE KEY-----`)
	return
}

/*
func GetPublickey() (result []byte, err error) {
	result = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDDI6d306Q8fIfCOaTXyiUeJHkr
IvYISRcc73s3vF1ZT7XN8RNPwJxo8pWaJMmvyTn9N4HQ632qJBVHf8sxHi/fEsra
prwCtzvzQETrNRwVxLO5jVmRGi60j8Ue1efIlzPXV9je9mkjzOmdssymZkh2QhUr
CmZYI/FCEa3/cNMW0QIDAQAB
-----END PUBLIC KEY-----`)
	return
}
*/

func GetPublickey() (result []byte, err error) {
	result = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDDI6d306Q8fIfCOaTXyiUeJHkrIvYISRcc73s3vF1ZT7XN8RNPwJxo8pWaJMmvyTn9N4HQ632qJBVHf8sxHi/fEsraprwCtzvzQETrNRwVxLO5jVmRGi60j8Ue1efIlzPXV9je9mkjzOmdssymZkh2QhUrCmZYI/FCEa3/cNMW0QIDAQAB
-----END PUBLIC KEY-----`)
	return
}

func Decrypt(ciphertext string) (result string, err error) {
	privateKey, err := GetClientPrivatekey()
	if err != nil {
		return
	}
	block, _ := pem.Decode(privateKey)
	if block == nil {
		err = errors.New("private key error")
		return
	}
	privInterface, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}

	resultTemp, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return
	}
	resultByte, err := rsa.DecryptPKCS1v15(rand.Reader, privInterface, resultTemp)
	if err != nil {
		return
	}
	result = string(resultByte)
	return
}

func Verify(origdata, ciphertext string) (status bool, err error) {
	publicKey, err := GetClientPublickey()
	if err != nil {
		return
	}
	block, _ := pem.Decode(publicKey)
	if block == nil {
		err = errors.New("public key error")
		return
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	pub := pubInterface.(*rsa.PublicKey)

	t := sha1.New()
	io.WriteString(t, string(origdata))
	digest := t.Sum(nil)

	resultTemp, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return
	}
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA1, digest, resultTemp)
	if err != nil {
		return
	}
	status = true
	return
}

func Sign(origdata string) (result string, err error) {
	privateKey, err := GetPrivatekey()
	if err != nil {
		return
	}
	block, e := pem.Decode(privateKey)
	if block == nil {
		err = fmt.Errorf("private key error! %v", e)
		return
	}
	privInterface, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}

	t := sha1.New()
	io.WriteString(t, string(origdata))
	digest := t.Sum(nil)

	resultByte, err := rsa.SignPKCS1v15(rand.Reader, privInterface, crypto.SHA1, digest)
	if err != nil {
		return
	}
	result = base64.StdEncoding.EncodeToString(resultByte)
	return
}
