package armservicenetworking_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicenetworking/armservicenetworking"
)

// Generated from example definition: 2025-03-01-preview/FrontendsGet.json
func ExampleFrontendsInterfaceClient_NewListByTrafficControllerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicenetworking.NewClientFactory("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFrontendsInterfaceClient().NewListByTrafficControllerPager("rg1", "tc1", nil)
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
		// page = armservicenetworking.FrontendsInterfaceClientListByTrafficControllerResponse{
		// 	FrontendListResult: armservicenetworking.FrontendListResult{
		// 		Value: []*armservicenetworking.Frontend{
		// 			{
		// 				Name: to.Ptr("fe1"),
		// 				Location: to.Ptr("NorthCentralUS"),
		// 				Type: to.Ptr("Microsoft.ServiceNetworking/trafficControllers/frontends"),
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ServiceNetworking/trafficControllers/tc1/frontends/fe1"),
		// 				Properties: &armservicenetworking.FrontendProperties{
		// 					Fqdn: to.Ptr("test.net"),
		// 					SecurityPolicyConfigurations: &armservicenetworking.SecurityPolicyConfigurations{
		// 						IPAccessRulesSecurityPolicy: &armservicenetworking.IPAccessRulesSecurityPolicy{
		// 							ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ServiceNetworking/trafficControllers/sample-tc/securityPolicies/ipAccessRules-0"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armservicenetworking.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
