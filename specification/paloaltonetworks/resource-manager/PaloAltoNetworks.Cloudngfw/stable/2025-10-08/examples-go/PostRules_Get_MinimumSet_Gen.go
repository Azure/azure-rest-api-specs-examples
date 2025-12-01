package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: 2025-10-08/PostRules_Get_MinimumSet_Gen.json
func ExamplePostRulesClient_Get_postRulesGetMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPostRulesClient().Get(ctx, "lrs1", "1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpanngfw.PostRulesClientGetResponse{
	// 	PostRulesResource: &armpanngfw.PostRulesResource{
	// 		ID: to.Ptr("/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/lrs1/postrules/1"),
	// 		Properties: &armpanngfw.RuleEntry{
	// 			RuleName: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 		},
	// 	},
	// }
}
