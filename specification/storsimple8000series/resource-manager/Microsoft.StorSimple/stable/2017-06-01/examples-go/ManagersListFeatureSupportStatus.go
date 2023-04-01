package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/ManagersListFeatureSupportStatus.json
func ExampleManagersClient_NewListFeatureSupportStatusPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagersClient().NewListFeatureSupportStatusPager("ResourceGroupForSDKTest", "ManagerForSDKTest1", &armstorsimple8000series.ManagersClientListFeatureSupportStatusOptions{Filter: nil})
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
		// page.FeatureList = armstorsimple8000series.FeatureList{
		// 	Value: []*armstorsimple8000series.Feature{
		// 		{
		// 			Name: to.Ptr("BandwidthThrottling"),
		// 			Status: to.Ptr(armstorsimple8000series.FeatureSupportStatusSupported),
		// 		},
		// 		{
		// 			Name: to.Ptr("OtherCloud"),
		// 			Status: to.Ptr(armstorsimple8000series.FeatureSupportStatusUnsupportedDeviceVersion),
		// 		},
		// 		{
		// 			Name: to.Ptr("DeviceUpdate"),
		// 			Status: to.Ptr(armstorsimple8000series.FeatureSupportStatusNotAvailable),
		// 		},
		// 		{
		// 			Name: to.Ptr("ControllerRestartFeature"),
		// 			Status: to.Ptr(armstorsimple8000series.FeatureSupportStatusSupported),
		// 		},
		// 		{
		// 			Name: to.Ptr("GetDRTargetDevicesApi"),
		// 			Status: to.Ptr(armstorsimple8000series.FeatureSupportStatusSupported),
		// 		},
		// 		{
		// 			Name: to.Ptr("Migration"),
		// 			Status: to.Ptr(armstorsimple8000series.FeatureSupportStatusSupported),
		// 		},
		// 		{
		// 			Name: to.Ptr("Update"),
		// 			Status: to.Ptr(armstorsimple8000series.FeatureSupportStatusSupported),
		// 		},
		// 		{
		// 			Name: to.Ptr("GetSupportedVirtualApplianceVersionInfoApi"),
		// 			Status: to.Ptr(armstorsimple8000series.FeatureSupportStatusSupported),
		// 		},
		// 		{
		// 			Name: to.Ptr("LocalOnlyVolume"),
		// 			Status: to.Ptr(armstorsimple8000series.FeatureSupportStatusSupported),
		// 		},
		// 		{
		// 			Name: to.Ptr("PremiumSVA"),
		// 			Status: to.Ptr(armstorsimple8000series.FeatureSupportStatusSupported),
		// 		},
		// 		{
		// 			Name: to.Ptr("VersionBasedARMApiAccess"),
		// 			Status: to.Ptr(armstorsimple8000series.FeatureSupportStatusNotAvailable),
		// 		},
		// 		{
		// 			Name: to.Ptr("SubscriptionBasedARMApiAccess"),
		// 			Status: to.Ptr(armstorsimple8000series.FeatureSupportStatusNotAvailable),
		// 	}},
		// }
	}
}
