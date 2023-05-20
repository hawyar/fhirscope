package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Context int
type Operation int

const (
	// declare with iota so that the values are auto incremented
	PatientContext Context = iota
	UserContext
	SystemContext

	CreateOperation Operation = iota
	ReadOperation
	UpdateOperation
	DeleteOperation
	SearchOperation
)

type Scope struct {
	Context    Context           `json:"context"`
	Operations []Operation       `json:"operations"`
	Resource   string            `json:"resource"`
	Params     map[string]string `json:"params"`
}

func main() {
	version := "0.0.1"
	usage := `Usage:
fhirscope <scope>
e.g. fhirscope patient/Observation.rs
	`

	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	raw := os.Args[1]

	if raw == "-h" || raw == "--help" {
		fmt.Println(usage)
		os.Exit(0)
	}

	if raw == "-v" || raw == "--version" {
		fmt.Println(version)
		os.Exit(0)
	}

	scope, err := Parse(raw)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonScope, err := json.Marshal(scope)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(jsonScope))
}

func Parse(scope string) (Scope, error) {
	parts := strings.Split(scope, "/")

	if len(parts) < 2 {
		return Scope{}, fmt.Errorf("invalid scope: %s", scope)
	}

	out := Scope{}

	ctx := parts[0]

	switch ctx {
	case "patient":
		out.Context = PatientContext
	case "user":
		out.Context = UserContext
	case "system":
		out.Context = SystemContext
	default:
		return Scope{}, fmt.Errorf("invalid context expected patient, user or system got: %s", ctx)
	}

	rest := strings.Split(parts[1], ".")

	if len(rest) < 2 {
		return Scope{}, fmt.Errorf("invalid resource or operation: %s", parts[1])
	}

	out.Resource = rest[0]

	restWithParams := strings.Split(rest[1], "?")

	if len(restWithParams) > 1 {
		fmt.Println("params", restWithParams[1])
		// pairs := strings.Split(params[1], "&")
	}

	ops := strings.Split(restWithParams[0], "")

	visited := make(map[string]bool)

	for _, op := range ops {
		if visited[op] {
			return Scope{}, fmt.Errorf("duplicate operation: %s", op)
		}
		switch op {
		case "c":
			out.Operations = append(out.Operations, CreateOperation)
		case "r":
			out.Operations = append(out.Operations, ReadOperation)
		case "u":
			out.Operations = append(out.Operations, UpdateOperation)
		case "d":
			out.Operations = append(out.Operations, DeleteOperation)
		case "s":
			out.Operations = append(out.Operations, SearchOperation)
		default:
			return Scope{}, fmt.Errorf("invalid operation expected c, r, u, d or s got: %s", op)
		}
	}

	return out, nil
}

func (c Context) MarshalJSON() ([]byte, error) {
	value := c.String()
	return json.Marshal(value)
}

func (o Operation) MarshalJSON() ([]byte, error) {
	value := o.String()
	return json.Marshal(value)
}

func (c Context) String() string {
	switch c {
	case PatientContext:
		return "patient"
	case UserContext:
		return "user"
	case SystemContext:
		return "system"
	default:
		return ""
	}
}

func (o Operation) String() string {
	switch o {
	case CreateOperation:
		return "c"
	case ReadOperation:
		return "r"
	case UpdateOperation:
		return "u"
	case DeleteOperation:
		return "d"
	case SearchOperation:
		return "s"
	default:
		return ""
	}
}
