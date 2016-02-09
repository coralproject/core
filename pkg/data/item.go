package data

// top level Item interface
type Item interface {

	ID() string
	Value(key interface{}) interface{}

}

type Stuffed struct {
	stuff map[interface{}] interface{},	
	Value(key interface{}) interface{},
}

type Comment struct {

	// getters 
	Body() string,


	// aspects
	Stuffed,
	Actionable,
}



func (c *Comment) Body() string {

	return c.Value("body").(string)

}


var c Comment;

fmt.Printf(c.Body())