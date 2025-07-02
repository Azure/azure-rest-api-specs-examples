package armresources_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4a2bb0762eaad11e725516708483598e0c12cabb/specification/resources/resource-manager/Microsoft.Resources/stable/2025-04-01/examples/PutTagsSubscription.json
func ExampleTagsClient_BeginCreateOrUpdateAtScope_updateTagsOnASubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresources.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTagsClient().BeginCreateOrUpdateAtScope(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", armresources.TagsResource{
		Properties: &armresources.Tags{
			Tags: map[string]*string{
				"tagKey1": to.Ptr("tag-value-1"),
				"tagKey2": to.Ptr("tag-value-2"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TagsResource = armresources.TagsResource{
	// 	ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000"),
	// 	Properties: &armresources.Tags{
	// 		Tags: map[string]*string{
	// 			"tagKey1": to.Ptr("tag-value-1"),
	// 			"tagKey2": to.Ptr("tag-value-2"),
	// 		},
	// 	},
	// }
}
