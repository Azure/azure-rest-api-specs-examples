package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6bbb6daca7175b2cab69b20b2dd01094e3c6a534/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/ListBlocklistItems.json
func ExampleRaiBlocklistItemsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRaiBlocklistItemsClient().NewListPager("resourceGroupName", "accountName", "raiBlocklistName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.RaiBlockListItemsResult = armcognitiveservices.RaiBlockListItemsResult{
		// 	Value: []*armcognitiveservices.RaiBlocklistItem{
		// 		{
		// 			Name: to.Ptr("raiBlocklistItemName"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/raiBlocklists/raiBlocklistName/raiBlocklistItems/raiBlocklistItemName"),
		// 			Properties: &armcognitiveservices.RaiBlocklistItemProperties{
		// 				IsRegex: to.Ptr(false),
		// 				Pattern: to.Ptr("Pattern To Block"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("raiBlocklistItemName2"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/raiBlocklists/raiBlocklistName/raiBlocklistItems/raiBlocklistItemName2"),
		// 			Properties: &armcognitiveservices.RaiBlocklistItemProperties{
		// 				IsRegex: to.Ptr(false),
		// 				Pattern: to.Ptr("Another Pattern To Block"),
		// 			},
		// 	}},
		// }
	}
}
