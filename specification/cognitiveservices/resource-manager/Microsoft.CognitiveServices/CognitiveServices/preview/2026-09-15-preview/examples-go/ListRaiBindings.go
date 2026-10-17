package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-09-15-preview/ListRaiBindings.json
func ExampleRaiBindingsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRaiBindingsClient().NewListPager("resource-group", "safety-account", &armcognitiveservices.RaiBindingsClientListOptions{
		Top: to.Ptr[int32](50)})
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
		// page = armcognitiveservices.RaiBindingsClientListResponse{
		// 	RaiBindingListResult: armcognitiveservices.RaiBindingListResult{
		// 		Value: []*armcognitiveservices.RaiBinding{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group/providers/Microsoft.CognitiveServices/accounts/safety-account/raiBindings/chat-guard"),
		// 				Name: to.Ptr("chat-guard"),
		// 				Type: to.Ptr("Microsoft.CognitiveServices/accounts/raiBindings"),
		// 				Etag: to.Ptr("\"00000000-0000-0000-0000-000000000002\""),
		// 				Properties: &armcognitiveservices.RaiBindingProperties{
		// 					BoundResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group/providers/Microsoft.CognitiveServices/accounts/aoai-account/deployments/chat-prod"),
		// 					TargetPolicyName: to.Ptr("agent-guard"),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group/providers/Microsoft.CognitiveServices/accounts/safety-account/raiBindings?api-version=2026-09-15-preview&%24top=50&%24skiptoken=Y2hhdC1ndWFyZA=="),
		// 	},
		// }
	}
}
