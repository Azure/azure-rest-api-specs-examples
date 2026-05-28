package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: 2015-05-01/AnnotationsList.json
func ExampleAnnotationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAnnotationsClient().NewListPager("my-resource-group", "my-component", "2018-02-05T00%253A30%253A00.000Z", "2018-02-06T00%253A33A00.000Z", nil)
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
		// page = armapplicationinsights.AnnotationsClientListResponse{
		// 	AnnotationsListResult: armapplicationinsights.AnnotationsListResult{
		// 		Value: []*armapplicationinsights.Annotation{
		// 			{
		// 				AnnotationName: to.Ptr("InsightsPortal-20180126-1"),
		// 				Category: to.Ptr("Deployment"),
		// 				EventTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-29T20:30:31+00:00"); return t}()),
		// 				ID: to.Ptr("4de4524f-fca5-44a8-b7e6-a67d5107796e"),
		// 				Properties: to.Ptr("{\"ReleaseDefinitionName\":\"InsightsPortal-PPEProd-KeyVault\",\"ReleaseRequestedFor\":\"6a970e9b-6220-47f3-a78c-b8be97506698\",\"TeamFoundationCollectionUri\":\"https://mseng.visualstudio.com/\",\"BuildNumber\":\"20180126.1\",\"ReleaseDescription\":\"Triggered by InsightsPortal_master_PROD1_vNext 20180126.1.\",\"ReleaseId\":\"31075\",\"ReleaseWebUrl\":\"https://mseng.visualstudio.com/96a62c4a-58c2-4dbb-94b6-5979ebc7f2af/_release?releaseId=31075\\u0026_a=release-summary\",\"SourceBranch\":\"refs/heads/master\",\"BuildRepositoryProvider\":\"TfsGit\",\"ReleaseEnvironmentName\":\"AIMON VIP SWAP\",\"BuildRepositoryName\":\"InsightsPortal\",\"ReleaseName\":\"InsightsPortal-20180126-1\"}"),
		// 			},
		// 			{
		// 				AnnotationName: to.Ptr("InsightsPortal-20180125-1"),
		// 				Category: to.Ptr("Deployment"),
		// 				EventTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-29T20:38:35+00:00"); return t}()),
		// 				ID: to.Ptr("9a82214f-8f98-4699-b5bc-cad6101ed223"),
		// 				Properties: to.Ptr("{\"ReleaseDefinitionName\":\"InsightsPortal-PPEProd-KeyVault\",\"ReleaseRequestedFor\":\"6a970e9b-6220-47f3-a78c-b8be97506698\",\"TeamFoundationCollectionUri\":\"https://mseng.visualstudio.com/\",\"BuildNumber\":\"20180125.1\",\"ReleaseDescription\":\"Triggered by InsightsPortal_master_PROD1_vNext 20180125.1.\",\"ReleaseId\":\"31035\",\"ReleaseWebUrl\":\"https://mseng.visualstudio.com/96a62c4a-58c2-4dbb-94b6-5979ebc7f2af/_release?releaseId=31035\\u0026_a=release-summary\",\"SourceBranch\":\"refs/heads/master\",\"BuildRepositoryProvider\":\"TfsGit\",\"ReleaseEnvironmentName\":\"Prod VIP Swap \",\"BuildRepositoryName\":\"InsightsPortal\",\"ReleaseName\":\"InsightsPortal-20180125-1\"}"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
