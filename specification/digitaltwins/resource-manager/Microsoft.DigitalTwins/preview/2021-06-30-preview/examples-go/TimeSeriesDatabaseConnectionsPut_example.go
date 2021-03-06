package armdigitaltwins_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/preview/2021-06-30-preview/examples/TimeSeriesDatabaseConnectionsPut_example.json
func ExampleTimeSeriesDatabaseConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdigitaltwins.NewTimeSeriesDatabaseConnectionsClient("50016170-c839-41ba-a724-51e9df440b9e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"resRg",
		"myDigitalTwinsService",
		"myConnection",
		armdigitaltwins.TimeSeriesDatabaseConnection{
			Properties: &armdigitaltwins.AzureDataExplorerConnectionProperties{
				ConnectionType:              to.Ptr(armdigitaltwins.ConnectionTypeAzureDataExplorer),
				AdxDatabaseName:             to.Ptr("myDatabase"),
				AdxEndpointURI:              to.Ptr("https://mycluster.kusto.windows.net"),
				AdxResourceID:               to.Ptr("/subscriptions/c493073e-2460-45ba-a403-f3e0df1e9feg/resourceGroups/testrg/providers/Microsoft.Kusto/clusters/mycluster"),
				AdxTableName:                to.Ptr("myTable"),
				EventHubEndpointURI:         to.Ptr("sb://myeh.servicebus.windows.net/"),
				EventHubEntityPath:          to.Ptr("myeh"),
				EventHubNamespaceResourceID: to.Ptr("/subscriptions/c493073e-2460-45ba-a403-f3e0df1e9feg/resourceGroups/testrg/providers/Microsoft.EventHub/namespaces/myeh"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
