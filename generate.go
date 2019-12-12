package jsonschema2go

import (
	"context"
	"github.com/jwilner/jsonschema2go/internal/cachingloader"
	"github.com/jwilner/jsonschema2go/internal/crawl"
	"github.com/jwilner/jsonschema2go/internal/planning"
	"github.com/jwilner/jsonschema2go/internal/print"
	"github.com/jwilner/jsonschema2go/pkg/gen"
	"text/template"
)

// Generate generates Go source code from the provided JSON schemas. Options can be provided to customize the
// output behavior
func Generate(ctx context.Context, uris []string, options ...Option) error {
	s := &settings{
		planner: planning.Composite,
		printer: print.New(nil),
		typer:   planning.DefaultTyper,
	}
	for _, o := range options {
		o(s)
	}

	if s.loader == nil {
		c := cachingloader.New(s.debug)
		defer func() {
			_ = c.Close()
		}()
		s.loader = c
	}
	ctx, cncl := context.WithCancel(ctx)
	defer cncl()

	if s.debug {
		ctx = gen.SetDebug(ctx)
	}

	grouped, err := crawl.Crawl(ctx, s.planner, s.loader, s.typer, uris)
	if err != nil {
		return err
	}

	return print.Print(ctx, s.printer, grouped, s.prefixes)
}

// Option controls the behavior of jsonschema2go, specifying an alternative to the default configuration
type Option func(s *settings)

// PrefixMap specifies where package prefixes should be mapped to.
func PrefixMap(pairs ...string) Option {
	prefixes := prefixPairs(pairs)
	return func(s *settings) {
		s.prefixes = prefixes
	}
}

// Debug enables debug logging
func Debug(opt bool) Option {
	return func(s *settings) {
		s.debug = opt
	}
}

// CustomTypeFunc registers a custom function for generating TypeInfo from a Schema.
func CustomTypeFunc(typeFunc func(schema *gen.Schema) gen.TypeInfo) Option {
	return func(s *settings) {
		s.typer.TypeFunc = typeFunc
	}
}

// CustomPrimitivesMap registers a custom map for mapping from JSONSchema simple types to go primitives.
func CustomPrimitivesMap(primitivesMap map[gen.SimpleType]string) Option {
	return func(s *settings) {
		s.typer.Primitives = primitivesMap
	}
}

// CustomPlanners permits providing an entirely custom list of planners, which will be jointed together.
func CustomPlanners(planners ...gen.Planner) Option {
	return func(s *settings) {
		s.planner = planning.CompositePlanner(planners)
	}
}

// TypeFromID defines how to map to type information from IDs
func TypeFromID(pairs ...string) Option {
	mapper := print.TypeFromId(prefixPairs(pairs))
	return func(s *settings) {
		s.typer.TypeFunc = func(schema *gen.Schema) gen.TypeInfo {
			if t := planning.DefaultTypeFunc(schema); !t.Unknown() {
				return t
			}
			if schema.CalcID != nil {
				if path, name := mapper(schema.CalcID.String()); name != "" {
					return gen.TypeInfo{GoPath: path, Name: name}
				}
			}
			return gen.TypeInfo{}
		}
	}
}

// CustomTemplate registers a custom top level template
func CustomTemplate(tmpl *template.Template) Option {
	return func(s *settings) {
		s.printer = print.New(tmpl)
	}
}

// CustomInitialisms registers a custom list of initialisms used in name generation
func CustomInitialisms(names ...string) Option {
	return func(s *settings) {
		s.typer.Namer = planning.NewNamer(append(names, "id", "http"))
	}
}

func prefixPairs(pairs []string) [][2]string {
	if len(pairs)%2 != 0 {
		panic("must be even list of prefixes")
	}
	var prefixes [][2]string
	for i := 0; i < len(pairs); i += 2 {
		prefixes = append(prefixes, [2]string{pairs[i], pairs[i+1]})
	}
	return prefixes
}

type settings struct {
	prefixes [][2]string
	typer    planning.Typer
	planner  gen.Planner
	printer  print.Printer
	loader   gen.Loader
	debug    bool
}