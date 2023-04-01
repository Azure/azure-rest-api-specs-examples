package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/VolumesListByVolumeContainer.json
func ExampleVolumesClient_NewListByVolumeContainerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVolumesClient().NewListByVolumeContainerPager("Device05ForSDKTest", "volumeContainerForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
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
		// page.VolumeList = armstorsimple8000series.VolumeList{
		// 	Value: []*armstorsimple8000series.Volume{
		// 		{
		// 			Name: to.Ptr("Clonedvolume1"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/volumeContainers/volumes"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest/volumes/Clonedvolume1"),
		// 			Kind: to.Ptr("Series8000"),
		// 			Properties: &armstorsimple8000series.VolumeProperties{
		// 				AccessControlRecordIDs: []*string{
		// 					to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/accessControlRecords/ACR2")},
		// 					BackupStatus: to.Ptr(armstorsimple8000series.BackupStatusDisabled),
		// 					MonitoringStatus: to.Ptr(armstorsimple8000series.MonitoringStatusDisabled),
		// 					OperationStatus: to.Ptr(armstorsimple8000series.OperationStatusNone),
		// 					SizeInBytes: to.Ptr[int64](10737418240),
		// 					VolumeContainerID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest"),
		// 					VolumeStatus: to.Ptr(armstorsimple8000series.VolumeStatusOnline),
		// 					VolumeType: to.Ptr(armstorsimple8000series.VolumeTypeTiered),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("volume1"),
		// 				Type: to.Ptr("Microsoft.StorSimple/managers/devices/volumeContainers/volumes"),
		// 				ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest/volumes/volume1"),
		// 				Kind: to.Ptr("Series8000"),
		// 				Properties: &armstorsimple8000series.VolumeProperties{
		// 					AccessControlRecordIDs: []*string{
		// 						to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/accessControlRecords/ACR2")},
		// 						BackupPolicyIDs: []*string{
		// 							to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backupPolicies/BkUpPolicyForSDKTest1032280949")},
		// 							BackupStatus: to.Ptr(armstorsimple8000series.BackupStatusDisabled),
		// 							MonitoringStatus: to.Ptr(armstorsimple8000series.MonitoringStatusEnabled),
		// 							OperationStatus: to.Ptr(armstorsimple8000series.OperationStatusNone),
		// 							SizeInBytes: to.Ptr[int64](10737418240),
		// 							VolumeContainerID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest"),
		// 							VolumeStatus: to.Ptr(armstorsimple8000series.VolumeStatusOnline),
		// 							VolumeType: to.Ptr(armstorsimple8000series.VolumeTypeTiered),
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("volume2"),
		// 						Type: to.Ptr("Microsoft.StorSimple/managers/devices/volumeContainers/volumes"),
		// 						ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest/volumes/volume2"),
		// 						Kind: to.Ptr("Series8000"),
		// 						Properties: &armstorsimple8000series.VolumeProperties{
		// 							AccessControlRecordIDs: []*string{
		// 								to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/accessControlRecords/ACR1"),
		// 								to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/accessControlRecords/ACR2")},
		// 								BackupPolicyIDs: []*string{
		// 									to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backupPolicies/BkUpPolicyForSDKTest1032280949")},
		// 									BackupStatus: to.Ptr(armstorsimple8000series.BackupStatusDisabled),
		// 									MonitoringStatus: to.Ptr(armstorsimple8000series.MonitoringStatusEnabled),
		// 									OperationStatus: to.Ptr(armstorsimple8000series.OperationStatusNone),
		// 									SizeInBytes: to.Ptr[int64](16106127360),
		// 									VolumeContainerID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest"),
		// 									VolumeStatus: to.Ptr(armstorsimple8000series.VolumeStatusOnline),
		// 									VolumeType: to.Ptr(armstorsimple8000series.VolumeTypeTiered),
		// 								},
		// 						}},
		// 					}
	}
}
