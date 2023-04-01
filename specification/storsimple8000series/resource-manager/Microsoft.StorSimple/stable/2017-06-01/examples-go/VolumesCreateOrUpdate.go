package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/VolumesCreateOrUpdate.json
func ExampleVolumesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumesClient().BeginCreateOrUpdate(ctx, "Device05ForSDKTest", "VolumeContainerForSDKTest", "Volume1ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", armstorsimple8000series.Volume{
		Properties: &armstorsimple8000series.VolumeProperties{
			AccessControlRecordIDs: []*string{
				to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/accessControlRecords/ACR2")},
			MonitoringStatus: to.Ptr(armstorsimple8000series.MonitoringStatusEnabled),
			SizeInBytes:      to.Ptr[int64](5368709120),
			VolumeStatus:     to.Ptr(armstorsimple8000series.VolumeStatusOffline),
			VolumeType:       to.Ptr(armstorsimple8000series.VolumeTypeTiered),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Volume = armstorsimple8000series.Volume{
	// 	Name: to.Ptr("Volume1ForSDKTest"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/volumeContainers/volumes"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/VolumeContainerForSDKTest/volumes/Volume1ForSDKTest"),
	// 	Kind: to.Ptr("Series8000"),
	// 	Properties: &armstorsimple8000series.VolumeProperties{
	// 		AccessControlRecordIDs: []*string{
	// 			to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/accessControlRecords/ACR2")},
	// 			BackupStatus: to.Ptr(armstorsimple8000series.BackupStatusDisabled),
	// 			MonitoringStatus: to.Ptr(armstorsimple8000series.MonitoringStatusEnabled),
	// 			OperationStatus: to.Ptr(armstorsimple8000series.OperationStatusNone),
	// 			SizeInBytes: to.Ptr[int64](5368709120),
	// 			VolumeContainerID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/VolumeContainerForSDKTest"),
	// 			VolumeStatus: to.Ptr(armstorsimple8000series.VolumeStatusOffline),
	// 			VolumeType: to.Ptr(armstorsimple8000series.VolumeTypeTiered),
	// 		},
	// 	}
}
