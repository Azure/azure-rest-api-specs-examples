package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-11-01/WorkspaceDiscoveries_Get_BaselineUnknown.json
func ExampleWorkspaceDiscoveriesClient_Get_getAHealthyContainerScanWithAStructurallyUnknownInventoryDenominator() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceDiscoveriesClient().Get(ctx, "exampleRG", "exampleWorkspace", "00000000-0000-0000-0000-000000000006", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armchaos.WorkspaceDiscoveriesClientGetResponse{
	// 	WorkspaceDiscovery: armchaos.WorkspaceDiscovery{
	// 		ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/discoveries/00000000-0000-0000-0000-000000000006"),
	// 		Name: to.Ptr("00000000-0000-0000-0000-000000000006"),
	// 		Type: to.Ptr("Microsoft.Chaos/workspaces/discoveries"),
	// 		Properties: &armchaos.WorkspaceDiscoveryProperties{
	// 			Status: to.Ptr(armchaos.WorkspaceDiscoveryStatusSucceeded),
	// 			StartTime: to.Ptr(time.Date(2026, time.November, 1, 10, 0, 0, 0, time.UTC)),
	// 			EndTime: to.Ptr(time.Date(2026, time.November, 1, 10, 1, 0, 0, time.UTC)),
	// 			WorkspaceID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace"),
	// 			ResourceSnapshotID: to.Ptr("22222222-2222-4222-8222-222222222222"),
	// 			Completeness: to.Ptr(armchaos.WorkspaceDiscoveryCompletenessBaselineUnknown),
	// 			CoverageSummary: &armchaos.WorkspaceDiscoveryCoverageSummary{
	// 				Scopes: &armchaos.WorkspaceDiscoveryScopeCoverageSummary{
	// 					TotalCount: to.Ptr[int64](1),
	// 					KnownExpectedSetCount: to.Ptr[int64](0),
	// 					UnknownExpectedSetCount: to.Ptr[int64](1),
	// 					EstablishedCount: to.Ptr[int64](0),
	// 					BaselineUnknownCount: to.Ptr[int64](1),
	// 					VerificationIncompleteCount: to.Ptr[int64](0),
	// 					KnownPartialCount: to.Ptr[int64](0),
	// 				},
	// 			},
	// 			Warnings: []armchaos.WorkspaceDiscoveryWarningClassification{
	// 			},
	// 			WarningsTotalCount: to.Ptr[int64](0),
	// 			WarningsTruncated: to.Ptr(false),
	// 			Errors: []*armchaos.OperationError{
	// 			},
	// 		},
	// 	},
	// }
}
