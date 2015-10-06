/*

Article is type externalReference

*/

package models

import (
	"github.com/coralproject/core/resource"
	"github.com/jinzhu/gorm"
)

// Article is an external reference to a piece of content
type Article struct {
	// To Do: The gorm ID should be set with asset ID
	gorm.Model        // ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `json:"name",sql:"size:255"`
	Publisher  string

	URL util.URL
}

func (a Article) validate() error {
	return nil
}
