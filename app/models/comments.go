package models

/*
Package resource/models/comments

Status: Stub

Comment is type Item

*/

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Comment models the comment with the content, user and comment childs.
type Comment struct {
	gorm.Model
	Content  []byte // it can not be nil
	User     User   // it can not be nil
	Comments []Comment
}

func (comment Comment) validate(db *gorm.DB) error {
	return nil
}

func (comment Comment) stringify() string {
	return fmt.Sprintf("%s", comment.Content)
}
