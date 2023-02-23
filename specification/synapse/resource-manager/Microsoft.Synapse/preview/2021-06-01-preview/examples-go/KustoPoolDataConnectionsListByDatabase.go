package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsListByDatabase.json
func ExampleKustoPoolDataConnectionsClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewKustoPoolDataConnectionsClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByDatabasePager("kustorptest", "synapseWorkspaceName", "kustoclusterrptest4", "KustoDatabase8", nil)
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
		// page.DataConnectionListResult = armsynapse.DataConnectionListResult{
		// 	Value: []armsynapse.DataConnectionClassification{
		// 		&armsynapse.EventHubDataConnection{
		// 			Name: to.Ptr("KustoClusterRPTest4/KustoDatabase8/KustoDataConnection8"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/DataConnections"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/synapseWorkspaceName/kustopools/KustoClusterRPTest4/Databases/KustoDatabase8/DataConnections/KustoDataConnection8"),
		// 			Kind: to.Ptr(armsynapse.DataConnectionKindEventHub),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armsynapse.EventHubConnectionProperties{
		// 				ConsumerGroup: to.Ptr("testConsumerGroup1"),
		// 				EventHubResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1"),
		// 			},
		// 		},
		// 		&armsynapse.EventHubDataConnection{
		// 			Name: to.Ptr("KustoClusterRPTest4/KustoDatabase9/KustoDataConnection9"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/DataConnections"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/synapseWorkspaceName/kustopools/KustoClusterRPTest4/Databases/KustoDatabase9/DataConnections/KustoDataConnection9"),
		// 			Kind: to.Ptr(armsynapse.DataConnectionKindEventHub),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armsynapse.EventHubConnectionProperties{
		// 				ConsumerGroup: to.Ptr("testConsumerGroup2"),
		// 				EventHubResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns2/eventhubs/eventhubTest2"),
		// 			},
		// 	}},
		// }
	}
}
