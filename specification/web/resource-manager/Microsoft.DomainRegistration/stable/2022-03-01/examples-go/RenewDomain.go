package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/web/resource-manager/Microsoft.DomainRegistration/stable/2022-03-01/examples/RenewDomain.json
func ExampleDomainsClient_Renew() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappservice.NewDomainsClient("3dddfa4f-cedf-4dc0-ba29-b6d1a69ab545", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Renew(ctx,
		"RG",
		"example.com",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
