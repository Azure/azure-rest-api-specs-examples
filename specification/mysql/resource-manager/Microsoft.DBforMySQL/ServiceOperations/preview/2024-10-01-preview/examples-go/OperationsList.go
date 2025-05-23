package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e26b89bcbec9eed5026c01416e481408b2a1ca1a/specification/mysql/resource-manager/Microsoft.DBforMySQL/ServiceOperations/preview/2024-10-01-preview/examples/OperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armmysqlflexibleservers.OperationListResult{
		// 	Value: []*armmysqlflexibleservers.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforMySQL/flexibleServers/firewallRules/read"),
		// 			Display: &armmysqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of firewall rules for a server or gets the properties for the specified firewall rule."),
		// 				Operation: to.Ptr("List/Get Firewall Rules"),
		// 				Provider: to.Ptr("Microsoft DB for MySQL"),
		// 				Resource: to.Ptr("Firewall Rules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforMySQL/flexibleServers/firewallRules/write"),
		// 			Display: &armmysqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates a firewall rule with the specified parameters or update an existing rule."),
		// 				Operation: to.Ptr("Create/Update Firewall Rule"),
		// 				Provider: to.Ptr("Microsoft DB for MySQL"),
		// 				Resource: to.Ptr("Firewall Rules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforMySQL/flexibleServers/firewallRules/delete"),
		// 			Display: &armmysqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing firewall rule."),
		// 				Operation: to.Ptr("Delete Firewall Rule"),
		// 				Provider: to.Ptr("Microsoft DB for MySQL"),
		// 				Resource: to.Ptr("Firewall Rules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforMySQL/flexibleServers/read"),
		// 			Display: &armmysqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of servers or gets the properties for the specified server."),
		// 				Operation: to.Ptr("List/Get MySQL Servers"),
		// 				Provider: to.Ptr("Microsoft DB for MySQL"),
		// 				Resource: to.Ptr("MySQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforMySQL/flexibleServers/write"),
		// 			Display: &armmysqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates a server with the specified parameters or update the properties or tags for the specified server."),
		// 				Operation: to.Ptr("Create/Update MySQL Server"),
		// 				Provider: to.Ptr("Microsoft DB for MySQL"),
		// 				Resource: to.Ptr("MySQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforMySQL/flexibleServers/delete"),
		// 			Display: &armmysqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing server."),
		// 				Operation: to.Ptr("Delete MySQL Server"),
		// 				Provider: to.Ptr("Microsoft DB for MySQL"),
		// 				Resource: to.Ptr("MySQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforMySQL/flexibleServers/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armmysqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return types of metrics that are available for databases"),
		// 				Operation: to.Ptr("Get database metric definitions"),
		// 				Provider: to.Ptr("Microsoft DB for MySQL"),
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
		// 							"name": "serverlog_storage_percent",
		// 							"aggregationType": "Average",
		// 							"displayDescription": "Server Log storage percent",
		// 							"displayName": "Server Log storage percent",
		// 							"unit": "Percent",
		// 						},
		// 						map[string]any{
		// 							"name": "serverlog_storage_usage",
		// 							"aggregationType": "Average",
		// 							"displayDescription": "Server Log storage used",
		// 							"displayName": "Server Log storage used",
		// 							"unit": "Bytes",
		// 						},
		// 						map[string]any{
		// 							"name": "serverlog_storage_limit",
		// 							"aggregationType": "Average",
		// 							"displayDescription": "Server Log storage limit",
		// 							"displayName": "Server Log storage limit",
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
		// 						map[string]any{
		// 							"name": "seconds_behind_master",
		// 							"aggregationType": "Average",
		// 							"displayDescription": "Replication lag in seconds",
		// 							"displayName": "Replication lag in seconds",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforMySQL/flexibleServers/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 			Display: &armmysqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets the disagnostic setting for the resource"),
		// 				Operation: to.Ptr("Read diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft DB for MySQL"),
		// 				Resource: to.Ptr("Database Metric Definition"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforMySQL/flexibleServers/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 			Display: &armmysqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the diagnostic setting for the resource"),
		// 				Operation: to.Ptr("Write diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft DB for MySQL"),
		// 				Resource: to.Ptr("Database Metric Definition"),
		// 			},
		// 	}},
		// }
	}
}
