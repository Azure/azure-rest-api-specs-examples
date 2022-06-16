package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2021-02-01/examples/GetTopLevelDomain.json
func ExampleTopLevelDomainsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewTopLevelDomainsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("TopLevelDomain.ID: %s\n", *res.ID)
}
