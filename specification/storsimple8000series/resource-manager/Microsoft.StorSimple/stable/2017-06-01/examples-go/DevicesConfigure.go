package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DevicesConfigure.json
func ExampleDevicesClient_BeginConfigure() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple8000series.NewDevicesClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginConfigure(ctx,
		"ResourceGroupForSDKTest",
		"ManagerForSDKTest1",
		armstorsimple8000series.ConfigureDeviceRequest{
			Properties: &armstorsimple8000series.ConfigureDeviceRequestProperties{
				CurrentDeviceName: to.Ptr("Device001ForSDKTest"),
				FriendlyName:      to.Ptr("Device001ForSDKTest"),
				NetworkInterfaceData0Settings: &armstorsimple8000series.NetworkInterfaceData0Settings{
					ControllerOneIP:  to.Ptr("10.168.220.228"),
					ControllerZeroIP: to.Ptr("10.168.220.227"),
				},
				TimeZone: to.Ptr("Pacific Standard Time"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
