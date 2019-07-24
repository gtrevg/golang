package object

import (
	"fmt"
	"strings"
)

// Returns the result of calling {@code toString} on the first
// argument if the first argument is not {@code null} and returns
// the second argument otherwise.
func ToString(o interface{}, nullDefault ...string) string {
	if o == nil && nullDefault != nil && len(nullDefault) != 0 {
		return strings.Join(nullDefault, "")
	}
	return fmt.Sprintf("%v", o)
}
