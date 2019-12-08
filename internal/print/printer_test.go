package print

import (
	"bytes"
	"context"
	"github.com/jwilner/jsonschema2go/internal/composite"
	"github.com/jwilner/jsonschema2go/internal/slice"
	"github.com/jwilner/jsonschema2go/pkg/generate"
	"github.com/stretchr/testify/require"
	"go/format"
	"testing"
)

func TestImports_List(t *testing.T) {
	tests := []struct {
		name          string
		currentGoPath string
		importGoPaths []string
		wantImports   []generate.Import
	}{
		{
			"empty",
			"github.com/jwilner/jsonschema2go",
			[]string{},
			nil,
		},
		{
			"alias",
			"github.com/jwilner/jsonschema2go",
			[]string{
				"github.com/jwilner/jsonschema2go/example",
				"github.com/jwilner/jsonschema2go/foo/example",
			},
			[]generate.Import{
				{"github.com/jwilner/jsonschema2go/example", ""},
				{"github.com/jwilner/jsonschema2go/foo/example", "example2"},
			},
		},
		{
			"multiple",
			"github.com/jwilner/jsonschema2go",
			[]string{"encoding/json", "encoding/json"},
			[]generate.Import{
				{"encoding/json", ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.wantImports, generate.NewImports(tt.currentGoPath, tt.importGoPaths).List())
		})
	}
}

func TestImports_QualName(t *testing.T) {
	tests := []struct {
		name          string
		currentGoPath string
		importGoPaths []string
		typeInfo      generate.TypeInfo
		want          string
	}{
		{
			"builtin",
			"github.com/jwilner/jsonschema2go",
			[]string{"github.com/jwilner/jsonschema2go/example", "github.com/jwilner/jsonschema2go/foo/example"},
			generate.TypeInfo{Name: "int"},
			"int",
		},
		{
			"external",
			"github.com/jwilner/jsonschema2go",
			[]string{"github.com/jwilner/jsonschema2go/example", "github.com/jwilner/jsonschema2go/foo/example"},
			generate.TypeInfo{GoPath: "github.com/jwilner/jsonschema2go", Name: "Bob"},
			"Bob",
		},
		{
			"external",
			"github.com/jwilner/jsonschema2go",
			[]string{"github.com/jwilner/jsonschema2go/example", "github.com/jwilner/jsonschema2go/foo/example"},
			generate.TypeInfo{GoPath: "github.com/jwilner/jsonschema2go/example", Name: "Bob"},
			"example.Bob",
		},
		{
			"external with alias",
			"github.com/jwilner/jsonschema2go",
			[]string{"github.com/jwilner/jsonschema2go/example", "github.com/jwilner/jsonschema2go/foo/example"},
			generate.TypeInfo{GoPath: "github.com/jwilner/jsonschema2go/foo/example", Name: "Bob"},
			"example2.Bob",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, generate.NewImports(tt.currentGoPath, tt.importGoPaths).QualName(tt.typeInfo))
		})
	}
}

