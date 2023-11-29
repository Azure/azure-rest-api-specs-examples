package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationProtectedItems_UpdateMobilityService.json
func ExampleReplicationProtectedItemsClient_BeginUpdateMobilityService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationProtectedItemsClient().BeginUpdateMobilityService(ctx, "WCUSVault", "wcusValidations", "WIN-JKKJ31QI8U2", "cloud_c6780228-83bd-4f3e-a70e-cb46b7da33a0", "79dd20ab-2b40-11e7-9791-0050568f387e", armrecoveryservicessiterecovery.UpdateMobilityServiceRequest{
		Properties: &armrecoveryservicessiterecovery.UpdateMobilityServiceRequestProperties{
			RunAsAccountID: to.Ptr("2"),
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
	// 	Name: to.Ptr("79dd20ab-2b40-11e7-9791-0050568f387e"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems"),
	// 	ID: to.Ptr("/Subscriptions/b364ed8d-4279-4bf8-8fd1-56f8fa0ae05c/resourceGroups/wcusValidations/providers/Microsoft.RecoveryServices/vaults/WCUSVault/replicationFabrics/d70b0326a201008a953505ef271dc908e5e23468bc7356862ea178696f5f15c7/replicationProtectionContainers/cloud_c6780228-83bd-4f3e-a70e-cb46b7da33a0/replicationProtectedItems/79dd20ab-2b40-11e7-9791-0050568f387e"),
	// 	Properties: &armrecoveryservicessiterecovery.ReplicationProtectedItemProperties{
	// 		ActiveLocation: to.Ptr("Primary"),
	// 		AllowedOperations: []*string{
	// 			to.Ptr("UnplannedFailover"),
	// 			to.Ptr("DisableProtection"),
	// 			to.Ptr("TestFailover")},
	// 			CurrentScenario: &armrecoveryservicessiterecovery.CurrentScenarioDetails{
	// 				JobID: to.Ptr("/Subscriptions/b364ed8d-4279-4bf8-8fd1-56f8fa0ae05c/resourceGroups/wcusValidations/providers/Microsoft.RecoveryServices/vaults/WCUSVault/replicationJobs/None"),
	// 				ScenarioName: to.Ptr("None"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1753-01-01T01:01:01.000Z"); return t}()),
	// 			},
	// 			FriendlyName: to.Ptr("MMR-LIN-V2A-3"),
	// 			PolicyFriendlyName: to.Ptr("MadhaviPolicyNew"),
	// 			PolicyID: to.Ptr("/Subscriptions/b364ed8d-4279-4bf8-8fd1-56f8fa0ae05c/resourceGroups/wcusValidations/providers/Microsoft.RecoveryServices/vaults/WCUSVault/replicationPolicies/MadhaviPolicyNew"),
	// 			PrimaryFabricFriendlyName: to.Ptr("WIN-JKKJ31QI8U2"),
	// 			PrimaryProtectionContainerFriendlyName: to.Ptr("WIN-JKKJ31QI8U2"),
	// 			ProtectableItemID: to.Ptr("/Subscriptions/b364ed8d-4279-4bf8-8fd1-56f8fa0ae05c/resourceGroups/wcusValidations/providers/Microsoft.RecoveryServices/vaults/WCUSVault/replicationFabrics/d70b0326a201008a953505ef271dc908e5e23468bc7356862ea178696f5f15c7/replicationProtectionContainers/cloud_c6780228-83bd-4f3e-a70e-cb46b7da33a0/replicationProtectableItems/79dd20ab-2b40-11e7-9791-0050568f387e"),
	// 			ProtectedItemType: to.Ptr(""),
	// 			ProtectionState: to.Ptr("Protected"),
	// 			ProtectionStateDescription: to.Ptr("Protected"),
	// 			ProviderSpecificDetails: &armrecoveryservicessiterecovery.InMageAzureV2ReplicationDetails{
	// 				InstanceType: to.Ptr("InMageAzureV2"),
	// 			},
	// 			RecoveryContainerID: to.Ptr("/Subscriptions/b364ed8d-4279-4bf8-8fd1-56f8fa0ae05c/resourceGroups/wcusValidations/providers/Microsoft.RecoveryServices/vaults/WCUSVault/replicationFabrics/d49858f157601230a6ac5862fbbc6e63bf38d23ecd96cf953767945d457fe9d5/replicationProtectionContainers/d38048d4-b460-4791-8ece-108395ee8478"),
	// 			RecoveryFabricFriendlyName: to.Ptr("Microsoft Azure"),
	// 			RecoveryFabricID: to.Ptr("Microsoft Azure"),
	// 			RecoveryProtectionContainerFriendlyName: to.Ptr("Microsoft Azure"),
	// 			RecoveryServicesProviderID: to.Ptr("/Subscriptions/b364ed8d-4279-4bf8-8fd1-56f8fa0ae05c/resourceGroups/wcusValidations/providers/Microsoft.RecoveryServices/vaults/WCUSVault/replicationFabrics/d70b0326a201008a953505ef271dc908e5e23468bc7356862ea178696f5f15c7/replicationRecoveryServicesProviders/c6780228-83bd-4f3e-a70e-cb46b7da33a0"),
	// 			ReplicationHealth: to.Ptr("Normal"),
	// 			TestFailoverState: to.Ptr("None"),
	// 			TestFailoverStateDescription: to.Ptr("None"),
	// 		},
	// 	}
}
