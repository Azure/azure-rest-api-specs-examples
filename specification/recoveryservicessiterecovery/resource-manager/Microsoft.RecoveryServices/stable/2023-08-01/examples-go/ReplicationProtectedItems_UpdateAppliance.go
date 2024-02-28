package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationProtectedItems_UpdateAppliance.json
func ExampleReplicationProtectedItemsClient_BeginUpdateAppliance() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationProtectedItemsClient().BeginUpdateAppliance(ctx, "Ayan-0106-SA-Vault", "Ayan-0106-SA-RG", "Ayan-0106-SA-Vaultreplicationfabric", "Ayan-0106-SA-Vaultreplicationcontainer", "idclab-vcen67_50158124-8857-3e08-0893-2ddf8ebb8c1f", armrecoveryservicessiterecovery.UpdateApplianceForReplicationProtectedItemInput{
		Properties: &armrecoveryservicessiterecovery.UpdateApplianceForReplicationProtectedItemInputProperties{
			ProviderSpecificDetails: &armrecoveryservicessiterecovery.InMageRcmUpdateApplianceForReplicationProtectedItemInput{
				InstanceType:   to.Ptr("InMageRcm"),
				RunAsAccountID: to.Ptr(""),
			},
			TargetApplianceID: to.Ptr(""),
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
	// 	Name: to.Ptr("idclab-vcen67_50158124-8857-3e08-0893-2ddf8ebb8c1f"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems"),
	// 	ID: to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/Ayan-0106-SA-RG/providers/Microsoft.RecoveryServices/vaults/Ayan-0106-SA-Vault/replicationFabrics/Ayan-0106-SA-Vaultreplicationfabric/replicationProtectionContainers/Ayan-0106-SA-Vaultreplicationcontainer/replicationProtectedItems/idclab-vcen67_50158124-8857-3e08-0893-2ddf8ebb8c1f"),
	// 	Properties: &armrecoveryservicessiterecovery.ReplicationProtectedItemProperties{
	// 		ActiveLocation: to.Ptr("Primary"),
	// 		AllowedOperations: []*string{
	// 			to.Ptr("UnplannedFailover"),
	// 			to.Ptr("DisableProtection"),
	// 			to.Ptr("TestFailover")},
	// 			CurrentScenario: &armrecoveryservicessiterecovery.CurrentScenarioDetails{
	// 				JobID: to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/Ayan-0106-SA-RG/providers/Microsoft.RecoveryServices/vaults/Ayan-0106-SA-Vault/replicationJobs/None"),
	// 				ScenarioName: to.Ptr("None"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1753-01-01T01:01:01.000Z"); return t}()),
	// 			},
	// 			EventCorrelationID: to.Ptr("fb40d161-cffd-44d9-a252-0b7978e1f73c"),
	// 			FailoverHealth: to.Ptr("Normal"),
	// 			FriendlyName: to.Ptr("Ayan-RHEL7-Test2"),
	// 			HealthErrors: []*armrecoveryservicessiterecovery.HealthError{
	// 			},
	// 			PolicyFriendlyName: to.Ptr("24-hour-replication-policy"),
	// 			PolicyID: to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/Ayan-0106-SA-RG/providers/Microsoft.RecoveryServices/vaults/Ayan-0106-SA-Vault/replicationPolicies/24-hour-replication-policy"),
	// 			PrimaryFabricFriendlyName: to.Ptr("Ayan-0106-SA-Vaultreplicationfabric"),
	// 			PrimaryFabricProvider: to.Ptr("InMageRcmFabric"),
	// 			PrimaryProtectionContainerFriendlyName: to.Ptr("Ayan-0106-SA-Vaultreplicationcontainer"),
	// 			ProtectedItemType: to.Ptr(""),
	// 			ProtectionState: to.Ptr("Protected"),
	// 			ProtectionStateDescription: to.Ptr("Protected"),
	// 			ProviderSpecificDetails: &armrecoveryservicessiterecovery.InMageRcmReplicationDetails{
	// 				InstanceType: to.Ptr("InMageRcm"),
	// 			},
	// 			RecoveryContainerID: to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/Ayan-0106-SA-RG/providers/Microsoft.RecoveryServices/vaults/Ayan-0106-SA-Vault/replicationFabrics/d49858f157601230a6ac5862fbbc6e63bf38d23ecd96cf953767945d457fe9d5/replicationProtectionContainers/d38048d4-b460-4791-8ece-108395ee8478"),
	// 			RecoveryFabricFriendlyName: to.Ptr("Microsoft Azure"),
	// 			RecoveryFabricID: to.Ptr("Microsoft Azure"),
	// 			RecoveryProtectionContainerFriendlyName: to.Ptr("Microsoft Azure"),
	// 			RecoveryServicesProviderID: to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/Ayan-0106-SA-RG/providers/Microsoft.RecoveryServices/vaults/Ayan-0106-SA-Vault/replicationFabrics/Ayan-0106-SA-Vaultreplicationfabric/replicationRecoveryServicesProviders/a552cf2d-bbb3-4d78-8145-e1992ecb31d1"),
	// 			ReplicationHealth: to.Ptr("Normal"),
	// 			TestFailoverState: to.Ptr("None"),
	// 			TestFailoverStateDescription: to.Ptr("None"),
	// 		},
	// 	}
}
