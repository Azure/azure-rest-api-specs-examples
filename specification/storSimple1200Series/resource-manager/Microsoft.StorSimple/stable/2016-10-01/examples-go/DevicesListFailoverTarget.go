package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/DevicesListFailoverTarget.json
func ExampleDevicesClient_NewListFailoverTargetPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDevicesClient().NewListFailoverTargetPager("HSDK-4XY4FI2IVG", "ResourceGroupForSDKTest", "hAzureSDKOperations", &armstorsimple1200series.DevicesClientListFailoverTargetOptions{Expand: nil})
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
		// page.DeviceList = armstorsimple1200series.DeviceList{
		// 	Value: []*armstorsimple1200series.Device{
		// 		{
		// 			Name: to.Ptr("HSDK-DMNJB2PET0"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-DMNJB2PET0"),
		// 			Properties: &armstorsimple1200series.DeviceProperties{
		// 				Type: to.Ptr(armstorsimple1200series.DeviceTypeSeries9000OnPremVirtualAppliance),
		// 				ActivationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-11T14:42:31.060Z"); return t}()),
		// 				AllowedDeviceOperations: []*armstorsimple1200series.DeviceOperation{
		// 					to.Ptr(armstorsimple1200series.DeviceOperationDelete),
		// 					to.Ptr(armstorsimple1200series.DeviceOperationDRTarget),
		// 					to.Ptr(armstorsimple1200series.DeviceOperationBrowsable)},
		// 					Culture: to.Ptr("en-US"),
		// 					DeviceCapabilities: []*armstorsimple1200series.SupportedDeviceCapabilities{
		// 						to.Ptr(armstorsimple1200series.SupportedDeviceCapabilitiesFileServer)},
		// 						DeviceConfigurationStatus: to.Ptr(armstorsimple1200series.DeviceConfigurationStatusPending),
		// 						DeviceSoftwareVersion: to.Ptr("10.0.10296.0"),
		// 						DomainName: to.Ptr("fareast.corp.microsoft.com"),
		// 						ModelDescription: to.Ptr("1200"),
		// 						Status: to.Ptr(armstorsimple1200series.DeviceStatusReadyToSetup),
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("HSDK-YYMYCY4NK0"),
		// 					Type: to.Ptr("Microsoft.StorSimple/managers/devices"),
		// 					ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-YYMYCY4NK0"),
		// 					Properties: &armstorsimple1200series.DeviceProperties{
		// 						Type: to.Ptr(armstorsimple1200series.DeviceTypeSeries9000OnPremVirtualAppliance),
		// 						ActivationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-11T15:19:40.825Z"); return t}()),
		// 						AllowedDeviceOperations: []*armstorsimple1200series.DeviceOperation{
		// 							to.Ptr(armstorsimple1200series.DeviceOperationDelete),
		// 							to.Ptr(armstorsimple1200series.DeviceOperationDRTarget),
		// 							to.Ptr(armstorsimple1200series.DeviceOperationBrowsable)},
		// 							Culture: to.Ptr("en-US"),
		// 							DeviceCapabilities: []*armstorsimple1200series.SupportedDeviceCapabilities{
		// 								to.Ptr(armstorsimple1200series.SupportedDeviceCapabilitiesFileServer)},
		// 								DeviceConfigurationStatus: to.Ptr(armstorsimple1200series.DeviceConfigurationStatusPending),
		// 								DeviceSoftwareVersion: to.Ptr("10.0.10296.0"),
		// 								DomainName: to.Ptr("fareast.corp.microsoft.com"),
		// 								ModelDescription: to.Ptr("1200"),
		// 								Status: to.Ptr(armstorsimple1200series.DeviceStatusReadyToSetup),
		// 							},
		// 					}},
		// 				}
	}
}
