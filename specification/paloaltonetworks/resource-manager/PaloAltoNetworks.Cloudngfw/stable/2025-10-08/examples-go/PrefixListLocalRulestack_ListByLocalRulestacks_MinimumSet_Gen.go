package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: 2025-10-08/PrefixListLocalRulestack_ListByLocalRulestacks_MinimumSet_Gen.json
func ExamplePrefixListLocalRulestackClient_NewListByLocalRulestacksPager_prefixListLocalRulestackListByLocalRulestacksMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("2bf4a339-294d-4c25-b0b2-ef649e9f5c27", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrefixListLocalRulestackClient().NewListByLocalRulestacksPager("rgopenapi", "lrs1", nil)
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
		// page = armpanngfw.PrefixListLocalRulestackClientListByLocalRulestacksResponse{
		// 	PrefixListResourceListResult: armpanngfw.PrefixListResourceListResult{
		// 		Value: []*armpanngfw.PrefixListResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourcegroups/rgopenapi/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/praval/prefixlists/prefixlists1"),
		// 				Properties: &armpanngfw.PrefixObject{
		// 					PrefixList: []*string{
		// 						to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
