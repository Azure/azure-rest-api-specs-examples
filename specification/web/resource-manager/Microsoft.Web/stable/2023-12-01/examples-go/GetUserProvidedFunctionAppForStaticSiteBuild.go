package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/de1f3772629b6f4d3ac01548a5f6d719bfb97c9e/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/GetUserProvidedFunctionAppForStaticSiteBuild.json
func ExampleStaticSitesClient_GetUserProvidedFunctionAppForStaticSiteBuild() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStaticSitesClient().GetUserProvidedFunctionAppForStaticSiteBuild(ctx, "rg", "testStaticSite0", "default", "testFunctionApp", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StaticSiteUserProvidedFunctionAppARMResource = armappservice.StaticSiteUserProvidedFunctionAppARMResource{
	// 	Name: to.Ptr("testFunctionApp"),
	// 	Type: to.Ptr("Microsoft.Web/staticSites/builds/userProvidedFunctionApps"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Web/staticSites/testStaticSite0/builds/default/userProvidedFunctionApps/testFunctionApp"),
	// 	Properties: &armappservice.StaticSiteUserProvidedFunctionAppARMResourceProperties{
	// 		CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-04T17:33:11.641Z"); return t}()),
	// 		FunctionAppRegion: to.Ptr("West US 2"),
	// 		FunctionAppResourceID: to.Ptr("/subscription/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/functionRG/providers/Microsoft.Web/sites/testFunctionApp"),
	// 	},
	// }
}
