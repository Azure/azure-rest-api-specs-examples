package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DevicesGet.json
func ExampleDevicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDevicesClient().Get(ctx, "Device001ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", &armstorsimple8000series.DevicesClientGetOptions{Expand: to.Ptr("details")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Device = armstorsimple8000series.Device{
	// 	Name: to.Ptr("Device001ForSDKTest"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device001ForSDKTest"),
	// 	Kind: to.Ptr("Series8000"),
	// 	Properties: &armstorsimple8000series.DeviceProperties{
	// 		ActivationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-22T09:36:20.963Z"); return t}()),
	// 		ActiveController: to.Ptr(armstorsimple8000series.ControllerIDController0),
	// 		AgentGroupVersion: to.Ptr[int32](7),
	// 		AvailableLocalStorageInBytes: to.Ptr[int64](43980464128),
	// 		AvailableTieredStorageInBytes: to.Ptr[int64](1099511627776),
	// 		Culture: to.Ptr("en-US"),
	// 		DeviceConfigurationStatus: to.Ptr(armstorsimple8000series.DeviceConfigurationStatusComplete),
	// 		DeviceDescription: to.Ptr("updated device description"),
	// 		DeviceSoftwareVersion: to.Ptr("6.3.9600.17802"),
	// 		DeviceType: to.Ptr(armstorsimple8000series.DeviceTypeSeries8000PhysicalAppliance),
	// 		FriendlyName: to.Ptr("Device001ForSDKTest"),
	// 		FriendlySoftwareVersion: to.Ptr("StorSimple 8000 Series Update 4.0"),
	// 		ModelDescription: to.Ptr("100"),
	// 		NetworkInterfaceCardCount: to.Ptr[int32](6),
	// 		ProvisionedLocalStorageInBytes: to.Ptr[int64](0),
	// 		ProvisionedTieredStorageInBytes: to.Ptr[int64](0),
	// 		ProvisionedVolumeSizeInBytes: to.Ptr[int64](0),
	// 		SerialNumber: to.Ptr("123456789"),
	// 		Status: to.Ptr(armstorsimple8000series.DeviceStatusOnline),
	// 		TargetIqn: to.Ptr("iqn.1991-05.com.microsoft:storsimple100-123456789-target"),
	// 		TotalTieredStorageInBytes: to.Ptr[int64](1099511627776),
	// 		UsingStorageInBytes: to.Ptr[int64](0),
	// 		Details: &armstorsimple8000series.DeviceDetails{
	// 			EndpointCount: to.Ptr[int32](0),
	// 			VolumeContainerCount: to.Ptr[int32](0),
	// 		},
	// 	},
	// }
}
