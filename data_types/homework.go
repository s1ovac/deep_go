package data_types

// go test -v homework_test.go

func ToLittleEndian(number uint32) uint32 {
	return ((number >> 24) & 0xff) |
		((number >> 8) & 0xff00) |
		((number << 8) & 0xff0000) |
		((number << 24) & 0xff000000)
}
