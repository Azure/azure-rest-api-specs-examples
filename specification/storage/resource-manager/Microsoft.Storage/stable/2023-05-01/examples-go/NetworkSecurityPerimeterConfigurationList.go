package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/220ad9c6554fc7d6d10a89bdb441c1e3b36e3285/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/NetworkSecurityPerimeterConfigurationList.json
func ExampleNetworkSecurityPerimeterConfigurationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkSecurityPerimeterConfigurationsClient().NewListPager("res4410", "sto8607", nil)
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
		// page.NetworkSecurityPerimeterConfigurationList = armstorage.NetworkSecurityPerimeterConfigurationList{
		// 	Value: []*armstorage.NetworkSecurityPerimeterConfiguration{
		// 		{
		// 			Name: to.Ptr("dbedb4e0-40e6-4145-81f3-f1314c150774.resourceAssociation1"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/networkSecurityPerimeterConfigurations"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/res4410/providers/Microsoft.Storage/storageAccounts/sto8607/networkSecurityPerimeterConfigurations/dbedb4e0-40e6-4145-81f3-f1314c150774.resourceAssociation1"),
		// 			Properties: &armstorage.NetworkSecurityPerimeterConfigurationProperties{
		// 				NetworkSecurityPerimeter: &armstorage.NetworkSecurityPerimeter{
		// 					ID: to.Ptr("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/res4794/providers/Microsoft.Network/networkSecurityPerimeters/nsp1"),
		// 					Location: to.Ptr("East US"),
		// 					PerimeterGUID: to.Ptr("ce2d5953-5c15-40ca-9d51-cc3f4a63b0f5"),
		// 				},
		// 				Profile: &armstorage.NetworkSecurityPerimeterConfigurationPropertiesProfile{
		// 					Name: to.Ptr("profile1"),
		// 					AccessRules: []*armstorage.NspAccessRule{
		// 						{
		// 							Name: to.Ptr("inVpnRule"),
		// 							Properties: &armstorage.NspAccessRuleProperties{
		// 								AddressPrefixes: []*string{
		// 									to.Ptr("148.0.0.0/8"),
		// 									to.Ptr("152.4.6.0/24")},
		// 									Direction: to.Ptr(armstorage.NspAccessRuleDirectionInbound),
		// 								},
		// 						}},
		// 						AccessRulesVersion: to.Ptr[float32](10),
		// 						DiagnosticSettingsVersion: to.Ptr[float32](5),
		// 						EnabledLogCategories: []*string{
		// 							to.Ptr("NspPublicInboundPerimeterRulesAllowed"),
		// 							to.Ptr("NspPublicInboundPerimeterRulesDenied")},
		// 						},
		// 						ProvisioningState: to.Ptr(armstorage.NetworkSecurityPerimeterConfigurationProvisioningStateSucceeded),
		// 						ResourceAssociation: &armstorage.NetworkSecurityPerimeterConfigurationPropertiesResourceAssociation{
		// 							Name: to.Ptr("association1"),
		// 							AccessMode: to.Ptr(armstorage.ResourceAssociationAccessModeEnforced),
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}
