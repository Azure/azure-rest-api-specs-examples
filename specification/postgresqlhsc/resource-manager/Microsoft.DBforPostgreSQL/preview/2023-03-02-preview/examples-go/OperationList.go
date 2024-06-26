package armcosmosforpostgresql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmosforpostgresql/armcosmosforpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-02-preview/examples/OperationList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmosforpostgresql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.OperationListResult = armcosmosforpostgresql.OperationListResult{
		// 	Value: []*armcosmosforpostgresql.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/read"),
		// 			Display: &armcosmosforpostgresql.OperationDisplay{
		// 				Description: to.Ptr("Return the list of clusters or gets the properties for the specified cluster."),
		// 				Operation: to.Ptr("List/Get PostgreSQL Clusters"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Cluster"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/write"),
		// 			Display: &armcosmosforpostgresql.OperationDisplay{
		// 				Description: to.Ptr("Creates a cluster with the specified parameters or update the properties or tags for the specified cluster."),
		// 				Operation: to.Ptr("Create/Update PostgreSQL Cluster"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Cluster"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/delete"),
		// 			Display: &armcosmosforpostgresql.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing cluster."),
		// 				Operation: to.Ptr("Delete PostgreSQL Cluster"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Cluster"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/servers/read"),
		// 			Display: &armcosmosforpostgresql.OperationDisplay{
		// 				Description: to.Ptr("Return the list of servers of cluster or gets the properties for the specified server from cluster."),
		// 				Operation: to.Ptr("List/Get PostgreSQL Servers of Cluster"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/servers/configurations/read"),
		// 			Display: &armcosmosforpostgresql.OperationDisplay{
		// 				Description: to.Ptr("Return the list of configurations for a server or gets the value for the specified configuration."),
		// 				Operation: to.Ptr("List/Get Configurations of Server"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Configurations"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/configurations/read"),
		// 			Display: &armcosmosforpostgresql.OperationDisplay{
		// 				Description: to.Ptr("Return the list of configurations for a cluster or gets the value for the specified configuration."),
		// 				Operation: to.Ptr("List/Get Configurations of Cluster"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Configurations"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/configurations/write"),
		// 			Display: &armcosmosforpostgresql.OperationDisplay{
		// 				Description: to.Ptr("Update an value of configurations."),
		// 				Operation: to.Ptr("Update Configurations of Cluster"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Configurations"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/firewallRules/read"),
		// 			Display: &armcosmosforpostgresql.OperationDisplay{
		// 				Description: to.Ptr("Return the list of firewall rules for a cluster or gets the properties for the specified firewall rule."),
		// 				Operation: to.Ptr("List/Get Firewall Rules"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Firewall Rules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/firewallRules/write"),
		// 			Display: &armcosmosforpostgresql.OperationDisplay{
		// 				Description: to.Ptr("Creates a firewall rule with the specified parameters or update an existing rule."),
		// 				Operation: to.Ptr("Create/Update Firewall Rule"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Firewall Rules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/firewallRules/delete"),
		// 			Display: &armcosmosforpostgresql.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing firewall rule."),
		// 				Operation: to.Ptr("Delete Firewall Rule"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Firewall Rules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armcosmosforpostgresql.OperationDisplay{
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
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 			Display: &armcosmosforpostgresql.OperationDisplay{
		// 				Description: to.Ptr("Gets the disagnostic setting for the resource"),
		// 				Operation: to.Ptr("Read diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Database Metric Definition"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 			Display: &armcosmosforpostgresql.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the diagnostic setting for the resource"),
		// 				Operation: to.Ptr("Write diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Database Metric Definition"),
		// 			},
		// 	}},
		// }
	}
}
