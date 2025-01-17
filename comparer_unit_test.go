package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_IsDST(t *testing.T) {
	assert := assert.New(t)

	tzWithDST, tzWithoutDST := "Australia/Sydney", "Australia/Brisbane"

	tests := []struct {
		input    string
		timezone string
		expected bool
	}{
		0: {"", tzWithDST, false},
		1: {"", tzWithoutDST, false},
		2: {"0", tzWithDST, false},
		3: {"0", tzWithoutDST, false},
		4: {"0000-00-00", tzWithDST, false},
		5: {"0000-00-00", tzWithoutDST, false},
		6: {"00:00:00", tzWithDST, false},
		7: {"00:00:00", tzWithoutDST, false},
		8: {"0000-00-00 00:00:00", tzWithDST, false},
		9: {"0000-00-00 00:00:00", tzWithoutDST, false},

		10: {"2009-01-01", tzWithDST, true},
		11: {"2009-01-01", tzWithoutDST, false},
	}

	for index, test := range tests {
		c := Parse(test.input, test.timezone)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsDST(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsZero(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", true},
		1: {"0", true},
		2: {"0000-00-00", true},
		3: {"00:00:00", true},
		4: {"0000-00-00 00:00:00", true},

		5: {"2020-08-05", false},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsZero(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsValid(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsValid(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsInvalid(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", true},
		1: {"0", true},
		2: {"0000-00-00", true},
		3: {"00:00:00", true},
		4: {"0000-00-00 00:00:00", true},

		5: {"2020-08-05", false},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsInvalid(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsNow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    Carbon
		expected bool
	}{
		0: {Parse(""), false},
		1: {Parse("0"), false},
		2: {Parse("0000-00-00"), false},
		3: {Parse("00:00:00"), false},
		4: {Parse("0000-00-00 00:00:00"), false},

		5: {Tomorrow(), false},
		6: {Now(), true},
		7: {Yesterday(), false},
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsNow(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsFuture(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    Carbon
		expected bool
	}{
		0: {Parse(""), false},
		1: {Parse("0"), false},
		2: {Parse("0000-00-00"), false},
		3: {Parse("00:00:00"), false},
		4: {Parse("0000-00-00 00:00:00"), false},

		5: {Tomorrow(), true},
		6: {Now(), false},
		7: {Yesterday(), false},
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsFuture(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsPast(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    Carbon
		expected bool
	}{
		0: {Parse(""), false},
		1: {Parse("0"), false},
		2: {Parse("0000-00-00"), false},
		3: {Parse("00:00:00"), false},
		4: {Parse("0000-00-00 00:00:00"), false},

		5: {Tomorrow(), false},
		6: {Now(), false},
		7: {Yesterday(), true},
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsPast(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsLeapYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2015-01-01", false},
		6:  {"2016-01-01", true},
		7:  {"2017-01-01", false},
		8:  {"2018-01-01", false},
		9:  {"2019-01-01", false},
		10: {"2020-01-01", true},
		11: {"2021-01-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsLeapYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsLongYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2015-01-01", true},
		6:  {"2016-01-01", false},
		7:  {"2017-01-01", false},
		8:  {"2018-01-01", false},
		9:  {"2019-01-01", false},
		10: {"2020-01-01", true},
		11: {"2021-01-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsLongYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsJanuary(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2020-01-01", true},
		6:  {"2020-02-01", false},
		7:  {"2020-03-01", false},
		8:  {"2020-04-01", false},
		9:  {"2020-05-01", false},
		10: {"2020-06-01", false},
		11: {"2020-07-01", false},
		12: {"2020-08-01", false},
		13: {"2020-09-01", false},
		14: {"2020-10-01", false},
		15: {"2020-11-01", false},
		16: {"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsJanuary(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsFebruary(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2020-01-01", false},
		6:  {"2020-02-01", true},
		7:  {"2020-03-01", false},
		8:  {"2020-04-01", false},
		9:  {"2020-05-01", false},
		10: {"2020-06-01", false},
		11: {"2020-07-01", false},
		12: {"2020-08-01", false},
		13: {"2020-09-01", false},
		14: {"2020-10-01", false},
		15: {"2020-11-01", false},
		16: {"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsFebruary(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsMarch(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2020-01-01", false},
		6:  {"2020-02-01", false},
		7:  {"2020-03-01", true},
		8:  {"2020-04-01", false},
		9:  {"2020-05-01", false},
		10: {"2020-06-01", false},
		11: {"2020-07-01", false},
		12: {"2020-08-01", false},
		13: {"2020-09-01", false},
		14: {"2020-10-01", false},
		15: {"2020-11-01", false},
		16: {"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsMarch(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsApril(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2020-01-01", false},
		6:  {"2020-02-01", false},
		7:  {"2020-03-01", false},
		8:  {"2020-04-01", true},
		9:  {"2020-05-01", false},
		10: {"2020-06-01", false},
		11: {"2020-07-01", false},
		12: {"2020-08-01", false},
		13: {"2020-09-01", false},
		14: {"2020-10-01", false},
		15: {"2020-11-01", false},
		16: {"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsApril(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsMay(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2020-01-01", false},
		6:  {"2020-02-01", false},
		7:  {"2020-03-01", false},
		8:  {"2020-04-01", false},
		9:  {"2020-05-01", true},
		10: {"2020-06-01", false},
		11: {"2020-07-01", false},
		12: {"2020-08-01", false},
		13: {"2020-09-01", false},
		14: {"2020-10-01", false},
		15: {"2020-11-01", false},
		16: {"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsMay(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsJune(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2020-01-01", false},
		6:  {"2020-02-01", false},
		7:  {"2020-03-01", false},
		8:  {"2020-04-01", false},
		9:  {"2020-05-01", false},
		10: {"2020-06-01", true},
		11: {"2020-07-01", false},
		12: {"2020-08-01", false},
		13: {"2020-09-01", false},
		14: {"2020-10-01", false},
		15: {"2020-11-01", false},
		16: {"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsJune(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsJuly(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2020-01-01", false},
		6:  {"2020-02-01", false},
		7:  {"2020-03-01", false},
		8:  {"2020-04-01", false},
		9:  {"2020-05-01", false},
		10: {"2020-06-01", false},
		11: {"2020-07-01", true},
		12: {"2020-08-01", false},
		13: {"2020-09-01", false},
		14: {"2020-10-01", false},
		15: {"2020-11-01", false},
		16: {"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsJuly(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsAugust(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2020-01-01", false},
		6:  {"2020-02-01", false},
		7:  {"2020-03-01", false},
		8:  {"2020-04-01", false},
		9:  {"2020-05-01", false},
		10: {"2020-06-01", false},
		11: {"2020-07-01", false},
		12: {"2020-08-01", true},
		13: {"2020-09-01", false},
		14: {"2020-10-01", false},
		15: {"2020-11-01", false},
		16: {"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsAugust(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSeptember(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2020-01-01", false},
		6:  {"2020-02-01", false},
		7:  {"2020-03-01", false},
		8:  {"2020-04-01", false},
		9:  {"2020-05-01", false},
		10: {"2020-06-01", false},
		11: {"2020-07-01", false},
		12: {"2020-08-01", false},
		13: {"2020-09-01", true},
		14: {"2020-10-01", false},
		15: {"2020-11-01", false},
		16: {"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsSeptember(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsOctober(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2020-01-01", false},
		6:  {"2020-02-01", false},
		7:  {"2020-03-01", false},
		8:  {"2020-04-01", false},
		9:  {"2020-05-01", false},
		10: {"2020-06-01", false},
		11: {"2020-07-01", false},
		12: {"2020-08-01", false},
		13: {"2020-09-01", false},
		14: {"2020-10-01", true},
		15: {"2020-11-01", false},
		16: {"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsOctober(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsNovember(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2020-01-01", false},
		6:  {"2020-02-01", false},
		7:  {"2020-03-01", false},
		8:  {"2020-04-01", false},
		9:  {"2020-05-01", false},
		10: {"2020-06-01", false},
		11: {"2020-07-01", false},
		12: {"2020-08-01", false},
		13: {"2020-09-01", false},
		14: {"2020-10-01", false},
		15: {"2020-11-01", true},
		16: {"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsNovember(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsDecember(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2020-01-01", false},
		6:  {"2020-02-01", false},
		7:  {"2020-03-01", false},
		8:  {"2020-04-01", false},
		9:  {"2020-05-01", false},
		10: {"2020-06-01", false},
		11: {"2020-07-01", false},
		12: {"2020-08-01", false},
		13: {"2020-09-01", false},
		14: {"2020-10-01", false},
		15: {"2020-11-01", false},
		16: {"2020-12-01", true},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsDecember(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsMonday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2020-10-05", true},
		6:  {"2020-10-06", false},
		7:  {"2020-10-07", false},
		8:  {"2020-10-08", false},
		9:  {"2020-10-09", false},
		10: {"2020-10-10", false},
		11: {"2020-10-11", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsMonday(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsTuesday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2020-10-05", false},
		6:  {"2020-10-06", true},
		7:  {"2020-10-07", false},
		8:  {"2020-10-08", false},
		9:  {"2020-10-09", false},
		10: {"2020-10-10", false},
		11: {"2020-10-11", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsTuesday(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsWednesday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2020-10-05", false},
		6:  {"2020-10-06", false},
		7:  {"2020-10-07", true},
		8:  {"2020-10-08", false},
		9:  {"2020-10-09", false},
		10: {"2020-10-10", false},
		11: {"2020-10-11", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsWednesday(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsThursday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2020-10-05", false},
		6:  {"2020-10-06", false},
		7:  {"2020-10-07", false},
		8:  {"2020-10-08", true},
		9:  {"2020-10-09", false},
		10: {"2020-10-10", false},
		11: {"2020-10-11", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsThursday(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsFriday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2020-10-05", false},
		6:  {"2020-10-06", false},
		7:  {"2020-10-07", false},
		8:  {"2020-10-08", false},
		9:  {"2020-10-09", true},
		10: {"2020-10-10", false},
		11: {"2020-10-11", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsFriday(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSaturday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2020-10-05", false},
		6:  {"2020-10-06", false},
		7:  {"2020-10-07", false},
		8:  {"2020-10-08", false},
		9:  {"2020-10-09", false},
		10: {"2020-10-10", true},
		11: {"2020-10-11", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsSaturday(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSunday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2020-10-05", false},
		6:  {"2020-10-06", false},
		7:  {"2020-10-07", false},
		8:  {"2020-10-08", false},
		9:  {"2020-10-09", false},
		10: {"2020-10-10", false},
		11: {"2020-10-11", true},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsSunday(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsWeekday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2020-10-05", true},
		6:  {"2020-10-06", true},
		7:  {"2020-10-07", true},
		8:  {"2020-10-08", true},
		9:  {"2020-10-09", true},
		10: {"2020-10-10", false},
		11: {"2020-10-11", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsWeekday(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsWeekend(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5:  {"2020-10-05", false},
		6:  {"2020-10-06", false},
		7:  {"2020-10-07", false},
		8:  {"2020-10-08", false},
		9:  {"2020-10-09", false},
		10: {"2020-10-10", true},
		11: {"2020-10-11", true},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsWeekend(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsYesterday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    Carbon
		expected bool
	}{
		0: {NewCarbon(), false},
		1: {Now(), false},
		2: {Yesterday(), true},
		3: {Tomorrow(), false},
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsYesterday(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsToday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    Carbon
		expected bool
	}{
		0: {NewCarbon(), false},
		1: {Now(), true},
		2: {Yesterday(), false},
		3: {Tomorrow(), false},
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsToday(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsTomorrow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    Carbon
		expected bool
	}{
		0: {NewCarbon(), false},
		1: {Now(), false},
		2: {Yesterday(), false},
		3: {Tomorrow(), true},
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsTomorrow(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSameCentury(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   Carbon
		input2   Carbon
		expected bool
	}{
		0: {Parse(""), Parse(""), false},
		1: {Parse("2020-08-05"), Parse("3020-08-05"), false},
		2: {Parse("2020-08-05"), Parse("2099-08-05"), true},
	}

	for index, test := range tests {
		assert.Nil(test.input1.Error)
		assert.Nil(test.input2.Error)
		assert.Equal(test.expected, test.input1.IsSameCentury(test.input2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSameDecade(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   Carbon
		input2   Carbon
		expected bool
	}{
		0: {Parse(""), Parse(""), false},
		1: {Parse("2020-08-05"), Parse("2030-08-05"), false},
		2: {Parse("2020-08-05"), Parse("2021-08-05"), true},
		3: {Parse("2020-01-01"), Parse("2120-01-31"), true},
	}

	for index, test := range tests {
		assert.Nil(test.input1.Error)
		assert.Nil(test.input2.Error)
		assert.Equal(test.expected, test.input1.IsSameDecade(test.input2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSameYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   Carbon
		input2   Carbon
		expected bool
	}{
		0: {Parse(""), Parse(""), false},
		1: {Parse("2020-08-05"), Parse("2021-08-05"), false},
		2: {Parse("2020-01-01"), Parse("2020-12-31"), true},
	}

	for index, test := range tests {
		assert.Nil(test.input1.Error)
		assert.Nil(test.input2.Error)
		assert.Equal(test.expected, test.input1.IsSameYear(test.input2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSameQuarter(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   Carbon
		input2   Carbon
		expected bool
	}{
		0: {Parse(""), Parse(""), false},
		1: {Parse("2020-08-05"), Parse("2020-01-05"), false},
		2: {Parse("2020-01-01"), Parse("2020-01-31"), true},
	}

	for index, test := range tests {
		assert.Nil(test.input1.Error)
		assert.Nil(test.input2.Error)
		assert.Equal(test.expected, test.input1.IsSameQuarter(test.input2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSameMonth(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   Carbon
		input2   Carbon
		expected bool
	}{
		0: {Parse(""), Parse(""), false},
		1: {Parse("2020-08-05"), Parse("2021-08-05"), false},
		2: {Parse("2020-01-01"), Parse("2020-01-31"), true},
	}

	for index, test := range tests {
		assert.Nil(test.input1.Error)
		assert.Nil(test.input2.Error)
		assert.Equal(test.expected, test.input1.IsSameMonth(test.input2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSameDay(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   Carbon
		input2   Carbon
		expected bool
	}{
		0: {Parse(""), Parse(""), false},
		1: {Parse("2020-08-05 13:14:15"), Parse("2021-08-05 13:14:15"), false},
		2: {Parse("2020-08-05 00:00:00"), Parse("2020-08-05 13:14:15"), true},
	}

	for index, test := range tests {
		assert.Nil(test.input1.Error)
		assert.Nil(test.input2.Error)
		assert.Equal(test.expected, test.input1.IsSameDay(test.input2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSameHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   Carbon
		input2   Carbon
		expected bool
	}{
		0: {Parse(""), Parse(""), false},
		1: {Parse("2020-08-05 13:14:15"), Parse("2021-08-05 13:14:15"), false},
		2: {Parse("2020-08-05 13:00:00"), Parse("2020-08-05 13:14:15"), true},
	}

	for index, test := range tests {
		assert.Nil(test.input1.Error)
		assert.Nil(test.input2.Error)
		assert.Equal(test.expected, test.input1.IsSameHour(test.input2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSameMinute(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   Carbon
		input2   Carbon
		expected bool
	}{
		0: {Parse(""), Parse(""), false},
		1: {Parse("2020-08-05 13:14:15"), Parse("2021-08-05 13:14:15"), false},
		2: {Parse("2020-08-05 13:14:00"), Parse("2020-08-05 13:14:15"), true},
	}

	for index, test := range tests {
		assert.Nil(test.input1.Error)
		assert.Nil(test.input2.Error)
		assert.Equal(test.expected, test.input1.IsSameMinute(test.input2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSameSecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   Carbon
		input2   Carbon
		expected bool
	}{
		0: {Parse(""), Parse(""), false},
		1: {Parse("2020-08-05 13:14:15"), Parse("2021-08-05 13:14:15"), false},
		2: {Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), true},
	}

	for index, test := range tests {
		assert.Nil(test.input1.Error)
		assert.Nil(test.input2.Error)
		assert.Equal(test.expected, test.input1.IsSameSecond(test.input2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Compare(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param1   string
		param2   string
		expected bool
	}{
		0:  {"2020-08-05", ">", "2020-08-04", true},
		1:  {"2020-08-05", "<", "2020-08-04", false},
		2:  {"2020-08-05", "<", "2020-08-06", true},
		3:  {"2020-08-05", ">", "2020-08-06", false},
		4:  {"2020-08-05", "=", "2020-08-05", true},
		5:  {"2020-08-05", ">=", "2020-08-05", true},
		6:  {"2020-08-05", "<=", "2020-08-05", true},
		7:  {"2020-08-05", "!=", "2020-08-05", false},
		8:  {"2020-08-05", "<>", "2020-08-05", false},
		9:  {"2020-08-05", "!=", "2020-08-04", true},
		10: {"2020-08-05", "<>", "2020-08-04", true},
		11: {"2020-08-05", "+", "2020-08-04", false},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.Compare(test.param1, c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Gt(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param    string
		expected bool
	}{
		0: {"2020-08-05", "2020-08-05", false},
		1: {"2020-08-05", "2020-08-04", true},
		2: {"2020-08-05", "2020-08-06", false},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.Gt(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Lt(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param    string
		expected bool
	}{
		0: {"2020-08-05", "2020-08-05", false},
		1: {"2020-08-05", "2020-08-04", false},
		2: {"2020-08-05", "2020-08-06", true},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.Lt(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Eq(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param    string
		expected bool
	}{
		0: {"2020-08-05", "2020-08-05", true},
		1: {"2020-08-05", "2020-08-04", false},
		2: {"2020-08-05", "2020-08-06", false},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.Eq(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Ne(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param    string
		expected bool
	}{
		0: {"2020-08-05", "2020-08-05", false},
		1: {"2020-08-05", "2020-08-04", true},
		2: {"2020-08-05", "2020-08-06", true},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.Ne(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Gte(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param    string
		expected bool
	}{
		0: {"2020-08-05", "2020-08-05", true},
		1: {"2020-08-05", "2020-08-04", true},
		2: {"2020-08-05", "2020-08-06", false},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.Gte(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Lte(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param    string
		expected bool
	}{
		0: {"2020-08-05", "2020-08-05", true},
		1: {"2020-08-05", "2020-08-04", false},
		2: {"2020-08-05", "2020-08-06", true},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.Lte(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Between(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param1   string
		param2   string
		expected bool
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-05 13:14:15", false},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-06 13:14:15", false},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-05 13:14:15", false},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-06 13:14:15", true},
	}

	for index, test := range tests {
		c1, c2, c3 := Parse(test.input), Parse(test.param1), Parse(test.param2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Nil(c3.Error)
		assert.Equal(test.expected, c1.Between(c2, c3), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_BetweenIncludedStart(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param1   string
		param2   string
		expected bool
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-05 13:14:15", false},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-06 13:14:15", true},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-05 13:14:15", false},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-06 13:14:15", true},
	}

	for index, test := range tests {
		c1, c2, c3 := Parse(test.input), Parse(test.param1), Parse(test.param2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Nil(c3.Error)
		assert.Equal(test.expected, c1.BetweenIncludedStart(c2, c3), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_BetweenIncludedEnd(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param1   string
		param2   string
		expected bool
	}{
		0: {"2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-05 13:14:15", false},
		1: {"2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-06 13:14:15", false},
		2: {"2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-05 13:14:15", true},
		3: {"2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-06 13:14:15", true},
	}

	for index, test := range tests {
		c1, c2, c3 := Parse(test.input), Parse(test.param1), Parse(test.param2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Nil(c3.Error)
		assert.Equal(test.expected, c1.BetweenIncludedEnd(c2, c3), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_BetweenIncludedBoth(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param1   string
		param2   string
		expected bool
	}{
		0: {"2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-05 13:14:15", true},
		1: {"2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-06 13:14:15", true},
		2: {"2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-05 13:14:15", true},
		3: {"2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-06 13:14:15", true},
		4: {"2020-08-05 13:14:15", "2020-08-06 13:14:15", "2020-08-06 13:14:15", false},
	}

	for index, test := range tests {
		c1, c2, c3 := Parse(test.input), Parse(test.param1), Parse(test.param2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Nil(c3.Error)
		assert.Equal(test.expected, c1.BetweenIncludedBoth(c2, c3), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_Comparer(t *testing.T) {
	time1, time2, time3, operator := "2020-13-50", "xxx", "xxx", ">"
	assert.True(t, Parse(time1).IsZero(), "It should catch an exception in IsZero()")
	assert.True(t, Parse(time1).IsInvalid(), "It should catch an exception in IsInvalid()")
	assert.True(t, Parse(time1).IsInvalid(), "It should catch an exception in IsInvalid()")
	assert.True(t, Parse(time1).Ne(Parse(time2)), "It should catch an exception in Ne()")
	assert.False(t, Parse(time1).Compare(operator, Parse(time2)), "It should catch an exception in Compare()")
	assert.False(t, Parse(time1).Gt(Parse(time2)), "It should catch an exception in Gt()")
	assert.False(t, Parse(time1).Lt(Parse(time2)), "It should catch an exception in Lt()")
	assert.False(t, Parse(time1).Eq(Parse(time2)), "It should catch an exception in Eq()")
	assert.False(t, Parse(time1).Gte(Parse(time2)), "It should catch an exception in Gte()")
	assert.False(t, Parse(time1).Lte(Parse(time2)), "It should catch an exception in Lte()")
	assert.False(t, Parse(time1).Between(Parse(time2), Parse(time3)), "It should catch an exception in Between()")
	assert.False(t, Parse(time1).BetweenIncludedStart(Parse(time2), Parse(time3)), "It should catch an exception in BetweenIncludedStart()")
	assert.False(t, Parse(time1).BetweenIncludedEnd(Parse(time2), Parse(time3)), "It should catch an exception in BetweenIncludedEnd()")
	assert.False(t, Parse(time1).BetweenIncludedBoth(Parse(time2), Parse(time3)), "It should catch an exception in BetweenIncludedBoth()")
}
