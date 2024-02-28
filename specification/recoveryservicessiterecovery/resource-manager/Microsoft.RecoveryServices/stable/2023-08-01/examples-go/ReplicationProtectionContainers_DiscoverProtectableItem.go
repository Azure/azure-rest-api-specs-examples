package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationProtectionContainers_DiscoverProtectableItem.json
func ExampleReplicationProtectionContainersClient_BeginDiscoverProtectableItem() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationProtectionContainersClient().BeginDiscoverProtectableItem(ctx, "MadhaviVault", "MadhaviVRG", "V2A-W2K12-660", "cloud_7328549c-5c37-4459-a3c2-e35f9ef6893c", armrecoveryservicessiterecovery.DiscoverProtectableItemRequest{
		Properties: &armrecoveryservicessiterecovery.DiscoverProtectableItemRequestProperties{
			FriendlyName: to.Ptr("Test"),
			IPAddress:    to.Ptr("10.150.2.3"),
			OSType:       to.Ptr("Windows"),
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
	// res.ProtectionContainer = armrecoveryservicessiterecovery.ProtectionContainer{
	// 	Name: to.Ptr("cloud_7328549c-5c37-4459-a3c2-e35f9ef6893c"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers"),
	// 	ID: to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/MadhaviVRG/providers/Microsoft.RecoveryServices/vaults/MadhaviVault/replicationFabrics/239f778f368e34f78216d81f030725cdf2033174b47879b9f2eeede06fdd9c4d/replicationProtectionContainers/cloud_7328549c-5c37-4459-a3c2-e35f9ef6893c"),
	// 	Properties: &armrecoveryservicessiterecovery.ProtectionContainerProperties{
	// 		FabricFriendlyName: to.Ptr("V2A-W2K12-660"),
	// 		FabricType: to.Ptr("VMware"),
	// 		FriendlyName: to.Ptr("V2A-W2K12-660"),
	// 		PairingStatus: to.Ptr("Paired"),
	// 		ProtectedItemCount: to.Ptr[int32](2),
	// 		Role: to.Ptr("Primary"),
	// 	},
	// }
}
