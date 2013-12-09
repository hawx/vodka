// Package types implements VType, the interface all vodka types must implement.
package types

type VType interface {

	// String returns the value as a string
	String()       string

	// Value returns the value
	Value()        interface{}

	// Type returns the name of the type as a string
	Type()         string

	// Compare compares the value with another, -1, 0, 1 being less, equal and
	// greater. -2 can be used to show not equal, but not less or greater.
	Compare(VType) int

}

type Rangeable interface {
	Next() Rangeable
	Prev() Rangeable
	VType
}
