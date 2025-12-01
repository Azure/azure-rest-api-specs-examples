package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: 2025-10-08/GlobalRulestack_List_MinimumSet_Gen.json
func ExampleGlobalRulestackClient_NewListPager_globalRulestackListMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGlobalRulestackClient().NewListPager(nil)
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
		// page = armpanngfw.GlobalRulestackClientListResponse{
		// 	GlobalRulestackResourceListResult: armpanngfw.GlobalRulestackResourceListResult{
		// 		Value: []*armpanngfw.GlobalRulestackResource{
		// 			{
		// 				ID: to.Ptr("/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/grs1"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armpanngfw.RulestackProperties{
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
