package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fdf43f2fdacf17bd78c0621df44a5c024b61db82/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/PostRules_refreshCounters_MaximumSet_Gen.json
func ExamplePostRulesClient_RefreshCounters_postRulesRefreshCountersMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewPostRulesClient().RefreshCounters(ctx, "lrs1", "1", &armpanngfw.PostRulesClientRefreshCountersOptions{FirewallName: to.Ptr("firewall1")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
