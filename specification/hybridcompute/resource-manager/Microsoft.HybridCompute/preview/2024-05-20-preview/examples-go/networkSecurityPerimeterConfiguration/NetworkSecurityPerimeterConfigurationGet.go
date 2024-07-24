package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b574e2a41acda14a90ef237006e8bbdda2b63c63/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-05-20-preview/examples/networkSecurityPerimeterConfiguration/NetworkSecurityPerimeterConfigurationGet.json
func ExampleNetworkSecurityPerimeterConfigurationsClient_GetByPrivateLinkScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkSecurityPerimeterConfigurationsClient().GetByPrivateLinkScope(ctx, "my-resource-group", "my-privatelinkscope", "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee.myAssociation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkSecurityPerimeterConfiguration = armhybridcompute.NetworkSecurityPerimeterConfiguration{
	// 	Name: to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee.myAssociation"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes/networkSecurityPerimeterConfigurations"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/my-resource-group/providers/Microsoft.HybridCompute/privateLinkScopes/my-privatelinkscope/networkSecurityPerimeterConfigurations/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee.myAssociation"),
	// 	Properties: &armhybridcompute.NetworkSecurityPerimeterConfigurationProperties{
	// 		NetworkSecurityPerimeter: &armhybridcompute.NetworkSecurityPerimeter{
	// 			ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/networkSecurityPerimeters/myPerimeter"),
	// 			Location: to.Ptr("westus"),
	// 			PerimeterGUID: to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
	// 		},
	// 		Profile: &armhybridcompute.NetworkSecurityPerimeterProfile{
	// 			Name: to.Ptr("myProfile"),
	// 			AccessRules: []*armhybridcompute.AccessRule{
	// 				{
	// 					Name: to.Ptr("myAccessRule"),
	// 					Properties: &armhybridcompute.AccessRuleProperties{
	// 						AddressPrefixes: []*string{
	// 							to.Ptr("148.0.0.0/8"),
	// 							to.Ptr("152.4.6.0/24"),
	// 							to.Ptr("...")},
	// 							Direction: to.Ptr(armhybridcompute.AccessRuleDirectionInbound),
	// 						},
	// 				}},
	// 				AccessRulesVersion: to.Ptr[int32](1),
	// 				DiagnosticSettingsVersion: to.Ptr[int32](1),
	// 				EnabledLogCategories: []*string{
	// 				},
	// 			},
	// 			ProvisioningIssues: []*armhybridcompute.ProvisioningIssue{
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			ResourceAssociation: &armhybridcompute.ResourceAssociation{
	// 				Name: to.Ptr("myAssociation"),
	// 				AccessMode: to.Ptr(armhybridcompute.AccessModeEnforced),
	// 			},
	// 		},
	// 	}
}
