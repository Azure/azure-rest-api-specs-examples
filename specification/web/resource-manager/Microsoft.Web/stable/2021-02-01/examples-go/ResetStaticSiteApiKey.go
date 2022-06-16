package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/ResetStaticSiteApiKey.json
func ExampleStaticSitesClient_ResetStaticSiteAPIKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewStaticSitesClient("<subscription-id>", cred, nil)
	_, err = client.ResetStaticSiteAPIKey(ctx,
		"<resource-group-name>",
		"<name>",
		armappservice.StaticSiteResetPropertiesARMResource{
			Properties: &armappservice.StaticSiteResetPropertiesARMResourceProperties{
				RepositoryToken:        to.StringPtr("<repository-token>"),
				ShouldUpdateRepository: to.BoolPtr(true),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
