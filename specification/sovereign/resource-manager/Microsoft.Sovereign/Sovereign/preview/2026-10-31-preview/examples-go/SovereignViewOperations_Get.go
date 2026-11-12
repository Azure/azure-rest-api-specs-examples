package armsovereign_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sovereign/armsovereign"
)

// Generated from example definition: 2026-10-31-preview/SovereignViewOperations_Get.json
func ExampleViewOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsovereign.NewClientFactory("22b62f96-c8d5-4131-97f8-f820a5a3c55f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewViewOperationsClient().Get(ctx, "ExampleResourceGroup", "ExampleSV", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsovereign.ViewOperationsClientGetResponse{
	// 	ViewResource: armsovereign.ViewResource{
	// 		Properties: &armsovereign.ViewResourceProperties{
	// 			Scope: &armsovereign.ViewScope{
	// 				ManagementGroups: []*armsovereign.ViewManagementGroupScope{
	// 					{
	// 						ID: to.Ptr("/providers/Microsoft.Management/managementGroups/ExampleManagementGroup"),
	// 						Name: to.Ptr("ExampleManagementGroup"),
	// 					},
	// 				},
	// 				Subscriptions: []*armsovereign.ViewSubscriptionScope{
	// 					{
	// 						ID: to.Ptr("22b62f96-c8d5-4131-97f8-f820a5a3c55f"),
	// 						Name: to.Ptr("ExampleSubscription1"),
	// 					},
	// 				},
	// 			},
	// 			PolicyInitiatives: []*armsovereign.ViewPolicyInitiativeDetail{
	// 				{
	// 					PolicyInitiativeID: to.Ptr("/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f/providers/Microsoft.Authorization/policySetDefinitions/ExamplePolicySet"),
	// 					Dimension: to.Ptr(armsovereign.ComplianceDimensionEncryptionAtRest),
	// 					PolicyVersion: to.Ptr("1.0.0"),
	// 					PolicyAssignmentIDs: []*string{
	// 						to.Ptr("/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f/providers/Microsoft.Authorization/policyAssignments/ExamplePolicyAssignment"),
	// 					},
	// 				},
	// 				{
	// 					PolicyInitiativeID: to.Ptr("/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f/providers/Microsoft.Authorization/policySetDefinitions/ExampleBusinessContinuityPolicySet"),
	// 					Dimension: to.Ptr(armsovereign.ComplianceDimensionBusinessContinuity),
	// 					PolicyVersion: to.Ptr("1.0.0"),
	// 					PolicyAssignmentIDs: []*string{
	// 						to.Ptr("/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f/providers/Microsoft.Authorization/policyAssignments/ExampleBusinessContinuityPolicyAssignment"),
	// 					},
	// 				},
	// 			},
	// 			RegulatoryDocumentName: to.Ptr(armsovereign.ViewRegulatoryDocumentNameSovereignControlBaseline),
	// 			IsWorkload: to.Ptr(true),
	// 			ProvisioningState: to.Ptr(armsovereign.ProvisioningStateSucceeded),
	// 			DimensionSettings: &armsovereign.ViewDimensionSettings{
	// 				DataResidency: &armsovereign.ViewDataResidencyDimensionSettings{
	// 					ResourceEvaluationScope: to.Ptr(armsovereign.DataResidencyResourceEvaluationScopeResidency360),
	// 					Parameters: &armsovereign.ViewDataResidencyParameters{
	// 						AllowedLocations: []*string{
	// 							to.Ptr("northeurope"),
	// 							to.Ptr("swedencentral"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Identity: &armsovereign.SystemAssignedServiceIdentity{
	// 			PrincipalID: to.Ptr("e7953ada-b215-4c82-9a4c-7b8dd6e09beb"),
	// 			TenantID: to.Ptr("8a7a0a60-587f-4774-af1f-4cbf9f1edb6b"),
	// 			Type: to.Ptr(armsovereign.SystemAssignedServiceIdentityTypeSystemAssigned),
	// 		},
	// 		Tags: map[string]*string{
	// 			"tag1": to.Ptr("MCFS"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f/resourceGroups/ExampleResourceGroup/providers/Microsoft.Sovereign/sovereignViews/ExampleSV"),
	// 		Name: to.Ptr("ExampleSV"),
	// 		Type: to.Ptr("Microsoft.Sovereign/sovereignViews"),
	// 		Location: to.Ptr("northeurope"),
	// 		SystemData: &armsovereign.SystemData{
	// 			CreatedBy: to.Ptr("example@example.com"),
	// 			CreatedByType: to.Ptr(armsovereign.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2026, time.March, 15, 10, 12, 0, 0, time.UTC)),
	// 			LastModifiedBy: to.Ptr("example@example.com"),
	// 			LastModifiedByType: to.Ptr(armsovereign.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2026, time.March, 15, 11, 45, 0, 0, time.UTC)),
	// 		},
	// 	},
	// }
}
