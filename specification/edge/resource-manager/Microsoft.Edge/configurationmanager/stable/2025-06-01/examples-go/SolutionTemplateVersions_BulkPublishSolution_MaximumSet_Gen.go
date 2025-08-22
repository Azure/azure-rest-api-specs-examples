package armworkloadorchestration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadorchestration/armworkloadorchestration"
)

// Generated from example definition: 2025-06-01/SolutionTemplateVersions_BulkPublishSolution_MaximumSet_Gen.json
func ExampleSolutionTemplateVersionsClient_BeginBulkPublishSolution() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadorchestration.NewClientFactory("9D54FE4C-00AF-4836-8F48-B6A9C4E47192", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSolutionTemplateVersionsClient().BeginBulkPublishSolution(ctx, "rgconfigurationmanager", "testname", "1.0.0", armworkloadorchestration.BulkPublishSolutionParameter{
		Targets: []*armworkloadorchestration.BulkPublishTargetDetails{
			{
				TargetID:             to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Edge/Targets/target"),
				SolutionInstanceName: to.Ptr("test-instance"),
			},
		},
		SolutionInstanceName: to.Ptr("test-instance"),
		SolutionDependencies: []*armworkloadorchestration.SolutionDependencyParameter{
			{
				SolutionVersionID:       to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Edge/Targets/target/Solutions/solution/Versions/solution-1.0.0.1"),
				SolutionTemplateID:      to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Edge/SolutionTemplates/st"),
				SolutionTemplateVersion: to.Ptr("1.0.0"),
				SolutionInstanceName:    to.Ptr("test-instance"),
				TargetID:                to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Edge/Targets/target"),
				Dependencies:            []*armworkloadorchestration.SolutionDependencyParameter{},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
