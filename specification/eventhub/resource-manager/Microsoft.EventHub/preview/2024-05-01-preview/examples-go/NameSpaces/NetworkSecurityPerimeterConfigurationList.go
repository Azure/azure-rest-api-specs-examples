package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5759c77eee2d57bdb9e47aa1805d0ffb61704f2d/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2024-05-01-preview/examples/NameSpaces/NetworkSecurityPerimeterConfigurationList.json
func ExampleNetworkSecurityPerimeterConfigurationClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkSecurityPerimeterConfigurationClient().List(ctx, "SDK-EventHub-4794", "sdk-Namespace-5828", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkSecurityPerimeterConfigurationList = armeventhub.NetworkSecurityPerimeterConfigurationList{
	// 	Value: []*armeventhub.NetworkSecurityPerimeterConfiguration{
	// 		{
	// 			Name: to.Ptr("resourceAssociation1"),
	// 			Type: to.Ptr("Microsoft.EventHub/Namespaces/networkSecurityPerimeterConfigurations"),
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SDK-EventHub-4794/providers/Microsoft.EventHub/namespaces/sdk-Namespace-5705-new/networkSecurityPerimeterConfigurations/resourceAssociation1"),
	// 			Properties: &armeventhub.NetworkSecurityPerimeterConfigurationProperties{
	// 				NetworkSecurityPerimeter: &armeventhub.NetworkSecurityPerimeter{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SDK-EventHub-4794/providers/Microsoft.Network/networkSecurityPerimeters/nsp1"),
	// 					Location: to.Ptr("East US"),
	// 					PerimeterGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				},
	// 				Profile: &armeventhub.NetworkSecurityPerimeterConfigurationPropertiesProfile{
	// 					Name: to.Ptr("devProfile"),
	// 					AccessRules: []*armeventhub.NspAccessRule{
	// 						{
	// 							Name: to.Ptr("inVpnRule"),
	// 							Properties: &armeventhub.NspAccessRuleProperties{
	// 								AddressPrefixes: []*string{
	// 									to.Ptr("148.0.0.0/8"),
	// 									to.Ptr("152.4.6.0/24")},
	// 									Direction: to.Ptr(armeventhub.NspAccessRuleDirectionInbound),
	// 								},
	// 						}},
	// 						AccessRulesVersion: to.Ptr("10"),
	// 					},
	// 					ProvisioningState: to.Ptr(armeventhub.NetworkSecurityPerimeterConfigurationProvisioningStateSucceeded),
	// 					ResourceAssociation: &armeventhub.NetworkSecurityPerimeterConfigurationPropertiesResourceAssociation{
	// 						Name: to.Ptr("association1"),
	// 					},
	// 				},
	// 		}},
	// 	}
}
