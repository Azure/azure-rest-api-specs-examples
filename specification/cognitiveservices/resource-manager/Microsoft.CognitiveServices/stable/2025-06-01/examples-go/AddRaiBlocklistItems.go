package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/44319b51c6f952fdc9543d3dc4fdd9959350d102/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/AddRaiBlocklistItems.json
func ExampleRaiBlocklistItemsClient_BatchAdd() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRaiBlocklistItemsClient().BatchAdd(ctx, "resourceGroupName", "accountName", "myblocklist", []*armcognitiveservices.RaiBlocklistItemBulkRequest{
		{
			Name: to.Ptr("myblocklistitem1"),
			Properties: &armcognitiveservices.RaiBlocklistItemProperties{
				IsRegex: to.Ptr(true),
				Pattern: to.Ptr("^[a-z0-9_-]{2,16}$"),
			},
		},
		{
			Name: to.Ptr("myblocklistitem2"),
			Properties: &armcognitiveservices.RaiBlocklistItemProperties{
				IsRegex: to.Ptr(false),
				Pattern: to.Ptr("blockwords"),
			},
		}}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RaiBlocklist = armcognitiveservices.RaiBlocklist{
	// 	Name: to.Ptr("myblocklist"),
	// 	Etag: to.Ptr("\"00000000-0000-0000-0000-000000000000\""),
	// 	Properties: &armcognitiveservices.RaiBlocklistProperties{
	// 		Description: to.Ptr("Brief description of the blocklist"),
	// 	},
	// }
}
