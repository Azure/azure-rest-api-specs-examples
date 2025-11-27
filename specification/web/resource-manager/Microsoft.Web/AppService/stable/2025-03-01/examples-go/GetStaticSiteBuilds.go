package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetStaticSiteBuilds.json
func ExampleStaticSitesClient_NewGetStaticSiteBuildsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStaticSitesClient().NewGetStaticSiteBuildsPager("rg", "testStaticSite0", nil)
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
		// page.StaticSiteBuildCollection = armappservice.StaticSiteBuildCollection{
		// 	Value: []*armappservice.StaticSiteBuildARMResource{
		// 		{
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Web/staticSites/testStaticSite0/builds/default"),
		// 			Properties: &armappservice.StaticSiteBuildARMResourceProperties{
		// 				BuildID: to.Ptr("default"),
		// 				CreatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-04T17:33:11.641Z"); return t}()),
		// 				Hostname: to.Ptr("happy-sea-15afae3e.azurestaticwebsites.net"),
		// 				LastUpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-04T17:33:11.641Z"); return t}()),
		// 				LinkedBackends: []*armappservice.StaticSiteLinkedBackend{
		// 				},
		// 				SourceBranch: to.Ptr("demo"),
		// 				Status: to.Ptr(armappservice.BuildStatusReady),
		// 			},
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Web/staticSites/testStaticSite0/builds/12"),
		// 			Properties: &armappservice.StaticSiteBuildARMResourceProperties{
		// 				BuildID: to.Ptr("12"),
		// 				CreatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-04T17:33:11.641Z"); return t}()),
		// 				Hostname: to.Ptr("happy-sea-15afae3e-12.westus2.azurestaticwebsites.net"),
		// 				LastUpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-04T17:33:11.641Z"); return t}()),
		// 				LinkedBackends: []*armappservice.StaticSiteLinkedBackend{
		// 					{
		// 						BackendResourceID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.ApiManagement/service/apimService0"),
		// 						CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-26T20:57:24.805Z"); return t}()),
		// 						ProvisioningState: to.Ptr("Succeeded"),
		// 				}},
		// 				PullRequestTitle: to.Ptr("Update README.md"),
		// 				SourceBranch: to.Ptr("demo-patch2"),
		// 				Status: to.Ptr(armappservice.BuildStatusReady),
		// 			},
		// 	}},
		// }
	}
}
