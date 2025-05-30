package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationProtectedItems_ListByReplicationProtectionContainers.json
func ExampleReplicationProtectedItemsClient_NewListByReplicationProtectionContainersPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicationProtectedItemsClient().NewListByReplicationProtectionContainersPager("resourceGroupPS1", "vault1", "cloud1", "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", nil)
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
		// page.ReplicationProtectedItemCollection = armrecoveryservicessiterecovery.ReplicationProtectedItemCollection{
		// 	Value: []*armrecoveryservicessiterecovery.ReplicationProtectedItem{
		// 		{
		// 			Name: to.Ptr("f8491e4f-817a-40dd-a90c-af773978c75b"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems"),
		// 			ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectedItems/f8491e4f-817a-40dd-a90c-af773978c75b"),
		// 			Properties: &armrecoveryservicessiterecovery.ReplicationProtectedItemProperties{
		// 				ActiveLocation: to.Ptr("Primary"),
		// 				AllowedOperations: []*string{
		// 					to.Ptr("PlannedFailover"),
		// 					to.Ptr("UnplannedFailover"),
		// 					to.Ptr("DisableProtection"),
		// 					to.Ptr("TestFailover")},
		// 					CurrentScenario: &armrecoveryservicessiterecovery.CurrentScenarioDetails{
		// 						JobID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationJobs/None"),
		// 						ScenarioName: to.Ptr("None"),
		// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1753-01-01T01:01:01.000Z"); return t}()),
		// 					},
		// 					FriendlyName: to.Ptr("vm1"),
		// 					PolicyFriendlyName: to.Ptr("protectionprofile1"),
		// 					PolicyID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/protectionprofile1"),
		// 					PrimaryFabricFriendlyName: to.Ptr("cloud1"),
		// 					PrimaryProtectionContainerFriendlyName: to.Ptr("cloud1"),
		// 					ProtectableItemID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectableItems/f8491e4f-817a-40dd-a90c-af773978c75b"),
		// 					ProtectedItemType: to.Ptr("HyperVVirtualMachine"),
		// 					ProtectionState: to.Ptr("Protected"),
		// 					ProtectionStateDescription: to.Ptr("Protected"),
		// 					ProviderSpecificDetails: &armrecoveryservicessiterecovery.HyperVReplicaAzureReplicationDetails{
		// 						InstanceType: to.Ptr("HyperVReplicaAzure"),
		// 					},
		// 					RecoveryContainerID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/d49858f157601230a6ac5862fbbc6e63bf38d23ecd96cf953767945d457fe9d5/replicationProtectionContainers/d38048d4-b460-4791-8ece-108395ee8478"),
		// 					RecoveryFabricFriendlyName: to.Ptr("Microsoft Azure"),
		// 					RecoveryFabricID: to.Ptr("Microsoft Azure"),
		// 					RecoveryProtectionContainerFriendlyName: to.Ptr("Microsoft Azure"),
		// 					RecoveryServicesProviderID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationRecoveryServicesProviders/241641e6-ee7b-4ee4-8141-821fadda43fa"),
		// 					ReplicationHealth: to.Ptr("Normal"),
		// 					TestFailoverState: to.Ptr("None"),
		// 					TestFailoverStateDescription: to.Ptr("None"),
		// 				},
		// 		}},
		// 	}
	}
}
