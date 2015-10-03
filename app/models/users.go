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

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `sql:"size:255"` // screen name
	// expertise
}
