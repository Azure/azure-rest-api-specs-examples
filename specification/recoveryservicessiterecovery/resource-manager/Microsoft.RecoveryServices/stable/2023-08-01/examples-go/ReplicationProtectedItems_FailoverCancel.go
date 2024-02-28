package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationProtectedItems_FailoverCancel.json
func ExampleReplicationProtectedItemsClient_BeginFailoverCancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationProtectedItemsClient().BeginFailoverCancel(ctx, "vault1", "resourceGroupPS1", "cloud1", "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", "f8491e4f-817a-40dd-a90c-af773978c75b", nil)
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
	// res.ReplicationProtectedItem = armrecoveryservicessiterecovery.ReplicationProtectedItem{
	// 	Name: to.Ptr("f8491e4f-817a-40dd-a90c-af773978c75b"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectedItems/f8491e4f-817a-40dd-a90c-af773978c75b"),
	// 	Properties: &armrecoveryservicessiterecovery.ReplicationProtectedItemProperties{
	// 		ActiveLocation: to.Ptr("Primary"),
	// 		AllowedOperations: []*string{
	// 			to.Ptr("PlannedFailover"),
	// 			to.Ptr("DisableProtection")},
	// 			CurrentScenario: &armrecoveryservicessiterecovery.CurrentScenarioDetails{
	// 				JobID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationJobs/None"),
	// 				ScenarioName: to.Ptr("None"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-24T07:17:23.135Z"); return t}()),
	// 			},
	// 			FailoverRecoveryPointID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectedItems/f8491e4f-817a-40dd-a90c-af773978c75b/recoveryPoints/b22134ea-620c-474b-9fa5-3c1cb47708e3"),
	// 			FriendlyName: to.Ptr("vm1"),
	// 			LastSuccessfulFailoverTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-24T06:37:23.157Z"); return t}()),
	// 			LastSuccessfulTestFailoverTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1601-01-01T00:00:00.000Z"); return t}()),
	// 			PolicyFriendlyName: to.Ptr("protectionprofile1"),
	// 			PolicyID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/protectionprofile1"),
	// 			PrimaryFabricFriendlyName: to.Ptr("cloud1"),
	// 			PrimaryProtectionContainerFriendlyName: to.Ptr("cloud1"),
	// 			ProtectableItemID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectableItems/f8491e4f-817a-40dd-a90c-af773978c75b"),
	// 			ProtectedItemType: to.Ptr(""),
	// 			ProtectionState: to.Ptr("Protected"),
	// 			ProtectionStateDescription: to.Ptr("Protected"),
	// 			ProviderSpecificDetails: &armrecoveryservicessiterecovery.InMageRcmFailbackReplicationDetails{
	// 				InstanceType: to.Ptr("InMageRcmFailback"),
	// 			},
	// 			RecoveryContainerID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179"),
	// 			RecoveryFabricFriendlyName: to.Ptr("cloud1"),
	// 			RecoveryFabricID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1"),
	// 			RecoveryProtectionContainerFriendlyName: to.Ptr("cloud1"),
	// 			RecoveryServicesProviderID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationRecoveryServicesProviders/241641e6-ee7b-4ee4-8141-821fadda43fa"),
	// 			ReplicationHealth: to.Ptr("Normal"),
	// 			TestFailoverState: to.Ptr("None"),
	// 			TestFailoverStateDescription: to.Ptr("None"),
	// 		},
	// 	}
}
