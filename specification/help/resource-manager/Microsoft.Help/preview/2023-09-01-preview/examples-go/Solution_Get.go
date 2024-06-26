package armselfhelp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/selfhelp/armselfhelp/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/help/resource-manager/Microsoft.Help/preview/2023-09-01-preview/examples/Solution_Get.json
func ExampleSolutionClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armselfhelp.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSolutionClient().Get(ctx, "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp", "SolutionResource1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SolutionResource = armselfhelp.SolutionResource{
	// 	Name: to.Ptr("SolutionResource1"),
	// 	Type: to.Ptr("Microsoft.Help/solutions"),
	// 	ID: to.Ptr("/subscriptions/mySubscription/resourceGroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp/providers/Microsoft.Help/solutions/SolutionResource1"),
	// 	Properties: &armselfhelp.SolutionResourceProperties{
	// 		Content: to.Ptr("<p>Sample content</p>"),
	// 		ProvisioningState: to.Ptr(armselfhelp.SolutionProvisioningStateSucceeded),
	// 		ReplacementMaps: &armselfhelp.ReplacementMaps{
	// 			Diagnostics: []*armselfhelp.SolutionsDiagnostic{
	// 				{
	// 					Insights: []*armselfhelp.Insight{
	// 						{
	// 							ID: to.Ptr("InsightArticleId"),
	// 							ImportanceLevel: to.Ptr(armselfhelp.ImportanceLevelCritical),
	// 							Results: to.Ptr("Article Content"),
	// 							Title: to.Ptr("An example title for an Insight"),
	// 					}},
	// 					ReplacementKey: to.Ptr("<!--12345678-BBBb-cCCCC-0000-123456789012-->"),
	// 					SolutionID: to.Ptr("sampleSolutionId3"),
	// 					Status: to.Ptr(armselfhelp.StatusSucceeded),
	// 					StatusDetails: to.Ptr(""),
	// 				},
	// 				{
	// 					Insights: []*armselfhelp.Insight{
	// 					},
	// 					ReplacementKey: to.Ptr("<!--12345678-BBBb-cCCCC-0000-123456789012-->"),
	// 					RequiredParameters: []*string{
	// 						to.Ptr("parameter1"),
	// 						to.Ptr("parameter2")},
	// 						SolutionID: to.Ptr("sampleSolutionId4"),
	// 						Status: to.Ptr(armselfhelp.StatusSucceeded),
	// 						StatusDetails: to.Ptr(""),
	// 				}},
	// 				MetricsBasedCharts: []*armselfhelp.MetricsBasedChart{
	// 					{
	// 						Name: to.Ptr("CPU_percent"),
	// 						AggregationType: to.Ptr(armselfhelp.AggregationTypeMax),
	// 						ReplacementKey: to.Ptr("<!--12345678-BBBb-cCCCC-0000-123456789012-->"),
	// 						TimeSpanDuration: to.Ptr("1d"),
	// 						Title: to.Ptr("CPU Usage in the last one day"),
	// 					},
	// 					{
	// 						Name: to.Ptr("memory_percent"),
	// 						AggregationType: to.Ptr(armselfhelp.AggregationTypeMax),
	// 						ReplacementKey: to.Ptr("<!--12345678-BBBb-cCCCC-0000-123456789012-->"),
	// 						TimeSpanDuration: to.Ptr("1d"),
	// 						Title: to.Ptr("Memory Usage in the last one day"),
	// 				}},
	// 				VideoGroups: []*armselfhelp.VideoGroup{
	// 					{
	// 						ReplacementKey: to.Ptr("<!--12345678-BBBb-cCCCC-0000-123456789012-->"),
	// 						Videos: []*armselfhelp.VideoGroupVideo{
	// 							{
	// 								Src: to.Ptr("sampleVideoSource2"),
	// 								Title: to.Ptr("widthtest"),
	// 						}},
	// 				}},
	// 				Videos: []*armselfhelp.Video{
	// 					{
	// 						Src: to.Ptr("sampleVideoSource"),
	// 						Title: to.Ptr("CI - CD with Azure DevOps"),
	// 						ReplacementKey: to.Ptr("<!--12345678-BBBb-cCCCC-0000-123456789012-->"),
	// 				}},
	// 				WebResults: []*armselfhelp.WebResult{
	// 					{
	// 						ReplacementKey: to.Ptr("<!--12345678-BBBb-cCCCC-0000-123456789012-->"),
	// 						SearchResults: []*armselfhelp.SearchResult{
	// 							{
	// 								Confidence: to.Ptr(armselfhelp.ConfidenceHigh),
	// 								Content: to.Ptr("I sysprep a VM and now I cannot connect to it via <b>RDP</b>"),
	// 								Link: to.Ptr("sampleLink"),
	// 								Rank: to.Ptr[int32](1),
	// 								ResultType: to.Ptr(armselfhelp.ResultTypeCommunity),
	// 								SolutionID: to.Ptr("sampleSolutionId2"),
	// 								Source: to.Ptr("sampleSource"),
	// 								Title: to.Ptr("Cannot RDP VM after SysPrep"),
	// 						}},
	// 				}},
	// 			},
	// 			Sections: []*armselfhelp.Section{
	// 				{
	// 					Content: to.Ptr("<p>sample content</p>"),
	// 					ReplacementMaps: &armselfhelp.ReplacementMaps{
	// 					},
	// 					Title: to.Ptr("RBAC Authentication Common Solutions"),
	// 			}},
	// 			SolutionID: to.Ptr("sampleSolutionId1"),
	// 			Title: to.Ptr("RBAC Authentication Common Solutions"),
	// 		},
	// 	}
}
