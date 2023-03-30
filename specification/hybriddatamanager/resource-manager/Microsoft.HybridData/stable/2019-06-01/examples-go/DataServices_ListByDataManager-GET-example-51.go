package armhybriddatamanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/DataServices_ListByDataManager-GET-example-51.json
func ExampleDataServicesClient_NewListByDataManagerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybriddatamanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataServicesClient().NewListByDataManagerPager("ResourceGroupForSDKTest", "TestAzureSDKOperations", nil)
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
		// page.DataServiceList = armhybriddatamanager.DataServiceList{
		// 	Value: []*armhybriddatamanager.DataService{
		// 		{
		// 			Name: to.Ptr("DataTransformation"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers/dataServices"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataServices/DataTransformation"),
		// 			Properties: &armhybriddatamanager.DataServiceProperties{
		// 				State: to.Ptr(armhybriddatamanager.StateEnabled),
		// 				SupportedDataSinkTypes: []*string{
		// 					to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/AzureStorageAccount"),
		// 					to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/AzureMediaServicesAccount")},
		// 					SupportedDataSourceTypes: []*string{
		// 						to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/StorSimple8000Series")},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("KeyRollover"),
		// 					Type: to.Ptr("Microsoft.HybridData/dataManagers/dataServices"),
		// 					ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataServices/KeyRollover"),
		// 					Properties: &armhybriddatamanager.DataServiceProperties{
		// 						State: to.Ptr(armhybriddatamanager.StateSupported),
		// 						SupportedDataSinkTypes: []*string{
		// 						},
		// 						SupportedDataSourceTypes: []*string{
		// 						},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("DataSecurity"),
		// 					Type: to.Ptr("Microsoft.HybridData/dataManagers/dataServices"),
		// 					ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataServices/DataSecurity"),
		// 					Properties: &armhybriddatamanager.DataServiceProperties{
		// 						State: to.Ptr(armhybriddatamanager.StateSupported),
		// 						SupportedDataSinkTypes: []*string{
		// 						},
		// 						SupportedDataSourceTypes: []*string{
		// 						},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("DataInsights"),
		// 					Type: to.Ptr("Microsoft.HybridData/dataManagers/dataServices"),
		// 					ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataServices/DataInsights"),
		// 					Properties: &armhybriddatamanager.DataServiceProperties{
		// 						State: to.Ptr(armhybriddatamanager.StateSupported),
		// 						SupportedDataSinkTypes: []*string{
		// 						},
		// 						SupportedDataSourceTypes: []*string{
		// 						},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("DataHealth"),
		// 					Type: to.Ptr("Microsoft.HybridData/dataManagers/dataServices"),
		// 					ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataServices/DataHealth"),
		// 					Properties: &armhybriddatamanager.DataServiceProperties{
		// 						State: to.Ptr(armhybriddatamanager.StateSupported),
		// 						SupportedDataSinkTypes: []*string{
		// 						},
		// 						SupportedDataSourceTypes: []*string{
		// 						},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("RunnerService"),
		// 					Type: to.Ptr("Microsoft.HybridData/dataManagers/dataServices"),
		// 					ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataServices/RunnerService"),
		// 					Properties: &armhybriddatamanager.DataServiceProperties{
		// 						State: to.Ptr(armhybriddatamanager.StateSupported),
		// 						SupportedDataSinkTypes: []*string{
		// 						},
		// 						SupportedDataSourceTypes: []*string{
		// 						},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("IndexingAndSearch"),
		// 					Type: to.Ptr("Microsoft.HybridData/dataManagers/dataServices"),
		// 					ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataServices/IndexingAndSearch"),
		// 					Properties: &armhybriddatamanager.DataServiceProperties{
		// 						State: to.Ptr(armhybriddatamanager.StateEnabled),
		// 						SupportedDataSinkTypes: []*string{
		// 							to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/SqlServer")},
		// 							SupportedDataSourceTypes: []*string{
		// 								to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/StorSimple8000Series"),
		// 								to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/AzureStorageAccount")},
		// 							},
		// 					}},
		// 				}
	}
}
