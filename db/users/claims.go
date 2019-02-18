package users

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CustomClaims struct with User
type CustomClaims struct {
	User *User
	jwt.StandardClaims
}

// ValidationError ValidationError
type ValidationError struct {
	jwt.ValidationError
}

func (e *ValidationError) valid() bool {
	return e.Errors == 0
}

// Validates time based claims "exp, iat, nbf".
// There is no accounting for clock skew.
// As well, if any of the above claims are not in the token, it will still
// be considered a valid claim.
func (c CustomClaims) Valid() error {
	vErr := new(ValidationError)
	now := jwt.TimeFunc().Unix()

	// The claims below are optional, by default, so if they are set to the
	// default value in Go, let's not fail the verification for them.
	if c.VerifyExpiresAt(now, false) == false {
		delta := time.Unix(now, 0).Sub(time.Unix(c.ExpiresAt, 0))
		vErr.Inner = fmt.Errorf("token is expired by %v", delta)
		vErr.Errors |= jwt.ValidationErrorExpired
	}

	if c.VerifyIssuedAt(now, false) == false {
		vErr.Inner = fmt.Errorf("Token used before issued")
		vErr.Errors |= jwt.ValidationErrorIssuedAt
	}

	if c.VerifyNotBefore(now, false) == false {
		vErr.Inner = fmt.Errorf("token is not valid yet")
		vErr.Errors |= jwt.ValidationErrorNotValidYet
	}
	valid, validUser := c.User.Validate()
	if valid || validUser.ID == "" {
		vErr.Inner = fmt.Errorf("User not exist in DB")
		vErr.Errors |= jwt.ValidationErrorNotValidYet
	}
	if vErr.valid() {
		return nil
	}
	return vErr
}

// Compares the aud claim against cmp.
// If required is false, this method will return true if the value matches or is unset
// func (c *CustomClaims) VerifyAudience(cmp string, req bool) bool {
// 	fmt.Println("VerifyAudience")
// 	fmt.Println(verifyAud(c.Audience, cmp, req))
// 	return verifyAud(c.Audience, cmp, req)
// }

// // Compares the exp claim against cmp.
// // If required is false, this method will return true if the value matches or is unset
// func (c *CustomClaims) VerifyExpiresAt(cmp int64, req bool) bool {
// 	fmt.Println("VerifyExpiresAt")
// 	fmt.Println(verifyExp(c.ExpiresAt, cmp, req))
// 	return verifyExp(c.ExpiresAt, cmp, req)
// }

// // Compares the iat claim against cmp.
// // If required is false, this method will return true if the value matches or is unset
// func (c *CustomClaims) VerifyIssuedAt(cmp int64, req bool) bool {
// 	fmt.Println("VerifyIssuedAt")
// 	fmt.Println(verifyIat(c.IssuedAt, cmp, req))
// 	return verifyIat(c.IssuedAt, cmp, req)
// }

// // Compares the iss claim against cmp.
// // If required is false, this method will return true if the value matches or is unset
// func (c *CustomClaims) VerifyIssuer(cmp string, req bool) bool {
// 	fmt.Println("verifyIss")
// 	fmt.Println(verifyIss(c.Issuer, cmp, req))
// 	return verifyIss(c.Issuer, cmp, req)
// }

// // Compares the nbf claim against cmp.
// // If required is false, this method will return true if the value matches or is unset
// func (c *CustomClaims) VerifyNotBefore(cmp int64, req bool) bool {
// 	fmt.Println("verifyNbf")
// 	fmt.Println(verifyNbf(c.NotBefore, cmp, req))
// 	return verifyNbf(c.NotBefore, cmp, req)
// }

// // ----- helpers

// func verifyAud(aud string, cmp string, required bool) bool {
// 	if aud == "" {
// 		return !required
// 	}
// 	if subtle.ConstantTimeCompare([]byte(aud), []byte(cmp)) != 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func verifyExp(exp int64, now int64, required bool) bool {
// 	if exp == 0 {
// 		return !required
// 	}
// 	return now <= exp
// }

// func verifyIat(iat int64, now int64, required bool) bool {
// 	if iat == 0 {
// 		return !required
// 	}
// 	return now >= iat
// }

// func verifyIss(iss string, cmp string, required bool) bool {
// 	if iss == "" {
// 		return !required
// 	}
// 	if subtle.ConstantTimeCompare([]byte(iss), []byte(cmp)) != 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func verifyNbf(nbf int64, now int64, required bool) bool {
// 	if nbf == 0 {
// 		return !required
// 	}
// 	return now >= nbf
// }
