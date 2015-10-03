/*
Package resource/models/comments

Status: Stub

Comment is type Item

*/
package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Content []byte
	User    User
}

func (comment Comment) validate(db *gorm.DB) {}

func (comment Comment) stringify() string {
	return fmt.Sprintf("%s", comment.Content)
}
