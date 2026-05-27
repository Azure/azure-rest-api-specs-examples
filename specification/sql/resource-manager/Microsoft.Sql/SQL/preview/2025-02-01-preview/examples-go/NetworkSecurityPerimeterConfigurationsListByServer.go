package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/NetworkSecurityPerimeterConfigurationsListByServer.json
func ExampleNetworkSecurityPerimeterConfigurationsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkSecurityPerimeterConfigurationsClient().NewListByServerPager("sqlcrudtest-7398", "sqlcrudtest-7398", nil)
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
		// page = armsql.NetworkSecurityPerimeterConfigurationsClientListByServerResponse{
		// 	NetworkSecurityPerimeterConfigurationListResult: armsql.NetworkSecurityPerimeterConfigurationListResult{
		// 		Value: []*armsql.NetworkSecurityPerimeterConfiguration{
		// 			{
		// 				Name: to.Ptr("00000001-2222-3333-4444-111144444444.assoc1"),
		// 				Type: to.Ptr("Microsoft.Sql/servers/networkSecurityPerimeterConfigurations"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/servers/sqlcrudtest-4645/networkSecurityPerimeterConfigurations/00000001-2222-3333-4444-111144444444.assoc1"),
		// 				Properties: &armsql.NetworkSecurityPerimeterConfigurationProperties{
		// 					NetworkSecurityPerimeter: &armsql.NSPConfigPerimeter{
		// 						ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/networkRG/providers/Microsoft.Network/networkSecurityPerimeters/perimeter1"),
		// 						Location: to.Ptr("japaneast"),
		// 					},
		// 					Profile: &armsql.NSPConfigProfile{
		// 						Name: to.Ptr("profile1"),
		// 						AccessRules: []*armsql.NSPConfigAccessRule{
		// 							{
		// 								Name: to.Ptr("rule1"),
		// 								Properties: &armsql.NSPConfigAccessRuleProperties{
		// 									AddressPrefixes: []*string{
		// 										to.Ptr("148.0.0.0/8"),
		// 										to.Ptr("152.4.6.0/24"),
		// 									},
		// 									Direction: to.Ptr("Inbound"),
		// 								},
		// 							},
		// 						},
		// 						AccessRulesVersion: to.Ptr("0"),
		// 					},
		// 					ProvisioningState: to.Ptr("Accepted"),
		// 					ResourceAssociation: &armsql.NSPConfigAssociation{
		// 						Name: to.Ptr("assoc1"),
		// 						AccessMode: to.Ptr("Enforced"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
