package testutil

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/parser"
)

var (
	Anderson  Employer
	Kate      Employer
	Suzan     Employer
	Jakob     Employer
	Fathe     Employer
	Rashid    Employer
	Elizabeth Employer

	//Department map[int]Department
	//Company    map[int]Company
	EmployerSchema graphql.Schema

	//Departments *graphql.Object
	//Companys    *graphql.Object
)

type Company struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}
type Department struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
	Type string `json:"type"`
}
type Employer struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Gender    string `json:"gender"`
	Position  string `json:"position"`
}

func init() {
	Anderson = Employer{
		ID:        100,
		FirstName: "Anderson",
		LastName:  "Stone",
		Gender:    "male",
		Position:  "Boss",
	}
	Kate = Employer{
		ID:        101,
		FirstName: "Kate",
		LastName:  "Lakritz",
		Gender:    "female",
		Position:  "Secreataty",
	}
	Suzan = Employer{
		ID:        102,
		FirstName: "Suzan",
		LastName:  "Berke",
		Gender:    "female",
		Position:  "Frontend",
	}
	Jakob = Employer{
		ID:        103,
		FirstName: "Jacob",
		LastName:  "Baloon",
		Gender:    "male",
		Position:  "backend",
	}
	Fathe = Employer{
		ID:        104,
		FirstName: "Fathe",
		LastName:  "Snow",
		Gender:    "male",
		Position:  "tester",
	}
	Rashid = Employer{
		ID:        105,
		FirstName: "Rahid",
		LastName:  "Spark",
		Gender:    "male",
		Position:  "tester",
	}
	Elizabeth = Employer{
		ID:        106,
		FirstName: "Elizabeth",
		LastName:  "Scram",
		Gender:    "female",
		Position:  "backend",
	}
	var employers = []Employer{Anderson, Kate, Suzan, Jakob, Fathe, Rashid, Elizabeth}

	var EmployerType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Employer",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"FirstName": &graphql.Field{
					Type: graphql.String,
				},
				"LastName": &graphql.Field{
					Type: graphql.String,
				},
				"Gender": &graphql.Field{
					Type: graphql.String,
				},
				"Position": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"employer": &graphql.Field{
				Type: EmployerType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "Employer id",
						Type:        graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						for _, employer := range employers {
							if int(employer.ID) == id {
								return employer, nil
							}
						}
					}
					return nil, nil
				},
			},
		},
	})
	EmployerSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}

// Test helper functions

func TestParse(t *testing.T, query string) *ast.Document {
	astDoc, err := parser.Parse(parser.ParseParams{
		Source: query,
		Options: parser.ParseOptions{
			// include source, for error reporting
			NoSource: false,
		},
	})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}
	return astDoc
}
func TestExecute(t *testing.T, ep graphql.ExecuteParams) *graphql.Result {
	return graphql.Execute(ep)
}

func Diff(want, got interface{}) []string {
	return []string{fmt.Sprintf("\ngot: %v", got), fmt.Sprintf("\nwant: %v\n", want)}
}

func ASTToJSON(t *testing.T, a ast.Node) interface{} {
	b, err := json.Marshal(a)
	if err != nil {
		t.Fatalf("Failed to marshal Node %v", err)
	}
	var f interface{}
	err = json.Unmarshal(b, &f)
	if err != nil {
		t.Fatalf("Failed to unmarshal Node %v", err)
	}
	return f
}

func ContainSubsetSlice(super []interface{}, sub []interface{}) bool {
	if len(sub) == 0 {
		return true
	}
subLoop:
	for _, subVal := range sub {
		found := false
	innerLoop:
		for _, superVal := range super {
			if subVal, ok := subVal.(map[string]interface{}); ok {
				if superVal, ok := superVal.(map[string]interface{}); ok {
					if ContainSubset(superVal, subVal) {
						found = true
						break innerLoop
					} else {
						continue
					}
				} else {
					return false
				}

			}
			if subVal, ok := subVal.([]interface{}); ok {
				if superVal, ok := superVal.([]interface{}); ok {
					if ContainSubsetSlice(superVal, subVal) {
						found = true
						break innerLoop
					} else {
						continue
					}
				} else {
					return false
				}
			}
			if reflect.DeepEqual(superVal, subVal) {
				found = true
				break innerLoop
			}
		}
		if !found {
			return false
		}
		continue subLoop
	}
	return true
}

func ContainSubset(super map[string]interface{}, sub map[string]interface{}) bool {
	if len(sub) == 0 {
		return true
	}
	for subKey, subVal := range sub {
		if superVal, ok := super[subKey]; ok {
			switch superVal := superVal.(type) {
			case []interface{}:
				if subVal, ok := subVal.([]interface{}); ok {
					if !ContainSubsetSlice(superVal, subVal) {
						return false
					}
				} else {
					return false
				}
			case map[string]interface{}:
				if subVal, ok := subVal.(map[string]interface{}); ok {
					if !ContainSubset(superVal, subVal) {
						return false
					}
				} else {
					return false
				}
			default:
				if !reflect.DeepEqual(superVal, subVal) {
					return false
				}
			}
		} else {
			return false
		}
	}
	return true
}

func EqualErrorMessage(expected, result *graphql.Result, i int) bool {
	return expected.Errors[i].Message == result.Errors[i].Message
}

func EqualFormattedError(exp, act gqlerrors.FormattedError) bool {
	if exp.Message != act.Message {
		return false
	}
	if !reflect.DeepEqual(exp.Locations, act.Locations) {
		return false
	}
	if !reflect.DeepEqual(exp.Path, act.Path) {
		return false
	}
	if !reflect.DeepEqual(exp.Extensions, act.Extensions) {
		return false
	}
	return true
}

func EqualFormattedErrors(expected, actual []gqlerrors.FormattedError) bool {
	if len(expected) != len(actual) {
		return false
	}
	for i := range expected {
		if !EqualFormattedError(expected[i], actual[i]) {
			return false
		}
	}
	return true
}

func EqualResults(expected, result *graphql.Result) bool {
	if !reflect.DeepEqual(expected.Data, result.Data) {
		return false
	}
	return EqualFormattedErrors(expected.Errors, result.Errors)
}
