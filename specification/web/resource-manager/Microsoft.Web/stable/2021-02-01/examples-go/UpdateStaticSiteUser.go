package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/UpdateStaticSiteUser.json
func ExampleStaticSitesClient_UpdateStaticSiteUser() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewStaticSitesClient("<subscription-id>", cred, nil)
	res, err := client.UpdateStaticSiteUser(ctx,
		"<resource-group-name>",
		"<name>",
		"<authprovider>",
		"<userid>",
		armappservice.StaticSiteUserARMResource{
			Properties: &armappservice.StaticSiteUserARMResourceProperties{
				Roles: to.StringPtr("<roles>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("StaticSiteUserARMResource.ID: %s\n", *res.ID)
}
