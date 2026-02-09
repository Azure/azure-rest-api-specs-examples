package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkManagerSecurityUserRuleGet.json
func ExampleSecurityUserRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSecurityUserRulesClient().Get(ctx, "rg1", "testNetworkManager", "myTestSecurityConfig", "testRuleCollection", "SampleUserRule", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SecurityUserRule = armnetwork.SecurityUserRule{
	// 	Name: to.Ptr("SampleUserRule"),
	// 	Type: to.Ptr("Microsoft.Network/networkManagers/SecurityUserConfigurations/ruleCollections/rules"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/SecurityUserConfigurations/myTestSecurityConfig/ruleCollections/testRuleCollection/rules/SampleUserRule"),
	// 	Properties: &armnetwork.SecurityUserRulePropertiesFormat{
	// 		Description: to.Ptr("Sample User Rule"),
	// 		DestinationPortRanges: []*string{
	// 			to.Ptr("22")},
	// 			Destinations: []*armnetwork.AddressPrefixItem{
	// 				{
	// 					AddressPrefix: to.Ptr("*"),
	// 					AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
	// 			}},
	// 			Direction: to.Ptr(armnetwork.SecurityConfigurationRuleDirectionInbound),
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			ResourceGUID: to.Ptr("f930553b-f78d-48c5-9445-6cf86b85e615"),
	// 			SourcePortRanges: []*string{
	// 				to.Ptr("0-65535")},
	// 				Sources: []*armnetwork.AddressPrefixItem{
	// 					{
	// 						AddressPrefix: to.Ptr("*"),
	// 						AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
	// 				}},
	// 				Protocol: to.Ptr(armnetwork.SecurityConfigurationRuleProtocolTCP),
	// 			},
	// 			SystemData: &armnetwork.SystemData{
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
	// 				CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 				CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
	// 				LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 				LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 			},
	// 		}
}
