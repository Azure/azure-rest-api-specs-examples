package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fb3217991ff57b5760525aeba1a0670bfe0880fa/specification/search/resource-manager/Microsoft.Search/stable/2025-05-01/examples/NetworkSecurityPerimeterConfigurationsListByService.json
func ExampleNetworkSecurityPerimeterConfigurationsClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsearch.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.NetworkSecurityPerimeterConfigurationListResult = armsearch.NetworkSecurityPerimeterConfigurationListResult{
		// 	Value: []*armsearch.NetworkSecurityPerimeterConfiguration{
		// 		{
		// 			Name: to.Ptr("00000001-2222-3333-4444-111144444444.assoc1"),
		// 			Type: to.Ptr("Microsoft.Search/searchServices/networkSecurityPerimeterConfigurations"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Search/searchServices/mysearchservice/networkSecurityPerimeterConfigurations/00000001-2222-3333-4444-111144444444.assoc1"),
		// 			Properties: &armsearch.NetworkSecurityPerimeterConfigurationProperties{
		// 				NetworkSecurityPerimeter: &armsearch.NetworkSecurityPerimeter{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/networkRG/providers/Microsoft.Network/networkSecurityPerimeters/perimeter1"),
		// 					Location: to.Ptr("westus"),
		// 				},
		// 				Profile: &armsearch.NetworkSecurityProfile{
		// 					Name: to.Ptr("profile1"),
		// 					AccessRules: []*armsearch.AccessRule{
		// 						{
		// 							Name: to.Ptr("rule1"),
		// 							Properties: &armsearch.AccessRuleProperties{
		// 								AddressPrefixes: []*string{
		// 									to.Ptr("148.0.0.0/8"),
		// 									to.Ptr("152.4.6.0/24")},
		// 									Direction: to.Ptr(armsearch.AccessRuleDirectionInbound),
		// 								},
		// 						}},
		// 						AccessRulesVersion: to.Ptr[int32](0),
		// 					},
		// 					ProvisioningState: to.Ptr(armsearch.NetworkSecurityPerimeterConfigurationProvisioningStateSucceeded),
		// 					ResourceAssociation: &armsearch.ResourceAssociation{
		// 						Name: to.Ptr("assoc1"),
		// 						AccessMode: to.Ptr(armsearch.ResourceAssociationAccessModeEnforced),
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}
