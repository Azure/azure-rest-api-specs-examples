package armpostgresql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/OperationList.json
func ExampleOperationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationsClient().List(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationListResult = armpostgresql.OperationListResult{
	// 	Value: []*armpostgresql.Operation{
	// 		{
	// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/locations/performanceTiers/read"),
	// 			Display: &armpostgresql.OperationDisplay{
	// 				Description: to.Ptr("Returns the list of Performance Tiers available."),
	// 				Operation: to.Ptr("List Performance Tiers"),
	// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
	// 				Resource: to.Ptr("Performance Tiers"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/read"),
	// 			Display: &armpostgresql.OperationDisplay{
	// 				Description: to.Ptr("Return the list of servers or gets the properties for the specified server."),
	// 				Operation: to.Ptr("List/Get PostgreSQL Servers"),
	// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
	// 				Resource: to.Ptr("PostgreSQL Server"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/write"),
	// 			Display: &armpostgresql.OperationDisplay{
	// 				Description: to.Ptr("Creates a server with the specified parameters or update the properties or tags for the specified server."),
	// 				Operation: to.Ptr("Create/Update PostgreSQL Server"),
	// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
	// 				Resource: to.Ptr("PostgreSQL Server"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/delete"),
	// 			Display: &armpostgresql.OperationDisplay{
	// 				Description: to.Ptr("Deletes an existing server."),
	// 				Operation: to.Ptr("Delete PostgreSQL Server"),
	// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
	// 				Resource: to.Ptr("PostgreSQL Server"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/performanceTiers/read"),
	// 			Display: &armpostgresql.OperationDisplay{
	// 				Description: to.Ptr("Returns the list of Performance Tiers available."),
	// 				Operation: to.Ptr("List Performance Tiers"),
	// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
	// 				Resource: to.Ptr("Performance Tiers"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/firewallRules/read"),
	// 			Display: &armpostgresql.OperationDisplay{
	// 				Description: to.Ptr("Return the list of firewall rules for a server or gets the properties for the specified firewall rule."),
	// 				Operation: to.Ptr("List/Get Firewall Rules"),
	// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
	// 				Resource: to.Ptr("Firewall Rules"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/firewallRules/write"),
	// 			Display: &armpostgresql.OperationDisplay{
	// 				Description: to.Ptr("Creates a firewall rule with the specified parameters or update an existing rule."),
	// 				Operation: to.Ptr("Create/Update Firewall Rule"),
	// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
	// 				Resource: to.Ptr("Firewall Rules"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/firewallRules/delete"),
	// 			Display: &armpostgresql.OperationDisplay{
	// 				Description: to.Ptr("Deletes an existing firewall rule."),
	// 				Operation: to.Ptr("Delete Firewall Rule"),
	// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
	// 				Resource: to.Ptr("Firewall Rules"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/providers/Microsoft.Insights/metricDefinitions/read"),
	// 			Display: &armpostgresql.OperationDisplay{
	// 				Description: to.Ptr("Return types of metrics that are available for databases"),
	// 				Operation: to.Ptr("Get database metric definitions"),
	// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
	// 				Resource: to.Ptr("Database Metric Definition"),
	// 			},
	// 			Properties: map[string]any{
	// 				"serviceSpecification": map[string]any{
	// 					"metricSpecifications":[]any{
	// 						map[string]any{
	// 							"name": "cpu_percent",
	// 							"aggregationType": "Average",
	// 							"displayDescription": "CPU percent",
	// 							"displayName": "CPU percent",
	// 							"fillGapWithZero": true,
	// 							"unit": "Percent",
	// 						},
	// 						map[string]any{
	// 							"name": "memory_percent",
	// 							"aggregationType": "Average",
	// 							"displayDescription": "Memory percent",
	// 							"displayName": "Memory percent",
	// 							"fillGapWithZero": true,
	// 							"unit": "Percent",
	// 						},
	// 						map[string]any{
	// 							"name": "io_consumption_percent",
	// 							"aggregationType": "Average",
	// 							"displayDescription": "IO percent",
	// 							"displayName": "IO percent",
	// 							"fillGapWithZero": true,
	// 							"unit": "Percent",
	// 						},
	// 						map[string]any{
	// 							"name": "storage_percent",
	// 							"aggregationType": "Average",
	// 							"displayDescription": "Storage percentage",
	// 							"displayName": "Storage percentage",
	// 							"unit": "Percent",
	// 						},
	// 						map[string]any{
	// 							"name": "storage_used",
	// 							"aggregationType": "Average",
	// 							"displayDescription": "Storage used",
	// 							"displayName": "Storage used",
	// 							"unit": "Bytes",
	// 						},
	// 						map[string]any{
	// 							"name": "storage_limit",
	// 							"aggregationType": "Average",
	// 							"displayDescription": "Storage limit",
	// 							"displayName": "Storage limit",
	// 							"unit": "Bytes",
	// 						},
	// 						map[string]any{
	// 							"name": "active_connections",
	// 							"aggregationType": "Average",
	// 							"displayDescription": "Total active connections",
	// 							"displayName": "Total active connections",
	// 							"fillGapWithZero": true,
	// 							"unit": "Count",
	// 						},
	// 						map[string]any{
	// 							"name": "connections_failed",
	// 							"aggregationType": "Average",
	// 							"displayDescription": "Total failed connections",
	// 							"displayName": "Total failed connections",
	// 							"fillGapWithZero": true,
	// 							"unit": "Count",
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/providers/Microsoft.Insights/diagnosticSettings/read"),
	// 			Display: &armpostgresql.OperationDisplay{
	// 				Description: to.Ptr("Gets the disagnostic setting for the resource"),
	// 				Operation: to.Ptr("Read diagnostic setting"),
	// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
	// 				Resource: to.Ptr("Database Metric Definition"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/providers/Microsoft.Insights/diagnosticSettings/write"),
	// 			Display: &armpostgresql.OperationDisplay{
	// 				Description: to.Ptr("Creates or updates the diagnostic setting for the resource"),
	// 				Operation: to.Ptr("Write diagnostic setting"),
	// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
	// 				Resource: to.Ptr("Database Metric Definition"),
	// 			},
	// 	}},
	// }
}
