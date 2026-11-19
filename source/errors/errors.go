package errors

// Error is the template for a language error.
type Error interface {
	String() string
}
