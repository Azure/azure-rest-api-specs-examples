package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c58fa033619b12c7cfa8a0ec5a9bf03bb18869ab/specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NetworkManagerDefaultAdminRuleGet.json
func ExampleAdminRulesClient_Get_getsSecurityDefaultAdminRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAdminRulesClient().Get(ctx, "rg1", "testNetworkManager", "myTestSecurityConfig", "testRuleCollection", "SampleDefaultAdminRule", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.AdminRulesClientGetResponse{
	// 	                            BaseAdminRuleClassification: &armnetwork.DefaultAdminRule{
	// 		Name: to.Ptr("SampleDefaultAdminRule"),
	// 		Type: to.Ptr("Microsoft.Network/networkManagers/securityAdminConfigurations/ruleCollections/rules"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/securityAdminConfigurations/myTestSecurityConfig/ruleCollections/testRuleCollection/rules/SampleDefaultAdminRule"),
	// 		Kind: to.Ptr(armnetwork.AdminRuleKindDefault),
	// 		SystemData: &armnetwork.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
	// 			CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 			CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 			LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 		},
	// 		Properties: &armnetwork.DefaultAdminPropertiesFormat{
	// 			Description: to.Ptr("This is Sample Default Admin Rule"),
	// 			Access: to.Ptr(armnetwork.SecurityConfigurationRuleAccessDeny),
	// 			DestinationPortRanges: []*string{
	// 				to.Ptr("22")},
	// 				Destinations: []*armnetwork.AddressPrefixItem{
	// 					{
	// 						AddressPrefix: to.Ptr("*"),
	// 						AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
	// 				}},
	// 				Direction: to.Ptr(armnetwork.SecurityConfigurationRuleDirectionInbound),
	// 				Flag: to.Ptr("AllowVnetInbound"),
	// 				Priority: to.Ptr[int32](1),
	// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				SourcePortRanges: []*string{
	// 					to.Ptr("0-65535")},
	// 					Sources: []*armnetwork.AddressPrefixItem{
	// 						{
	// 							AddressPrefix: to.Ptr("*"),
	// 							AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
	// 					}},
	// 					Protocol: to.Ptr(armnetwork.SecurityConfigurationRuleProtocolTCP),
	// 				},
	// 			},
	// 			                        }
}
