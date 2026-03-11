package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v3"
)

// Generated from example definition: 2025-08-01/ReplicationProtectedItems_Delete.json
func ExampleReplicationProtectedItemsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationProtectedItemsClient().BeginDelete(ctx, "resourceGroupPS1", "vault1", "cloud1", "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", "c0c14913-3d7a-48ea-9531-cc99e0e686e6", armrecoveryservicessiterecovery.DisableProtectionInput{
		Properties: &armrecoveryservicessiterecovery.DisableProtectionInputProperties{
			ReplicationProviderInput: &armrecoveryservicessiterecovery.InMageDisableProtectionProviderSpecificInput{
				InstanceType: to.Ptr("DisableProtectionProviderSpecificInput"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
