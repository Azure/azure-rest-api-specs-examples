package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/BackupsListByDevice.json
func ExampleBackupsClient_NewListByDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupsClient().NewListByDevicePager("HSDK-4XY4FI2IVG", "ResourceGroupForSDKTest", "hAzureSDKOperations", &armstorsimple1200series.BackupsClientListByDeviceOptions{ForFailover: nil,
		Filter: to.Ptr("initiatedBy%20eq%20'Manual'"),
	})
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
		// page.BackupList = armstorsimple1200series.BackupList{
		// 	Value: []*armstorsimple1200series.Backup{
		// 		{
		// 			Name: to.Ptr("58d33025-e837-41cc-b15f-7c85ced64aab"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/backups"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/hsdk-4xy4fi2ivg/backups/58d33025-e837-41cc-b15f-7c85ced64aab"),
		// 			Properties: &armstorsimple1200series.BackupProperties{
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-09T16:19:09.653Z"); return t}()),
		// 				DeviceID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/hsdk-4xy4fi2ivg"),
		// 				Elements: []*armstorsimple1200series.BackupElement{
		// 					{
		// 						Name: to.Ptr("2304968f-91af-4f59-8b79-31e321eee40e"),
		// 						Type: to.Ptr("Microsoft.StorSimple/managers/devices/backups/elements"),
		// 						ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/hsdk-4xy4fi2ivg/backups/58d33025-e837-41cc-b15f-7c85ced64aab/elements/2304968f-91af-4f59-8b79-31e321eee40e"),
		// 						Properties: &armstorsimple1200series.BackupElementProperties{
		// 							DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyLocal),
		// 							EndpointName: to.Ptr("Auto-TestFileShare2"),
		// 							SizeInBytes: to.Ptr[int64](536870912000),
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("c5661246-17b7-4daf-a82a-6cc86c68a1dc"),
		// 						Type: to.Ptr("Microsoft.StorSimple/managers/devices/backups/elements"),
		// 						ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/hsdk-4xy4fi2ivg/backups/58d33025-e837-41cc-b15f-7c85ced64aab/elements/c5661246-17b7-4daf-a82a-6cc86c68a1dc"),
		// 						Properties: &armstorsimple1200series.BackupElementProperties{
		// 							DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
		// 							EndpointName: to.Ptr("Auto-TestFileShare1"),
		// 							SizeInBytes: to.Ptr[int64](536870912000),
		// 						},
		// 				}},
		// 				ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "9999-12-31T23:59:59.999Z"); return t}()),
		// 				InitiatedBy: to.Ptr(armstorsimple1200series.InitiatedByManual),
		// 				SizeInBytes: to.Ptr[int64](1073741824000),
		// 				TargetID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/hsdk-4xy4fi2ivg/fileServers/HSDK-4XY4FI2IVG"),
		// 				TargetType: to.Ptr("Microsoft.StorSimple/managers/devices/fileServers"),
		// 			},
		// 	}},
		// }
	}
}
