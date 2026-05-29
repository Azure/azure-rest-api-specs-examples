package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: 2025-02-14/KustoDataConnectionsEventHubWithManagedIdentityGet.json
func ExampleDataConnectionsClient_Get_kustoDataConnectionsEventHubWithManagedIdentityGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectionsClient().Get(ctx, "kustorptest", "kustoCluster", "KustoDatabase8", "dataConnectionTest", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armkusto.DataConnectionsClientGetResponse{
	// 	DataConnectionClassification: &armkusto.EventHubDataConnectionWithManagedIdentity{
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/KustoDatabase8/DataConnections/dataConnectionTest"),
	// 		Name: to.Ptr("kustoCluster/KustoDatabase8/dataConnectionTest"),
	// 		Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/DataConnections"),
	// 		Location: to.Ptr("westus"),
	// 		Kind: to.Ptr(armkusto.DataConnectionKindEventHubWithManagedIdentity),
	// 		Properties: &armkusto.EventHubConnectionWithManagedIdentityProperties{
	// 			EventHubResourceIDForManagedIdentity: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1"),
	// 			ConsumerGroup: to.Ptr("$Default"),
	// 			TableName: to.Ptr("TestTable"),
	// 			MappingRuleName: to.Ptr("TestMapping"),
	// 			DataFormat: to.Ptr(armkusto.EventHubDataFormatMULTIJSON),
	// 			EventSystemProperties: []*string{
	// 			},
	// 			Compression: to.Ptr(armkusto.CompressionNone),
	// 			ManagedIdentityResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1"),
	// 			ManagedIdentityObjectID: to.Ptr("87654321-1234-1234-1234-123456789123"),
	// 			DatabaseRouting: to.Ptr(armkusto.DatabaseRoutingSingle),
	// 			RetrievalStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-29T12:00:00.6554616Z"); return t}()),
	// 			ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
