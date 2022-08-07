package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerAdminRulePut.json
func ExampleAdminRulesClient_CreateOrUpdate_networkManagerAdminRulePut() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewAdminRulesClient("subId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "rg1", "testNetworkManager", "myTestSecurityConfig", "testRuleCollection", "SampleAdminRule", &armnetwork.AdminRule{
		Kind: to.Ptr(armnetwork.AdminRuleKindCustom),
		Properties: &armnetwork.AdminPropertiesFormat{
			Description: to.Ptr("This is Sample Admin Rule"),
			Access:      to.Ptr(armnetwork.SecurityConfigurationRuleAccessDeny),
			DestinationPortRanges: []*string{
				to.Ptr("22")},
			Destinations: []*armnetwork.AddressPrefixItem{
				{
					AddressPrefix:     to.Ptr("*"),
					AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
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
	// TODO: use response item
	_ = res
}
