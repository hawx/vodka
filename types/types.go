package types

/*import (
	"strconv"
	"strings"
	p "../parser"
	"../stack"
)
*/


type VType interface {
	// Returns the value as a string
	String()       string
	// Returns the value
	Value()        interface{}
	// Returns the name of the type as a string
	Type()         string
	// Compares the value with another, -1, 0, 1 being less, equal and greater. -2
	// can be used to show not equal, but not less or greater.
	Compare(VType) int
}
