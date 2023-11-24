package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/DevicesGet.json
func ExampleDevicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDevicesClient().Get(ctx, "HSDK-ARCSX4MVKZ", "ResourceGroupForSDKTest", "hAzureSDKOperations", &armstorsimple1200series.DevicesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Device = armstorsimple1200series.Device{
	// 	Name: to.Ptr("HSDK-ARCSX4MVKZ"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices"),
	// 	ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ"),
	// 	Properties: &armstorsimple1200series.DeviceProperties{
	// 		Type: to.Ptr(armstorsimple1200series.DeviceTypeSeries9000OnPremVirtualAppliance),
	// 		ActivationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T10:59:26.165Z"); return t}()),
	// 		AllowedDeviceOperations: []*armstorsimple1200series.DeviceOperation{
	// 			to.Ptr(armstorsimple1200series.DeviceOperationDeactivate),
	// 			to.Ptr(armstorsimple1200series.DeviceOperationBrowsable)},
	// 			Culture: to.Ptr("en-US"),
	// 			DeviceCapabilities: []*armstorsimple1200series.SupportedDeviceCapabilities{
	// 				to.Ptr(armstorsimple1200series.SupportedDeviceCapabilitiesFileServer)},
	// 				DeviceConfigurationStatus: to.Ptr(armstorsimple1200series.DeviceConfigurationStatusComplete),
	// 				DeviceSoftwareVersion: to.Ptr("10.0.10296.0"),
	// 				DomainName: to.Ptr("fareast.corp.microsoft.com"),
	// 				ModelDescription: to.Ptr("1200"),
	// 				Status: to.Ptr(armstorsimple1200series.DeviceStatusOnline),
	// 			},
	// 		}
}
