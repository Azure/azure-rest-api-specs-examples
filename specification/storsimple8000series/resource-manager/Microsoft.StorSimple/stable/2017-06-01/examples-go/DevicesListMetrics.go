package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DevicesListMetrics.json
func ExampleDevicesClient_NewListMetricsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDevicesClient().NewListMetricsPager("Device05ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", "name/value%20eq%20'PrimaryStorageTieredUsed'%20and%20timeGrain%20eq%20'PT1H'%20and%20startTime%20ge%20'2017-06-17T18:30:00Z'%20and%20endTime%20le%20'2017-06-21T18:30:00Z'%20and%20category%20eq%20'CapacityUtilization'", nil)
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
		// page.MetricList = armstorsimple8000series.MetricList{
		// 	Value: []*armstorsimple8000series.Metrics{
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("Primary Tiered Storage Used"),
		// 				Value: to.Ptr("PrimaryStorageTieredUsed"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metrics"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Device"),
		// 					Value: to.Ptr("Device05ForSDKTest"),
		// 			}},
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-21T18:30:00.000Z"); return t}()),
		// 			PrimaryAggregation: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-17T18:30:00.000Z"); return t}()),
		// 			TimeGrain: to.Ptr("01:00:00"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitBytes),
		// 			Values: []*armstorsimple8000series.MetricData{
		// 			},
		// 	}},
		// }
	}
}
