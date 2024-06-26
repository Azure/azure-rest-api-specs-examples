package armstoragepool_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagepool/armstoragepool"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragepool.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationListResult = armstoragepool.OperationListResult{
		// 	Value: []*armstoragepool.RPOperation{
		// 		{
		// 			Name: to.Ptr("Microsoft.StoragePool/diskPools/read"),
		// 			Display: &armstoragepool.OperationDisplay{
		// 				Description: to.Ptr("Read Disk Pool"),
		// 				Operation: to.Ptr("Read Microsoft.StoragePool/diskPools"),
		// 				Provider: to.Ptr("Microsoft.StoragePool"),
		// 				Resource: to.Ptr("Disk Pool"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StoragePool/diskPools/write"),
		// 			Display: &armstoragepool.OperationDisplay{
		// 				Description: to.Ptr("Create or Update Disk Pool"),
		// 				Operation: to.Ptr("Create or Update Microsoft.StoragePool/diskPools"),
		// 				Provider: to.Ptr("Microsoft.StoragePool"),
		// 				Resource: to.Ptr("Disk Pool"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StoragePool/diskPools/delete"),
		// 			Display: &armstoragepool.OperationDisplay{
		// 				Description: to.Ptr("Delete Disk Pool"),
		// 				Operation: to.Ptr("Delete Microsoft.StoragePool/diskPools"),
		// 				Provider: to.Ptr("Microsoft.StoragePool"),
		// 				Resource: to.Ptr("Microsoft.StoragePool/diskPools"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StoragePool/diskPools/start/action"),
		// 			Display: &armstoragepool.OperationDisplay{
		// 				Description: to.Ptr("Start Disk Pool"),
		// 				Operation: to.Ptr("Start Microsoft.StoragePool/diskPools"),
		// 				Provider: to.Ptr("Microsoft.StoragePool"),
		// 				Resource: to.Ptr("Microsoft.StoragePool/diskPools"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StoragePool/diskPools/deallocate/action"),
		// 			Display: &armstoragepool.OperationDisplay{
		// 				Description: to.Ptr("Deallocate Disk Pool"),
		// 				Operation: to.Ptr("Deallocate Microsoft.StoragePool/diskPools"),
		// 				Provider: to.Ptr("Microsoft.StoragePool"),
		// 				Resource: to.Ptr("Microsoft.StoragePool/diskPools"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StoragePool/diskPools/iscsiTargets/read"),
		// 			Display: &armstoragepool.OperationDisplay{
		// 				Description: to.Ptr("Read iSCSI targets"),
		// 				Operation: to.Ptr("Read Microsoft.StoragePool/diskPools/iscsiTargets"),
		// 				Provider: to.Ptr("Microsoft.StoragePool"),
		// 				Resource: to.Ptr("iSCSI targets"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StoragePool/diskPools/iscsiTargets/write"),
		// 			Display: &armstoragepool.OperationDisplay{
		// 				Description: to.Ptr("Create or Update iSCSI targets"),
		// 				Operation: to.Ptr("Create or Update Microsoft.StoragePool/diskPools/iscsiTargets"),
		// 				Provider: to.Ptr("Microsoft.StoragePool"),
		// 				Resource: to.Ptr("iSCSI targets"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StoragePool/diskPools/iscsiTargets/delete"),
		// 			Display: &armstoragepool.OperationDisplay{
		// 				Description: to.Ptr("Delete iSCSI targets"),
		// 				Operation: to.Ptr("Delete Microsoft.StoragePool/diskPools/iscsiTargets"),
		// 				Provider: to.Ptr("Microsoft.StoragePool"),
		// 				Resource: to.Ptr("iSCSI targets"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 	}},
		// }
	}
}
