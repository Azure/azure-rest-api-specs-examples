package armhybriddatamanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/DataStores_ListByDataManager-GET-example-151.json
func ExampleDataStoresClient_NewListByDataManagerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybriddatamanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataStoresClient().NewListByDataManagerPager("ResourceGroupForSDKTest", "TestAzureSDKOperations", &armhybriddatamanager.DataStoresClientListByDataManagerOptions{Filter: nil})
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
		// page.DataStoreList = armhybriddatamanager.DataStoreList{
		// 	Value: []*armhybriddatamanager.DataStore{
		// 		{
		// 			Name: to.Ptr("TestAzureStorage1"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers/dataStores"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStores/TestAzureStorage1"),
		// 			Properties: &armhybriddatamanager.DataStoreProperties{
		// 				DataStoreTypeID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/AzureStorageAccount"),
		// 				ExtendedProperties: map[string]any{
		// 					"StorageAccountNameForQueue": "dmsdatasink",
		// 				},
		// 				RepositoryID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.Storage/storageAccounts/dmsdatasink"),
		// 				State: to.Ptr(armhybriddatamanager.StateEnabled),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("TestStorSimpleSource1"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers/dataStores"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStores/TestStorSimpleSource1"),
		// 			Properties: &armhybriddatamanager.DataStoreProperties{
		// 				DataStoreTypeID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/StorSimple8000Series"),
		// 				ExtendedProperties: map[string]any{
		// 					"extendedSaKey": nil,
		// 					"resourceId": "/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ForDMS/providers/Microsoft.StorSimple/managers/BLR8600",
		// 				},
		// 				RepositoryID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ForDMS/providers/Microsoft.StorSimple/managers/BLR8600"),
		// 				State: to.Ptr(armhybriddatamanager.StateEnabled),
		// 			},
		// 	}},
		// }
	}
}
