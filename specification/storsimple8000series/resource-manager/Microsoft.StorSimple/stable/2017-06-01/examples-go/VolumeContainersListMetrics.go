package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/VolumeContainersListMetrics.json
func ExampleVolumeContainersClient_NewListMetricsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVolumeContainersClient().NewListMetricsPager("Device05ForSDKTest", "vcForOdataFilterTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", "name/value%20eq%20'CloudConsumedStorage'%20and%20timeGrain%20eq%20'PT1M'%20and%20startTime%20ge%20'2017-06-17T18:30:00Z'%20and%20endTime%20le%20'2017-06-21T18:30:00Z'%20and%20category%20eq%20'CapacityUtilization'", nil)
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
		// 				LocalizedValue: to.Ptr("Cloud Storage Used"),
		// 				Value: to.Ptr("CloudConsumedStorage"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/volumeContainers/metrics"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("VolumeContainer"),
		// 					Value: to.Ptr("vcForOdataFilterTest"),
		// 			}},
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-21T18:30:00.000Z"); return t}()),
		// 			PrimaryAggregation: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/vcForOdataFilterTest"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-17T18:30:00.000Z"); return t}()),
		// 			TimeGrain: to.Ptr("00:01:00"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitBytes),
		// 			Values: []*armstorsimple8000series.MetricData{
		// 			},
		// 	}},
		// }
	}
}
