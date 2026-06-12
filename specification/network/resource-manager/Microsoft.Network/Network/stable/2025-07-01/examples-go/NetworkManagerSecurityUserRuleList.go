package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NetworkManagerSecurityUserRuleList.json
func ExampleSecurityUserRulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSecurityUserRulesClient().NewListPager("rg1", "testNetworkManager", "myTestSecurityConfig", "testRuleCollection", nil)
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
		// page = armnetwork.SecurityUserRulesClientListResponse{
		// 	SecurityUserRuleListResult: armnetwork.SecurityUserRuleListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/22222222-2222-2222-2222-222222222222/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/securityUserConfigurations/myTestSecurityConfig/ruleCollections/testRuleCollection/rules?api-version=2025-07-01&$skipToken=10"),
		// 		Value: []*armnetwork.SecurityUserRule{
		// 			{
		// 				Name: to.Ptr("SampleUserRule"),
		// 				Type: to.Ptr("Microsoft.Network/networkManagers/securityConfigurations/ruleCollections/rules"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkmanagers/testNetworkManager/securityUserConfigurations/myTestSecurityConfig/ruleCollections/testRuleCollection/rules/SampleUserRule"),
		// 				Properties: &armnetwork.SecurityUserRulePropertiesFormat{
		// 					Description: to.Ptr("Sample User Rule"),
		// 					DestinationPortRanges: []*string{
		// 						to.Ptr("22"),
		// 					},
		// 					Destinations: []*armnetwork.AddressPrefixItem{
		// 						{
		// 							AddressPrefix: to.Ptr("*"),
		// 							AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
		// 						},
		// 					},
		// 					Direction: to.Ptr(armnetwork.SecurityConfigurationRuleDirectionInbound),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("f930553b-f78d-48c5-9445-6cf86b85e615"),
		// 					SourcePortRanges: []*string{
		// 						to.Ptr("0-65535"),
		// 					},
		// 					Sources: []*armnetwork.AddressPrefixItem{
		// 						{
		// 							AddressPrefix: to.Ptr("*"),
		// 							AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
		// 						},
		// 					},
		// 					Protocol: to.Ptr(armnetwork.SecurityConfigurationRuleProtocolTCP),
		// 				},
		// 				SystemData: &armnetwork.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
		// 					CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
		// 					CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
		// 					LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
