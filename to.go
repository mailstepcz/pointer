package pointer

// To returns a pointer to its argument.
func To[T any](x T) *T {
	return &x
}
