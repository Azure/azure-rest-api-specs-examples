package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: 2025-02-14/KustoDataConnectionsEventHubWithManagedIdentityUpdate.json
func ExampleDataConnectionsClient_BeginUpdate_kustoDataConnectionsEventHubWithManagedIdentityUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDataConnectionsClient().BeginUpdate(ctx, "kustorptest", "kustoCluster", "KustoDatabase8", "dataConnectionTest", &armkusto.EventHubDataConnectionWithManagedIdentity{
		Location: to.Ptr("westus"),
		Kind:     to.Ptr(armkusto.DataConnectionKindEventHubWithManagedIdentity),
		Properties: &armkusto.EventHubConnectionWithManagedIdentityProperties{
			EventHubResourceIDForManagedIdentity: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1"),
			ManagedIdentityResourceID:            to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1"),
			ConsumerGroup:                        to.Ptr("testConsumerGroup1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armkusto.DataConnectionsClientUpdateResponse{
	// 	DataConnectionClassification: &armkusto.EventHubDataConnectionWithManagedIdentity{
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/KustoDatabase8/DataConnections/dataConnectionTest"),
	// 		Name: to.Ptr("kustoCluster/KustoDatabase8/dataConnectionTest"),
	// 		Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/DataConnections"),
	// 		Location: to.Ptr("westus"),
	// 		Kind: to.Ptr(armkusto.DataConnectionKindEventHubWithManagedIdentity),
	// 		Properties: &armkusto.EventHubConnectionWithManagedIdentityProperties{
	// 			EventHubResourceIDForManagedIdentity: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1"),
	// 			ManagedIdentityResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1"),
	// 			ConsumerGroup: to.Ptr("testConsumerGroup1"),
	// 			Compression: to.Ptr(armkusto.CompressionNone),
	// 		},
	// 	},
	// }
}
