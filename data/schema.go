package data

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

var userType *graphql.Object
var visitType *graphql.Object

var nodeDefinitions *relay.NodeDefinitions
var visitConnection *relay.GraphQLConnectionDefinitions

//Schema stuff
var Schema graphql.Schema

func init() {

	/**
	 * We get the node interface and field from the Relay library.
	 *
	 * The first method defines the way we resolve an ID to its object.
	 * The second defines the way we resolve an object to its GraphQL type.
	 */
	nodeDefinitions = relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher: func(id string, info graphql.ResolveInfo) interface{} {
			resolvedID := relay.FromGlobalID(id)
			if resolvedID.Type == "User" {
				return GetUser(resolvedID.ID)
			}
			if resolvedID.Type == "Visit" {
				return GetVisit(resolvedID.ID)
			}
			return nil
		},
		TypeResolve: func(value interface{}, info graphql.ResolveInfo) *graphql.Object {
			switch value.(type) {
			case *User:
				return userType
			case *Visit:
				return visitType
			}
			return nil
		},
	})

	/**
	 * Define your own types here
	 */
	visitType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Visit",
		Description: "Visits to mystery shops",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Visit", nil),
			"location": &graphql.Field{
				Description: "The location of the Visit",
				Type:        graphql.String,
			},
		},
		Interfaces: []*graphql.Interface{
			nodeDefinitions.NodeInterface,
		},
	})
	visitConnection = relay.ConnectionDefinitions(relay.ConnectionConfig{
		Name:     "VisitConnection",
		NodeType: visitType,
	})

	userType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: "A person who uses our app",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("User", nil),
			"name": &graphql.Field{
				Description: "The user's name",
				Type:        graphql.String,
			},
			"email": &graphql.Field{
				Description: "The user's email address",
				Type:        graphql.String,
			},
			"visits": &graphql.Field{
				Type:        visitConnection.ConnectionType,
				Description: "A person's past mystery shop visits",
				Args:        relay.ConnectionArgs,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					args := relay.NewConnectionArguments(p.Args)
					dataSlice := VisitsToInterfaceSlice(GetVisits()...)

					return relay.ConnectionFromArray(dataSlice, args), nil
				},
			},
		},
		Interfaces: []*graphql.Interface{
			nodeDefinitions.NodeInterface,
		},
	})

	/**
	 * This is the type that will be the root of our query,
	 * and the entry point into our schema.
	 */
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"node": nodeDefinitions.NodeField,

			// Add you own root fields here
			"viewer": &graphql.Field{
				Type: userType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetViewer(), nil
				},
			},
		},
	})

	/**
	 * This is the type that will be the root of our mutations,
	 * and the entry point into performing writes in our schema.
	 */
	//	mutationType := graphql.NewObject(graphql.ObjectConfig{
	//		Name: "Mutation",
	//		Fields: graphql.Fields{
	//			// Add you own mutations here
	//		},
	//	})

	/**
	* Finally, we construct our schema (whose starting query type is the query
	* type we defined above) and export it.
	 */
	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	if err != nil {
		panic(err)
	}

}
