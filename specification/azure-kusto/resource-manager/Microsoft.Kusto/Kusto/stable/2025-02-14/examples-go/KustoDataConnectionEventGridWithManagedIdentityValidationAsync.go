package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: 2025-02-14/KustoDataConnectionEventGridWithManagedIdentityValidationAsync.json
func ExampleDataConnectionsClient_BeginDataConnectionValidation_kustoDataConnectionEventGridWithManagedIdentityValidationAsync() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDataConnectionsClient().BeginDataConnectionValidation(ctx, "kustorptest", "kustoCluster", "KustoDatabase8", armkusto.DataConnectionValidation{
		DataConnectionName: to.Ptr("dataConnectionTest"),
		Properties: &armkusto.EventGridDataConnectionWithManagedIdentity{
			Kind: to.Ptr(armkusto.DataConnectionKindEventGridWithManagedIdentity),
			Properties: &armkusto.EventGridConnectionWithManagedIdentityProperties{
				ConsumerGroup:                        to.Ptr("$Default"),
				EventHubResourceIDForManagedIdentity: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1"),
				EventGridResourceID:                  to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Storage/storageAccounts/teststorageaccount/providers/Microsoft.EventGrid/eventSubscriptions/eventSubscriptionTest"),
				TableName:                            to.Ptr("TestTable"),
				MappingRuleName:                      to.Ptr("TestMapping"),
				DataFormat:                           to.Ptr(armkusto.EventGridDataFormatMULTIJSON),
				StorageAccountResourceIDForManagedIdentity: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Storage/storageAccounts/teststorageaccount"),
				IgnoreFirstRecord:                          to.Ptr(false),
				BlobStorageEventType:                       to.Ptr(armkusto.BlobStorageEventTypeMicrosoftStorageBlobCreated),
				ManagedIdentityResourceID:                  to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1"),
				DatabaseRouting:                            to.Ptr(armkusto.DatabaseRoutingSingle),
			},
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
	// res = armkusto.DataConnectionsClientDataConnectionValidationResponse{
	// 	DataConnectionValidationListResult: armkusto.DataConnectionValidationListResult{
	// 		Value: []*armkusto.DataConnectionValidationResult{
	// 			{
	// 				ErrorMessage: to.Ptr("Database does not exist"),
	// 			},
	// 		},
	// 	},
	// }
}
