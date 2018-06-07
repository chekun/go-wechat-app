package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"encoding/json"

	"github.com/chekun/go-wechat-app/wechat/model"
)

//Decryptor controls all decrypt features
type Decryptor struct {
	appID      string
	sessionKey []byte
	Err        error
}

//New creates a decryptor
func New(appID, sessionKey string) (*Decryptor, error) {

	aesKey, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}

	return &Decryptor{
		appID:      appID,
		sessionKey: []byte(aesKey),
	}, nil
}

func (decryptor *Decryptor) base64Decode(encodedStr string) []byte {
	if decryptor.Err != nil {
		return nil
	}

	result, err := base64.StdEncoding.DecodeString(encodedStr)

	if err != nil {
		decryptor.Err = err
		return nil
	}

	return result
}

func (decryptor *Decryptor) decrypt(encryptedData, iv string) []byte {
	decryptor.Err = nil
	cipherBytes := decryptor.base64Decode(encryptedData)
	ivBytes := decryptor.base64Decode(iv)

	if decryptor.Err != nil {
		return nil
	}

	block, err := aes.NewCipher(decryptor.sessionKey)
	if err != nil {
		decryptor.Err = err
		return nil
	}
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	mode.CryptBlocks(cipherBytes, cipherBytes)

	return pkcs7Unpad(cipherBytes)
}

//Profile decrypt the encryped profile
func (decryptor *Decryptor) Profile(encryptedData, iv string) *model.Profile {
	result := decryptor.decrypt(encryptedData, iv)
	if decryptor.Err != nil {
		return nil
	}
	var profile model.Profile
	decryptor.Err = json.Unmarshal(result, &profile)
	return &profile
}

//Share decrypt the encryped share ticket
func (decryptor *Decryptor) Share(encryptedData, iv string) *model.Share {
	result := decryptor.decrypt(encryptedData, iv)
	if decryptor.Err != nil {
		return nil
	}
	var share model.Share
	decryptor.Err = json.Unmarshal(result, &share)
	return &share
}

//Run decrypt the encryped rundata
func (decryptor *Decryptor) Run(encryptedData, iv string) *model.Run {
	result := decryptor.decrypt(encryptedData, iv)
	if decryptor.Err != nil {
		return nil
	}
	var run model.Run
	decryptor.Err = json.Unmarshal(result, &run)
	return &run
}

func pkcs7Unpad(padBytes []byte) []byte {
	length := len(padBytes)
	padLen := int(padBytes[length-1])
	if padLen < 1 || padLen > 32 {
		padLen = 0
	}
	return padBytes[:(length - padLen)]
}
