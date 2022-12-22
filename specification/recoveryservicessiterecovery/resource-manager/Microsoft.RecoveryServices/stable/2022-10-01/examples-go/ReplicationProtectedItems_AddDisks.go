package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationProtectedItems_AddDisks.json
func ExampleReplicationProtectedItemsClient_BeginAddDisks() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1", "resourceGroupPS1", "c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginAddDisks(ctx, "cloud1", "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", "f8491e4f-817a-40dd-a90c-af773978c75b", armrecoveryservicessiterecovery.AddDisksInput{
		Properties: &armrecoveryservicessiterecovery.AddDisksInputProperties{
			ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2AAddDisksInput{
				InstanceType: to.Ptr("A2A"),
				VMDisks: []*armrecoveryservicessiterecovery.A2AVMDiskInputDetails{
					{
						DiskURI:                             to.Ptr("https://vmstorage.blob.core.windows.net/vhds/datadisk1.vhd"),
						PrimaryStagingAzureStorageAccountID: to.Ptr("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourcegroups/primaryResource/providers/Microsoft.Storage/storageAccounts/vmcachestorage"),
						RecoveryAzureStorageAccountID:       to.Ptr("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourcegroups/recoveryResource/providers/Microsoft.Storage/storageAccounts/recoverystorage"),
					}},
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
