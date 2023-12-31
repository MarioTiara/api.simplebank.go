package token

import "time"

type Maker interface {
	//createToken creates a new token for s specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	//VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
