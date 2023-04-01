package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DevicesListFailoverTargets.json
func ExampleDevicesClient_NewListFailoverTargetsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDevicesClient().NewListFailoverTargetsPager("Device05ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", armstorsimple8000series.ListFailoverTargetsRequest{
		VolumeContainers: []*string{
			to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/vcforsdktest")},
	}, nil)
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
		// page.FailoverTargetsList = armstorsimple8000series.FailoverTargetsList{
		// 	Value: []*armstorsimple8000series.FailoverTarget{
		// 		{
		// 			AvailableLocalStorageInBytes: to.Ptr[int64](43980464128),
		// 			AvailableTieredStorageInBytes: to.Ptr[int64](1099511627776),
		// 			DataContainersCount: to.Ptr[int32](0),
		// 			DeviceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/sugattdeviceforSDK"),
		// 			DeviceSoftwareVersion: to.Ptr("6.3.9600.17802"),
		// 			DeviceStatus: to.Ptr(armstorsimple8000series.DeviceStatusOffline),
		// 			EligibilityResult: &armstorsimple8000series.TargetEligibilityResult{
		// 				EligibilityStatus: to.Ptr(armstorsimple8000series.TargetEligibilityStatusNotEligible),
		// 				Messages: []*armstorsimple8000series.TargetEligibilityErrorMessage{
		// 					{
		// 						Message: to.Ptr("The selected target device is in 'Offline' state. The device configuration status is 'True'. Only online devices with the configuration status as complete can be failed over."),
		// 						Resolution: to.Ptr("Ensure that the specified target device is online and the device configuration is complete before you try the device failover."),
		// 						ResultCode: to.Ptr(armstorsimple8000series.TargetEligibilityResultCodeTargetIsNotOnlineError),
		// 				}},
		// 			},
		// 			FriendlyDeviceSoftwareVersion: to.Ptr("StorSimple 8000 Series Update 4.0"),
		// 			VolumesCount: to.Ptr[int32](0),
		// 		},
		// 		{
		// 			AvailableLocalStorageInBytes: to.Ptr[int64](43980464128),
		// 			AvailableTieredStorageInBytes: to.Ptr[int64](1099511627776),
		// 			DataContainersCount: to.Ptr[int32](0),
		// 			DeviceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/jemdeviceforsdk"),
		// 			DeviceSoftwareVersion: to.Ptr("6.3.9600.17802"),
		// 			DeviceStatus: to.Ptr(armstorsimple8000series.DeviceStatusOnline),
		// 			EligibilityResult: &armstorsimple8000series.TargetEligibilityResult{
		// 				EligibilityStatus: to.Ptr(armstorsimple8000series.TargetEligibilityStatusEligible),
		// 			},
		// 			FriendlyDeviceSoftwareVersion: to.Ptr("StorSimple 8000 Series Update 4.0"),
		// 			VolumesCount: to.Ptr[int32](0),
		// 		},
		// 		{
		// 			AvailableLocalStorageInBytes: to.Ptr[int64](40458590976),
		// 			AvailableTieredStorageInBytes: to.Ptr[int64](1011464798208),
		// 			DataContainersCount: to.Ptr[int32](2),
		// 			DeviceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			DeviceSoftwareVersion: to.Ptr("6.3.9600.17802"),
		// 			DeviceStatus: to.Ptr(armstorsimple8000series.DeviceStatusOnline),
		// 			EligibilityResult: &armstorsimple8000series.TargetEligibilityResult{
		// 				EligibilityStatus: to.Ptr(armstorsimple8000series.TargetEligibilityStatusNotEligible),
		// 				Messages: []*armstorsimple8000series.TargetEligibilityErrorMessage{
		// 					{
		// 						Message: to.Ptr("The specified target device is the same as the source device. This is not supported. "),
		// 						Resolution: to.Ptr("Select a different target device and then retry the failover. "),
		// 						ResultCode: to.Ptr(armstorsimple8000series.TargetEligibilityResultCodeTargetAndSourceCannotBeSameError),
		// 				}},
		// 			},
		// 			FriendlyDeviceSoftwareVersion: to.Ptr("StorSimple 8000 Series Update 4.0"),
		// 			VolumesCount: to.Ptr[int32](7),
		// 	}},
		// }
	}
}
