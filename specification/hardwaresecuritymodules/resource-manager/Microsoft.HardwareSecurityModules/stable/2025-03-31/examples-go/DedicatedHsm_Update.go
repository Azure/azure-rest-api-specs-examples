package armhardwaresecuritymodules_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules/v2"
)

// Generated from example definition: 2025-03-31/DedicatedHsm_Update.json
func ExampleDedicatedHsmClient_BeginUpdate_updateAnExistingDedicatedHsm() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhardwaresecuritymodules.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDedicatedHsmClient().BeginUpdate(ctx, "hsm-group", "hsm1", armhardwaresecuritymodules.DedicatedHsmPatchParameters{
		Tags: map[string]*string{
			"Dept":        to.Ptr("hsm"),
			"Environment": to.Ptr("dogfood"),
			"Slice":       to.Ptr("A"),
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
	// res = armhardwaresecuritymodules.DedicatedHsmClientUpdateResponse{
	// 	DedicatedHsm: &armhardwaresecuritymodules.DedicatedHsm{
	// 		Name: to.Ptr("hsm1"),
	// 		Type: to.Ptr("Microsoft.HardwareSecurityModules/dedicatedHSMs"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/hsm1"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armhardwaresecuritymodules.DedicatedHsmProperties{
	// 			NetworkProfile: &armhardwaresecuritymodules.NetworkProfile{
	// 				NetworkInterfaces: []*armhardwaresecuritymodules.NetworkInterface{
	// 					{
	// 						PrivateIPAddress: to.Ptr("1.0.0.1"),
	// 						ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/networkInterfaces/hsm_vnic"),
	// 					},
	// 				},
	// 				Subnet: &armhardwaresecuritymodules.APIEntityReference{
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/virtualNetworks/stamp01/subnets/stamp01"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armhardwaresecuritymodules.JSONWebKeyTypeSucceeded),
	// 			StampID: to.Ptr("stamp01"),
	// 			StatusMessage: to.Ptr("DedicatedHsm device is functional."),
	// 		},
	// 		SKU: &armhardwaresecuritymodules.SKU{
	// 			Name: to.Ptr(armhardwaresecuritymodules.SKUNameSafeNetLunaNetworkHSMA790),
	// 		},
	// 		Tags: map[string]*string{
	// 			"Dept": to.Ptr("hsm"),
	// 			"Environment": to.Ptr("dogfood"),
	// 			"Slice": to.Ptr("A"),
	// 		},
	// 	},
	// }
}
