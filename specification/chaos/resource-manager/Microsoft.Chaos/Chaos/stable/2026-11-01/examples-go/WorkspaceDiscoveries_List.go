package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-11-01/WorkspaceDiscoveries_List.json
func ExampleWorkspaceDiscoveriesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspaceDiscoveriesClient().NewListPager("exampleRG", "exampleWorkspace", nil)
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
		// page = armchaos.WorkspaceDiscoveriesClientListResponse{
		// 	WorkspaceDiscoveryListResult: armchaos.WorkspaceDiscoveryListResult{
		// 		Value: []*armchaos.WorkspaceDiscovery{
		// 			{
		// 				ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/discoveries/00000000-0000-0000-0000-000000000002"),
		// 				Name: to.Ptr("00000000-0000-0000-0000-000000000002"),
		// 				Type: to.Ptr("Microsoft.Chaos/workspaces/discoveries"),
		// 				Properties: &armchaos.WorkspaceDiscoveryProperties{
		// 					Status: to.Ptr(armchaos.WorkspaceDiscoveryStatusSucceeded),
		// 					StartTime: to.Ptr(time.Date(2026, time.November, 1, 10, 0, 0, 0, time.UTC)),
		// 					EndTime: to.Ptr(time.Date(2026, time.November, 1, 10, 1, 0, 0, time.UTC)),
		// 					WorkspaceID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace"),
		// 					ResourceSnapshotID: to.Ptr("11111111-1111-4111-8111-111111111111"),
		// 					Completeness: to.Ptr(armchaos.WorkspaceDiscoveryCompletenessEstablished),
		// 					CoverageSummary: &armchaos.WorkspaceDiscoveryCoverageSummary{
		// 						Scopes: &armchaos.WorkspaceDiscoveryScopeCoverageSummary{
		// 							TotalCount: to.Ptr[int64](1),
		// 							KnownExpectedSetCount: to.Ptr[int64](1),
		// 							UnknownExpectedSetCount: to.Ptr[int64](0),
		// 							EstablishedCount: to.Ptr[int64](1),
		// 							BaselineUnknownCount: to.Ptr[int64](0),
		// 							VerificationIncompleteCount: to.Ptr[int64](0),
		// 							KnownPartialCount: to.Ptr[int64](0),
		// 						},
		// 						Resources: &armchaos.WorkspaceDiscoveryResourceCoverageSummary{
		// 							ExpectedCount: to.Ptr[int64](2),
		// 							ClassifiedCount: to.Ptr[int64](2),
		// 							UnclassifiedCount: to.Ptr[int64](0),
		// 							CoveredCount: to.Ptr[int64](2),
		// 						},
		// 					},
		// 					Warnings: []armchaos.WorkspaceDiscoveryWarningClassification{
		// 					},
		// 					WarningsTotalCount: to.Ptr[int64](0),
		// 					WarningsTruncated: to.Ptr(false),
		// 					Errors: []*armchaos.OperationError{
		// 					},
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/discoveries?api-version=2026-11-01&continuationToken=opaque-demo-token"),
		// 	},
		// }
	}
}
