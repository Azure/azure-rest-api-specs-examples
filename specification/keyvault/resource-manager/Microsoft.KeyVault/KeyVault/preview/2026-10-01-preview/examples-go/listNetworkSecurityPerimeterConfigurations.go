package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault/v2"
)

// Generated from example definition: 2026-10-01-preview/listNetworkSecurityPerimeterConfigurations.json
func ExampleNetworkSecurityPerimeterConfigurationsClient_NewListConfigurationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkSecurityPerimeterConfigurationsClient().NewListConfigurationsPager("sample-resource-group", "sample-vault", nil)
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
		// page = armkeyvault.NetworkSecurityPerimeterConfigurationsClientListConfigurationsResponse{
		// 	NetworkSecurityPerimeterConfigurationListResult: armkeyvault.NetworkSecurityPerimeterConfigurationListResult{
		// 		Value: []*armkeyvault.NetworkSecurityPerimeterConfiguration{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-resource-group/providers/Microsoft.KeyVault/vaults/sample-vault/networkSecurityPerimeterConfigurations/00000000-0000-0000-0000-000000000000.sampleAssociation"),
		// 				Name: to.Ptr("00000000-0000-0000-0000-000000000000.sampleAssociation"),
		// 				Type: to.Ptr("Microsoft.KeyVault/vaults/networkSecurityPerimeterConfigurations"),
		// 				Properties: &armkeyvault.NetworkSecurityPerimeterConfigurationProperties{
		// 					ProvisioningState: to.Ptr(armkeyvault.NetworkSecurityPerimeterConfigurationProvisioningStateSucceeded),
		// 					NetworkSecurityPerimeter: &armkeyvault.NetworkSecurityPerimeter{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-resource-group/providers/Microsoft.Network/networkSecurityPerimeters/sampleNsp"),
		// 						PerimeterGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						Location: to.Ptr("centraluseuap"),
		// 					},
		// 					ResourceAssociation: &armkeyvault.ResourceAssociation{
		// 						Name: to.Ptr("sampleAssociation"),
		// 						AccessMode: to.Ptr(armkeyvault.ResourceAssociationAccessModeLearning),
		// 					},
		// 					Profile: &armkeyvault.NetworkSecurityProfile{
		// 						Name: to.Ptr("defaultprofile"),
		// 						AccessRulesVersion: to.Ptr[int32](1),
		// 						AccessRules: []*armkeyvault.AccessRule{
		// 							{
		// 								Name: to.Ptr("sampleRule"),
		// 								Properties: &armkeyvault.AccessRuleProperties{
		// 									Direction: to.Ptr(armkeyvault.AccessRuleDirectionInbound),
		// 									AddressPrefixes: []*string{
		// 										to.Ptr("192.0.2.0/24"),
		// 									},
		// 									Subscriptions: []*armkeyvault.AccessRulePropertiesSubscription{
		// 									},
		// 									FullyQualifiedDomainNames: []*string{
		// 									},
		// 								},
		// 							},
		// 						},
		// 						DiagnosticSettingsVersion: to.Ptr[int32](0),
		// 						EnabledLogCategories: []*string{
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
