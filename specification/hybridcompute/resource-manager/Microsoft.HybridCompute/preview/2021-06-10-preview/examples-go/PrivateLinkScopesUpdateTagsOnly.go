package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
)

// x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-06-10-preview/examples/PrivateLinkScopesUpdateTagsOnly.json
func ExamplePrivateLinkScopesClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridcompute.NewPrivateLinkScopesClient("<subscription-id>", cred, nil)
	res, err := client.UpdateTags(ctx,
		"<resource-group-name>",
		"<scope-name>",
		armhybridcompute.TagsResource{
			Tags: map[string]*string{
				"Tag1": to.StringPtr("Value1"),
				"Tag2": to.StringPtr("Value2"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateLinkScopesClientUpdateTagsResult)
}
