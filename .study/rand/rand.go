package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func genKeys() (*rsa.PrivateKey, *rsa.PublicKey) {
  priv, err := rsa.GenerateKey(rand.Reader, 2048)
  if err != nil {
    return nil, nil
  }
  return priv, &priv.PublicKey
}

func generateJTI() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func signJWT(priv *rsa.PrivateKey) (string, error) {
  token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
    "sub": "1234567890",
    "name": "John Doe",
    "iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"jti":  generateJTI(),
  })

  return token.SignedString(priv)
}

func verifyJWT(tokenString string, publicKey *rsa.PublicKey) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func main() {
  priv, pub := genKeys()

  // fmt.Println("Private Key: ", priv)
  fmt.Println("Public Key: ", pub)

  token, err := signJWT(priv)
  if err != nil {
    fmt.Println("Error: ", err)
  }

  fmt.Println("Token: ", token)

  claims, err := verifyJWT(token, pub)
	if err != nil {
		fmt.Printf("JWTの検証に失敗: %v\n", err)
		return
	}
	fmt.Printf("検証済みのペイロード: %v\n", claims)
}
