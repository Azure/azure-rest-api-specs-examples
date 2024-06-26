package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/ManagersList.json
func ExampleManagersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.ManagerList = armstorsimple1200series.ManagerList{
		// 	Value: []*armstorsimple1200series.Manager{
		// 		{
		// 			Name: to.Ptr("AuTo1117160133591"),
		// 			Type: to.Ptr("Microsoft.StorSimple/Managers"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/GroupForArmAutomation/providers/Microsoft.StorSimple/Managers/AuTo1117160133591"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"datetime'2017-03-22T09%3A58%3A03.0020106Z'\""),
		// 			Properties: &armstorsimple1200series.ManagerProperties{
		// 				CisIntrinsicSettings: &armstorsimple1200series.ManagerIntrinsicSettings{
		// 					Type: to.Ptr(armstorsimple1200series.ManagerTypeHelsinkiV1),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SKU: &armstorsimple1200series.ManagerSKU{
		// 					Name: to.Ptr("Standard"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("GAuTosumanths-pc10227181824221"),
		// 			Type: to.Ptr("Microsoft.StorSimple/Managers"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/GroupForArmAutomation/providers/Microsoft.StorSimple/Managers/GAuTosumanths-pc10227181824221"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"datetime'2018-02-27T12%3A56%3A02.0580397Z'\""),
		// 			Properties: &armstorsimple1200series.ManagerProperties{
		// 				CisIntrinsicSettings: &armstorsimple1200series.ManagerIntrinsicSettings{
		// 					Type: to.Ptr(armstorsimple1200series.ManagerTypeGardaV1),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SKU: &armstorsimple1200series.ManagerSKU{
		// 					Name: to.Ptr("Standard"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
