package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/StorageSyncServices_ListByResourceGroup.json
func ExampleServicesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServicesClient().NewListByResourceGroupPager("SampleResourceGroup_1", nil)
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
		// page.ServiceArray = armstoragesync.ServiceArray{
		// 	Value: []*armstoragesync.Service{
		// 		{
		// 			Name: to.Ptr("SampleStorageSyncService_1"),
		// 			Type: to.Ptr("Microsoft.StorageSync/storageSyncServices"),
		// 			ID: to.Ptr("/subscriptions/3a048283-338f-4002-a9dd-a50fdadcb392/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1"),
		// 			Location: to.Ptr("WestUS"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armstoragesync.ServiceProperties{
		// 				IncomingTrafficPolicy: to.Ptr(armstoragesync.IncomingTrafficPolicyAllowAllTraffic),
		// 				StorageSyncServiceStatus: to.Ptr[int32](0),
		// 				StorageSyncServiceUID: to.Ptr("\"3d1bf292-0f2a-4cc1-a3e1-60f35800e40c\""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SampleStorageSyncService_2"),
		// 			Type: to.Ptr("Microsoft.StorageSync/storageSyncServices"),
		// 			ID: to.Ptr("/subscriptions/3a048283-338f-4002-a9dd-a50fdadcb392/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_2"),
		// 			Location: to.Ptr("WestUS"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armstoragesync.ServiceProperties{
		// 				IncomingTrafficPolicy: to.Ptr(armstoragesync.IncomingTrafficPolicyAllowAllTraffic),
		// 				StorageSyncServiceStatus: to.Ptr[int32](0),
		// 				StorageSyncServiceUID: to.Ptr("\"2de01144-72da-4d7f-9d0c-e858855114a8\""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SampleStorageSyncService_3"),
		// 			Type: to.Ptr("Microsoft.StorageSync/storageSyncServices"),
		// 			ID: to.Ptr("/subscriptions/3a048283-338f-4002-a9dd-a50fdadcb392/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_3"),
		// 			Location: to.Ptr("WestUS"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armstoragesync.ServiceProperties{
		// 				IncomingTrafficPolicy: to.Ptr(armstoragesync.IncomingTrafficPolicyAllowAllTraffic),
		// 				StorageSyncServiceStatus: to.Ptr[int32](0),
		// 				StorageSyncServiceUID: to.Ptr("\"b2c58ee5-933e-462c-8a9e-b30f2bdd8fa3\""),
		// 			},
		// 	}},
		// }
	}
}
