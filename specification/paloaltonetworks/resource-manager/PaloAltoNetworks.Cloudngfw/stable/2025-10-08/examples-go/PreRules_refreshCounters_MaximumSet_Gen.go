package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: 2025-10-08/PreRules_refreshCounters_MaximumSet_Gen.json
func ExamplePreRulesClient_RefreshCounters_preRulesRefreshCountersMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewPreRulesClient().RefreshCounters(ctx, "lrs1", "1", &armpanngfw.PreRulesClientRefreshCountersOptions{
		FirewallName: to.Ptr("firewall1")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
