// Code generated by "gcraft_stringer -type=FieldType"; DO NOT EDIT.

package update

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Uint32-0]
	_ = x[Uint32Array-1]
	_ = x[Int32-2]
	_ = x[Int32Array-3]
	_ = x[Float32-4]
	_ = x[Float32Array-5]
	_ = x[GUID-6]
	_ = x[GUIDArray-7]
	_ = x[ArrayType-8]
	_ = x[Uint8-9]
	_ = x[Bit-10]
	_ = x[Pad-11]
}

const _FieldType_name = "Uint32Uint32ArrayInt32Int32ArrayFloat32Float32ArrayGUIDGUIDArrayArrayTypeUint8BitPad"

var _FieldType_index = [...]uint8{0, 6, 17, 22, 32, 39, 51, 55, 64, 73, 78, 81, 84}

func (i FieldType) String() string {
	if i >= FieldType(len(_FieldType_index)-1) {
		return "FieldType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FieldType_name[_FieldType_index[i]:_FieldType_index[i+1]]
}
