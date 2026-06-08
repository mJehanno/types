package result_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mJehanno/types/option"
	"github.com/mJehanno/types/result"
)

func Test_UnwrapOrDefault(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		def      any
		err      error
		expected any
	}{
		{
			name:     "unwrap ok int",
			value:    7,
			expected: 7,
		},
		{
			name:     "unwrap ok string",
			value:    "hello",
			expected: "hello",
		},
		{
			name: "unwrap ok complex type",
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
			name:     "unwrap err string",
			err:      errors.New("boom"),
			def:      "default value",
			expected: "default value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err == nil {
				v := result.NewOk(tt.value)
				got := v.UnwrapOrDefault(10)
				if got != tt.expected {
					t.Errorf("test %q failed, got %v, expected %v", tt.name, got, tt.expected)
				}
			} else {
				v := result.NewErr[any](tt.err)
				got := v.UnwrapOrDefault(tt.def)
				t.Logf("%#v \n", v)
				if got != tt.expected {
					t.Errorf("test %q failed, got %v, expected %v", tt.name, got, tt.expected)
				}
				if v.GetError() != tt.err {
					t.Errorf("test %q `GetError` failed, got %v, expected %v", tt.name, v.GetError(), tt.err)
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
		err      error
		expected any
	}{
		{
			name:     "unwrap ok int",
			value:    7,
			expected: 7,
		},
		{
			name:     "unwrap ok string",
			value:    "hello",
			expected: "hello",
		},
		{
			name: "unwrap ok complex type",
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
			name:     "unwrap err string",
			err:      errors.New("boom"),
			def:      func() any { return "default value" },
			expected: "default value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err == nil {
				v := result.NewOk(tt.value)
				got := v.UnwrapOrElse(func() any { return 10 })
				if got != tt.expected {
					t.Errorf("test %q failed, got %v, expected %v", tt.name, got, tt.expected)
				}
			} else {
				v := result.NewErr[any](tt.err)
				got := v.UnwrapOrElse(tt.def)
				if got != tt.expected {
					t.Errorf("test %q failed, got %v, expected %v", tt.name, got, tt.expected)
				}

				if v.GetError() != tt.err {
					t.Errorf("test %q `GetError` failed, got %v, expected %v", tt.name, v.GetError(), tt.err)
				}

			}
		})
	}
}

// Examples
func ExampleSome_UnwrapOrDefault() {
	some := option.NewSome(7)

	fmt.Println(some.UnwrapOrDefault(10))
	// Output: 7
}

func ExampleNone_UnwrapOrDefault() {
	some := option.NewNone[int]()

	fmt.Println(some.UnwrapOrDefault(10))
	// Output: 10
}

func ExampleSome_UnwrapOrElse() {
	some := option.NewSome(7)

	fmt.Println(some.UnwrapOrElse(func() int { return 9 }))
	// Output: 7
}

func ExampleNone_UnwrapOrElse() {
	some := option.NewNone[int]()

	fmt.Println(some.UnwrapOrElse(func() int { return 10 }))
	// Output: 10
}
