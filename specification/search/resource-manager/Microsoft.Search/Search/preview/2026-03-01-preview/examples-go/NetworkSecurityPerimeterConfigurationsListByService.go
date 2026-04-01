package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch/v2"
)

// Generated from example definition: 2026-03-01-preview/NetworkSecurityPerimeterConfigurationsListByService.json
func ExampleNetworkSecurityPerimeterConfigurationsClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsearch.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkSecurityPerimeterConfigurationsClient().NewListByServicePager("rg1", "mysearchservice", nil)
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
		// page = armsearch.NetworkSecurityPerimeterConfigurationsClientListByServiceResponse{
		// 	NetworkSecurityPerimeterConfigurationListResult: armsearch.NetworkSecurityPerimeterConfigurationListResult{
		// 		Value: []*armsearch.NetworkSecurityPerimeterConfiguration{
		// 			{
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Search/searchServices/mysearchservice/networkSecurityPerimeterConfigurations/00000001-2222-3333-4444-111144444444.assoc1"),
		// 				Name: to.Ptr("00000001-2222-3333-4444-111144444444.assoc1"),
		// 				Type: to.Ptr("Microsoft.Search/searchServices/networkSecurityPerimeterConfigurations"),
		// 				Properties: &armsearch.NetworkSecurityPerimeterConfigurationProperties{
		// 					ProvisioningState: to.Ptr(armsearch.NetworkSecurityPerimeterConfigurationProvisioningStateAccepted),
		// 					NetworkSecurityPerimeter: &armsearch.NetworkSecurityPerimeter{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/networkRG/providers/Microsoft.Network/networkSecurityPerimeters/perimeter1"),
		// 						Location: to.Ptr("westus"),
		// 					},
		// 					ResourceAssociation: &armsearch.ResourceAssociation{
		// 						Name: to.Ptr("assoc1"),
		// 						AccessMode: to.Ptr(armsearch.ResourceAssociationAccessModeEnforced),
		// 					},
		// 					Profile: &armsearch.NetworkSecurityProfile{
		// 						Name: to.Ptr("profile1"),
		// 						AccessRulesVersion: to.Ptr[int32](0),
		// 						AccessRules: []*armsearch.AccessRule{
		// 							{
		// 								Name: to.Ptr("rule1"),
		// 								Properties: &armsearch.AccessRuleProperties{
		// 									Direction: to.Ptr(armsearch.AccessRuleDirectionInbound),
		// 									AddressPrefixes: []*string{
		// 										to.Ptr("148.0.0.0/8"),
		// 										to.Ptr("152.4.6.0/24"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
