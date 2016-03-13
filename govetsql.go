package vetsql

// INTS
const minU = 0

const minInt = -2147483648
const maxInt = 2147483647
const maxUInt = 4294967295

const minTinyInt = -128
const maxTinyInt = 127
const maxUTinyInt = 255

const minSmallInt = -32768
const maxSmallInt = 32767
const maxUSmallInt = 65535

const minMediumInt = -8388608
const maxMediumInt = 8388607
const maxUMediumInt = 16777215

const minBigInt = -9223372036854775808
const maxBigInt = 9223372036854775807
const maxUBigInt = 18446744073709551615

// ValidSQLInt ...
func ValidSQLInt(value int) bool {
	if value <= maxInt && value >= minInt {
		return true
	}
	return false
}

// ValidSQLUInt ...
func ValidSQLUInt(value int) bool {
	if value <= maxUInt && value >= minU {
		return true
	}
	return false
}

// ValidSQLTinyInt ...
func ValidSQLTinyInt(value int) bool {
	if value <= maxTinyInt && value >= minTinyInt {
		return true
	}
	return false
}

// ValidSQLUTinyInt ...
func ValidSQLUTinyInt(value int) bool {
	if value <= maxUTinyInt && value >= minU {
		return true
	}
	return false
}

// ValidSQLSmallInt ...
func ValidSQLSmallInt(value int) bool {
	if value <= maxSmallInt && value >= minSmallInt {
		return true
	}
	return false
}

// ValidSQLUSmallInt ...
func ValidSQLUSmallInt(value int) bool {
	if value <= maxUSmallInt && value >= minU {
		return true
	}
	return false
}

// ValidSQLMediumInt ...
func ValidSQLMediumInt(value int) bool {
	if value <= maxMediumInt && value >= minMediumInt {
		return true
	}
	return false
}

// ValidSQLUMediumInt ...
func ValidSQLUMediumInt(value int) bool {
	if value <= maxUMediumInt && value >= minU {
		return true
	}
	return false
}

// ValidSQLBigInt ...
func ValidSQLBigInt(value int) bool {
	if value <= maxBigInt && value >= minBigInt {
		return true
	}
	return false
}

// ValidSQLUBigInt ...
func ValidSQLUBigInt(value uint64) bool {
	if value <= maxUBigInt && value >= minU {
		return true
	}
	return false
}

// strings

const blobLimit = 65535
const tinyBlobLimit = 255
const mediumBlobLimit = 16777215
const longBlobLimit = 4294967295

// ExceedsCharLimit ...
func ExceedsCharLimit(str string, limit int) bool {
	if len(str) > limit {
		return true
	}
	return false
}

// ExceedsBlobLimit ...
func ExceedsBlobLimit(str string) bool {
	return ExceedsCharLimit(str, blobLimit)
}

// ExceedsTinyBlobLimit ...
func ExceedsTinyBlobLimit(str string) bool {
	return ExceedsCharLimit(str, tinyBlobLimit)
}

// ExceedsMediumBlobLimit ...
func ExceedsMediumBlobLimit(str string) bool {
	return ExceedsCharLimit(str, mediumBlobLimit)
}

// ExceedsLongBlobLimit ...
func ExceedsLongBlobLimit(str string) bool {
	return ExceedsCharLimit(str, longBlobLimit)
}

// ExceedsTextLimit ...
func ExceedsTextLimit(str string) bool {
	return ExceedsBlobLimit(str)
}

// ExceedsTinyTextLimit ...
func ExceedsTinyTextLimit(str string) bool {
	return ExceedsTinyBlobLimit(str)
}

// ExceedsMediumTextLimit ...
func ExceedsMediumTextLimit(str string) bool {
	return ExceedsMediumBlobLimit(str)
}

// ExceedsLongTextLimit ...
func ExceedsLongTextLimit(str string) bool {
	return ExceedsLongBlobLimit(str)
}
