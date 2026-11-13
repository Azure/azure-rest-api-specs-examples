package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-11-01/WorkspaceDiscoveries_Get_PartialMixed.json
func ExampleWorkspaceDiscoveriesClient_Get_getAPartiallySucceededResultWithClassifiedGuidanceSystemUncertaintyAndATruncatedWarningSample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceDiscoveriesClient().Get(ctx, "exampleRG", "exampleWorkspace", "00000000-0000-0000-0000-000000000008", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armchaos.WorkspaceDiscoveriesClientGetResponse{
	// 	WorkspaceDiscovery: armchaos.WorkspaceDiscovery{
	// 		ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/discoveries/00000000-0000-0000-0000-000000000008"),
	// 		Name: to.Ptr("00000000-0000-0000-0000-000000000008"),
	// 		Type: to.Ptr("Microsoft.Chaos/workspaces/discoveries"),
	// 		Properties: &armchaos.WorkspaceDiscoveryProperties{
	// 			Status: to.Ptr(armchaos.WorkspaceDiscoveryStatusPartiallySucceeded),
	// 			WorkspaceID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace"),
	// 			ResourceSnapshotID: to.Ptr("44444444-4444-4444-8444-444444444444"),
	// 			Completeness: to.Ptr(armchaos.WorkspaceDiscoveryCompletenessVerificationIncomplete),
	// 			CoverageSummary: &armchaos.WorkspaceDiscoveryCoverageSummary{
	// 				Scopes: &armchaos.WorkspaceDiscoveryScopeCoverageSummary{
	// 					TotalCount: to.Ptr[int64](3),
	// 					KnownExpectedSetCount: to.Ptr[int64](1),
	// 					UnknownExpectedSetCount: to.Ptr[int64](2),
	// 					EstablishedCount: to.Ptr[int64](0),
	// 					BaselineUnknownCount: to.Ptr[int64](1),
	// 					VerificationIncompleteCount: to.Ptr[int64](1),
	// 					KnownPartialCount: to.Ptr[int64](1),
	// 				},
	// 				Resources: &armchaos.WorkspaceDiscoveryResourceCoverageSummary{
	// 					ExpectedCount: to.Ptr[int64](6),
	// 					ClassifiedCount: to.Ptr[int64](6),
	// 					UnclassifiedCount: to.Ptr[int64](0),
	// 					CoveredCount: to.Ptr[int64](1),
	// 				},
	// 			},
	// 			Warnings: []armchaos.WorkspaceDiscoveryWarningClassification{
	// 				&armchaos.WorkspaceDiscoveryPermissionWarning{
	// 					ScopeID: to.Ptr("/providers/Microsoft.Management/serviceGroups/exampleGroup"),
	// 					ScopeKind: to.Ptr("ServiceGroup"),
	// 					Reason: to.Ptr(armchaos.WorkspaceDiscoveryWarningReasonRBACValidationFailed),
	// 					Message: to.Ptr("Grant the required read permission, wait for propagation, and retry discovery."),
	// 					Permission: &armchaos.PermissionError{
	// 						ResourceID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/targets/providers/Microsoft.Compute/virtualMachines/deniedVm"),
	// 						MissingPermissions: []*string{
	// 							to.Ptr("Microsoft.Compute/virtualMachines/read"),
	// 						},
	// 						RequiredPermissions: []*string{
	// 							to.Ptr("Microsoft.Compute/virtualMachines/read"),
	// 						},
	// 						RecommendedRoles: []*string{
	// 							to.Ptr("Reader"),
	// 						},
	// 						Identity: &armchaos.EntraIdentity{
	// 							ObjectID: to.Ptr("55555555-5555-4555-8555-555555555555"),
	// 							TenantID: to.Ptr("66666666-6666-4666-8666-666666666666"),
	// 						},
	// 					},
	// 				},
	// 				&armchaos.WorkspaceDiscoveryResourceWarning{
	// 					ScopeID: to.Ptr("/providers/Microsoft.Management/serviceGroups/exampleGroup"),
	// 					ScopeKind: to.Ptr("ServiceGroup"),
	// 					Reason: to.Ptr(armchaos.WorkspaceDiscoveryWarningReasonTargetNotFound),
	// 					Message: to.Ptr("Remove the confirmed missing resource from this service-group membership."),
	// 					ResourceID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/targets/providers/Microsoft.Compute/virtualMachines/missingVm"),
	// 				},
	// 				&armchaos.WorkspaceDiscoveryAccessWarning{
	// 					ScopeID: to.Ptr("/providers/Microsoft.Management/serviceGroups/exampleGroup"),
	// 					ScopeKind: to.Ptr("ServiceGroup"),
	// 					Reason: to.Ptr(armchaos.WorkspaceDiscoveryWarningReasonAccessIndeterminate),
	// 					Message: to.Ptr("Grant access, or remove this membership if the resource no longer exists, then retry discovery."),
	// 					ResourceID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/targets/providers/Microsoft.Compute/virtualMachines/indeterminateVm"),
	// 				},
	// 			},
	// 			WarningsTotalCount: to.Ptr[int64](5),
	// 			WarningsTruncated: to.Ptr(true),
	// 			Errors: []*armchaos.OperationError{
	// 				{
	// 					ErrorCode: to.Ptr("ResourceDiscoveryScopeNotCompletedError"),
	// 					ErrorMessage: to.Ptr("Required child enumeration could not be completed; the accepted result could not establish all attainable coverage."),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
