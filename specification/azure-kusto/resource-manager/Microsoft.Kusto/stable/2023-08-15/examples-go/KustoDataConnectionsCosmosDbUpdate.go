package armkusto_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDataConnectionsCosmosDbUpdate.json
func ExampleDataConnectionsClient_BeginUpdate_kustoDataConnectionsCosmosDbUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDataConnectionsClient().BeginUpdate(ctx, "kustorptest", "kustoCluster", "KustoDatabase1", "dataConnectionTest", &armkusto.CosmosDbDataConnection{
		Kind:     to.Ptr(armkusto.DataConnectionKindCosmosDb),
		Location: to.Ptr("westus"),
		Properties: &armkusto.CosmosDbDataConnectionProperties{
			CosmosDbAccountResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.DocumentDb/databaseAccounts/cosmosDbAccountTest1"),
			CosmosDbContainer:         to.Ptr("cosmosDbContainerTest"),
			CosmosDbDatabase:          to.Ptr("cosmosDbDatabaseTest"),
			ManagedIdentityResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1"),
			MappingRuleName:           to.Ptr("TestMapping"),
			RetrievalStartDate:        to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-29T12:00:00.655Z"); return t }()),
			TableName:                 to.Ptr("TestTable"),
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
	// res = armkusto.DataConnectionsClientUpdateResponse{
	// 	                            DataConnectionClassification: &armkusto.CosmosDbDataConnection{
	// 		Name: to.Ptr("kustoCluster/KustoDatabase1/dataConnectionTest"),
	// 		Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/DataConnections"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/KustoDatabase1/DataConnections/dataConnectionTest"),
	// 		Kind: to.Ptr(armkusto.DataConnectionKindCosmosDb),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armkusto.CosmosDbDataConnectionProperties{
	// 			CosmosDbAccountResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.DocumentDb/databaseAccounts/cosmosDbAccountTest1"),
	// 			CosmosDbContainer: to.Ptr("cosmosDbContainerTest"),
	// 			CosmosDbDatabase: to.Ptr("cosmosDbDatabaseTest"),
	// 			ManagedIdentityObjectID: to.Ptr("87654321-1234-1234-1234-123456789123"),
	// 			ManagedIdentityResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1"),
	// 			MappingRuleName: to.Ptr("TestMapping"),
	// 			RetrievalStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-29T12:00:00.655Z"); return t}()),
	// 			TableName: to.Ptr("TestTable"),
	// 		},
	// 	},
	// 	                        }
}
