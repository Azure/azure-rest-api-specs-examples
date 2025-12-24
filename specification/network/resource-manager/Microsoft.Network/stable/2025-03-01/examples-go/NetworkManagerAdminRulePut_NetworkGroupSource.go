package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkManagerAdminRulePut_NetworkGroupSource.json
func ExampleAdminRulesClient_CreateOrUpdate_createAAdminRuleWithNetworkGroupAsSourceOrDestination() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAdminRulesClient().CreateOrUpdate(ctx, "rg1", "testNetworkManager", "myTestSecurityConfig", "testRuleCollection", "SampleAdminRule", &armnetwork.AdminRule{
		Kind: to.Ptr(armnetwork.AdminRuleKindCustom),
		Properties: &armnetwork.AdminPropertiesFormat{
			Description: to.Ptr("This is Sample Admin Rule"),
			Access:      to.Ptr(armnetwork.SecurityConfigurationRuleAccessDeny),
			DestinationPortRanges: []*string{
				to.Ptr("22")},
			Destinations: []*armnetwork.AddressPrefixItem{
				{
					AddressPrefix:     to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/ng1"),
					AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeNetworkGroup),
				}},
			Direction: to.Ptr(armnetwork.SecurityConfigurationRuleDirectionInbound),
			Priority:  to.Ptr[int32](1),
			SourcePortRanges: []*string{
				to.Ptr("0-65535")},
			Sources: []*armnetwork.AddressPrefixItem{
				{
					AddressPrefix:     to.Ptr("Internet"),
					AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeServiceTag),
				}},
			Protocol: to.Ptr(armnetwork.SecurityConfigurationRuleProtocolTCP),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.AdminRulesClientCreateOrUpdateResponse{
	// 	                            BaseAdminRuleClassification: &armnetwork.AdminRule{
	// 		Name: to.Ptr("SampleAdminRule"),
	// 		Type: to.Ptr("Microsoft.Network/networkManagers/securityAdminConfigurations/ruleCollections/rules"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/securityAdminConfigurations/myTestSecurityConfig/ruleCollections/testRuleCollection/rules/SampleAdminRule"),
	// 		Kind: to.Ptr(armnetwork.AdminRuleKindCustom),
	// 		SystemData: &armnetwork.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
	// 			CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 			CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 			LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 		},
	// 		Properties: &armnetwork.AdminPropertiesFormat{
	// 			Description: to.Ptr("This is Sample Admin Rule using a network group as a source and destination."),
	// 			Access: to.Ptr(armnetwork.SecurityConfigurationRuleAccessDeny),
	// 			DestinationPortRanges: []*string{
	// 				to.Ptr("22")},
	// 				Destinations: []*armnetwork.AddressPrefixItem{
	// 					{
	// 						AddressPrefix: to.Ptr("*"),
	// 						AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
	// 				}},
	// 				Direction: to.Ptr(armnetwork.SecurityConfigurationRuleDirectionInbound),
	// 				Priority: to.Ptr[int32](1),
	// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				SourcePortRanges: []*string{
	// 					to.Ptr("0-65535")},
	// 					Sources: []*armnetwork.AddressPrefixItem{
	// 						{
	// 							AddressPrefix: to.Ptr("Internet"),
	// 							AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeServiceTag),
	// 					}},
	// 					Protocol: to.Ptr(armnetwork.SecurityConfigurationRuleProtocolTCP),
	// 				},
	// 			},
	// 			                        }
}
