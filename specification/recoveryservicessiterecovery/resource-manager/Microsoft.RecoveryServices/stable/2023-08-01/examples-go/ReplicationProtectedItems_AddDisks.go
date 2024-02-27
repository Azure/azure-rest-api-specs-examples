package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationProtectedItems_AddDisks.json
func ExampleReplicationProtectedItemsClient_BeginAddDisks() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationProtectedItemsClient().BeginAddDisks(ctx, "vault1", "resourceGroupPS1", "cloud1", "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", "f8491e4f-817a-40dd-a90c-af773978c75b", armrecoveryservicessiterecovery.AddDisksInput{
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
	// 			to.Ptr("RepairReplication"),
	// 			to.Ptr("UnplannedFailover"),
	// 			to.Ptr("DisableProtection"),
	// 			to.Ptr("TestFailover")},
	// 			CurrentScenario: &armrecoveryservicessiterecovery.CurrentScenarioDetails{
	// 				JobID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationJobs/None"),
	// 				ScenarioName: to.Ptr("None"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1752-12-31T19:31:01.000Z"); return t}()),
	// 			},
	// 			FailoverHealth: to.Ptr("Normal"),
	// 			FriendlyName: to.Ptr("f8491e4f-817a-40dd-a90c-af773978c75b"),
	// 			HealthErrors: []*armrecoveryservicessiterecovery.HealthError{
	// 			},
	// 			PolicyFriendlyName: to.Ptr("A2APolicy"),
	// 			PolicyID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/A2APolicy"),
	// 			PrimaryFabricFriendlyName: to.Ptr("cloud1"),
	// 			PrimaryFabricProvider: to.Ptr("AzureFabric"),
	// 			PrimaryProtectionContainerFriendlyName: to.Ptr("cloud_6d224fc6-f326-5d35-96de-fbf51efb3179"),
	// 			ProtectedItemType: to.Ptr(""),
	// 			ProtectionState: to.Ptr("Protected"),
	// 			ProtectionStateDescription: to.Ptr("Protected"),
	// 			ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2AReplicationDetails{
	// 				InstanceType: to.Ptr("A2A"),
	// 			},
	// 			RecoveryContainerID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud2/replicationProtectionContainers/cloud_81224fc6-f326-5d35-96de-fbf51efb3188"),
	// 			RecoveryFabricFriendlyName: to.Ptr("cloud2"),
	// 			RecoveryFabricID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud2"),
	// 			RecoveryProtectionContainerFriendlyName: to.Ptr("cloud_81224fc6-f326-5d35-96de-fbf51efb3188"),
	// 			RecoveryServicesProviderID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationRecoveryServicesProviders/6d2940f9-4c34-5989-9f56-1243a6e76ecf"),
	// 			ReplicationHealth: to.Ptr("Normal"),
	// 			TestFailoverState: to.Ptr("None"),
	// 			TestFailoverStateDescription: to.Ptr("None"),
	// 		},
	// 	}
}
