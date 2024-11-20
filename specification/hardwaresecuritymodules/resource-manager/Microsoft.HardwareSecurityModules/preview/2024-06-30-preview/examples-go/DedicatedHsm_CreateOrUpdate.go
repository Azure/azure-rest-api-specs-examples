package armhardwaresecuritymodules_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e838027e88cca634c1545e744630de9262a6e72a/specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2024-06-30-preview/examples/DedicatedHsm_CreateOrUpdate.json
func ExampleDedicatedHsmClient_BeginCreateOrUpdate_createANewOrUpdateAnExistingDedicatedHsm() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhardwaresecuritymodules.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDedicatedHsmClient().BeginCreateOrUpdate(ctx, "hsm-group", "hsm1", armhardwaresecuritymodules.DedicatedHsm{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"Dept":        to.Ptr("hsm"),
			"Environment": to.Ptr("dogfood"),
		},
		Properties: &armhardwaresecuritymodules.DedicatedHsmProperties{
			NetworkProfile: &armhardwaresecuritymodules.NetworkProfile{
				NetworkInterfaces: []*armhardwaresecuritymodules.NetworkInterface{
					{
						PrivateIPAddress: to.Ptr("1.0.0.1"),
					}},
				Subnet: &armhardwaresecuritymodules.APIEntityReference{
					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/virtualNetworks/stamp01/subnets/stamp01"),
				},
			},
			StampID: to.Ptr("stamp01"),
		},
		SKU: &armhardwaresecuritymodules.SKU{
			Name: to.Ptr(armhardwaresecuritymodules.SKUNameSafeNetLunaNetworkHSMA790),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DedicatedHsm = armhardwaresecuritymodules.DedicatedHsm{
	// 	Name: to.Ptr("hsm1"),
	// 	Type: to.Ptr("Microsoft.HardwareSecurityModules/dedicatedHSMs"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/hsm1"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"Dept": to.Ptr("hsm"),
	// 		"Environment": to.Ptr("dogfood"),
	// 	},
	// 	Properties: &armhardwaresecuritymodules.DedicatedHsmProperties{
	// 		NetworkProfile: &armhardwaresecuritymodules.NetworkProfile{
	// 			NetworkInterfaces: []*armhardwaresecuritymodules.NetworkInterface{
	// 				{
	// 					PrivateIPAddress: to.Ptr("1.0.0.1"),
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/networkInterfaces/hsm_vnic"),
	// 			}},
	// 			Subnet: &armhardwaresecuritymodules.APIEntityReference{
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/virtualNetworks/stamp01/subnets/stamp01"),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armhardwaresecuritymodules.JSONWebKeyTypeSucceeded),
	// 		StampID: to.Ptr("stamp01"),
	// 		StatusMessage: to.Ptr("DedicatedHsm device is functional."),
	// 	},
	// 	SKU: &armhardwaresecuritymodules.SKU{
	// 		Name: to.Ptr(armhardwaresecuritymodules.SKUNameSafeNetLunaNetworkHSMA790),
	// 	},
	// }
}
