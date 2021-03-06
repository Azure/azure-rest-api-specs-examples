package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/CreateUserRolesInvitationLink.json
func ExampleStaticSitesClient_CreateUserRolesInvitationLink() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappservice.NewStaticSitesClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateUserRolesInvitationLink(ctx,
		"rg",
		"testStaticSite0",
		armappservice.StaticSiteUserInvitationRequestResource{
			Properties: &armappservice.StaticSiteUserInvitationRequestResourceProperties{
				Domain:               to.Ptr("happy-sea-15afae3e.azurestaticwebsites.net"),
				NumHoursToExpiration: to.Ptr[int32](1),
				Provider:             to.Ptr("aad"),
				Roles:                to.Ptr("admin,contributor"),
				UserDetails:          to.Ptr("username"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
