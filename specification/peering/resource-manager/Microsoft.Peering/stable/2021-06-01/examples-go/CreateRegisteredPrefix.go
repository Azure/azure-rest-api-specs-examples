package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/CreateRegisteredPrefix.json
func ExampleRegisteredPrefixesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpeering.NewRegisteredPrefixesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<peering-name>",
		"<registered-prefix-name>",
		armpeering.RegisteredPrefix{
			Properties: &armpeering.RegisteredPrefixProperties{
				Prefix: to.StringPtr("<prefix>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RegisteredPrefixesClientCreateOrUpdateResult)
}
