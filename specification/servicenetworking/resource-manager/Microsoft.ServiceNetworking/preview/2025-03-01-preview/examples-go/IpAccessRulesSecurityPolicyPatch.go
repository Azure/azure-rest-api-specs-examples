package armservicenetworking_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicenetworking/armservicenetworking"
)

// Generated from example definition: 2025-03-01-preview/IpAccessRulesSecurityPolicyPatch.json
func ExampleSecurityPoliciesInterfaceClient_Update_updateIPAccessRulesSecurityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicenetworking.NewClientFactory("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSecurityPoliciesInterfaceClient().Update(ctx, "rg1", "tc1", "sp1", armservicenetworking.SecurityPolicyUpdate{
		Properties: &armservicenetworking.SecurityPolicyUpdateProperties{
			IPAccessRulesPolicy: &armservicenetworking.IPAccessRulesPolicy{
				Rules: []*armservicenetworking.IPAccessRule{},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicenetworking.SecurityPoliciesInterfaceClientUpdateResponse{
	// 	SecurityPolicy: &armservicenetworking.SecurityPolicy{
	// 		Name: to.Ptr("ipAccessRules-0"),
	// 		Type: to.Ptr("Microsoft.ServiceNetworking/trafficControllers/securityPolicies"),
	// 		ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ServiceNetworking/trafficControllers/sample-tc/securityPolicies/ipAccessRules-0"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armservicenetworking.SecurityPolicyProperties{
	// 			PolicyType: to.Ptr(armservicenetworking.PolicyTypeIPAccessRules),
	// 			IPAccessRulesPolicy: &armservicenetworking.IPAccessRulesPolicy{
	// 				Rules: []*armservicenetworking.IPAccessRule{
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armservicenetworking.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
