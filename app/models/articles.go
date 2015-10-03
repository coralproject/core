/*

Article is type externalReference

*/

package models

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Name      string
	Publisher string
}
