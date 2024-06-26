package armselfhelp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/selfhelp/armselfhelp/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/help/resource-manager/Microsoft.Help/preview/2023-09-01-preview/examples/ListDiscoverySolutionsAtSubscriptionScope.json
func ExampleDiscoverySolutionClient_NewListPager_listDiscoverySolutionsAtSubscriptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armselfhelp.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDiscoverySolutionClient().NewListPager("subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6", &armselfhelp.DiscoverySolutionClientListOptions{Filter: to.Ptr("ProblemClassificationId eq 'SampleProblemClassificationId1'"),
		Skiptoken: nil,
	})
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
		// page.DiscoveryResponse = armselfhelp.DiscoveryResponse{
		// 	Value: []*armselfhelp.SolutionMetadataResource{
		// 		{
		// 			Name: to.Ptr("SampleProblemClassificationId1"),
		// 			Type: to.Ptr("Microsoft.Help/discoverySolutions"),
		// 			ID: to.Ptr("/subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6/providers/Microsoft.Help/discoverySolutions/SampleProblemClassificationId1"),
		// 			Properties: &armselfhelp.Solutions{
		// 				Solutions: []*armselfhelp.SolutionMetadataProperties{
		// 					{
		// 						Description: to.Ptr("This is an azure solution to troubleshoot subscription issues."),
		// 						RequiredInputs: []*string{
		// 							to.Ptr("SubscriptionId")},
		// 							SolutionID: to.Ptr("SampleSolutionId1"),
		// 							SolutionType: to.Ptr(armselfhelp.SolutionTypeDiagnostics),
		// 						},
		// 						{
		// 							Description: to.Ptr("This is an azure solution to troubleshoot subscription issues."),
		// 							RequiredInputs: []*string{
		// 								to.Ptr("SubscriptionId")},
		// 								SolutionID: to.Ptr("SampleSolutionId2"),
		// 								SolutionType: to.Ptr(armselfhelp.SolutionTypeSolutions),
		// 						}},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("SampleProblemClassificationId2"),
		// 					Type: to.Ptr("Microsoft.Help/discoverySolutions"),
		// 					ID: to.Ptr("/subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6/providers/Microsoft.Help/discoverySolutions/SampleProblemClassificationId2"),
		// 					Properties: &armselfhelp.Solutions{
		// 						Solutions: []*armselfhelp.SolutionMetadataProperties{
		// 							{
		// 								Description: to.Ptr("This is an azure solution to troubleshoot subscription issues."),
		// 								RequiredInputs: []*string{
		// 									to.Ptr("SubscriptionId")},
		// 									SolutionID: to.Ptr("SampleSolutionId3"),
		// 									SolutionType: to.Ptr(armselfhelp.SolutionTypeDiagnostics),
		// 							}},
		// 						},
		// 				}},
		// 			}
	}
}
