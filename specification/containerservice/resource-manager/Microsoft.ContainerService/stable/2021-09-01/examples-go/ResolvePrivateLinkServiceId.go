package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
)

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/ResolvePrivateLinkServiceId.json
func ExampleResolvePrivateLinkServiceIDClient_POST() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewResolvePrivateLinkServiceIDClient("<subscription-id>", cred, nil)
	res, err := client.POST(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armcontainerservice.PrivateLinkResource{
			Name: to.StringPtr("<name>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PrivateLinkResource.ID: %s\n", *res.ID)
}
