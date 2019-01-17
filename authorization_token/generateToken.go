package authorization_token

import (
	"fmt"
	"io/ioutil"
	"encoding/pem"
	"crypto/x509"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	PrivateKey string
	TeamId string
	KeyId string
	IssuedAt int64
	Bearer string
}


func GenerateAuthorizationToken(p8FilePath, keyId, teamId string) (string, error)  {

	fmt.Println("---------------------------Generating Authorization Token----------------------------")

	//First read p8 file
	privateKey, err := readP8File(p8FilePath)
	if err!=nil{

		return "", err
	}

	if privateKey == nil{
		return "", errors.New("Error in p8 file")
	}

	iss := teamId
	iat := time.Now().Unix()
	alg := "ES256"

	myToken := new(Token)

	myToken.IssuedAt = time.Now().Unix()

	jwtToken := jwt.New(jwt.SigningMethodES256)

	jwtToken.Claims = jwt.MapClaims{
		"iss" : iss,
		"iat" : iat,
	}

	jwtToken.Header["alg"] = alg
	jwtToken.Header["kid"] = keyId

	fmt.Println("--------------------------Authorization Token Generated----------------------------")

	return jwtToken.SignedString(privateKey)

}

func readP8File(p8FilePath string) (interface{}, error)  {

	bytes, err := ioutil.ReadFile(p8FilePath)
	if err!= nil{
		fmt.Println("---------------------------P8 file not found----------------------------")
		return nil, err
	}

	block, _ := pem.Decode(bytes)
	if block == nil{
		fmt.Println("---------------------------Error while decoding P8 file----------------------------")
		return nil, errors.New("Error while decoding P8 file")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)

	if err!=nil{

		fmt.Println("---------------------------Error while parsing P8 file----------------------------")
		return nil, errors.New("Error while parsing P8 file")
	}

	return privateKey, nil
}
