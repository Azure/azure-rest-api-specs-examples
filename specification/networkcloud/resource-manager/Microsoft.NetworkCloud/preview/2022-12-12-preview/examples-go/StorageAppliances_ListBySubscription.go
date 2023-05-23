package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/StorageAppliances_ListBySubscription.json
func ExampleStorageAppliancesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStorageAppliancesClient().NewListBySubscriptionPager(nil)
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
		// page.StorageApplianceList = armnetworkcloud.StorageApplianceList{
		// 	Value: []*armnetworkcloud.StorageAppliance{
		// 		{
		// 			Name: to.Ptr("storageApplianceName"),
		// 			Type: to.Ptr("Microsoft.NetworkCloud/storageAppliances"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/storageAppliances/storageApplianceName"),
		// 			SystemData: &armnetworkcloud.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
		// 				CreatedBy: to.Ptr("identityA"),
		// 				CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("identityB"),
		// 				LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("location"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("myvalue1"),
		// 				"key2": to.Ptr("myvalue2"),
		// 			},
		// 			ExtendedLocation: &armnetworkcloud.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
		// 				Type: to.Ptr("CustomLocation"),
		// 			},
		// 			Properties: &armnetworkcloud.StorageApplianceProperties{
		// 				AdministratorCredentials: &armnetworkcloud.AdministrativeCredentials{
		// 					Username: to.Ptr("adminUser"),
		// 				},
		// 				Capacity: to.Ptr[int64](893),
		// 				CapacityUsed: to.Ptr[int64](500),
		// 				ClusterID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName"),
		// 				DetailedStatus: to.Ptr(armnetworkcloud.StorageApplianceDetailedStatusAvailable),
		// 				DetailedStatusMessage: to.Ptr("Storage appliance is up and running"),
		// 				ManagementIPv4Address: to.Ptr("192.0.2.2"),
		// 				ProvisioningState: to.Ptr(armnetworkcloud.StorageApplianceProvisioningStateSucceeded),
		// 				RackID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/racks/rackName"),
		// 				RackSlot: to.Ptr[int64](1),
		// 				RemoteVendorManagementFeature: to.Ptr(armnetworkcloud.RemoteVendorManagementFeatureSupported),
		// 				RemoteVendorManagementStatus: to.Ptr(armnetworkcloud.RemoteVendorManagementStatusEnabled),
		// 				SerialNumber: to.Ptr("BM1219XXX"),
		// 				StorageApplianceSKUID: to.Ptr("684E-3B16-399E"),
		// 			},
		// 	}},
		// }
	}
}
