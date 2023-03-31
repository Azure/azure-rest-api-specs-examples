package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/VolumeContainersListByDevice.json
func ExampleVolumeContainersClient_NewListByDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVolumeContainersClient().NewListByDevicePager("Device05ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
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
		// page.VolumeContainerList = armstorsimple8000series.VolumeContainerList{
		// 	Value: []*armstorsimple8000series.VolumeContainer{
		// 		{
		// 			Name: to.Ptr("vcforsdktest"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/volumeContainers"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/vcforsdktest"),
		// 			Kind: to.Ptr("Series8000"),
		// 			Properties: &armstorsimple8000series.VolumeContainerProperties{
		// 				BandWidthRateInMbps: to.Ptr[int32](0),
		// 				EncryptionStatus: to.Ptr(armstorsimple8000series.EncryptionStatusDisabled),
		// 				OwnerShipStatus: to.Ptr(armstorsimple8000series.OwnerShipStatusOwned),
		// 				StorageAccountCredentialID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/storageAccountCredentials/safortestrecording"),
		// 				TotalCloudStorageUsageInBytes: to.Ptr[int64](33839216),
		// 				VolumeCount: to.Ptr[int32](1),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("vcForOdataFilterTest"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/volumeContainers"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/vcForOdataFilterTest"),
		// 			Kind: to.Ptr("Series8000"),
		// 			Properties: &armstorsimple8000series.VolumeContainerProperties{
		// 				BandWidthRateInMbps: to.Ptr[int32](0),
		// 				EncryptionStatus: to.Ptr(armstorsimple8000series.EncryptionStatusDisabled),
		// 				OwnerShipStatus: to.Ptr(armstorsimple8000series.OwnerShipStatusOwned),
		// 				StorageAccountCredentialID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/storageAccountCredentials/safortestrecording"),
		// 				TotalCloudStorageUsageInBytes: to.Ptr[int64](200277),
		// 				VolumeCount: to.Ptr[int32](4),
		// 			},
		// 	}},
		// }
	}
}
