package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fdf43f2fdacf17bd78c0621df44a5c024b61db82/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/FqdnListGlobalRulestack_List_MinimumSet_Gen.json
func ExampleFqdnListGlobalRulestackClient_NewListPager_fqdnListGlobalRulestackListMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFqdnListGlobalRulestackClient().NewListPager("praval", nil)
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
		// page.FqdnListGlobalRulestackResourceListResult = armpanngfw.FqdnListGlobalRulestackResourceListResult{
		// 	Value: []*armpanngfw.FqdnListGlobalRulestackResource{
		// 		{
		// 			ID: to.Ptr("/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/praval/fqdnlists/fqdnlists1"),
		// 			Properties: &armpanngfw.FqdnObject{
		// 				FqdnList: []*string{
		// 					to.Ptr("string1"),
		// 					to.Ptr("string2")},
		// 				},
		// 		}},
		// 	}
	}
}
