package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-11-01/WorkspaceEvaluations_List.json
func ExampleWorkspaceEvaluationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspaceEvaluationsClient().NewListPager("exampleRG", "exampleWorkspace", nil)
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
		// page = armchaos.WorkspaceEvaluationsClientListResponse{
		// 	WorkspaceEvaluationListResult: armchaos.WorkspaceEvaluationListResult{
		// 		Value: []*armchaos.WorkspaceEvaluation{
		// 			{
		// 				ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/evaluations/00000000-0000-0000-0000-000000000003"),
		// 				Name: to.Ptr("00000000-0000-0000-0000-000000000003"),
		// 				Type: to.Ptr("Microsoft.Chaos/workspaces/evaluations"),
		// 				Properties: &armchaos.WorkspaceEvaluationProperties{
		// 					Status: to.Ptr(armchaos.WorkspaceEvaluationStatusSucceeded),
		// 					StartTime: to.Ptr(time.Date(2026, time.November, 1, 10, 2, 0, 0, time.UTC)),
		// 					EndTime: to.Ptr(time.Date(2026, time.November, 1, 10, 3, 0, 0, time.UTC)),
		// 					WorkspaceID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace"),
		// 					ResourceSnapshotID: to.Ptr("11111111-1111-4111-8111-111111111111"),
		// 					DiscoveryResultID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/discoveries/00000000-0000-0000-0000-000000000002"),
		// 					NumTemplatesToEvaluate: to.Ptr[int32](2),
		// 					NumTemplatesEvaluatedSucceeded: to.Ptr[int32](2),
		// 					NumTemplatesEvaluatedFailed: to.Ptr[int32](0),
		// 					NumTemplatesEvaluatedCancelled: to.Ptr[int32](0),
		// 					Results: []*armchaos.TemplateEvaluationResultItem{
		// 						{
		// 							TemplateID: to.Ptr("/providers/Microsoft.Chaos/scenarioTemplates/CpuPressure/versions/1.0"),
		// 							TemplateName: to.Ptr("CpuPressure"),
		// 							EvaluationResult: to.Ptr(armchaos.RecommendationStatusNotApplicable),
		// 						},
		// 						{
		// 							TemplateID: to.Ptr("/providers/Microsoft.Chaos/scenarioTemplates/DnsOutage/versions/1.0"),
		// 							TemplateName: to.Ptr("DnsOutage"),
		// 							EvaluationResult: to.Ptr(armchaos.RecommendationStatusRecommended),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/evaluations?api-version=2026-11-01&continuationToken=opaque-demo-token"),
		// 	},
		// }
	}
}
