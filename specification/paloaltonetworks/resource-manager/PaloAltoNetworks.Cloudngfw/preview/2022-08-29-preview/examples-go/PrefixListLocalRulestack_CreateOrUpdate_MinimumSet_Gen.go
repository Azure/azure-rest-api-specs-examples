package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fdf43f2fdacf17bd78c0621df44a5c024b61db82/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/PrefixListLocalRulestack_CreateOrUpdate_MinimumSet_Gen.json
func ExamplePrefixListLocalRulestackClient_BeginCreateOrUpdate_prefixListLocalRulestackCreateOrUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrefixListLocalRulestackClient().BeginCreateOrUpdate(ctx, "rgopenapi", "lrs1", "armid1", armpanngfw.PrefixListResource{
		Properties: &armpanngfw.PrefixObject{
			PrefixList: []*string{
				to.Ptr("1.0.0.0/24")},
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
	// res.PrefixListResource = armpanngfw.PrefixListResource{
	// 	ID: to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourcegroups/rgopenapi/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/praval/prefixlists/armid1"),
	// 	Properties: &armpanngfw.PrefixObject{
	// 		PrefixList: []*string{
	// 			to.Ptr("1.0.0.0/24")},
	// 		},
	// 	}
}
