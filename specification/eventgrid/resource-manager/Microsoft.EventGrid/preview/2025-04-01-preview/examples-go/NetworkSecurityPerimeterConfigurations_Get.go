package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/NetworkSecurityPerimeterConfigurations_Get.json
func ExampleNetworkSecurityPerimeterConfigurationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkSecurityPerimeterConfigurationsClient().Get(ctx, "examplerg", armeventgrid.NetworkSecurityPerimeterResourceTypeTopics, "exampleResourceName", "8f6b6269-84f2-4d09-9e31-1127efcd1e40perimeter", "someAssociation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkSecurityPerimeterConfiguration = armeventgrid.NetworkSecurityPerimeterConfiguration{
	// 	Name: to.Ptr("8f6b6269-84f2-4d09-9e31-1127efcd1e40perimeter.someAssociation"),
	// 	Type: to.Ptr("Microsoft.EventGrid/topics/networkSecurityPerimeterConfigurations"),
	// 	ID: to.Ptr("/subscriptions//resourceGroups/paasrg/providers/Microsoft.EventGrid/topics/egtopic/networkSecurityPerimeterConfigurations/8f6b6269-84f2-4d09-9e31-1127efcd1e40perimeter.someAssociation"),
	// 	Properties: &armeventgrid.NetworkSecurityPerimeterConfigurationProperties{
	// 		NetworkSecurityPerimeter: &armeventgrid.NetworkSecurityPerimeterInfo{
	// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/perimeterrg/providers/Microsoft.Network/networkSecurityPerimeters/somePerimeter"),
	// 			Location: to.Ptr("West US"),
	// 			PerimeterGUID: to.Ptr("8f6b6269-84f2-4d09-9e31-1127efcd1e40perimeter"),
	// 		},
	// 		Profile: &armeventgrid.NetworkSecurityPerimeterConfigurationProfile{
	// 			Name: to.Ptr("someProfile"),
	// 			AccessRules: []*armeventgrid.NetworkSecurityPerimeterProfileAccessRule{
	// 				{
	// 					Name: to.Ptr("ipRule"),
	// 					Properties: &armeventgrid.NetworkSecurityPerimeterProfileAccessRuleProperties{
	// 						AddressPrefixes: []*string{
	// 							to.Ptr("148.0.0.0/8")},
	// 							Direction: to.Ptr(armeventgrid.NetworkSecurityPerimeterProfileAccessRuleDirectionInbound),
	// 						},
	// 				}},
	// 				AccessRulesVersion: to.Ptr("1"),
	// 				DiagnosticSettingsVersion: to.Ptr("1"),
	// 				EnabledLogCategories: []*string{
	// 					to.Ptr("NspOutbound")},
	// 				},
	// 				ProvisioningState: to.Ptr(armeventgrid.NetworkSecurityPerimeterConfigProvisioningStateSucceeded),
	// 				ResourceAssociation: &armeventgrid.ResourceAssociation{
	// 					Name: to.Ptr("someAssociation"),
	// 					AccessMode: to.Ptr(armeventgrid.NetworkSecurityPerimeterAssociationAccessModeEnforced),
	// 				},
	// 			},
	// 		}
}
