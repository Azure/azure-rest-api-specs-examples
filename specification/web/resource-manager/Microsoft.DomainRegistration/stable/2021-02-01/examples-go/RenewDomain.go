package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2021-02-01/examples/RenewDomain.json
func ExampleDomainsClient_Renew() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewDomainsClient("<subscription-id>", cred, nil)
	_, err = client.Renew(ctx,
		"<resource-group-name>",
		"<domain-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
