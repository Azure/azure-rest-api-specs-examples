package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDataConnectionsListByDatabase.json
func ExampleDataConnectionsClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataConnectionsClient().NewListByDatabasePager("kustorptest", "kustoCluster", "KustoDatabase8", nil)
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
		// page.DataConnectionListResult = armkusto.DataConnectionListResult{
		// 	Value: []armkusto.DataConnectionClassification{
		// 		&armkusto.EventHubDataConnection{
		// 			Name: to.Ptr("kustoCluster/KustoDatabase8/KustoDataConnection1"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/DataConnections"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/KustoDatabase8/DataConnections/KustoDataConnection1"),
		// 			Kind: to.Ptr(armkusto.DataConnectionKindEventHub),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armkusto.EventHubConnectionProperties{
		// 				Compression: to.Ptr(armkusto.CompressionNone),
		// 				ConsumerGroup: to.Ptr("$Default"),
		// 				DataFormat: to.Ptr(armkusto.EventHubDataFormatJSON),
		// 				DatabaseRouting: to.Ptr(armkusto.DatabaseRoutingSingle),
		// 				EventHubResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1"),
		// 				EventSystemProperties: []*string{
		// 				},
		// 				ManagedIdentityObjectID: to.Ptr("87654321-1234-1234-1234-123456789123"),
		// 				ManagedIdentityResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1"),
		// 				MappingRuleName: to.Ptr("TestMapping"),
		// 				ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 				RetrievalStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-29T12:00:00.655Z"); return t}()),
		// 				TableName: to.Ptr("TestTable"),
		// 			},
		// 		},
		// 		&armkusto.EventGridDataConnection{
		// 			Name: to.Ptr("kustoCluster/KustoDatabase8/KustoDataConnection2"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/DataConnections"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/KustoDatabase8/DataConnections/KustoDataConnection2"),
		// 			Kind: to.Ptr(armkusto.DataConnectionKindEventGrid),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armkusto.EventGridConnectionProperties{
		// 				BlobStorageEventType: to.Ptr(armkusto.BlobStorageEventTypeMicrosoftStorageBlobCreated),
		// 				ConsumerGroup: to.Ptr("$Default"),
		// 				DataFormat: to.Ptr(armkusto.EventGridDataFormatJSON),
		// 				DatabaseRouting: to.Ptr(armkusto.DatabaseRoutingSingle),
		// 				EventGridResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Storage/storageAccounts/teststorageaccount/providers/Microsoft.EventGrid/eventSubscriptions/eventSubscriptionTest"),
		// 				EventHubResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest2"),
		// 				IgnoreFirstRecord: to.Ptr(false),
		// 				ManagedIdentityObjectID: to.Ptr("87654321-1234-1234-1234-123456789123"),
		// 				ManagedIdentityResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1"),
		// 				MappingRuleName: to.Ptr("TestMapping"),
		// 				ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 				StorageAccountResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Storage/storageAccounts/teststorageaccount"),
		// 				TableName: to.Ptr("TestTable"),
		// 			},
		// 		},
		// 		&armkusto.CosmosDbDataConnection{
		// 			Name: to.Ptr("kustoCluster/KustoDatabase8/KustoDataConnection3"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/DataConnections"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/KustoDatabase8/DataConnections/KustoDataConnection3"),
		// 			Kind: to.Ptr(armkusto.DataConnectionKindCosmosDb),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armkusto.CosmosDbDataConnectionProperties{
		// 				CosmosDbAccountResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.DocumentDb/databaseAccounts/cosmosDbAccountTest1"),
		// 				CosmosDbContainer: to.Ptr("cosmosDbContainerTest"),
		// 				CosmosDbDatabase: to.Ptr("cosmosDbDatabaseTest"),
		// 				ManagedIdentityObjectID: to.Ptr("87654321-1234-1234-1234-123456789123"),
		// 				ManagedIdentityResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1"),
		// 				MappingRuleName: to.Ptr("TestMapping"),
		// 				ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 				RetrievalStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-29T12:00:00.655Z"); return t}()),
		// 				TableName: to.Ptr("TestTable"),
		// 			},
		// 	}},
		// }
	}
}
