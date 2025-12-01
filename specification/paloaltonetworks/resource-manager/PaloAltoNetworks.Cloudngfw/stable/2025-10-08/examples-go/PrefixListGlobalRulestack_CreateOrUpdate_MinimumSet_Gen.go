package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: 2025-10-08/PrefixListGlobalRulestack_CreateOrUpdate_MinimumSet_Gen.json
func ExamplePrefixListGlobalRulestackClient_BeginCreateOrUpdate_prefixListGlobalRulestackCreateOrUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrefixListGlobalRulestackClient().BeginCreateOrUpdate(ctx, "praval", "armid1", armpanngfw.PrefixListGlobalRulestackResource{
		Properties: &armpanngfw.PrefixObject{
			PrefixList: []*string{
				to.Ptr("1.0.0.0/24"),
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
	// res = armpanngfw.PrefixListGlobalRulestackClientCreateOrUpdateResponse{
	// 	PrefixListGlobalRulestackResource: &armpanngfw.PrefixListGlobalRulestackResource{
	// 		ID: to.Ptr("/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/praval/prefixlists/armid1"),
	// 		Properties: &armpanngfw.PrefixObject{
	// 			PrefixList: []*string{
	// 				to.Ptr("1.0.0.0/24"),
	// 			},
	// 		},
	// 	},
	// }
}
