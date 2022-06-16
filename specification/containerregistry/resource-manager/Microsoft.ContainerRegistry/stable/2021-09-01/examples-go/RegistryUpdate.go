package armcontainerregistry_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2021-09-01/examples/RegistryUpdate.json
func ExampleRegistriesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerregistry.NewRegistriesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<registry-name>",
		armcontainerregistry.RegistryUpdateParameters{
			Properties: &armcontainerregistry.RegistryPropertiesUpdateParameters{
				AdminUserEnabled: to.BoolPtr(true),
			},
			SKU: &armcontainerregistry.SKU{
				Name: armcontainerregistry.SKUName("Standard").ToPtr(),
			},
			Tags: map[string]*string{
				"key": to.StringPtr("value"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RegistriesClientUpdateResult)
}
