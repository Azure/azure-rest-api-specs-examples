package armhybriddatamanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/DataStoreTypes_ListByDataManager-GET-example-171.json
func ExampleDataStoreTypesClient_NewListByDataManagerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybriddatamanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataStoreTypesClient().NewListByDataManagerPager("ResourceGroupForSDKTest", "TestAzureSDKOperations", nil)
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
		// page.DataStoreTypeList = armhybriddatamanager.DataStoreTypeList{
		// 	Value: []*armhybriddatamanager.DataStoreType{
		// 		{
		// 			Name: to.Ptr("StorSimple8000Series"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers/dataStoreTypes"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/StorSimple8000Series"),
		// 			Properties: &armhybriddatamanager.DataStoreTypeProperties{
		// 				RepositoryType: to.Ptr("Microsoft.StorSimple/managers"),
		// 				State: to.Ptr(armhybriddatamanager.StateEnabled),
		// 				SupportedDataServicesAsSink: []*string{
		// 				},
		// 				SupportedDataServicesAsSource: []*string{
		// 					to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataServices/DataTransformation"),
		// 					to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataServices/IndexingAndSearch")},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("AzureStorageAccount"),
		// 				Type: to.Ptr("Microsoft.HybridData/dataManagers/dataStoreTypes"),
		// 				ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/AzureStorageAccount"),
		// 				Properties: &armhybriddatamanager.DataStoreTypeProperties{
		// 					RepositoryType: to.Ptr("Microsoft.Storage/storageAccounts"),
		// 					State: to.Ptr(armhybriddatamanager.StateEnabled),
		// 					SupportedDataServicesAsSink: []*string{
		// 						to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataServices/DataTransformation")},
		// 						SupportedDataServicesAsSource: []*string{
		// 							to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataServices/IndexingAndSearch")},
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("AzureMediaServicesAccount"),
		// 						Type: to.Ptr("Microsoft.HybridData/dataManagers/dataStoreTypes"),
		// 						ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/AzureMediaServicesAccount"),
		// 						Properties: &armhybriddatamanager.DataStoreTypeProperties{
		// 							RepositoryType: to.Ptr("Microsoft.Media/mediaservices"),
		// 							State: to.Ptr(armhybriddatamanager.StateEnabled),
		// 							SupportedDataServicesAsSink: []*string{
		// 								to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataServices/DataTransformation")},
		// 								SupportedDataServicesAsSource: []*string{
		// 								},
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("RunnerDataSource"),
		// 							Type: to.Ptr("Microsoft.HybridData/dataManagers/dataStoreTypes"),
		// 							ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/RunnerDataSource"),
		// 							Properties: &armhybriddatamanager.DataStoreTypeProperties{
		// 								State: to.Ptr(armhybriddatamanager.StateEnabled),
		// 								SupportedDataServicesAsSink: []*string{
		// 								},
		// 								SupportedDataServicesAsSource: []*string{
		// 								},
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("RunnerDataSink"),
		// 							Type: to.Ptr("Microsoft.HybridData/dataManagers/dataStoreTypes"),
		// 							ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/RunnerDataSink"),
		// 							Properties: &armhybriddatamanager.DataStoreTypeProperties{
		// 								State: to.Ptr(armhybriddatamanager.StateEnabled),
		// 								SupportedDataServicesAsSink: []*string{
		// 								},
		// 								SupportedDataServicesAsSource: []*string{
		// 								},
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("SqlServer"),
		// 							Type: to.Ptr("Microsoft.HybridData/dataManagers/dataStoreTypes"),
		// 							ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/SqlServer"),
		// 							Properties: &armhybriddatamanager.DataStoreTypeProperties{
		// 								RepositoryType: to.Ptr("Microsoft.Sql/servers"),
		// 								State: to.Ptr(armhybriddatamanager.StateEnabled),
		// 								SupportedDataServicesAsSink: []*string{
		// 									to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataServices/IndexingAndSearch")},
		// 									SupportedDataServicesAsSource: []*string{
		// 									},
		// 								},
		// 						}},
		// 					}
	}
}
