package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v3"
)

// Generated from example definition: 2025-06-01/SecurityPolicies_ListByProfile.json
func ExampleSecurityPoliciesClient_NewListByProfilePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSecurityPoliciesClient().NewListByProfilePager("RG", "profile1", nil)
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
		// page = armcdn.SecurityPoliciesClientListByProfileResponse{
		// 	SecurityPolicyListResult: armcdn.SecurityPolicyListResult{
		// 		Value: []*armcdn.SecurityPolicy{
		// 			{
		// 				Name: to.Ptr("securityPolicy1"),
		// 				Type: to.Ptr("Microsoft.Cdn/profiles/securitypolicies"),
		// 				ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/securitypolicies/securityPolicy1"),
		// 				Properties: &armcdn.SecurityPolicyProperties{
		// 					DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
		// 					Parameters: &armcdn.SecurityPolicyWebApplicationFirewallParameters{
		// 						Type: to.Ptr(armcdn.SecurityPolicyTypeWebApplicationFirewall),
		// 						Associations: []*armcdn.SecurityPolicyWebApplicationFirewallAssociation{
		// 							{
		// 								Domains: []*armcdn.ActivatedResourceReference{
		// 									{
		// 										ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/customdomains/testdomain1"),
		// 									},
		// 									{
		// 										ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/customdomains/testdomain2"),
		// 									},
		// 								},
		// 								PatternsToMatch: []*string{
		// 									to.Ptr("/*"),
		// 								},
		// 							},
		// 						},
		// 						WafPolicy: &armcdn.ResourceReference{
		// 							ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Network/frontdoorwebapplicationfirewallpolicies/wafTest"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
