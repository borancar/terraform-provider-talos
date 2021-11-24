// Code generated by "stringer -type=Family -linecomment -output family_string_linux.go"; DO NOT EDIT.

package nethelpers

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FamilyInet4-2]
	_ = x[FamilyInet6-10]
}

const (
	_Family_name_0 = "inet4"
	_Family_name_1 = "inet6"
)

func (i Family) String() string {
	switch {
	case i == 2:
		return _Family_name_0
	case i == 10:
		return _Family_name_1
	default:
		return "Family(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
