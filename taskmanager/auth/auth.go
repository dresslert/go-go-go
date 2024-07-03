package auth

import (
    "time"
    "github.com/golang-jwt/jwt"
    "os"
)

func GenerateJWT(userID string) (string, error) {
    jwtSecret := os.Getenv("JWT_SECRET")

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString([]byte(jwtSecret))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
    jwtSecret := os.Getenv("JWT_SECRET")

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, jwt.ErrInvalidKey
        }
        return []byte(jwtSecret), nil
    })

    if err != nil {
        return nil, err
    }

    return token, nil
}
