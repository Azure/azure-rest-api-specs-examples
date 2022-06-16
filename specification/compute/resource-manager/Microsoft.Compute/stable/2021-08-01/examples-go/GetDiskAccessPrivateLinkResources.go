package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-08-01/examples/GetDiskAccessPrivateLinkResources.json
func ExampleDiskAccessesClient_GetPrivateLinkResources() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewDiskAccessesClient("<subscription-id>", cred, nil)
	res, err := client.GetPrivateLinkResources(ctx,
		"<resource-group-name>",
		"<disk-access-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DiskAccessesClientGetPrivateLinkResourcesResult)
}
