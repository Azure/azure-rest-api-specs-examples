package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v3"
)

// Generated from example definition: 2025-08-01/ReplicationProtectionContainerMappings_Delete.json
func ExampleReplicationProtectionContainerMappingsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationProtectionContainerMappingsClient().BeginDelete(ctx, "resourceGroupPS1", "vault1", "cloud1", "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", "cloud1protectionprofile1", armrecoveryservicessiterecovery.RemoveProtectionContainerMappingInput{
		Properties: &armrecoveryservicessiterecovery.RemoveProtectionContainerMappingInputProperties{
			ProviderSpecificInput: &armrecoveryservicessiterecovery.ReplicationProviderContainerUnmappingInput{},
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
