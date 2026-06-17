package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-04-01-preview/RecoveryPlanActions_ValidateForTestFailover_MaximumSet_Gen.json
func ExampleRecoveryPlanActionsClient_BeginValidateForTestFailover() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresiliencemanagement.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRecoveryPlanActionsClient().BeginValidateForTestFailover(ctx, "sampleServiceGroupName", "<operationID>", "samplePlanName", armresiliencemanagement.FailoverRequest{
		FailoverDirection: to.Ptr(armresiliencemanagement.FailoverDirectionTypesFromSpecificLocations),
		FailoverRequestProperties: &armresiliencemanagement.FailoverRequestProperties{
			SourceLocations: []*string{
				to.Ptr("westus"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armresiliencemanagement.RecoveryPlanActionsClientValidateForTestFailoverResponse{
	// 	ValidateForRecoveryOperationBaseResponse: armresiliencemanagement.ValidateForRecoveryOperationBaseResponse{
	// 		RecoveryResourceQualifications: []*armresiliencemanagement.RecoveryResourceQualification{
	// 			{
	// 				RecoveryResource: &armresiliencemanagement.RecoveryResource{
	// 					Properties: &armresiliencemanagement.RecoveryResourceProperties{
	// 						ProvisioningState: to.Ptr(armresiliencemanagement.ProvisioningStateSucceeded),
	// 						ResourceID: to.Ptr("/subscriptions/ad261018-e582-488a-815d-c2ebe28ca544/resourceGroups/sampleResourceGroupName/providers/Microsoft.Compute/virtualMachines/iaasvm-1"),
	// 						ResourceLocation: to.Ptr("westus"),
	// 						RecoveryResourceUniqueID: to.Ptr("e2a7b8d1-4c3f-4e2b-9a1c-7f6e2d8b5c4a"),
	// 						ProtectionStatus: to.Ptr(armresiliencemanagement.ResourceProtectionStatusUnknown),
	// 						InclusionState: to.Ptr(armresiliencemanagement.ResourceInclusionStateIncluded),
	// 						NeedsAttention: to.Ptr(true),
	// 						AttentionReasons: []*string{
	// 							to.Ptr("ResourceInNotProtectedState"),
	// 						},
	// 						RecoveryGroupID: to.Ptr("11111111-1111-1111-1111-123456789012"),
	// 						AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 							Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 							UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 						},
	// 						ErrorDetails: &armresiliencemanagement.ErrorDetail{
	// 							Code: to.Ptr("ResourceInNotProtectedState"),
	// 							Message: to.Ptr("Resource in not protected with any recovery solution."),
	// 							Target: to.Ptr("Please make sure resource is protected with any recovery solution."),
	// 							Details: []*armresiliencemanagement.ErrorDetail{
	// 							},
	// 							AdditionalInfo: []*armresiliencemanagement.ErrorAdditionalInfo{
	// 								{
	// 									Type: to.Ptr("Unable to detect resource protections with any recovery solution."),
	// 									Info: map[string]any{
	// 									},
	// 								},
	// 							},
	// 						},
	// 						ResourceProtectionSolutions: []*armresiliencemanagement.ResourceProtectionSolutionSettings{
	// 							{
	// 								ProtectionSolutionType: to.Ptr(armresiliencemanagement.ResourceProtectionSolutionTypeNone),
	// 								ProtectionStatus: to.Ptr(armresiliencemanagement.ResourceProtectionStatusUnknown),
	// 								ResourceID: to.Ptr("/subscriptions/ad261018-e582-488a-815d-c2ebe28ca544/resourceGroups/sampleResourceGroupName/providers/Microsoft.Compute/virtualMachines/iaasvm-1"),
	// 								ActiveLocation: to.Ptr("westus"),
	// 								RecoveryLocations: []*string{
	// 									to.Ptr("eastus"),
	// 								},
	// 								ReplicationRole: to.Ptr(armresiliencemanagement.ResourceReplicationRolePrimary),
	// 								PrimaryResource: to.Ptr("/subscriptions/ad261018-e582-488a-815d-c2ebe28ca544/resourceGroups/sampleResourceGroupName/providers/Microsoft.Compute/virtualMachines/iaasvm-1"),
	// 								ReplicaResources: []*string{
	// 									to.Ptr("/subscriptions/ad261018-e582-488a-815d-c2ebe28ca544/resourceGroups/sampleResourceGroupName/providers/Microsoft.Compute/virtualMachines/iaasvm-2"),
	// 								},
	// 								IsAutoFailover: to.Ptr(true),
	// 								FailoverState: to.Ptr(armresiliencemanagement.FailoverStateNone),
	// 								TestFailoverState: to.Ptr(armresiliencemanagement.TestFailoverStateNone),
	// 								ActiveLocations: []*string{
	// 									to.Ptr("eastus-zone2"),
	// 								},
	// 								ActivePhysicalZones: []*string{
	// 									to.Ptr("eastus2-az3"),
	// 								},
	// 							},
	// 						},
	// 						SelectedProtectionSolutionType: to.Ptr(armresiliencemanagement.ResourceProtectionSolutionTypeAzureNative),
	// 						SelectedProtectionSolutionSetting: &armresiliencemanagement.ResourceNativeProtectionSolutionSetting{
	// 							ProtectionSolutionType: to.Ptr(armresiliencemanagement.ResourceProtectionSolutionTypeAzureNative),
	// 						},
	// 						ResourcePhysicalZones: []*string{
	// 							to.Ptr("eastus-az3"),
	// 						},
	// 					},
	// 					ID: to.Ptr("/providers/Microsoft.Management/serviceGroups/sampleServiceGroupName/providers/Microsoft.AzureResilienceManagement/recoveryPlans/samplePlanName/recoveryResources/12345678-9012-3456-7890-123456789012"),
	// 					Name: to.Ptr("12345678-9012-3456-7890-123456789012"),
	// 					Type: to.Ptr("Microsoft.AzureResilienceManagement/recoveryPlans/recoveryResources"),
	// 					SystemData: &armresiliencemanagement.SystemData{
	// 						CreatedBy: to.Ptr("wmfonl"),
	// 						CreatedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-15T09:19:54.175Z"); return t}()),
	// 						LastModifiedBy: to.Ptr("paiugykk"),
	// 						LastModifiedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-15T09:19:54.176Z"); return t}()),
	// 					},
	// 				},
	// 				OperationQualificationDetails: &armresiliencemanagement.OperationQualificationDetails{
	// 					QualificationState: to.Ptr(armresiliencemanagement.QualificationStateQualified),
	// 					NotQualifiedReasons: []*string{
	// 						to.Ptr("ResourceInNotProtectedState"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
