package diagnosticism_test

import (
	d "github.com/synesissoftware/Diagnosticism.Go"

	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ItoaThousands_int(t *testing.T) {

	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{
			name:     "0",
			input:    0,
			expected: "0",
		},
		{
			name:     "1",
			input:    1,
			expected: "1",
		},
		{
			name:     "123",
			input:    123,
			expected: "123",
		},
		{
			name:     "1234",
			input:    1_234,
			expected: "1,234",
		},
		{
			name:     "12345",
			input:    12_345,
			expected: "12,345",
		},
		{
			name:     "123456",
			input:    123_456,
			expected: "123,456",
		},
		{
			name:     "1234567",
			input:    1_234_567,
			expected: "1,234,567",
		},
		{
			name:     "-1",
			input:    -1,
			expected: "-1",
		},
		{
			name:     "-123",
			input:    -123,
			expected: "-123",
		},
		{
			name:     "-1234",
			input:    -1_234,
			expected: "-1,234",
		},
		{
			name:     "-12345",
			input:    -12_345,
			expected: "-12,345",
		},
		{
			name:     "-123456",
			input:    -123_456,
			expected: "-123,456",
		},
		{
			name:     "-1234567",
			input:    -1_234_567,
			expected: "-1,234,567",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			actual := d.ItoaThousands(tt.input)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_ItoaThousands_int64(t *testing.T) {

	tests := []struct {
		name     string
		input    int64
		expected string
	}{
		{
			name:     "0",
			input:    0,
			expected: "0",
		},
		{
			name:     "1",
			input:    1,
			expected: "1",
		},
		{
			name:     "123",
			input:    123,
			expected: "123",
		},
		{
			name:     "1234",
			input:    1_234,
			expected: "1,234",
		},
		{
			name:     "12345",
			input:    12_345,
			expected: "12,345",
		},
		{
			name:     "123456",
			input:    123_456,
			expected: "123,456",
		},
		{
			name:     "1234567",
			input:    1_234_567,
			expected: "1,234,567",
		},
		{
			name:     "-1",
			input:    -1,
			expected: "-1",
		},
		{
			name:     "-123",
			input:    -123,
			expected: "-123",
		},
		{
			name:     "-1234",
			input:    -1_234,
			expected: "-1,234",
		},
		{
			name:     "-12345",
			input:    -12_345,
			expected: "-12,345",
		},
		{
			name:     "-123456",
			input:    -123_456,
			expected: "-123,456",
		},
		{
			name:     "-1234567",
			input:    -1_234_567,
			expected: "-1,234,567",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			actual := d.ItoaThousands(tt.input)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_ItoaThousands_uint32(t *testing.T) {

	tests := []struct {
		name     string
		input    uint32
		expected string
	}{
		{
			name:     "0",
			input:    0,
			expected: "0",
		},
		{
			name:     "1",
			input:    1,
			expected: "1",
		},
		{
			name:     "123",
			input:    123,
			expected: "123",
		},
		{
			name:     "1234",
			input:    1_234,
			expected: "1,234",
		},
		{
			name:     "12345",
			input:    12_345,
			expected: "12,345",
		},
		{
			name:     "123456",
			input:    123_456,
			expected: "123,456",
		},
		{
			name:     "1234567",
			input:    1_234_567,
			expected: "1,234,567",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			actual := d.ItoaThousands(tt.input)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
