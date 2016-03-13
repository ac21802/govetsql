package vetsql

import (
	"strings"
	"testing"
)

var IntTests = []struct {
	n        int
	expected bool
}{
	{-2147483649, false},
	{-2147483648, true},
	{2147483647, true},
	{2147483648, false},
}

func TestValidSQLInt(t *testing.T) {
	for _, tt := range IntTests {
		actual := ValidSQLInt(tt.n)
		if actual != tt.expected {
			t.Errorf("ValidSQLInt(%d): expected %t, actual %t", tt.n, tt.expected, actual)
		}
	}
}

var UIntTests = []struct {
	n        int
	expected bool
}{
	{-1, false},
	{0, true},
	{4294967295, true},
	{4294967296, false},
}

func TestValidSQLUInt(t *testing.T) {
	for _, tt := range UIntTests {
		actual := ValidSQLUInt(tt.n)
		if actual != tt.expected {
			t.Errorf("ValidSQLUInt(%d): expected %t, actual %t", tt.n, tt.expected, actual)
		}
	}
}

var TinyIntTests = []struct {
	n        int
	expected bool
}{
	{-129, false},
	{-128, true},
	{127, true},
	{128, false},
}

func TestValidSQLTinyInt(t *testing.T) {
	for _, tt := range TinyIntTests {
		actual := ValidSQLTinyInt(tt.n)
		if actual != tt.expected {
			t.Errorf("ValidSQLTinyInt(%d): expected %t, actual %t", tt.n, tt.expected, actual)
		}
	}
}

var UTinyIntTests = []struct {
	n        int
	expected bool
}{
	{-1, false},
	{0, true},
	{255, true},
	{256, false},
}

func TestValidSQLUTinyInt(t *testing.T) {
	for _, tt := range UTinyIntTests {
		actual := ValidSQLUTinyInt(tt.n)
		if actual != tt.expected {
			t.Errorf("ValidSQLUTinyInt(%d): expected %t, actual %t", tt.n, tt.expected, actual)
		}
	}
}

var SmallIntTests = []struct {
	n        int
	expected bool
}{
	{-32769, false},
	{-32768, true},
	{32767, true},
	{32768, false},
}

func TestValidSQLSmallInt(t *testing.T) {
	for _, tt := range SmallIntTests {
		actual := ValidSQLSmallInt(tt.n)
		if actual != tt.expected {
			t.Errorf("ValidSQLSmallInt(%d): expected %t, actual %t", tt.n, tt.expected, actual)
		}
	}
}

var USmallIntTests = []struct {
	n        int
	expected bool
}{
	{-1, false},
	{0, true},
	{65535, true},
	{65536, false},
}

func TestValidSQLUSmallInt(t *testing.T) {
	for _, tt := range USmallIntTests {
		actual := ValidSQLUSmallInt(tt.n)
		if actual != tt.expected {
			t.Errorf("ValidSQLUSmallInt(%d): expected %t, actual %t", tt.n, tt.expected, actual)
		}
	}
}

var MediumIntTests = []struct {
	n        int
	expected bool
}{
	{-8388609, false},
	{-8388608, true},
	{8388607, true},
	{8388608, false},
}

func TestValidSQLMediumInt(t *testing.T) {
	for _, tt := range MediumIntTests {
		actual := ValidSQLMediumInt(tt.n)
		if actual != tt.expected {
			t.Errorf("ValidSQLMediumInt(%d): expected %t, actual %t", tt.n, tt.expected, actual)
		}
	}
}

var UMediumIntTests = []struct {
	n        int
	expected bool
}{
	{-1, false},
	{0, true},
	{16777215, true},
	{16777216, false},
}

func TestValidSQLUMediumInt(t *testing.T) {
	for _, tt := range UMediumIntTests {
		actual := ValidSQLUMediumInt(tt.n)
		if actual != tt.expected {
			t.Errorf("ValidSQLUMediumInt(%d): expected %t, actual %t", tt.n, tt.expected, actual)
		}
	}
}

var BigIntTests = []struct {
	n        int
	expected bool
}{
	{-9223372036854775808, true},
	{9223372036854775807, true},
}

func TestValidSQLBigInt(t *testing.T) {
	for _, tt := range BigIntTests {
		actual := ValidSQLBigInt(tt.n)
		if actual != tt.expected {
			t.Errorf("ValidSQLBigInt(%d): expected %t, actual %t", tt.n, tt.expected, actual)
		}
	}
}

var UBigIntTests = []struct {
	n        uint64
	expected bool
}{
	{0, true},
	{18446744073709551615, true},
}

func TestValidSQLUBigInt(t *testing.T) {
	for _, tt := range UBigIntTests {
		actual := ValidSQLUBigInt(tt.n)
		if actual != tt.expected {
			t.Errorf("ValidSQLUBigInt(%d): expected %t, actual %t", tt.n, tt.expected, actual)
		}
	}
}

var textLimitTests = []struct {
	s        string
	expected bool
}{
	{strings.Repeat("g", 65535), false},
	{strings.Repeat("b", 65536), true},
}

func TestExceedsTextLimit(t *testing.T) {
	for _, tt := range textLimitTests {
		actual := ExceedsTextLimit(tt.s)
		if actual != tt.expected {
			t.Errorf("ExceedsTextLimit(%s): expected %t, actual %t", tt.s, tt.expected, actual)
		}
	}
}

var tinyTextLimitTests = []struct {
	s        string
	expected bool
}{
	{strings.Repeat("g", 255), false},
	{strings.Repeat("b", 256), true},
}

func TestExceedsTinyTextLimit(t *testing.T) {
	for _, tt := range tinyTextLimitTests {
		actual := ExceedsTinyTextLimit(tt.s)
		if actual != tt.expected {
			t.Errorf("ExceedsTinyTextLimit(%s): expected %t, actual %t", tt.s, tt.expected, actual)
		}
	}
}

var mediumTextLimitTests = []struct {
	s        string
	expected bool
}{
	{strings.Repeat("g", 16777215), false},
	{strings.Repeat("b", 16777216), true},
}

func TestExceedsMediumTextLimit(t *testing.T) {
	for _, tt := range mediumTextLimitTests {
		actual := ExceedsMediumTextLimit(tt.s)
		if actual != tt.expected {
			t.Errorf("ExceedsMediumTextLimit(%s): expected %t, actual %t", tt.s, tt.expected, actual)
		}
	}
}

var longTextLimitTests = []struct {
	s        string
	expected bool
}{
	{strings.Repeat("g", 4294967295), false},
	{strings.Repeat("b", 4294967296), true},
}

func TestExceedsLongTextLimit(t *testing.T) {
	for _, tt := range longTextLimitTests {
		actual := ExceedsLongTextLimit(tt.s)
		if actual != tt.expected {
			t.Errorf("ExceedsLongTextLimit(%s): expected %t, actual %t", tt.s, tt.expected, actual)
		}
	}
}
