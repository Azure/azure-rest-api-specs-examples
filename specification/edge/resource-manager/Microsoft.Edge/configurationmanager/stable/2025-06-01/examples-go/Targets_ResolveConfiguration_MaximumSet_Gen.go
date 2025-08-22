package armworkloadorchestration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadorchestration/armworkloadorchestration"
)

// Generated from example definition: 2025-06-01/Targets_ResolveConfiguration_MaximumSet_Gen.json
func ExampleTargetsClient_BeginResolveConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadorchestration.NewClientFactory("9D54FE4C-00AF-4836-8F48-B6A9C4E47192", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTargetsClient().BeginResolveConfiguration(ctx, "rgconfigurationmanager", "testname", armworkloadorchestration.SolutionTemplateParameter{
		SolutionTemplateVersionID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}/{resourceType}/{resourceName}"),
		SolutionInstanceName:      to.Ptr("testname"),
		SolutionDependencies: []*armworkloadorchestration.SolutionDependencyParameter{
			{
				SolutionVersionID:       to.Ptr("cydzqntmjlqtksbavjwteru"),
				SolutionTemplateID:      to.Ptr("liqauthxnscodbiwktwfwrrsg"),
				SolutionTemplateVersion: to.Ptr("gordjasyxxrj"),
				SolutionInstanceName:    to.Ptr("testname"),
				TargetID:                to.Ptr("steadvphxtyhjokqicrtg"),
				Dependencies:            []*armworkloadorchestration.SolutionDependencyParameter{},
			},
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
	// res = armworkloadorchestration.TargetsClientResolveConfigurationResponse{
	// 	ResolvedConfiguration: &armworkloadorchestration.ResolvedConfiguration{
	// 		Configuration: to.Ptr("vjxrtzgboqltqnkbcq"),
	// 	},
	// }
}
