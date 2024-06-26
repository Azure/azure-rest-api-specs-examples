package armselfhelp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/selfhelp/armselfhelp/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c3cc9abe085093ba880ee3eeb792edb4fa789553/specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/DiscoverSolutionsAtTenantScope.json
func ExampleDiscoverySolutionNLPClient_DiscoverSolutions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armselfhelp.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDiscoverySolutionNLPClient().DiscoverSolutions(ctx, &armselfhelp.DiscoverySolutionNLPClientDiscoverSolutionsOptions{DiscoverSolutionRequest: &armselfhelp.DiscoveryNlpRequest{
		IssueSummary: to.Ptr("how to retrieve certs from deleted keyvault."),
		ServiceID:    to.Ptr("0d0fcd2e-c4fd-4349-8497-200edb39s3ca"),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiscoveryNlpResponse = armselfhelp.DiscoveryNlpResponse{
	// 	Value: []*armselfhelp.SolutionNlpMetadataResource{
	// 		{
	// 			Name: to.Ptr("SampleProblemClassificationId1"),
	// 			Type: to.Ptr("Microsoft.Help/discoverySolutions"),
	// 			ID: to.Ptr("/subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6/providers/Microsoft.Help/discoverSolutions/SampleProblemClassificationId1"),
	// 			Properties: &armselfhelp.NlpSolutions{
	// 				ProblemClassificationID: to.Ptr("SampleProblemClassificationId1"),
	// 				ProblemDescription: to.Ptr("SampleDescription1"),
	// 				ProblemTitle: to.Ptr("SampleTitle1"),
	// 				RelatedServices: []*armselfhelp.ClassificationService{
	// 					{
	// 						DisplayName: to.Ptr("SQL Server in VM - Linux"),
	// 						ResourceTypes: []*string{
	// 							to.Ptr("MICROSOFT.CLASSICCOMPUTE/VIRTUALMACHINES"),
	// 							to.Ptr("MICROSOFT.COMPUTE/VIRTUALMACHINES")},
	// 							ServiceID: to.Ptr("/providers/Microsoft.Support/services/40ef020e-8ae7-8d57-b538-9153c47cee69"),
	// 					}},
	// 					ServiceID: to.Ptr("SampleServiceId1"),
	// 					Solutions: []*armselfhelp.SolutionMetadataProperties{
	// 						{
	// 							Description: to.Ptr("This is an azure solution to troubleshoot subscription issues."),
	// 							RequiredInputs: []*string{
	// 								to.Ptr("SubscriptionId")},
	// 								SolutionID: to.Ptr("SampleSolutionId1"),
	// 								SolutionType: to.Ptr(armselfhelp.SolutionTypeDiagnostics),
	// 							},
	// 							{
	// 								Description: to.Ptr("This is an azure solution to troubleshoot keyvault resource."),
	// 								RequiredInputs: []*string{
	// 									to.Ptr("SubscriptionId")},
	// 									SolutionID: to.Ptr("SampleSolutionId2"),
	// 									SolutionType: to.Ptr(armselfhelp.SolutionTypeSolutions),
	// 							}},
	// 						},
	// 				}},
	// 			}
}
