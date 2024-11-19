package armhardwaresecuritymodules_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e838027e88cca634c1545e744630de9262a6e72a/specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2024-06-30-preview/examples/PaymentHsm_Get.json
func ExampleDedicatedHsmClient_Get_getAPaymentHsm() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhardwaresecuritymodules.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDedicatedHsmClient().Get(ctx, "hsm-group", "hsm1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DedicatedHsm = armhardwaresecuritymodules.DedicatedHsm{
	// 	Name: to.Ptr("hsm1"),
	// 	Type: to.Ptr("Microsoft.HardwareSecurityModules/dedicatedHSMs"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/hsm1"),
	// 	SystemData: &armhardwaresecuritymodules.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:00:00.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armhardwaresecuritymodules.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:00:00.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armhardwaresecuritymodules.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"Dept": to.Ptr("hsm"),
	// 		"Environment": to.Ptr("dogfood"),
	// 		"Slice": to.Ptr("A"),
	// 	},
	// 	Properties: &armhardwaresecuritymodules.DedicatedHsmProperties{
	// 		ManagementNetworkProfile: &armhardwaresecuritymodules.NetworkProfile{
	// 			NetworkInterfaces: []*armhardwaresecuritymodules.NetworkInterface{
	// 				{
	// 					PrivateIPAddress: to.Ptr("1.0.0.2"),
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/networkInterfaces/hsm1_HSMMgmtNic"),
	// 			}},
	// 			Subnet: &armhardwaresecuritymodules.APIEntityReference{
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/virtualNetworks/stamp01/subnets/stamp01"),
	// 			},
	// 		},
	// 		NetworkProfile: &armhardwaresecuritymodules.NetworkProfile{
	// 			NetworkInterfaces: []*armhardwaresecuritymodules.NetworkInterface{
	// 				{
	// 					PrivateIPAddress: to.Ptr("1.0.0.1"),
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/networkInterfaces/hsm1_HSMHost1Nic"),
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
	// 		Name: to.Ptr(armhardwaresecuritymodules.SKUNamePayShield10KLMK1CPS60),
	// 	},
	// }
}
