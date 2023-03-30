package armhybriddatamanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/DataManagers_ListByResourceGroup-GET-example-31.json
func ExampleDataManagersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybriddatamanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataManagersClient().NewListByResourceGroupPager("ResourceGroupForSDKTest", nil)
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
		// page.DataManagerList = armhybriddatamanager.DataManagerList{
		// 	Value: []*armhybriddatamanager.DataManager{
		// 		{
		// 			Name: to.Ptr("AzureSDKOperations"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/AzureSDKOperations"),
		// 			Location: to.Ptr("westus2"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-02-17T14%3A50%3A37.866739Z'\"_W/\"datetime'2019-02-17T14%3A50%3A38.038859Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("AzSDKOps"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/AzSDKOps"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 				"helL00000": to.Ptr("dlrow"),
		// 				"hello": to.Ptr("World"),
		// 				"new": to.Ptr("true"),
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-03-30T06%3A35%3A01.1816182Z'\"_W/\"datetime'2019-03-30T06%3A35%3A01.2846913Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("TestAzureSDKOperations"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2020-02-05T04%3A22%3A40.6354864Z'\"_W/\"datetime'2020-02-05T04%3A22%3A40.7912864Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("EcyTestDMSRes"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/EcyTestDMSRes"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-08-12T10%3A20%3A40.4679832Z'\"_W/\"datetime'2019-08-12T10%3A20%3A40.6030796Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("CcyTestDMSRes2"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/CcyTestDMSRes2"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-08-13T04%3A47%3A07.5063631Z'\"_W/\"datetime'2019-08-13T04%3A47%3A07.5113667Z'\""),
		// 	}},
		// }
	}
}
