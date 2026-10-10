package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-10-01/RecoveryResources_Get_CrossZoneVMRecovery.json
func ExampleRecoveryResourcesClient_Get_recoveryResourcesGetCrossZoneVMRecovery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresiliencemanagement.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecoveryResourcesClient().Get(ctx, "sampleServiceGroupName", "samplePlanName", "12345678-9012-3456-7890-123456789012", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armresiliencemanagement.RecoveryResourcesClientGetResponse{
	// 	RecoveryResource: armresiliencemanagement.RecoveryResource{
	// 		Properties: &armresiliencemanagement.RecoveryResourceProperties{
	// 			RecoveryResourceUniqueID: to.Ptr("e2a7b8d1-4c3f-4e2b-9a1c-7f6e2d8b5c4a"),
	// 			ProtectionStatus: to.Ptr(armresiliencemanagement.ResourceProtectionStatusProtected),
	// 			SelectedProtectionSolutionType: to.Ptr(armresiliencemanagement.ResourceProtectionSolutionTypeCrossZoneVMRecovery),
	// 			SelectedProtectionSolutionSetting: &armresiliencemanagement.ResourceCrossZoneVMRecoveryProtectionSetting{
	// 				ProtectionSolutionType: to.Ptr(armresiliencemanagement.ResourceProtectionSolutionTypeCrossZoneVMRecovery),
	// 				TargetZone: to.Ptr("2"),
	// 				CapacityReservationGroupID: to.Ptr("/subscriptions/ad261018-e582-488a-815d-c2ebe28ca544/resourceGroups/sampleResourceGroupName/providers/Microsoft.Compute/capacityReservationGroups/sampleCrg"),
	// 			},
	// 			RecoveryGroupID: to.Ptr("11111111-1111-1111-1111-123456789012"),
	// 			AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 				Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 				UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 			},
	// 			ProvisioningState: to.Ptr(armresiliencemanagement.ProvisioningStateSucceeded),
	// 			ResourceID: to.Ptr("/subscriptions/ad261018-e582-488a-815d-c2ebe28ca544/resourceGroups/sampleResourceGroupName/providers/Microsoft.Compute/virtualMachines/iaasvm-1"),
	// 			ResourceLocation: to.Ptr("eastus2"),
	// 			ResourcePhysicalZones: []*string{
	// 				to.Ptr("eastus-az1"),
	// 				to.Ptr("eastus-az2"),
	// 			},
	// 			InclusionState: to.Ptr(armresiliencemanagement.ResourceInclusionStateIncluded),
	// 			NeedsAttention: to.Ptr(false),
	// 			ResourceProtectionSolutions: []*armresiliencemanagement.ResourceProtectionSolutionSettings{
	// 				{
	// 					ProtectionSolutionType: to.Ptr(armresiliencemanagement.ResourceProtectionSolutionTypeCrossZoneVMRecovery),
	// 					ProtectionStatus: to.Ptr(armresiliencemanagement.ResourceProtectionStatusProtected),
	// 					ResourceID: to.Ptr("/subscriptions/ad261018-e582-488a-815d-c2ebe28ca544/resourceGroups/sampleResourceGroupName/providers/Microsoft.Compute/virtualMachines/iaasvm-1"),
	// 					ActiveLocation: to.Ptr("eastus-zone1"),
	// 					RecoveryLocations: []*string{
	// 						to.Ptr("eastus-zone2"),
	// 					},
	// 					ReplicationRole: to.Ptr(armresiliencemanagement.ResourceReplicationRolePrimary),
	// 					IsAutoFailover: to.Ptr(true),
	// 					FailoverState: to.Ptr(armresiliencemanagement.FailoverStateNone),
	// 					TestFailoverState: to.Ptr(armresiliencemanagement.TestFailoverStateNone),
	// 					ActivePhysicalZones: []*string{
	// 						to.Ptr("eastus-az1"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/providers/Microsoft.Management/serviceGroups/sampleServiceGroupName/providers/Microsoft.AzureResilienceManagement/recoveryPlans/samplePlanName/recoveryResources/12345678-9012-3456-7890-123456789012"),
	// 		Name: to.Ptr("12345678-9012-3456-7890-123456789012"),
	// 		Type: to.Ptr("Microsoft.AzureResilienceManagement/recoveryPlans/recoveryResources"),
	// 		SystemData: &armresiliencemanagement.SystemData{
	// 			CreatedBy: to.Ptr("wmfonl"),
	// 			CreatedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2025, time.February, 15, 9, 19, 54, 175000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("paiugykk"),
	// 			LastModifiedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2025, time.February, 15, 9, 19, 54, 176000000, time.UTC)),
	// 		},
	// 	},
	// }
}
