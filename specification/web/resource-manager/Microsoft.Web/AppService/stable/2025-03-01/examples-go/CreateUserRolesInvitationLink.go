package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/CreateUserRolesInvitationLink.json
func ExampleStaticSitesClient_CreateUserRolesInvitationLink() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStaticSitesClient().CreateUserRolesInvitationLink(ctx, "rg", "testStaticSite0", armappservice.StaticSiteUserInvitationRequestResource{
		Properties: &armappservice.StaticSiteUserInvitationRequestResourceProperties{
			Domain:               to.Ptr("happy-sea-15afae3e.azurestaticwebsites.net"),
			NumHoursToExpiration: to.Ptr[int32](1),
			Provider:             to.Ptr("aad"),
			Roles:                to.Ptr("admin,contributor"),
			UserDetails:          to.Ptr("username"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StaticSiteUserInvitationResponseResource = armappservice.StaticSiteUserInvitationResponseResource{
	// 	Properties: &armappservice.StaticSiteUserInvitationResponseResourceProperties{
	// 		ExpiresOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-28T16:53:56.211Z"); return t}()),
	// 		InvitationURL: to.Ptr("https://happy-sea-15afae3e.azurestaticwebsites.net?invite=asdf"),
	// 	},
	// }
}
