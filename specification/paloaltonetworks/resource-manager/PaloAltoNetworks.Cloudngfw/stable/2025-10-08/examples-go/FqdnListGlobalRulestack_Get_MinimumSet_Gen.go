package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: 2025-10-08/FqdnListGlobalRulestack_Get_MinimumSet_Gen.json
func ExampleFqdnListGlobalRulestackClient_Get_fqdnListGlobalRulestackGetMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFqdnListGlobalRulestackClient().Get(ctx, "praval", "armid1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpanngfw.FqdnListGlobalRulestackClientGetResponse{
	// 	FqdnListGlobalRulestackResource: &armpanngfw.FqdnListGlobalRulestackResource{
	// 		ID: to.Ptr("/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/praval/fqdnlists/armid1"),
	// 		Properties: &armpanngfw.FqdnObject{
	// 			FqdnList: []*string{
	// 				to.Ptr("string1"),
	// 				to.Ptr("string2"),
	// 			},
	// 		},
	// 	},
	// }
}
