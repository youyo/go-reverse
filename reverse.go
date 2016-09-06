package reverse

// Bool inverts the value.
//  newBool := reverse.Bool(bool)
func Bool(b bool) bool {
	return func() bool {
		if b {
			return false
		} else {
			return true
		}
	}()
}
