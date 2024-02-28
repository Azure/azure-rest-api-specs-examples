package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationProtectableItems_ListByReplicationProtectionContainers.json
func ExampleReplicationProtectableItemsClient_NewListByReplicationProtectionContainersPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicationProtectableItemsClient().NewListByReplicationProtectionContainersPager("vault1", "resourceGroupPS1", "cloud1", "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", &armrecoveryservicessiterecovery.ReplicationProtectableItemsClientListByReplicationProtectionContainersOptions{Filter: nil,
		Take:      nil,
		SkipToken: nil,
	})
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
		// page.ProtectableItemCollection = armrecoveryservicessiterecovery.ProtectableItemCollection{
		// 	Value: []*armrecoveryservicessiterecovery.ProtectableItem{
		// 		{
		// 			Name: to.Ptr("c0c14913-3d7a-48ea-9531-cc99e0e686e6"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectableItems"),
		// 			ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectableItems/c0c14913-3d7a-48ea-9531-cc99e0e686e6"),
		// 			Properties: &armrecoveryservicessiterecovery.ProtectableItemProperties{
		// 				CustomDetails: &armrecoveryservicessiterecovery.ReplicationGroupDetails{
		// 					InstanceType: to.Ptr("ReplicationGroupDetails"),
		// 				},
		// 				FriendlyName: to.Ptr("vm2"),
		// 				ProtectionReadinessErrors: []*string{
		// 				},
		// 				ProtectionStatus: to.Ptr("Unprotected"),
		// 				RecoveryServicesProviderID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationRecoveryServicesProviders/241641e6-ee7b-4ee4-8141-821fadda43fa"),
		// 				SupportedReplicationProviders: []*string{
		// 					to.Ptr("HyperVReplicaAzure")},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("f8491e4f-817a-40dd-a90c-af773978c75b"),
		// 				Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectableItems"),
		// 				ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectableItems/f8491e4f-817a-40dd-a90c-af773978c75b"),
		// 				Properties: &armrecoveryservicessiterecovery.ProtectableItemProperties{
		// 					CustomDetails: &armrecoveryservicessiterecovery.ReplicationGroupDetails{
		// 						InstanceType: to.Ptr("ReplicationGroupDetails"),
		// 					},
		// 					FriendlyName: to.Ptr("vm1"),
		// 					ProtectionReadinessErrors: []*string{
		// 					},
		// 					ProtectionStatus: to.Ptr("Unprotected"),
		// 					RecoveryServicesProviderID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationRecoveryServicesProviders/241641e6-ee7b-4ee4-8141-821fadda43fa"),
		// 					SupportedReplicationProviders: []*string{
		// 						to.Ptr("HyperVReplicaAzure")},
		// 					},
		// 			}},
		// 		}
	}
}
