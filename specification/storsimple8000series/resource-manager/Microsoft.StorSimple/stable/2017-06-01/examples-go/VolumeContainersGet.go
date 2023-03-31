package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/VolumeContainersGet.json
func ExampleVolumeContainersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVolumeContainersClient().Get(ctx, "Device05ForSDKTest", "VolumeContainerForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VolumeContainer = armstorsimple8000series.VolumeContainer{
	// 	Name: to.Ptr("VolumeContainerForSDKTest"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/volumeContainers"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/VolumeContainerForSDKTest"),
	// 	Kind: to.Ptr("Series8000"),
	// 	Properties: &armstorsimple8000series.VolumeContainerProperties{
	// 		BandwidthSettingID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/bandwidthSettings/bandwidthSetting1"),
	// 		EncryptionStatus: to.Ptr(armstorsimple8000series.EncryptionStatusEnabled),
	// 		OwnerShipStatus: to.Ptr(armstorsimple8000series.OwnerShipStatusOwned),
	// 		StorageAccountCredentialID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/storageAccountCredentials/safortestrecording"),
	// 		TotalCloudStorageUsageInBytes: to.Ptr[int64](0),
	// 		VolumeCount: to.Ptr[int32](1),
	// 	},
	// }
}
