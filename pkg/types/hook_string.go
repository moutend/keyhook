// Code generated by "stringer -type=Hook"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[WH_KEYBOARD_LL-13]
	_ = x[WH_MOUSE_LL-14]
}

const _Hook_name = "WH_KEYBOARD_LLWH_MOUSE_LL"

var _Hook_index = [...]uint8{0, 14, 25}

func (i Hook) String() string {
	i -= 13
	if i >= Hook(len(_Hook_index)-1) {
		return "Hook(" + strconv.FormatInt(int64(i+13), 10) + ")"
	}
	return _Hook_name[_Hook_index[i]:_Hook_index[i+1]]
}