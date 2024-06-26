package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkDevices_ListBySubscription_MaximumSet_Gen.json
func ExampleNetworkDevicesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkDevicesClient().NewListBySubscriptionPager(nil)
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
		// 			Name: to.Ptr("networkDeviceName"),
		// 			Type: to.Ptr("microsoft.managednetworkfabric/networkdevices"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkDevices/networkDeviceName"),
		// 			SystemData: &armmanagednetworkfabric.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-18T07:58:04.840Z"); return t}()),
		// 				CreatedBy: to.Ptr("d1bd24c7-b27f-477e-86dd-939e107873d7"),
		// 				CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-18T07:58:04.840Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("email@address.com"),
		// 				LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"keyID": to.Ptr("keyValue"),
		// 			},
		// 			Properties: &armmanagednetworkfabric.NetworkDeviceProperties{
		// 				Annotation: to.Ptr("null"),
		// 				HostName: to.Ptr("networkDeviceName"),
		// 				SerialNumber: to.Ptr("Arista;DCS-7280PR3-24;12.05;JPE21330382"),
		// 				NetworkDeviceRole: to.Ptr(armmanagednetworkfabric.NetworkDeviceRoleTypesCE),
		// 				NetworkDeviceSKU: to.Ptr("DefaultSku"),
		// 				NetworkRackID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkRacks/networkRackName"),
		// 				ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
		// 				Version: to.Ptr("null"),
		// 			},
		// 	}},
		// }
	}
}
