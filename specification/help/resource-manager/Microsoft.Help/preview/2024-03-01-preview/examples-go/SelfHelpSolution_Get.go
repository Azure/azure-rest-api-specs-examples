package armselfhelp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/selfhelp/armselfhelp/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c3cc9abe085093ba880ee3eeb792edb4fa789553/specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/SelfHelpSolution_Get.json
func ExampleSolutionSelfHelpClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armselfhelp.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSolutionSelfHelpClient().Get(ctx, "SolutionId1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SolutionResourceSelfHelp = armselfhelp.SolutionResourceSelfHelp{
	// 	Name: to.Ptr("SolutionId1"),
	// 	Type: to.Ptr("Microsoft.Help/SelfHelp"),
	// 	ID: to.Ptr("/providers/Microsoft.KeyVault/vaults/test-keyvault-rp/providers/Microsoft.Help/selfHelp/SolutionId1"),
	// 	Properties: &armselfhelp.SolutionsResourcePropertiesSelfHelp{
	// 		Content: to.Ptr("<p>sample content</p>"),
	// 		ReplacementMaps: &armselfhelp.ReplacementMapsSelfHelp{
	// 			VideoGroups: []*armselfhelp.VideoGroup{
	// 				{
	// 					ReplacementKey: to.Ptr("<!--12345678-BBBb-cCCCC-0000-123456789012-->"),
	// 					Videos: []*armselfhelp.VideoGroupVideo{
	// 						{
	// 							Src: to.Ptr("sampleLink"),
	// 							Title: to.Ptr("widthtest"),
	// 					}},
	// 			}},
	// 			Videos: []*armselfhelp.Video{
	// 				{
	// 					Src: to.Ptr("sampleLink"),
	// 					Title: to.Ptr("CI - CD with Azure DevOps"),
	// 					ReplacementKey: to.Ptr("<!--12345678-BBBb-cCCCC-0000-123456789012-->"),
	// 			}},
	// 			WebResults: []*armselfhelp.WebResult{
	// 				{
	// 					ReplacementKey: to.Ptr("<!--12345678-BBBb-cCCCC-0000-123456789012-->"),
	// 					SearchResults: []*armselfhelp.SearchResult{
	// 						{
	// 							Confidence: to.Ptr(armselfhelp.ConfidenceHigh),
	// 							Content: to.Ptr("I sysprep a VM and now I cannot connect to it via <b>RDP</b>"),
	// 							Link: to.Ptr("sampleLink"),
	// 							Rank: to.Ptr[int32](1),
	// 							ResultType: to.Ptr(armselfhelp.ResultTypeCommunity),
	// 							SolutionID: to.Ptr("SolutionId1"),
	// 							Source: to.Ptr("sampleSource"),
	// 							Title: to.Ptr("Cannot RDP VM after SysPrep"),
	// 					}},
	// 			}},
	// 		},
	// 		Sections: []*armselfhelp.SectionSelfHelp{
	// 			{
	// 				Content: to.Ptr("<p>sample content</p>"),
	// 				ReplacementMaps: &armselfhelp.ReplacementMapsSelfHelp{
	// 				},
	// 				Title: to.Ptr("RBAC Authentication Common Solutions"),
	// 		}},
	// 		SolutionID: to.Ptr("sampleSolutionId"),
	// 		Title: to.Ptr("RBAC Authentication Common Solutions"),
	// 	},
	// }
}
