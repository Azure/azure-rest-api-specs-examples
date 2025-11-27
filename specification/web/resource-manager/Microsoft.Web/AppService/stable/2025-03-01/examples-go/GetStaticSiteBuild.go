package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetStaticSiteBuild.json
func ExampleStaticSitesClient_GetStaticSiteBuild() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStaticSitesClient().GetStaticSiteBuild(ctx, "rg", "testStaticSite0", "12", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StaticSiteBuildARMResource = armappservice.StaticSiteBuildARMResource{
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Web/staticSites/testStaticSite0/builds/12"),
	// 	Properties: &armappservice.StaticSiteBuildARMResourceProperties{
	// 		BuildID: to.Ptr("12"),
	// 		CreatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-04T17:33:11.641Z"); return t}()),
	// 		Hostname: to.Ptr("happy-sea-15afae3e-12.westus2.azurestaticwebsites.net"),
	// 		LastUpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-04T17:33:11.641Z"); return t}()),
	// 		LinkedBackends: []*armappservice.StaticSiteLinkedBackend{
	// 			{
	// 				BackendResourceID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.ApiManagement/service/apimService0"),
	// 				CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-26T20:57:24.805Z"); return t}()),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 		}},
	// 		PullRequestTitle: to.Ptr("Update README.md"),
	// 		SourceBranch: to.Ptr("pr-branch"),
	// 		Status: to.Ptr(armappservice.BuildStatusUploading),
	// 	},
	// }
}
