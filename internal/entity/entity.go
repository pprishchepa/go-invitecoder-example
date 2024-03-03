package entity

import "fmt"

var ErrNotAvailable = fmt.Errorf("not available")
var ErrAlreadyExists = fmt.Errorf("already exists")

type InvitedUser struct {
	Email      string
	InvitedVia string
}
