package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/NetworkFunctionSkuDetailsGet.json
func ExampleNetworkFunctionVendorSKUsClient_NewListBySKUPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkFunctionVendorSKUsClient().NewListBySKUPager("testVendor", "testSku", nil)
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
		// page.NetworkFunctionSKUDetails = armhybridnetwork.NetworkFunctionSKUDetails{
		// 	SKUType: to.Ptr(armhybridnetwork.SKUTypeSDWAN),
		// 	Value: []*armhybridnetwork.NetworkFunctionSKURoleDetails{
		// 		{
		// 			NetworkInterfaces: []*armhybridnetwork.NetworkInterface{
		// 				{
		// 					IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
		// 						{
		// 							Gateway: to.Ptr(""),
		// 							IPAddress: to.Ptr(""),
		// 							IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
		// 							IPVersion: to.Ptr(armhybridnetwork.IPVersionIPv4),
		// 							Subnet: to.Ptr(""),
		// 					}},
		// 					MacAddress: to.Ptr(""),
		// 					NetworkInterfaceName: to.Ptr("nic1"),
		// 					VMSwitchType: to.Ptr(armhybridnetwork.VMSwitchTypeManagement),
		// 				},
		// 				{
		// 					IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
		// 						{
		// 							Gateway: to.Ptr(""),
		// 							IPAddress: to.Ptr(""),
		// 							IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
		// 							IPVersion: to.Ptr(armhybridnetwork.IPVersionIPv4),
		// 							Subnet: to.Ptr(""),
		// 					}},
		// 					MacAddress: to.Ptr(""),
		// 					NetworkInterfaceName: to.Ptr("nic2"),
		// 					VMSwitchType: to.Ptr(armhybridnetwork.VMSwitchTypeWan),
		// 			}},
		// 			RoleName: to.Ptr("testRole"),
		// 			UserDataParameters: map[string]any{
		// 			},
		// 			UserDataTemplate: map[string]any{
		// 			},
		// 	}},
		// }
	}
}
