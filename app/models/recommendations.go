/*
Package resource/models/comments

Status: Stub

Comment is type Item

*/
package models

import "github.com/jinzhu/gorm"

type Recommendation struct {
	gorm.Model
	user    User
	comment Comment
}
