package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DeviceSettingsCreateOrUpdateAlertSettings.json
func ExampleDeviceSettingsClient_BeginCreateOrUpdateAlertSettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple8000series.NewDeviceSettingsClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdateAlertSettings(ctx,
		"Device05ForSDKTest",
		"ResourceGroupForSDKTest",
		"ManagerForSDKTest1",
		armstorsimple8000series.AlertSettings{
			Properties: &armstorsimple8000series.AlertNotificationProperties{
				AdditionalRecipientEmailList: []*string{},
				AlertNotificationCulture:     to.Ptr("en-US"),
				EmailNotification:            to.Ptr(armstorsimple8000series.AlertEmailNotificationStatusEnabled),
				NotificationToServiceOwners:  to.Ptr(armstorsimple8000series.AlertEmailNotificationStatusEnabled),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
