package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/DevicesGetUpdateSummary.json
func ExampleDevicesClient_GetUpdateSummary() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDevicesClient().GetUpdateSummary(ctx, "HBVT-02X525X2W0", "ResourceGroupForSDKTest", "hManagerForSDKTest4", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Updates = armstorsimple1200series.Updates{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/updateSummary"),
	// 	ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hManagerForSDKTest4/devices/HBVT-02X525X2W0/updateSummary/default"),
	// 	Properties: &armstorsimple1200series.UpdatesProperties{
	// 		DeviceLastScannedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-02T05:46:21.681Z"); return t}()),
	// 		DeviceVersion: to.Ptr("10.0.10296.0"),
	// 		InProgressDownloadJobID: to.Ptr(""),
	// 		InProgressInstallJobID: to.Ptr(""),
	// 		LastCompletedScanTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-02T05:46:22.368Z"); return t}()),
	// 		RebootRequiredForInstall: to.Ptr(false),
	// 		RegularUpdatesAvailable: to.Ptr(false),
	// 		Status: to.Ptr(armstorsimple1200series.UpdateOperationIdle),
	// 		TotalItemsPendingForDownload: to.Ptr[int32](0),
	// 		TotalItemsPendingForInstall: to.Ptr[int32](0),
	// 	},
	// }
}
