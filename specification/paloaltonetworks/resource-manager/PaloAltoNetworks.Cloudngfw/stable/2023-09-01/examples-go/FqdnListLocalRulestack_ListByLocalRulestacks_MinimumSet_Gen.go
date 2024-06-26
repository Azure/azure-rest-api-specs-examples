package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/FqdnListLocalRulestack_ListByLocalRulestacks_MinimumSet_Gen.json
func ExampleFqdnListLocalRulestackClient_NewListByLocalRulestacksPager_fqdnListLocalRulestackListByLocalRulestacksMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFqdnListLocalRulestackClient().NewListByLocalRulestacksPager("rgopenapi", "lrs1", nil)
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
		// page.FqdnListLocalRulestackResourceListResult = armpanngfw.FqdnListLocalRulestackResourceListResult{
		// 	Value: []*armpanngfw.FqdnListLocalRulestackResource{
		// 		{
		// 			ID: to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourcegroups/rgopenapi/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/praval/fqdnlists/fqdnlists1"),
		// 			Properties: &armpanngfw.FqdnObject{
		// 				FqdnList: []*string{
		// 					to.Ptr("string1"),
		// 					to.Ptr("string2")},
		// 				},
		// 		}},
		// 	}
	}
}
