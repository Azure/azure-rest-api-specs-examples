package armdigitaltwins_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/TimeSeriesDatabaseConnectionsPut_example.json
func ExampleTimeSeriesDatabaseConnectionsClient_BeginCreateOrUpdate_createOrReplaceATimeSeriesDatabaseConnectionForADigitalTwinsInstance() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdigitaltwins.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTimeSeriesDatabaseConnectionsClient().BeginCreateOrUpdate(ctx, "resRg", "myDigitalTwinsService", "myConnection", armdigitaltwins.TimeSeriesDatabaseConnection{
		Properties: &armdigitaltwins.AzureDataExplorerConnectionProperties{
			ConnectionType:                          to.Ptr(armdigitaltwins.ConnectionTypeAzureDataExplorer),
			AdxDatabaseName:                         to.Ptr("myDatabase"),
			AdxEndpointURI:                          to.Ptr("https://mycluster.kusto.windows.net"),
			AdxRelationshipLifecycleEventsTableName: to.Ptr("myRelationshipLifecycleEventsTable"),
			AdxResourceID:                           to.Ptr("/subscriptions/c493073e-2460-45ba-a403-f3e0df1e9feg/resourceGroups/testrg/providers/Microsoft.Kusto/clusters/mycluster"),
			AdxTableName:                            to.Ptr("myPropertyUpdatesTable"),
			AdxTwinLifecycleEventsTableName:         to.Ptr("myTwinLifecycleEventsTable"),
			EventHubEndpointURI:                     to.Ptr("sb://myeh.servicebus.windows.net/"),
			EventHubEntityPath:                      to.Ptr("myeh"),
			EventHubNamespaceResourceID:             to.Ptr("/subscriptions/c493073e-2460-45ba-a403-f3e0df1e9feg/resourceGroups/testrg/providers/Microsoft.EventHub/namespaces/myeh"),
			RecordPropertyAndItemRemovals:           to.Ptr(armdigitaltwins.RecordPropertyAndItemRemovalsTrue),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TimeSeriesDatabaseConnection = armdigitaltwins.TimeSeriesDatabaseConnection{
	// 	Name: to.Ptr("myConnection"),
	// 	Type: to.Ptr("Microsoft.DigitalTwins/digitalTwinsInstances/timeSeriesDatabaseConnections"),
	// 	ID: to.Ptr("/subscriptions/50016170-c839-41ba-a724-51e9df440b9e/resourcegroups/resRg/providers/Microsoft.DigitalTwins/digitalTwinsInstances/myDigitalTwinsService/timeSeriesDatabaseConnections/myConnection"),
	// 	Properties: &armdigitaltwins.AzureDataExplorerConnectionProperties{
	// 		ConnectionType: to.Ptr(armdigitaltwins.ConnectionTypeAzureDataExplorer),
	// 		ProvisioningState: to.Ptr(armdigitaltwins.TimeSeriesDatabaseConnectionStateSucceeded),
	// 		AdxDatabaseName: to.Ptr("myDatabase"),
	// 		AdxEndpointURI: to.Ptr("https://mycluster.kusto.windows.net"),
	// 		AdxRelationshipLifecycleEventsTableName: to.Ptr("myRelationshipLifecycleEventsTable"),
	// 		AdxResourceID: to.Ptr("/subscriptions/c493073e-2460-45ba-a403-f3e0df1e9feg/resourceGroups/testrg/providers/Microsoft.Kusto/clusters/mycluster"),
	// 		AdxTableName: to.Ptr("myPropertyUpdatesTable"),
	// 		AdxTwinLifecycleEventsTableName: to.Ptr("myTwinLifecycleEventsTable"),
	// 		EventHubConsumerGroup: to.Ptr("$Default"),
	// 		EventHubEndpointURI: to.Ptr("sb://myeh.servicebus.windows.net/"),
	// 		EventHubEntityPath: to.Ptr("myeh"),
	// 		EventHubNamespaceResourceID: to.Ptr("/subscriptions/c493073e-2460-45ba-a403-f3e0df1e9feg/resourceGroups/testrg/providers/Microsoft.EventHub/namespaces/myeh"),
	// 		RecordPropertyAndItemRemovals: to.Ptr(armdigitaltwins.RecordPropertyAndItemRemovalsTrue),
	// 	},
	// }
}
