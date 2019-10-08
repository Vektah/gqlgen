// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package introspection

import (
	"fmt"
	"io"
	"strconv"
)

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
}

type UserCreateInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateInput struct {
	ID       string  `json:"id"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

type Role string

const (
	RoleAdmin     Role = "Admin"
	RoleUser      Role = "User"
	RoleAnonymous Role = "Anonymous"
)

var AllRole = []Role{
	RoleAdmin,
	RoleUser,
	RoleAnonymous,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleUser, RoleAnonymous:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
