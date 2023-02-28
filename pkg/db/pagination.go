package db

const (
	DefaultOffset = 0
	DefaultFirst  = 10
)

// LoadOffset ...
func LoadOffset(first, offset *int) int {
	ret := DefaultOffset

	if offset == nil {
		return DefaultOffset
	}

	if first != nil {
		ret = (*first) * (*offset)
	} else {
		ret = DefaultFirst * (*offset)
	}

	return ret
}

// LoadLimit ...
func LoadLimit(first *int) int {
	ret := DefaultFirst

	if first != nil {
		ret = *first
	}

	return ret
}
