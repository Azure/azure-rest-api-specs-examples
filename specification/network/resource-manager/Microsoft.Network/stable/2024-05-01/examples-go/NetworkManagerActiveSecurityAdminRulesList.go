package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ab04533261eff228f28e08900445d0edef3eb70c/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/NetworkManagerActiveSecurityAdminRulesList.json
func ExampleManagementClient_ListActiveSecurityAdminRules() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementClient().ListActiveSecurityAdminRules(ctx, "myResourceGroup", "testNetworkManager", armnetwork.ActiveConfigurationParameter{
		Regions: []*string{
			to.Ptr("westus")},
		SkipToken: to.Ptr("fakeSkipTokenCode"),
	}, &armnetwork.ManagementClientListActiveSecurityAdminRulesOptions{Top: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ActiveSecurityAdminRulesListResult = armnetwork.ActiveSecurityAdminRulesListResult{
	// 	SkipToken: to.Ptr("FakeSkipTokenCode"),
	// 	Value: []armnetwork.ActiveBaseSecurityAdminRuleClassification{
	// 		&armnetwork.ActiveDefaultSecurityAdminRule{
	// 			CommitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-10T12:33:22.257Z"); return t}()),
	// 			ConfigurationDescription: to.Ptr("SampleDescription"),
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/securityAdminConfigurations/myTestSecurityConfig/ruleCollections/testRuleCollection/rules/SampleAdminRule"),
	// 			Kind: to.Ptr(armnetwork.EffectiveAdminRuleKindDefault),
	// 			Region: to.Ptr("westus"),
	// 			RuleCollectionAppliesToGroups: []*armnetwork.ManagerSecurityGroupItem{
	// 				{
	// 					NetworkGroupID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroup/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/group1"),
	// 			}},
	// 			RuleCollectionDescription: to.Ptr("SampleRuleCollectionDescription"),
	// 			RuleGroups: []*armnetwork.ConfigurationGroup{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroup/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/group1"),
	// 					Properties: &armnetwork.GroupProperties{
	// 						Description: to.Ptr("A group for all test Virtual Networks"),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					},
	// 			}},
	// 			Properties: &armnetwork.DefaultAdminPropertiesFormat{
	// 				Description: to.Ptr("Sample Admin Rule"),
	// 				Access: to.Ptr(armnetwork.SecurityConfigurationRuleAccessDeny),
	// 				DestinationPortRanges: []*string{
	// 					to.Ptr("22")},
	// 					Destinations: []*armnetwork.AddressPrefixItem{
	// 						{
	// 							AddressPrefix: to.Ptr("*"),
	// 							AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
	// 					}},
	// 					Direction: to.Ptr(armnetwork.SecurityConfigurationRuleDirectionInbound),
	// 					Flag: to.Ptr("AllowVnetInbound"),
	// 					Priority: to.Ptr[int32](1),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					SourcePortRanges: []*string{
	// 						to.Ptr("0-65535")},
	// 						Sources: []*armnetwork.AddressPrefixItem{
	// 							{
	// 								AddressPrefix: to.Ptr("*"),
	// 								AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
	// 						}},
	// 						Protocol: to.Ptr(armnetwork.SecurityConfigurationRuleProtocolTCP),
	// 					},
	// 			}},
	// 		}
}
