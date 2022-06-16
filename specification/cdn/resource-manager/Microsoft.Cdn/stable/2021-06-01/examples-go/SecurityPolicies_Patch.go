package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/SecurityPolicies_Patch.json
func ExampleSecurityPoliciesClient_BeginPatch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcdn.NewSecurityPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginPatch(ctx,
		"RG",
		"profile1",
		"securityPolicy1",
		armcdn.SecurityPolicyUpdateParameters{
			Properties: &armcdn.SecurityPolicyUpdateProperties{
				Parameters: &armcdn.SecurityPolicyWebApplicationFirewallParameters{
					Type: to.Ptr(armcdn.SecurityPolicyTypeWebApplicationFirewall),
					Associations: []*armcdn.SecurityPolicyWebApplicationFirewallAssociation{
						{
							Domains: []*armcdn.ActivatedResourceReference{
								{
									ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/customdomains/testdomain1"),
								},
								{
									ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/customdomains/testdomain2"),
								}},
							PatternsToMatch: []*string{
								to.Ptr("/*")},
						}},
					WafPolicy: &armcdn.ResourceReference{
						ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Network/frontdoorwebapplicationfirewallpolicies/wafTest"),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
