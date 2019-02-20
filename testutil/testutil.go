package testutil

import (
	"github.com/graphql-go/graphql"
)

var (
	Anderson  Employer
	Kate      Employer
	Suzan     Employer
	Jakob     Employer
	Fathe     Employer
	Rashid    Employer
	Elizabeth Employer

	EmployerSchema graphql.Schema
)

type Employer struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Gender    string `json:"gender"`
	Position  string `json:"position"`
	Salary    int    `json:"salary"`
}

func init() {
	Anderson = Employer{
		ID:        100,
		FirstName: "Anderson",
		LastName:  "Stone",
		Gender:    "male",
		Position:  "Boss",
		Salary:    4000,
	}
	Kate = Employer{
		ID:        101,
		FirstName: "Kate",
		LastName:  "Lakritz",
		Gender:    "female",
		Position:  "Secreataty",
		Salary:    2000,
	}
	Suzan = Employer{
		ID:        102,
		FirstName: "Suzan",
		LastName:  "Berke",
		Gender:    "female",
		Position:  "Frontend",
		Salary:    3000,
	}
	Jakob = Employer{
		ID:        103,
		FirstName: "Jacob",
		LastName:  "Baloon",
		Gender:    "male",
		Position:  "backend",
		Salary:    3000,
	}
	Fathe = Employer{
		ID:        104,
		FirstName: "Fathe",
		LastName:  "Snow",
		Gender:    "male",
		Position:  "tester",
		Salary:    3000,
	}
	Rashid = Employer{
		ID:        105,
		FirstName: "Rahid",
		LastName:  "Spark",
		Gender:    "male",
		Position:  "tester",
		Salary:    2000,
	}
	Elizabeth = Employer{
		ID:        106,
		FirstName: "Elizabeth",
		LastName:  "Scram",
		Gender:    "female",
		Position:  "backend",
		Salary:    3000,
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
				"Salary": &graphql.Field{
					Type: graphql.Int,
				},
			},
		},
	)
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			/*
				http://localhost:8080/product?query={employer(id:106){id,FirstName,LastName,Gender,Position,Salary}}
			*/
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
			/*
				http://localhost:8080/product?query={list{id,FirstName,LastName,Gender,Position,Salary}}
			*/

			"list": &graphql.Field{
				Type:        graphql.NewList(EmployerType),
				Description: "Get Emploeyr list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return employers, nil
				},
			},
		},
	})
	var mutationType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			/* Create new employer item
			http://localhost:8080/?query=mutation+{create(FirstName:"Inca",LastName:"Kola",Gender:"other",Position:"CEO",Salary:6000){id,FirstName,LastName,Gender, Position, Salary}}
			*/
			"create": &graphql.Field{
				Type:        EmployerType,
				Description: "Create new employer",
				Args: graphql.FieldConfigArgument{
					"FirstName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"LastName": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Gender": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Position": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Salary": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					employer := Employer{
						ID:        employers[len(employers)-1].ID + 1,
						FirstName: params.Args["FirstName"].(string),
						LastName:  params.Args["LastName"].(string),
						Gender:    params.Args["Gender"].(string),
						Position:  params.Args["Position"].(string),
						Salary:    params.Args["Salary"].(int),
					}
					employers = append(employers, employer)
					return employer, nil
				},
			},

			/* Update employer by id
			   http://localhost:8080/?query=mutation{update(id:101,Salary:9000){id,FirstName,LastName,Gender,Position,Salary}}
			*/
			"update": &graphql.Field{
				Type:        EmployerType,
				Description: "Update employer by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"FirstName": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"LastName": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Gender": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Position": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Salary": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)
					firstName, firstNameOk := params.Args["FirstName"].(string)
					lastName, lastNameOk := params.Args["LastName"].(string)
					gender, genderOk := params.Args["Gender"].(string)
					position, positionOk := params.Args["Position"].(string)
					salary, salaryOk := params.Args["Salary"].(int)
					employer := Employer{}
					for i, p := range employers {
						if int(id) == p.ID {
							if firstNameOk {
								employers[i].FirstName = firstName
							}
							if lastNameOk {
								employers[i].LastName = lastName
							}
							if genderOk {
								employers[i].Gender = gender
							}
							if positionOk {
								employers[i].Position = position
							}
							if salaryOk {
								employers[i].Salary = salary
							}
							employer = employers[i]
							break
						}
					}
					return employer, nil
				},
			},
			/* Delete product by id
			   http://localhost:8080/product?query=mutation{delete(id:106){id,FirstName,LastName,Gender,Position,Salary}}
			*/
			"delete": &graphql.Field{
				Type:        EmployerType,
				Description: "Employer by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)
					employer := Employer{}
					for i, p := range employers {
						if int(id) == p.ID {
							employer = employers[i]
							// Remove from employer list
							employers = append(employers[:i], employers[i+1:]...)
						}
					}

					return employer, nil
				},
			},
		},
	})

	EmployerSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
}
