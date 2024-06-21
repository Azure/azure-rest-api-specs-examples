package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/de1f3772629b6f4d3ac01548a5f6d719bfb97c9e/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/UpdateStaticSiteUser.json
func ExampleStaticSitesClient_UpdateStaticSiteUser() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStaticSitesClient().UpdateStaticSiteUser(ctx, "rg", "testStaticSite0", "aad", "1234", armappservice.StaticSiteUserARMResource{
		Properties: &armappservice.StaticSiteUserARMResourceProperties{
			Roles: to.Ptr("contributor"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StaticSiteUserARMResource = armappservice.StaticSiteUserARMResource{
	// 	Properties: &armappservice.StaticSiteUserARMResourceProperties{
	// 		DisplayName: to.Ptr("username"),
	// 		Provider: to.Ptr("aad"),
	// 		Roles: to.Ptr("contributor,anonymous,authenticated"),
	// 		UserID: to.Ptr("1234"),
	// 	},
	// }
}
