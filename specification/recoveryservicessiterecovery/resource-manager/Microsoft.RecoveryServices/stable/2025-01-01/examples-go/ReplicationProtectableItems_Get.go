package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationProtectableItems_Get.json
func ExampleReplicationProtectableItemsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReplicationProtectableItemsClient().Get(ctx, "resourceGroupPS1", "vault1", "cloud1", "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", "c0c14913-3d7a-48ea-9531-cc99e0e686e6", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProtectableItem = armrecoveryservicessiterecovery.ProtectableItem{
	// 	Name: to.Ptr("c0c14913-3d7a-48ea-9531-cc99e0e686e6"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectableItems"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectableItems/c0c14913-3d7a-48ea-9531-cc99e0e686e6"),
	// 	Properties: &armrecoveryservicessiterecovery.ProtectableItemProperties{
	// 		CustomDetails: &armrecoveryservicessiterecovery.ReplicationGroupDetails{
	// 			InstanceType: to.Ptr("ReplicationGroupDetails"),
	// 		},
	// 		FriendlyName: to.Ptr("vm2"),
	// 		ProtectionReadinessErrors: []*string{
	// 		},
	// 		ProtectionStatus: to.Ptr("Unprotected"),
	// 		RecoveryServicesProviderID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationRecoveryServicesProviders/241641e6-ee7b-4ee4-8141-821fadda43fa"),
	// 		SupportedReplicationProviders: []*string{
	// 			to.Ptr("HyperVReplicaAzure")},
	// 		},
	// 	}
}
