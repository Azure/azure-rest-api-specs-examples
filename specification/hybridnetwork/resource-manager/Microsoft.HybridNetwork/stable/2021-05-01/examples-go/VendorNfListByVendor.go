package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/VendorNfListByVendor.json
func ExampleVendorNetworkFunctionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVendorNetworkFunctionsClient().NewListPager("eastus", "testVendor", &armhybridnetwork.VendorNetworkFunctionsClientListOptions{Filter: nil})
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
		// page.VendorNetworkFunctionListResult = armhybridnetwork.VendorNetworkFunctionListResult{
		// 	Value: []*armhybridnetwork.VendorNetworkFunction{
		// 		{
		// 			Name: to.Ptr("TestServiceKey"),
		// 			Type: to.Ptr("Microsoft.HybridNetwork/locations/vendors/networkFunctions"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.HybridNetwork/locations/eastus/vendors/testVendor/networkFunctions/testServiceKey"),
		// 			Properties: &armhybridnetwork.VendorNetworkFunctionPropertiesFormat{
		// 				NetworkFunctionVendorConfigurations: []*armhybridnetwork.NetworkFunctionVendorConfiguration{
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
		// 						OSProfile: &armhybridnetwork.OsProfile{
		// 							AdminUsername: to.Ptr("dummyuser"),
		// 							CustomData: to.Ptr("base-64 encoded string of custom data"),
		// 							LinuxConfiguration: &armhybridnetwork.LinuxConfiguration{
		// 								SSH: &armhybridnetwork.SSHConfiguration{
		// 									PublicKeys: []*armhybridnetwork.SSHPublicKey{
		// 										{
		// 											Path: to.Ptr("home/user/.ssh/authorized_keys"),
		// 											KeyData: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAgEAwrr66r8n6B8Y0zMF3dOpXEapIQD9DiYQ6D6/zwor9o39jSkHNiMMER/GETBbzP83LOcekm02aRjo55ArO7gPPVvCXbrirJu9pkm4AC4BBre5xSLS= user@constoso-DSH"),
		// 									}},
		// 								},
		// 							},
		// 						},
		// 						RoleName: to.Ptr("testRole"),
		// 						UserDataParameters: map[string]any{
		// 						},
		// 				}},
		// 				ProvisioningState: to.Ptr(armhybridnetwork.ProvisioningStateSucceeded),
		// 				SKUName: to.Ptr("testSku"),
		// 				SKUType: to.Ptr(armhybridnetwork.SKUTypeSDWAN),
		// 				VendorProvisioningState: to.Ptr(armhybridnetwork.VendorProvisioningStateProvisioning),
		// 			},
		// 	}},
		// }
	}
}
