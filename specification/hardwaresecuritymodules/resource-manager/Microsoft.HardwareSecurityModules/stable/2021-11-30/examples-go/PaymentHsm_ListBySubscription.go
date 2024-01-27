package armhardwaresecuritymodules_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/18b5c820705ab69735b7e1e2e0da5e37ca6e1969/specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/stable/2021-11-30/examples/PaymentHsm_ListBySubscription.json
func ExampleDedicatedHsmClient_NewListBySubscriptionPager_listDedicatedHsmDevicesInASubscriptionIncludingPaymentHsm() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhardwaresecuritymodules.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDedicatedHsmClient().NewListBySubscriptionPager(&armhardwaresecuritymodules.DedicatedHsmClientListBySubscriptionOptions{Top: nil})
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
		// page.DedicatedHsmListResult = armhardwaresecuritymodules.DedicatedHsmListResult{
		// 	Value: []*armhardwaresecuritymodules.DedicatedHsm{
		// 		{
		// 			Name: to.Ptr("hsm1"),
		// 			Type: to.Ptr("Microsoft.HardwareSecurityModules/dedicatedHSMs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/hsm1"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhardwaresecuritymodules.SKU{
		// 				Name: to.Ptr(armhardwaresecuritymodules.SKUNamePayShield10KLMK1CPS60),
		// 			},
		// 			Tags: map[string]*string{
		// 				"Dept": to.Ptr("hsm1"),
		// 				"Environment": to.Ptr("dogfood"),
		// 				"Slice": to.Ptr("A"),
		// 			},
		// 			Properties: &armhardwaresecuritymodules.DedicatedHsmProperties{
		// 				ManagementNetworkProfile: &armhardwaresecuritymodules.NetworkProfile{
		// 					NetworkInterfaces: []*armhardwaresecuritymodules.NetworkInterface{
		// 						{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/networkInterfaces/hsm1_HSMMgmtNic"),
		// 							PrivateIPAddress: to.Ptr("1.0.0.2"),
		// 					}},
		// 					Subnet: &armhardwaresecuritymodules.APIEntityReference{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/virtualNetworks/stamp01/subnets/stamp01"),
		// 					},
		// 				},
		// 				NetworkProfile: &armhardwaresecuritymodules.NetworkProfile{
		// 					NetworkInterfaces: []*armhardwaresecuritymodules.NetworkInterface{
		// 						{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/networkInterfaces/hsm1_HSMHost1Nic"),
		// 							PrivateIPAddress: to.Ptr("1.0.0.1"),
		// 					}},
		// 					Subnet: &armhardwaresecuritymodules.APIEntityReference{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/virtualNetworks/stamp01/subnets/stamp01"),
		// 					},
		// 				},
		// 				ProvisioningState: to.Ptr(armhardwaresecuritymodules.JSONWebKeyTypeSucceeded),
		// 				StampID: to.Ptr("stamp01"),
		// 				StatusMessage: to.Ptr("DedicatedHsm device is functional."),
		// 			},
		// 			SystemData: &armhardwaresecuritymodules.DedicatedHsmSystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:00:00.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armhardwaresecuritymodules.IdentityTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:00:00.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armhardwaresecuritymodules.IdentityTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("hsm2"),
		// 			Type: to.Ptr("Microsoft.HardwareSecurityModules/dedicatedHSMs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/hsm2"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhardwaresecuritymodules.SKU{
		// 				Name: to.Ptr(armhardwaresecuritymodules.SKUNamePayShield10KLMK1CPS60),
		// 			},
		// 			Tags: map[string]*string{
		// 				"Dept": to.Ptr("hsm"),
		// 				"Environment": to.Ptr("dogfood"),
		// 				"Slice": to.Ptr("B"),
		// 			},
		// 			Properties: &armhardwaresecuritymodules.DedicatedHsmProperties{
		// 				ManagementNetworkProfile: &armhardwaresecuritymodules.NetworkProfile{
		// 					NetworkInterfaces: []*armhardwaresecuritymodules.NetworkInterface{
		// 						{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/networkInterfaces/hsm2_HSMMgmtNic"),
		// 							PrivateIPAddress: to.Ptr("1.0.0.4"),
		// 					}},
		// 					Subnet: &armhardwaresecuritymodules.APIEntityReference{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/virtualNetworks/stamp01/subnets/stamp01"),
		// 					},
		// 				},
		// 				NetworkProfile: &armhardwaresecuritymodules.NetworkProfile{
		// 					NetworkInterfaces: []*armhardwaresecuritymodules.NetworkInterface{
		// 						{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/networkInterfaces/hsm2_HSMHost1Nic"),
		// 							PrivateIPAddress: to.Ptr("1.0.0.3"),
		// 					}},
		// 					Subnet: &armhardwaresecuritymodules.APIEntityReference{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/virtualNetworks/stamp01/subnets/stamp01"),
		// 					},
		// 				},
		// 				ProvisioningState: to.Ptr(armhardwaresecuritymodules.JSONWebKeyTypeSucceeded),
		// 				StampID: to.Ptr("stamp01"),
		// 				StatusMessage: to.Ptr("DedicatedHsm device is functional."),
		// 			},
		// 			SystemData: &armhardwaresecuritymodules.DedicatedHsmSystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:00:00.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armhardwaresecuritymodules.IdentityTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:00:00.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armhardwaresecuritymodules.IdentityTypeUser),
		// 			},
		// 	}},
		// }
	}
}
