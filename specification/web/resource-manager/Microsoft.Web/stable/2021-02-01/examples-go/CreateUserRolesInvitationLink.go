package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/CreateUserRolesInvitationLink.json
func ExampleStaticSitesClient_CreateUserRolesInvitationLink() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewStaticSitesClient("<subscription-id>", cred, nil)
	res, err := client.CreateUserRolesInvitationLink(ctx,
		"<resource-group-name>",
		"<name>",
		armappservice.StaticSiteUserInvitationRequestResource{
			Properties: &armappservice.StaticSiteUserInvitationRequestResourceProperties{
				Domain:               to.StringPtr("<domain>"),
				NumHoursToExpiration: to.Int32Ptr(1),
				Provider:             to.StringPtr("<provider>"),
				Roles:                to.StringPtr("<roles>"),
				UserDetails:          to.StringPtr("<user-details>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("StaticSiteUserInvitationResponseResource.ID: %s\n", *res.ID)
}
