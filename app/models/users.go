/*
Package resource/models/contributors

Status: Stub

TO DO

* GORM migrations
* Validations
* Testing

* Fields for the Trustonator
  - algorithms being used to

*/

package models

import (
	"fmt"
	"net/mail"

	"github.com/coralproject/core/resource"
	"github.com/jinzhu/gorm"
)

// User is the reader, author or moderator
type User struct {
	gorm.Model
	Name     string       `sql:"size:255"` // screen name
	Email    mail.Address // struct { Proper name (may be empty) , user@domain }
	Location string       // for now only a string (may be empty)
	Avatar   util.Avatar  // struct { URL to the image } (may be empty)
	Comments []Comment    // all comments on this comment (may be empty)
}

func (u *User) stringify() string {
	return fmt.Sprintf("%s", u.Name)
}
