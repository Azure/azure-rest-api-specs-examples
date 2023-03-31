package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/BandwidthSettingsGet.json
func ExampleBandwidthSettingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBandwidthSettingsClient().Get(ctx, "bandwidthSetting1", "ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BandwidthSetting = armstorsimple8000series.BandwidthSetting{
	// 	Name: to.Ptr("bandwidthSetting1"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/bandwidthSettings"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/bandwidthSettings/bandwidthSetting1"),
	// 	Kind: to.Ptr("Series8000"),
	// 	Properties: &armstorsimple8000series.BandwidthRateSettingProperties{
	// 		Schedules: []*armstorsimple8000series.BandwidthSchedule{
	// 			{
	// 				Days: []*armstorsimple8000series.DayOfWeek{
	// 					to.Ptr(armstorsimple8000series.DayOfWeekMonday),
	// 					to.Ptr(armstorsimple8000series.DayOfWeekTuesday),
	// 					to.Ptr(armstorsimple8000series.DayOfWeekWednesday),
	// 					to.Ptr(armstorsimple8000series.DayOfWeekThursday),
	// 					to.Ptr(armstorsimple8000series.DayOfWeekFriday),
	// 					to.Ptr(armstorsimple8000series.DayOfWeekSaturday),
	// 					to.Ptr(armstorsimple8000series.DayOfWeekSunday)},
	// 					RateInMbps: to.Ptr[int32](150),
	// 					Start: &armstorsimple8000series.Time{
	// 						Hours: to.Ptr[int32](10),
	// 						Minutes: to.Ptr[int32](0),
	// 						Seconds: to.Ptr[int32](0),
	// 					},
	// 					Stop: &armstorsimple8000series.Time{
	// 						Hours: to.Ptr[int32](22),
	// 						Minutes: to.Ptr[int32](0),
	// 						Seconds: to.Ptr[int32](0),
	// 					},
	// 				},
	// 				{
	// 					Days: []*armstorsimple8000series.DayOfWeek{
	// 						to.Ptr(armstorsimple8000series.DayOfWeekSaturday),
	// 						to.Ptr(armstorsimple8000series.DayOfWeekSunday)},
	// 						RateInMbps: to.Ptr[int32](100),
	// 						Start: &armstorsimple8000series.Time{
	// 							Hours: to.Ptr[int32](8),
	// 							Minutes: to.Ptr[int32](0),
	// 							Seconds: to.Ptr[int32](0),
	// 						},
	// 						Stop: &armstorsimple8000series.Time{
	// 							Hours: to.Ptr[int32](10),
	// 							Minutes: to.Ptr[int32](0),
	// 							Seconds: to.Ptr[int32](0),
	// 						},
	// 				}},
	// 				VolumeCount: to.Ptr[int32](1),
	// 			},
	// 		}
}
