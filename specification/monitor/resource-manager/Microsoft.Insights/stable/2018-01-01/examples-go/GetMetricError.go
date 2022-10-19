package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-01-01/examples/GetMetricError.json
func ExampleMetricsClient_List_getMetricWithError() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewMetricsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.List(ctx, "subscriptions/ac41e21f-afd6-4a79-8070-f01eba278f97/resourceGroups/todking/providers/Microsoft.DocumentDb/databaseAccounts/tk-cosmos-mongo", &armmonitor.MetricsClientListOptions{Timespan: to.Ptr("2021-06-07T21:51:00Z/2021-06-08T01:51:00Z"),
		Interval:        to.Ptr("FULL"),
		Metricnames:     to.Ptr("MongoRequestsCount,MongoRequests"),
		Aggregation:     to.Ptr("average"),
		Top:             nil,
		Orderby:         nil,
		Filter:          nil,
		ResultType:      nil,
		Metricnamespace: to.Ptr("microsoft.documentdb/databaseaccounts"),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
