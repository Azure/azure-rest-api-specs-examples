package armpurview_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview"
)

// x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/PrivateLinkResources_GetByGroupId.json
func ExamplePrivateLinkResourcesClient_GetByGroupID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpurview.NewPrivateLinkResourcesClient("<subscription-id>", cred, nil)
	res, err := client.GetByGroupID(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<group-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PrivateLinkResource.ID: %s\n", *res.ID)
}
