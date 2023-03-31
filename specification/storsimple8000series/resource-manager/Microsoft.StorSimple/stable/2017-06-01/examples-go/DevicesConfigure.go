package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DevicesConfigure.json
func ExampleDevicesClient_BeginConfigure() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDevicesClient().BeginConfigure(ctx, "ResourceGroupForSDKTest", "ManagerForSDKTest1", armstorsimple8000series.ConfigureDeviceRequest{
		Properties: &armstorsimple8000series.ConfigureDeviceRequestProperties{
			CurrentDeviceName: to.Ptr("Device001ForSDKTest"),
			FriendlyName:      to.Ptr("Device001ForSDKTest"),
			NetworkInterfaceData0Settings: &armstorsimple8000series.NetworkInterfaceData0Settings{
				ControllerOneIP:  to.Ptr("10.168.220.228"),
				ControllerZeroIP: to.Ptr("10.168.220.227"),
			},
			TimeZone: to.Ptr("Pacific Standard Time"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
