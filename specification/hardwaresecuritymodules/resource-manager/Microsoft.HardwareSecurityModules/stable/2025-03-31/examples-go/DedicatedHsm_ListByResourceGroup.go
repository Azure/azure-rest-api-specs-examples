package armhardwaresecuritymodules_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules/v2"
)

// Generated from example definition: 2025-03-31/DedicatedHsm_ListByResourceGroup.json
func ExampleDedicatedHsmClient_NewListByResourceGroupPager_listDedicatedHsmDevicesInAResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhardwaresecuritymodules.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDedicatedHsmClient().NewListByResourceGroupPager("hsm-group", nil)
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
		// page = armhardwaresecuritymodules.DedicatedHsmClientListByResourceGroupResponse{
		// 	DedicatedHsmListResult: armhardwaresecuritymodules.DedicatedHsmListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs?api-version=2018-10-31-preview&$skiptoken=dmF1bHQtcGVza3ktanVyeS03MzA3Ng=="),
		// 		Value: []*armhardwaresecuritymodules.DedicatedHsm{
		// 			{
		// 				Name: to.Ptr("hsm1"),
		// 				Type: to.Ptr("Microsoft.HardwareSecurityModules/dedicatedHSMs"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/hsm1"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armhardwaresecuritymodules.DedicatedHsmProperties{
		// 					NetworkProfile: &armhardwaresecuritymodules.NetworkProfile{
		// 						NetworkInterfaces: []*armhardwaresecuritymodules.NetworkInterface{
		// 							{
		// 								PrivateIPAddress: to.Ptr("1.0.0.1"),
		// 								ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/networkInterfaces/hsm_vnic"),
		// 							},
		// 						},
		// 						Subnet: &armhardwaresecuritymodules.APIEntityReference{
		// 							ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/virtualNetworks/stamp01/subnets/stamp01"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armhardwaresecuritymodules.JSONWebKeyTypeSucceeded),
		// 					StampID: to.Ptr("stamp01"),
		// 					StatusMessage: to.Ptr("DedicatedHsm device is functional."),
		// 				},
		// 				SKU: &armhardwaresecuritymodules.SKU{
		// 					Name: to.Ptr(armhardwaresecuritymodules.SKUNameSafeNetLunaNetworkHSMA790),
		// 				},
		// 				Tags: map[string]*string{
		// 					"Dept": to.Ptr("hsm"),
		// 					"Environment": to.Ptr("dogfood"),
		// 					"Slice": to.Ptr("A"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("hsm1"),
		// 				Type: to.Ptr("Microsoft.HardwareSecurityModules/dedicatedHSMs"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/hsm2"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armhardwaresecuritymodules.DedicatedHsmProperties{
		// 					NetworkProfile: &armhardwaresecuritymodules.NetworkProfile{
		// 						NetworkInterfaces: []*armhardwaresecuritymodules.NetworkInterface{
		// 							{
		// 								PrivateIPAddress: to.Ptr("1.0.0.2"),
		// 								ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/networkInterfaces/hsm_vnic"),
		// 							},
		// 						},
		// 						Subnet: &armhardwaresecuritymodules.APIEntityReference{
		// 							ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/virtualNetworks/stamp01/subnets/stamp01"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armhardwaresecuritymodules.JSONWebKeyTypeSucceeded),
		// 					StampID: to.Ptr("stamp01"),
		// 					StatusMessage: to.Ptr("DedicatedHsm device is functional."),
		// 				},
		// 				SKU: &armhardwaresecuritymodules.SKU{
		// 					Name: to.Ptr(armhardwaresecuritymodules.SKUNameSafeNetLunaNetworkHSMA790),
		// 				},
		// 				Tags: map[string]*string{
		// 					"Dept": to.Ptr("hsm"),
		// 					"Environment": to.Ptr("dogfood"),
		// 					"Slice": to.Ptr("B"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
