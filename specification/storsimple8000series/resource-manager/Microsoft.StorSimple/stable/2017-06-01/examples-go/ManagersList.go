package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/ManagersList.json
func ExampleManagersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagersClient().NewListPager(nil)
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
		// page.ManagerList = armstorsimple8000series.ManagerList{
		// 	Value: []*armstorsimple8000series.Manager{
		// 		{
		// 			Name: to.Ptr("Manager1"),
		// 			Type: to.Ptr("Microsoft.StorSimple/Managers"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroup1/providers/Microsoft.StorSimple/Managers/ManagerForSDKTest"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"datetime'2017-06-05T09%3A36%3A00.8822928Z'\""),
		// 			Properties: &armstorsimple8000series.ManagerProperties{
		// 				CisIntrinsicSettings: &armstorsimple8000series.ManagerIntrinsicSettings{
		// 					Type: to.Ptr(armstorsimple8000series.ManagerTypeGardaV1),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SKU: &armstorsimple8000series.ManagerSKU{
		// 					Name: to.Ptr("Standard"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ManagerForSDKTest"),
		// 			Type: to.Ptr("Microsoft.StorSimple/Managers"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/Managers/ManagerForSDKTest"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"datetime'2017-06-05T09%3A36%3A00.8822928Z'\""),
		// 			Properties: &armstorsimple8000series.ManagerProperties{
		// 				CisIntrinsicSettings: &armstorsimple8000series.ManagerIntrinsicSettings{
		// 					Type: to.Ptr(armstorsimple8000series.ManagerTypeGardaV1),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SKU: &armstorsimple8000series.ManagerSKU{
		// 					Name: to.Ptr("Standard"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ManagerForSDKTest1"),
		// 			Type: to.Ptr("Microsoft.StorSimple/Managers"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/Managers/ManagerForSDKTest1"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"datetime'2017-06-14T11%3A27%3A23.8667962Z'\""),
		// 			Properties: &armstorsimple8000series.ManagerProperties{
		// 				CisIntrinsicSettings: &armstorsimple8000series.ManagerIntrinsicSettings{
		// 					Type: to.Ptr(armstorsimple8000series.ManagerTypeGardaV1),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SKU: &armstorsimple8000series.ManagerSKU{
		// 					Name: to.Ptr("Standard"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ManagerForSDKTest2"),
		// 			Type: to.Ptr("Microsoft.StorSimple/Managers"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/Managers/ManagerForSDKTest2"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"TagName": to.Ptr("ForSDKTest"),
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2017-06-18T16%3A13%3A43.8476579Z'\""),
		// 			Properties: &armstorsimple8000series.ManagerProperties{
		// 				CisIntrinsicSettings: &armstorsimple8000series.ManagerIntrinsicSettings{
		// 					Type: to.Ptr(armstorsimple8000series.ManagerTypeGardaV1),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SKU: &armstorsimple8000series.ManagerSKU{
		// 					Name: to.Ptr("Standard"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
