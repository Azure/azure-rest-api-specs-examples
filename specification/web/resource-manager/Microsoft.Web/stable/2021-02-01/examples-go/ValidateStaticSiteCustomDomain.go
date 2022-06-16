package armappservice_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/ValidateStaticSiteCustomDomain.json
func ExampleStaticSitesClient_BeginValidateCustomDomainCanBeAddedToStaticSite() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewStaticSitesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginValidateCustomDomainCanBeAddedToStaticSite(ctx,
		"<resource-group-name>",
		"<name>",
		"<domain-name>",
		armappservice.StaticSiteCustomDomainRequestPropertiesARMResource{
			Properties: &armappservice.StaticSiteCustomDomainRequestPropertiesARMResourceProperties{},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
