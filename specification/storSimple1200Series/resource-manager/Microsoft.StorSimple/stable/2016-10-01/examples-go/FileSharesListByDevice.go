package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/FileSharesListByDevice.json
func ExampleFileSharesClient_NewListByDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSharesClient().NewListByDevicePager("HSDK-4XY4FI2IVG", "ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
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
		// page.FileShareList = armstorsimple1200series.FileShareList{
		// 	Value: []*armstorsimple1200series.FileShare{
		// 		{
		// 			Name: to.Ptr("Auto-TestFileShare2"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/fileServers/shares"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/hsdk-4xy4fi2ivg/fileServers/HSDK-4XY4FI2IVG/shares/Auto-TestFileShare2"),
		// 			Properties: &armstorsimple1200series.FileShareProperties{
		// 				Description: to.Ptr("Demo FileShare for SDK Test Local"),
		// 				AdminUser: to.Ptr("fareast\\idcdlslb"),
		// 				DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyLocal),
		// 				LocalUsedCapacityInBytes: to.Ptr[int64](0),
		// 				MonitoringStatus: to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
		// 				ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
		// 				ShareStatus: to.Ptr(armstorsimple1200series.ShareStatusOnline),
		// 				UsedCapacityInBytes: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Auto-TestFileShare1"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/fileServers/shares"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/hsdk-4xy4fi2ivg/fileServers/HSDK-4XY4FI2IVG/shares/Auto-TestFileShare1"),
		// 			Properties: &armstorsimple1200series.FileShareProperties{
		// 				Description: to.Ptr("Updated: Demo FileShare for SDK Test Tiered"),
		// 				AdminUser: to.Ptr("fareast\\idcdlslb"),
		// 				DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
		// 				LocalUsedCapacityInBytes: to.Ptr[int64](0),
		// 				MonitoringStatus: to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
		// 				ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
		// 				ShareStatus: to.Ptr(armstorsimple1200series.ShareStatusOffline),
		// 				UsedCapacityInBytes: to.Ptr[int64](0),
		// 			},
		// 	}},
		// }
	}
}
