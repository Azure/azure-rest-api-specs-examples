package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: 2021-10-01/NSPForDataCollectionEndpoints_Get.json
func ExampleDataCollectionEndpointsClient_GetNSP() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataCollectionEndpointsClient().GetNSP(ctx, "exampleRG", "someDataCollectionEndpoint", "somePerimeterConfiguration", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmonitor.DataCollectionEndpointsClientGetNSPResponse{
	// 	NetworkSecurityPerimeterConfiguration: armmonitor.NetworkSecurityPerimeterConfiguration{
	// 		Name: to.Ptr("somePerimeterConfiguration"),
	// 		Type: to.Ptr("Microsoft.Insights/dataCollectionEndpoints/networkSecurityPerimeterConfigurations"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/exampleRG/providers/Microsoft.Insights/dataCollectionEndpoints/someDataCollectionEndpoint/networkSecurityPerimeterConfigurations/somePerimeterConfiguration"),
	// 		Properties: &armmonitor.NetworkSecurityPerimeterConfigurationProperties{
	// 			NetworkSecurityPerimeter: &armmonitor.NetworkSecurityPerimeter{
	// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/networkRG/providers/Microsoft.Network/networkSecurityPerimeters/perimeter1"),
	// 				Location: to.Ptr("japaneast"),
	// 			},
	// 			Profile: &armmonitor.NetworkSecurityProfile{
	// 				Name: to.Ptr("profile1"),
	// 				AccessRules: []*armmonitor.AccessRule{
	// 					{
	// 						Name: to.Ptr("rule1"),
	// 						Properties: &armmonitor.AccessRuleProperties{
	// 							AddressPrefixes: []*string{
	// 								to.Ptr("148.0.0.0/8"),
	// 								to.Ptr("152.4.6.0/24"),
	// 							},
	// 							Direction: to.Ptr(armmonitor.AccessRuleDirectionInbound),
	// 						},
	// 					},
	// 				},
	// 				AccessRulesVersion: to.Ptr[int32](0),
	// 			},
	// 			ProvisioningState: to.Ptr(armmonitor.NetworkSecurityPerimeterConfigurationProvisioningStateAccepted),
	// 			ResourceAssociation: &armmonitor.ResourceAssociation{
	// 				Name: to.Ptr("assoc1"),
	// 				AccessMode: to.Ptr(armmonitor.ResourceAssociationAccessModeEnforced),
	// 			},
	// 		},
	// 	},
	// }
}
