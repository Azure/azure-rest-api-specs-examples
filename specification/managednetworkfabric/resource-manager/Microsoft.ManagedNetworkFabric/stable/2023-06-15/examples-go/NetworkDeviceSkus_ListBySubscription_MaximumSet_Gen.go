package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkDeviceSkus_ListBySubscription_MaximumSet_Gen.json
func ExampleNetworkDeviceSKUsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkDeviceSKUsClient().NewListBySubscriptionPager(nil)
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
		// page.NetworkDeviceSKUsListResult = armmanagednetworkfabric.NetworkDeviceSKUsListResult{
		// 	Value: []*armmanagednetworkfabric.NetworkDeviceSKU{
		// 		{
		// 			Name: to.Ptr("example-deviceSku"),
		// 			Type: to.Ptr("microsoft.managednetworkfabric/networkDeviceSkus"),
		// 			ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/providers/Microsoft.ManagedNetworkFabric/networkDeviceSkus/example-deviceSku"),
		// 			SystemData: &armmanagednetworkfabric.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T16:55:43.467Z"); return t}()),
		// 				CreatedBy: to.Ptr("email@address.com"),
		// 				CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T16:55:43.467Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@mail.com"),
		// 				LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 			},
		// 			Properties: &armmanagednetworkfabric.NetworkDeviceSKUProperties{
		// 				Interfaces: []*armmanagednetworkfabric.DeviceInterfaceProperties{
		// 					{
		// 						Identifier: to.Ptr("HundredGigE0/0"),
		// 						InterfaceType: to.Ptr("type1"),
		// 						SupportedConnectorTypes: []*armmanagednetworkfabric.SupportedConnectorProperties{
		// 							{
		// 								ConnectorType: to.Ptr("Optical"),
		// 								MaxSpeedInMbps: to.Ptr[int32](100),
		// 						}},
		// 				}},
		// 				Manufacturer: to.Ptr("Arista"),
		// 				Model: to.Ptr("model1"),
		// 				ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
		// 				SupportedRoleTypes: []*armmanagednetworkfabric.NetworkDeviceRoleName{
		// 					to.Ptr(armmanagednetworkfabric.NetworkDeviceRoleNameCE)},
		// 					SupportedVersions: []*armmanagednetworkfabric.SupportedVersionProperties{
		// 						{
		// 							IsDefault: to.Ptr(armmanagednetworkfabric.BooleanEnumPropertyTrue),
		// 							VendorFirmwareVersion: to.Ptr("11.2"),
		// 							VendorOsVersion: to.Ptr("2.0"),
		// 							Version: to.Ptr("1.0.0"),
		// 					}},
		// 				},
		// 		}},
		// 	}
	}
}
