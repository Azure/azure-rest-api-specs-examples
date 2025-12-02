package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: 2025-10-08/PrefixListLocalRulestack_Get_MinimumSet_Gen.json
func ExamplePrefixListLocalRulestackClient_Get_prefixListLocalRulestackGetMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("2bf4a339-294d-4c25-b0b2-ef649e9f5c27", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrefixListLocalRulestackClient().Get(ctx, "rgopenapi", "lrs1", "armid1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpanngfw.PrefixListLocalRulestackClientGetResponse{
	// 	PrefixListResource: &armpanngfw.PrefixListResource{
	// 		ID: to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourcegroups/rgopenapi/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/praval/prefixlists/armid1"),
	// 		Properties: &armpanngfw.PrefixObject{
	// 			PrefixList: []*string{
	// 				to.Ptr("1.0.0.0/24"),
	// 			},
	// 		},
	// 	},
	// }
}
