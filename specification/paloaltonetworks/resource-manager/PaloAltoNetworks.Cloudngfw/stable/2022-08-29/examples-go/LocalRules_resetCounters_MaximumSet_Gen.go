package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/LocalRules_resetCounters_MaximumSet_Gen.json
func ExampleLocalRulesClient_ResetCounters_localRulesResetCountersMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocalRulesClient().ResetCounters(ctx, "firewall-rg", "lrs1", "1", &armpanngfw.LocalRulesClientResetCountersOptions{FirewallName: to.Ptr("firewall1")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RuleCounterReset = armpanngfw.RuleCounterReset{
	// 	FirewallName: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 	Priority: to.Ptr("aaaaaaa"),
	// 	RuleListName: to.Ptr("aaaaa"),
	// 	RuleName: to.Ptr("aaaaa"),
	// 	RuleStackName: to.Ptr("aa"),
	// }
}
