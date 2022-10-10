package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkProfileCreateConfigOnly.json
func ExampleProfilesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewProfilesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "rg1", "networkProfile1", armnetwork.Profile{
		Location: to.Ptr("westus"),
		Properties: &armnetwork.ProfilePropertiesFormat{
			ContainerNetworkInterfaceConfigurations: []*armnetwork.ContainerNetworkInterfaceConfiguration{
				{
					Name: to.Ptr("eth1"),
					Properties: &armnetwork.ContainerNetworkInterfaceConfigurationPropertiesFormat{
						IPConfigurations: []*armnetwork.IPConfigurationProfile{
							{
								Name: to.Ptr("ipconfig1"),
								Properties: &armnetwork.IPConfigurationProfilePropertiesFormat{
									Subnet: &armnetwork.Subnet{
										ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/networkProfileVnet/subnets/networkProfileSubnet1"),
									},
								},
							}},
					},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
