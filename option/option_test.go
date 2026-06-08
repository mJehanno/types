package option_test

import (
	"testing"

	"github.com/mJehanno/types/option"
)

func Test_UnwrapOrDefault(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		def      any
		expected any
	}{
		{
			name:     "unwrap some int",
			value:    7,
			expected: 7,
		},
		{
			name:     "unwrap some string",
			value:    "hello",
			expected: "hello",
		},
		{
			name: "unwrap some complex type",
			value: struct {
				x int
				y float64
				z string
			}{
				x: 7,
				y: 5.6,
				z: "hello",
			},
			expected: struct {
				x int
				y float64
				z string
			}{
				x: 7,
				y: 5.6,
				z: "hello",
			},
		},
		{
			name:     "unwrap none string",
			value:    nil,
			def:      "default value",
			expected: "default value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != nil {
				v := option.NewSome(tt.value)
				got := v.UnwrapOrDefault(10)
				if got != tt.expected {
					t.Errorf("test %q failed, got %v, expected %v", tt.name, got, tt.expected)
				}
			} else {
				v := option.NewNone[any]()
				got := v.UnwrapOrDefault(tt.def)
				if got != tt.expected {
					t.Errorf("test %q failed, got %v, expected %v", tt.name, got, tt.expected)
				}
			}
		})
	}
}

func Test_UnwrapOrElse(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		def      func() any
		expected any
	}{
		{
			name:     "unwrap some int",
			value:    7,
			expected: 7,
		},
		{
			name:     "unwrap some string",
			value:    "hello",
			expected: "hello",
		},
		{
			name: "unwrap some complex type",
			value: struct {
				x int
				y float64
				z string
			}{
				x: 7,
				y: 5.6,
				z: "hello",
			},
			expected: struct {
				x int
				y float64
				z string
			}{
				x: 7,
				y: 5.6,
				z: "hello",
			},
		},
		{
			name:     "unwrap none string",
			value:    nil,
			def:      func() any { return "default value" },
			expected: "default value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != nil {
				v := option.NewSome(tt.value)
				got := v.UnwrapOrElse(func() any { return 10 })
				if got != tt.expected {
					t.Errorf("test %q failed, got %v, expected %v", tt.name, got, tt.expected)
				}
			} else {
				v := option.NewNone[any]()
				got := v.UnwrapOrElse(tt.def)
				if got != tt.expected {
					t.Errorf("test %q failed, got %v, expected %v", tt.name, got, tt.expected)
				}
			}
		})
	}
}
