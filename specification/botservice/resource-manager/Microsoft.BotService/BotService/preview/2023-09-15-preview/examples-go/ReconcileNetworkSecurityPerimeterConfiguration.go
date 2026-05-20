package armbotservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice/v2"
)

// Generated from example definition: 2023-09-15-preview/ReconcileNetworkSecurityPerimeterConfiguration.json
func ExampleNetworkSecurityPerimeterConfigurationsClient_BeginReconcile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbotservice.NewClientFactory("subId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkSecurityPerimeterConfigurationsClient().BeginReconcile(ctx, "rgName", "botId", "00000000-0000-0000-0000-000000000000.associationName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbotservice.NetworkSecurityPerimeterConfigurationsClientReconcileResponse{
	// 	NetworkSecurityPerimeterConfiguration: &armbotservice.NetworkSecurityPerimeterConfiguration{
	// 		Name: to.Ptr("00000000-0000-0000-0000-000000000000.associationName"),
	// 		Type: to.Ptr("Microsoft.BotService/botServices/networkSecurityPerimeterConfigurations"),
	// 		ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.BotService/botServices/botId/networkSecurityPerimeterConfigurations/00000000-0000-0000-0000-000000000000.associationName"),
	// 		Properties: &armbotservice.NetworkSecurityPerimeterConfigurationProperties{
	// 			NetworkSecurityPerimeter: &armbotservice.NetworkSecurityPerimeter{
	// 				ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Network/networkSecurityPerimeters/nspName"),
	// 				Location: to.Ptr("westcentralus"),
	// 				PerimeterGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			},
	// 			Profile: &armbotservice.Profile{
	// 				Name: to.Ptr("defaultProfile"),
	// 				AccessRules: []*armbotservice.NspAccessRule{
	// 					{
	// 						Name: to.Ptr("Inbound1"),
	// 						Properties: &armbotservice.NspAccessRuleProperties{
	// 							AddressPrefixes: []*string{
	// 								to.Ptr("10.0.0.0/16"),
	// 							},
	// 							Direction: to.Ptr(armbotservice.NspAccessRuleDirectionInbound),
	// 							EmailAddresses: []*string{
	// 							},
	// 							FullyQualifiedDomainNames: []*string{
	// 							},
	// 							NetworkSecurityPerimeters: []*armbotservice.NetworkSecurityPerimeter{
	// 							},
	// 							PhoneNumbers: []*string{
	// 							},
	// 							Subscriptions: []*armbotservice.NspAccessRulePropertiesSubscriptionsItem{
	// 							},
	// 						},
	// 					},
	// 					{
	// 						Name: to.Ptr("Outbound1"),
	// 						Properties: &armbotservice.NspAccessRuleProperties{
	// 							AddressPrefixes: []*string{
	// 							},
	// 							Direction: to.Ptr(armbotservice.NspAccessRuleDirectionOutbound),
	// 							EmailAddresses: []*string{
	// 							},
	// 							FullyQualifiedDomainNames: []*string{
	// 								to.Ptr("*"),
	// 							},
	// 							NetworkSecurityPerimeters: []*armbotservice.NetworkSecurityPerimeter{
	// 							},
	// 							PhoneNumbers: []*string{
	// 							},
	// 							Subscriptions: []*armbotservice.NspAccessRulePropertiesSubscriptionsItem{
	// 							},
	// 						},
	// 					},
	// 				},
	// 				AccessRulesVersion: to.Ptr[int64](2),
	// 				DiagnosticSettingsVersion: to.Ptr[int64](0),
	// 				EnabledLogCategories: []*string{
	// 				},
	// 			},
	// 			ProvisioningIssues: []*armbotservice.ProvisioningIssue{
	// 			},
	// 			ProvisioningState: to.Ptr(armbotservice.ProvisioningStateSucceeded),
	// 			ResourceAssociation: &armbotservice.ResourceAssociation{
	// 				Name: to.Ptr("associationName"),
	// 				AccessMode: to.Ptr(armbotservice.AccessModeLearning),
	// 			},
	// 		},
	// 	},
	// }
}
