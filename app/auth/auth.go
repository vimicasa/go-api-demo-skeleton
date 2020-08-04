package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

type AccessDetails struct {
	TokenID string
	UserId  string
}

type TokenDetails struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenID      string `json:"token_id"`
	RefreshID    string `json:"refresh_id"`
	AtExpires    int64  `json:"at_expires_in"`
	RtExpires    int64  `json:"rt_expires_in"`
}

func CreateToken(userId string) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 30).Unix() // TODO: Change it if needed
	td.TokenID = uuid.NewV4().String()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	//TODO: Add if necessary claims: guid,iss,aud,amr,auth_source,auth_time
	atClaims["jti"] = td.TokenID
	atClaims["user_id"] = userId
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte("ACCESS_SECRET")) // TODO: change this value. os.Getenv("ACCESS_SECRET")
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token TODO: Store Refresh Token
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix() // TODO: Change it if needed
	td.RefreshID = td.TokenID + "++" + userId

	rtClaims := jwt.MapClaims{}
	rtClaims["jti"] = td.RefreshID
	rtClaims["user_id"] = userId
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	td.RefreshToken, err = rt.SignedString([]byte("REFRESH_SECRET")) // TODO: change this value. os.Getenv("REFRESH_SECRET")
	if err != nil {
		return nil, err
	}
	return td, nil
}

// VerifyToken check validity of token

func GetValidToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := verifyToken(tokenStr)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, err
	}
	return claims, nil
}

func verifyToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("ACCESS_SECRET"), nil // TODO: change this value. os.Getenv("REFRESH_SECRET")
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
