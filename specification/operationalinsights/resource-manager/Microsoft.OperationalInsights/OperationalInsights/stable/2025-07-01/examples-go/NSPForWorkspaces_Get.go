package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53d56e4ec74156c450d1e51745a971d3f2031dd7/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/NSPForWorkspaces_Get.json
func ExampleWorkspacesClient_GetNSP() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().GetNSP(ctx, "exampleRG", "someWorkspace", "somePerimeterConfiguration", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkSecurityPerimeterConfiguration = armoperationalinsights.NetworkSecurityPerimeterConfiguration{
	// 	Name: to.Ptr("somePerimeterConfiguration"),
	// 	Type: to.Ptr("Microsoft.OperationalInsights/workspaces/networkSecurityPerimeterConfigurations"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/exampleRG/providers/Microsoft.OperationalInsights/workspaces/someWorkspace/networkSecurityPerimeterConfigurations/somePerimeterConfiguration"),
	// 	Properties: &armoperationalinsights.NetworkSecurityPerimeterConfigurationProperties{
	// 		NetworkSecurityPerimeter: &armoperationalinsights.NetworkSecurityPerimeter{
	// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/networkRG/providers/Microsoft.Network/networkSecurityPerimeters/perimeter1"),
	// 			Location: to.Ptr("japaneast"),
	// 		},
	// 		Profile: &armoperationalinsights.NetworkSecurityProfile{
	// 			Name: to.Ptr("profile1"),
	// 			AccessRules: []*armoperationalinsights.AccessRule{
	// 				{
	// 					Name: to.Ptr("rule1"),
	// 					Properties: &armoperationalinsights.AccessRuleProperties{
	// 						AddressPrefixes: []*string{
	// 							to.Ptr("148.0.0.0/8"),
	// 							to.Ptr("152.4.6.0/24")},
	// 							Direction: to.Ptr(armoperationalinsights.AccessRuleDirectionInbound),
	// 						},
	// 				}},
	// 				AccessRulesVersion: to.Ptr[int32](0),
	// 			},
	// 			ProvisioningState: to.Ptr(armoperationalinsights.NetworkSecurityPerimeterConfigurationProvisioningStateSucceeded),
	// 			ResourceAssociation: &armoperationalinsights.ResourceAssociation{
	// 				Name: to.Ptr("assoc1"),
	// 				AccessMode: to.Ptr(armoperationalinsights.ResourceAssociationAccessModeEnforced),
	// 			},
	// 		},
	// 	}
}
