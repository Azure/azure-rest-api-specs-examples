package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkDevices_ListByResourceGroup_MaximumSet_Gen.json
func ExampleNetworkDevicesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkDevicesClient().NewListByResourceGroupPager("example-rg", nil)
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
		// page.NetworkDevicesListResult = armmanagednetworkfabric.NetworkDevicesListResult{
		// 	Value: []*armmanagednetworkfabric.NetworkDevice{
		// 		{
		// 			Name: to.Ptr("example-device"),
		// 			Type: to.Ptr("microsoft.managednetworkfabric/networkdevices"),
		// 			ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-device"),
		// 			SystemData: &armmanagednetworkfabric.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T16:02:19.538Z"); return t}()),
		// 				CreatedBy: to.Ptr("email@address.com"),
		// 				CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T16:02:19.538Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("email@address.com"),
		// 				LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastuseuap"),
		// 			Tags: map[string]*string{
		// 				"keyID": to.Ptr("KeyValue"),
		// 			},
		// 			Properties: &armmanagednetworkfabric.NetworkDeviceProperties{
		// 				Annotation: to.Ptr("annotation"),
		// 				HostName: to.Ptr("NFA-Device"),
		// 				SerialNumber: to.Ptr("Vendor;DCS-7280XXX-24;12.05;JPE2111XXXX"),
		// 				AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
		// 				ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
		// 				ManagementIPv4Address: to.Ptr("10.10.10.10"),
		// 				ManagementIPv6Address: to.Ptr("2F::1/100"),
		// 				NetworkDeviceRole: to.Ptr(armmanagednetworkfabric.NetworkDeviceRoleCE),
		// 				NetworkDeviceSKU: to.Ptr("DeviceSku"),
		// 				NetworkRackID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkRacks/example-rack"),
		// 				ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
		// 				Version: to.Ptr("1.0"),
		// 			},
		// 	}},
		// }
	}
}
