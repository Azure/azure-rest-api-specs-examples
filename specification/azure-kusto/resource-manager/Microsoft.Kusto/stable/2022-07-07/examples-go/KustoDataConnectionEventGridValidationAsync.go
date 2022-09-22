package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-07-07/examples/KustoDataConnectionEventGridValidationAsync.json
func ExampleDataConnectionsClient_BeginDataConnectionValidation_kustoDataConnectionEventGridValidation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armkusto.NewDataConnectionsClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDataConnectionValidation(ctx, "kustorptest", "kustoCluster", "KustoDatabase8", armkusto.DataConnectionValidation{
		DataConnectionName: to.Ptr("dataConnectionTest"),
		Properties: &armkusto.EventGridDataConnection{
			Kind: to.Ptr(armkusto.DataConnectionKindEventGrid),
			Properties: &armkusto.EventGridConnectionProperties{
				BlobStorageEventType:      to.Ptr(armkusto.BlobStorageEventTypeMicrosoftStorageBlobCreated),
				ConsumerGroup:             to.Ptr("$Default"),
				DataFormat:                to.Ptr(armkusto.EventGridDataFormatJSON),
				DatabaseRouting:           to.Ptr(armkusto.DatabaseRoutingSingle),
				EventGridResourceID:       to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Storage/storageAccounts/teststorageaccount/providers/Microsoft.EventGrid/eventSubscriptions/eventSubscriptionTest"),
				EventHubResourceID:        to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1"),
				IgnoreFirstRecord:         to.Ptr(false),
				ManagedIdentityResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1"),
				MappingRuleName:           to.Ptr("TestMapping"),
				StorageAccountResourceID:  to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Storage/storageAccounts/teststorageaccount"),
				TableName:                 to.Ptr("TestTable"),
			},
		},
	}, nil)
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
