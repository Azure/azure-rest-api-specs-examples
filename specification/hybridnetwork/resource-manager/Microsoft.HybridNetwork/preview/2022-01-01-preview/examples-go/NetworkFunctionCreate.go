package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/preview/2022-01-01-preview/examples/NetworkFunctionCreate.json
func ExampleNetworkFunctionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhybridnetwork.NewNetworkFunctionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg",
		"testNf",
		armhybridnetwork.NetworkFunction{
			Location: to.Ptr("eastus"),
			Properties: &armhybridnetwork.NetworkFunctionPropertiesFormat{
				Device: &armhybridnetwork.SubResource{
					ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/devices/testDevice"),
				},
				ManagedApplicationParameters: map[string]interface{}{},
				NetworkFunctionUserConfigurations: []*armhybridnetwork.NetworkFunctionUserConfiguration{
					{
						NetworkInterfaces: []*armhybridnetwork.NetworkInterface{
							{
								IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
									{
										Gateway:            to.Ptr(""),
										IPAddress:          to.Ptr(""),
										IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
										IPVersion:          to.Ptr(armhybridnetwork.IPVersionIPv4),
										Subnet:             to.Ptr(""),
									}},
								MacAddress:           to.Ptr(""),
								NetworkInterfaceName: to.Ptr("nic1"),
								VMSwitchType:         to.Ptr(armhybridnetwork.VMSwitchTypeManagement),
							},
							{
								IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
									{
										Gateway:            to.Ptr(""),
										IPAddress:          to.Ptr(""),
										IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
										IPVersion:          to.Ptr(armhybridnetwork.IPVersionIPv4),
										Subnet:             to.Ptr(""),
									}},
								MacAddress:           to.Ptr("DC-97-F8-79-16-7D"),
								NetworkInterfaceName: to.Ptr("nic2"),
								VMSwitchType:         to.Ptr(armhybridnetwork.VMSwitchTypeWan),
							}},
						RoleName:           to.Ptr("testRole"),
						UserDataParameters: map[string]interface{}{},
					}},
				SKUName:    to.Ptr("testSku"),
				SKUType:    to.Ptr(armhybridnetwork.SKUTypeSDWAN),
				VendorName: to.Ptr("testVendor"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
