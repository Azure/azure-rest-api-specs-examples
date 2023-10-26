package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7b551033155a63739b6d28f79b9c07569f6179b8/specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/SecurityPolicies_ListByProfile.json
func ExampleSecurityPoliciesClient_NewListByProfilePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.SecurityPolicyListResult = armcdn.SecurityPolicyListResult{
		// 	Value: []*armcdn.SecurityPolicy{
		// 		{
		// 			Name: to.Ptr("securityPolicy1"),
		// 			Type: to.Ptr("Microsoft.Cdn/profiles/securitypolicies"),
		// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/securitypolicies/securityPolicy1"),
		// 			Properties: &armcdn.SecurityPolicyProperties{
		// 				DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
		// 				ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
		// 				Parameters: &armcdn.SecurityPolicyWebApplicationFirewallParameters{
		// 					Type: to.Ptr(armcdn.SecurityPolicyTypeWebApplicationFirewall),
		// 					Associations: []*armcdn.SecurityPolicyWebApplicationFirewallAssociation{
		// 						{
		// 							Domains: []*armcdn.ActivatedResourceReference{
		// 								{
		// 									ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/customdomains/testdomain1"),
		// 								},
		// 								{
		// 									ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/customdomains/testdomain2"),
		// 							}},
		// 							PatternsToMatch: []*string{
		// 								to.Ptr("/*")},
		// 						}},
		// 						WafPolicy: &armcdn.ResourceReference{
		// 							ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Network/frontdoorwebapplicationfirewallpolicies/wafTest"),
		// 						},
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}
