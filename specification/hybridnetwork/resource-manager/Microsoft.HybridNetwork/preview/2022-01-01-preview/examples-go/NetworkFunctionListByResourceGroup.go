package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/792db17291c758b2bfdbbc0d35d0e2f5b5a1bd05/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/preview/2022-01-01-preview/examples/NetworkFunctionListByResourceGroup.json
func ExampleNetworkFunctionsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkFunctionsClient().NewListByResourceGroupPager("rg", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.NetworkFunctionListResult = armhybridnetwork.NetworkFunctionListResult{
		// 	Value: []*armhybridnetwork.NetworkFunction{
		// 		{
		// 			Name: to.Ptr("testNf"),
		// 			Type: to.Ptr("Microsoft.HybridNetwork/networkFunctions"),
		// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/networkFunctions/testNf"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armhybridnetwork.NetworkFunctionPropertiesFormat{
		// 				Device: &armhybridnetwork.SubResource{
		// 					ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/devices/testDevice"),
		// 				},
		// 				ManagedApplication: &armhybridnetwork.SubResource{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/testServiceKey"),
		// 				},
		// 				ManagedApplicationParameters: map[string]any{
		// 				},
		// 				NetworkFunctionUserConfigurations: []*armhybridnetwork.NetworkFunctionUserConfiguration{
		// 					{
		// 						NetworkInterfaces: []*armhybridnetwork.NetworkInterface{
		// 							{
		// 								IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
		// 									{
		// 										Gateway: to.Ptr(""),
		// 										IPAddress: to.Ptr(""),
		// 										IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
		// 										IPVersion: to.Ptr(armhybridnetwork.IPVersionIPv4),
		// 										Subnet: to.Ptr(""),
		// 								}},
		// 								MacAddress: to.Ptr(""),
		// 								NetworkInterfaceName: to.Ptr("nic1"),
		// 								VMSwitchType: to.Ptr(armhybridnetwork.VMSwitchTypeManagement),
		// 							},
		// 							{
		// 								IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
		// 									{
		// 										Gateway: to.Ptr(""),
		// 										IPAddress: to.Ptr(""),
		// 										IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
		// 										IPVersion: to.Ptr(armhybridnetwork.IPVersionIPv4),
		// 										Subnet: to.Ptr(""),
		// 								}},
		// 								MacAddress: to.Ptr("DC-97-F8-79-16-7D"),
		// 								NetworkInterfaceName: to.Ptr("nic2"),
		// 								VMSwitchType: to.Ptr(armhybridnetwork.VMSwitchTypeWan),
		// 						}},
		// 						RoleName: to.Ptr("testRole"),
		// 						UserDataParameters: map[string]any{
		// 						},
		// 				}},
		// 				ProvisioningState: to.Ptr(armhybridnetwork.ProvisioningStateSucceeded),
		// 				ServiceKey: to.Ptr("testServiceKey"),
		// 				SKUName: to.Ptr("testSku"),
		// 				SKUType: to.Ptr(armhybridnetwork.SKUTypeSDWAN),
		// 				VendorName: to.Ptr("testVendor"),
		// 				VendorProvisioningState: to.Ptr(armhybridnetwork.VendorProvisioningStateNotProvisioned),
		// 			},
		// 	}},
		// }
	}
}
