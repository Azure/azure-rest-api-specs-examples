package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2021-09-01/examples/RegistryGetPrivateLinkResource.json
func ExampleRegistriesClient_GetPrivateLinkResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerregistry.NewRegistriesClient("<subscription-id>", cred, nil)
	res, err := client.GetPrivateLinkResource(ctx,
		"<resource-group-name>",
		"<registry-name>",
		"<group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RegistriesClientGetPrivateLinkResourceResult)
}
