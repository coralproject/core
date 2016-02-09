package policy

import (
	"github.com/coralproject/core/pkg/data"
)

type PolicyItem struct {
	item Item
}

func (i *PolicyItem) Value(key interface{}) interface{} {

	// get value from i.item

	// check policy

	// transform item value if necessary

	// return value

}
