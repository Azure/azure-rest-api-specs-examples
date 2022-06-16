package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance"
)

// x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-09-01/examples/ContainerGroupsGet_Failed.json
func ExampleContainerGroupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerinstance.NewContainerGroupsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<container-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ContainerGroup.ID: %s\n", *res.ID)
}