func TestPrintFile(t *testing.T) {
	errBit := `
type valErr interface {
	ErrType() string
	JSONPath() []interface{}
	Path() []interface{}
	Message() string
}

type validationError struct {
	errType, message string
	jsonPath, path   []interface{}
}

func (e *validationError) ErrType() string {
	return e.errType
}

func (e *validationError) JSONPath() []interface{} {
	return e.jsonPath
}

func (e *validationError) Path() []interface{} {
	return e.path
}

func (e *validationError) Message() string {
	return e.message
}

func (e *validationError) Error() string {
	return fmt.Sprintf("%v: %v", e.path, e.message)
}

var _ valErr = new(validationError)
`

	tests := []struct {
		name    string
		goPath  string
		plans   []generate.Plan
		wantW   string
		wantErr bool
	}{
		{
			name:   "simple struct",
			goPath: "github.com/jwilner/jsonschema2go",
			plans: []generate.Plan{
				&composite.StructPlan{
					Comment: "Bob does lots of cool things",
					Fields: []composite.StructField{
						{Name: "Count", Type: generate.TypeInfo{Name: "int"}, Tag: `json:"count,omitempty"`},
					},
					TypeInfo: generate.TypeInfo{
						Name: "Bob",
					},
				},
			},
			wantW: `
// Code generated by jsonschema2go. DO NOT EDIT.
package jsonschema2go

import (
	"fmt"
)

// Bob does lots of cool things
type Bob struct {
	Count int ` + "`" + `json:"count,omitempty"` + "`" + `
}

func (m *Bob) Validate() error {
	return nil
}
` + errBit,
		},
		{
			name:   "struct with qualified field",
			goPath: "github.com/jwilner/jsonschema2go",
			plans: []generate.Plan{
				&composite.StructPlan{
					Comment: "Bob does lots of cool things",
					Fields: []composite.StructField{
						{Name: "Count", Type: generate.TypeInfo{Name: "int"}, Tag: `json:"count,omitempty"`},
						{
							Name: "Other",
							Type: generate.TypeInfo{
								GoPath: "github.com/jwilner/jsonschema2go/blah",
								Name:   "OtherType",
							},
							Tag: `json:"other,omitempty"`,
						},
					},
					TypeInfo: generate.TypeInfo{
						Name: "Bob",
					},
				},
			},
			wantW: `
// Code generated by jsonschema2go. DO NOT EDIT.
package jsonschema2go

import (
	"fmt"
	"github.com/jwilner/jsonschema2go/blah"
)

// Bob does lots of cool things
type Bob struct {
	Count int 				` + "`" + `json:"count,omitempty"` + "`" + `
	Other blah.OtherType 	` + "`" + `json:"other,omitempty"` + "`" + `
}

func (m *Bob) Validate() error {
	return nil
}

` + errBit,
		},
		{
			name:   "struct with aliased import",
			goPath: "github.com/jwilner/jsonschema2go",
			plans: []generate.Plan{
				&composite.StructPlan{
					Comment: "Bob does lots of cool things",
					Fields: []composite.StructField{
						{Name: "Count", Type: generate.TypeInfo{Name: "int"}, Tag: `json:"count,omitempty"`},
						{
							Name: "Other",
							Type: generate.TypeInfo{
								GoPath:  "github.com/jwilner/jsonschema2go/blah",
								Name:    "OtherType",
								Pointer: true,
							},
							Tag: `json:"other,omitempty"`,
						},
						{
							Name: "OtherOther",
							Type: generate.TypeInfo{
								GoPath: "github.com/jwilner/jsonschema2go/bob/blah",
								Name:   "AnotherType",
							},
							Tag: `json:"another,omitempty"`,
						},
					},
					TypeInfo: generate.TypeInfo{
						Name: "Bob",
					},
				},
			},
			wantW: `
// Code generated by jsonschema2go. DO NOT EDIT.
package jsonschema2go

import (
	"fmt"
	"github.com/jwilner/jsonschema2go/blah"
	blah2 "github.com/jwilner/jsonschema2go/bob/blah"
)

// Bob does lots of cool things
type Bob struct {
	Count 		int 				` + "`" + `json:"count,omitempty"` + "`" + `
	Other 		*blah.OtherType 	` + "`" + `json:"other,omitempty"` + "`" + `
	OtherOther 	blah2.AnotherType 	` + "`" + `json:"another,omitempty"` + "`" + `
}

func (m *Bob) Validate() error {
	return nil
}

` + errBit,
		},
		{
			name:   "struct with embedded",
			goPath: "github.com/jwilner/jsonschema2go",
			plans: []generate.Plan{
				&composite.StructPlan{
					Comment: "Bob does lots of cool things",
					Fields: []composite.StructField{
						{
							Type: generate.TypeInfo{
								GoPath: "github.com/jwilner/jsonschema2go/blah",
								Name:   "OtherType",
							},
						},
					},
					TypeInfo: generate.TypeInfo{
						Name: "Bob",
					},
				},
			},
			wantW: `
// Code generated by jsonschema2go. DO NOT EDIT.
package jsonschema2go

import (
	"fmt"
	"github.com/jwilner/jsonschema2go/blah"
)

// Bob does lots of cool things
type Bob struct {
	blah.OtherType
}

func (m *Bob) Validate() error {
	return nil
}

` + errBit,
		},
		{
			name:   "struct with embedded",
			goPath: "github.com/jwilner/jsonschema2go",
			plans: []generate.Plan{
				&composite.StructPlan{
					Comment: "Bob does lots of cool things",
					Fields: []composite.StructField{
						{
							Type: generate.TypeInfo{
								GoPath: "github.com/jwilner/jsonschema2go",
								Name:   "OtherType",
							},
						},
					},
					TypeInfo: generate.TypeInfo{
						Name: "Bob",
					},
				},
				&composite.StructPlan{
					Comment: "OtherType does lots of cool things",
					Fields: []composite.StructField{
						{Type: generate.TypeInfo{Name: "int"}, Name: "Count", Tag: `json:"count,omitempty"`},
					},
					TypeInfo: generate.TypeInfo{
						Name: "OtherType",
					},
				},
			},
			wantW: `
// Code generated by jsonschema2go. DO NOT EDIT.
package jsonschema2go

import (
	"fmt"
)

// Bob does lots of cool things
type Bob struct {
	OtherType
}

func (m *Bob) Validate() error {
	return nil
}

// OtherType does lots of cool things
type OtherType struct {
	Count int ` + "`" + `json:"count,omitempty"` + "`" + `
}

func (m *OtherType) Validate() error {
	return nil
}

` + errBit,
		},
		{
			name:   "array with struct",
			goPath: "github.com/jwilner/jsonschema2go",
			plans: []generate.Plan{
				&slice.SlicePlan{
					TypeInfo: generate.TypeInfo{
						Name: "Bob",
					},
					Comment: "Bob does lots of cool things",
					ItemType: generate.TypeInfo{
						GoPath: "github.com/jwilner/jsonschema2go",
						Name:   "OtherType",
					},
				},
				&composite.StructPlan{
					Comment: "OtherType does lots of cool things",
					Fields: []composite.StructField{
						{Type: generate.TypeInfo{Name: "int"}, Name: "Count", Tag: `json:"count,omitempty"`},
					},
					TypeInfo: generate.TypeInfo{
						Name: "OtherType",
					},
				},
			},
			wantW: `
// Code generated by jsonschema2go. DO NOT EDIT.
package jsonschema2go

import (
	"encoding/json"
	"fmt"
)

// OtherType does lots of cool things
type OtherType struct {
	Count int ` + "`" + `json:"count,omitempty"` + "`" + `
}

func (m *OtherType) Validate() error {
	return nil
}

// Bob does lots of cool things
type Bob []OtherType

func (m Bob) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte(` + "`[]`" + `), nil
	}
	return json.Marshal([]OtherType(m))
}

func (m Bob) Validate() error {
	return nil
}

` + errBit,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var w bytes.Buffer
			err := New(nil).Print(context.Background(), &w, tt.goPath, tt.plans)
			if (err != nil) != tt.wantErr {
				t.Fatalf("printStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
			formatted, err := format.Source(w.Bytes())
			if err != nil {
				t.Fatalf("unable to format: %v", err)
			}
			formattedWant, err := format.Source([]byte(tt.wantW))
			if err != nil {
				t.Fatalf("unable to format wanted: %v", err)
			}
			require.Equal(t, string(formattedWant), string(formatted))
		})
	}
}

func Test_mapPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		prefixes [][2]string
		want     string
	}{
		{"empty", "blah", nil, "blah"},
		{"one", "github.com/jsonschema2go/foo/bar", [][2]string{{"github.com/jsonschema2go", "code"}}, "code/foo/bar"},
		{
			"greater",
			"github.com/jsonschema2go/foo/bar",
			[][2]string{{"github.com/jsonschema2go", "code"}, {"github.com/otherpath", "blob"}},
			"code/foo/bar",
		},
		{
			"less",
			"github.com/jsonschema2go/foo/bar",
			[][2]string{{"github.com/a", "other"}, {"github.com/jsonschema2go", "code"}},
			"code/foo/bar",
		},
		{
			"takes longest",
			"github.com/jsonschema2go/foo/bar",
			[][2]string{{"github.com/", "other"}, {"github.com/jsonschema2go", "code"}},
			"code/foo/bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PathMapper(tt.prefixes)(tt.path); got != tt.want {
				t.Errorf("mapPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_typeFromID(t *testing.T) {
	for _, tt := range []struct {
		name                   string
		pairs                  [][2]string
		id, wantPath, wantName string
	}{
		{
			name:     "maps",
			pairs:    [][2]string{{"https://example.com/v1/", "github.com/example/"}},
			id:       "https://example.com/v1/blah/bar.json",
			wantPath: "github.com/example/blah",
			wantName: "bar",
		},
		{
			name:     "maps no extension",
			pairs:    [][2]string{{"https://example.com/v1/", "github.com/example/"}},
			id:       "https://example.com/v1/blah/bar",
			wantPath: "github.com/example/blah",
			wantName: "bar",
		},
		{
			name:     "maps no pairs",
			pairs:    [][2]string{},
			id:       "https://example.com/v1/blah/bar",
			wantPath: "example.com/v1/blah",
			wantName: "bar",
		},
		{
			name:     "maps no scheme",
			pairs:    [][2]string{},
			id:       "example.com/v1/blah/bar",
			wantPath: "example.com/v1/blah",
			wantName: "bar",
		},
		{
			name:     "maps empty fragment",
			pairs:    [][2]string{{"https://example.com/v1/", "github.com/example/"}},
			id:       "https://example.com/v1/blah/bar.json#",
			wantPath: "github.com/example/blah",
			wantName: "bar",
		},
		{
			name:     "maps properties fragment",
			pairs:    [][2]string{{"https://example.com/v1/", "github.com/example/"}},
			id:       "https://example.com/v1/blah/bar.json#/properties/baz",
			wantPath: "github.com/example/blah",
			wantName: "barBaz",
		},
		{
			name:     "maps extended fragment",
			pairs:    [][2]string{{"https://example.com/v1/", "github.com/example/"}},
			id:       "https://example.com/v1/blah/bar.json#/properties/baz/items/2/properties/hello",
			wantPath: "github.com/example/blah",
			wantName: "barBazItems2Hello",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if path, name := TypeFromId(tt.pairs)(tt.id); tt.wantName != name || tt.wantPath != path {
				t.Errorf("wanted (%q, %q) got (%q, %q)", tt.wantPath, tt.wantName, path, name)
			}
		})
	}
}
