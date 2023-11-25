package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiServersListMetrics.json
func ExampleIscsiServersClient_NewListMetricsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIscsiServersClient().NewListMetricsPager("HSDK-UGU4PITWNI", "HSDK-UGU4PITWNI", "ResourceGroupForSDKTest", "hAzureSDKOperations", &armstorsimple1200series.IscsiServersClientListMetricsOptions{Filter: to.Ptr("startTime%20ge%20'2018-08-10T18:30:00Z'%20and%20endTime%20le%20'2018-08-11T18:30:00Z'")})
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
		// page.MetricList = armstorsimple1200series.MetricList{
		// 	Value: []*armstorsimple1200series.Metrics{
		// 		{
		// 			Name: &armstorsimple1200series.MetricName{
		// 				LocalizedValue: to.Ptr("Cloud Storage Used"),
		// 				Value: to.Ptr("CloudConsumedStorage"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiServers/metrics"),
		// 			Dimensions: []*armstorsimple1200series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("IscsiServer"),
		// 					Value: to.Ptr("HSDK-UGU4PITWNI"),
		// 			}},
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-11T18:30:00.000Z"); return t}()),
		// 			PrimaryAggregation: to.Ptr(armstorsimple1200series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("https://pod01-cis2.sea.storsimple.windowsazure.com/managers/4239154091695873374/devices/HSDK-UGU4PITWNI/iscsiservers/HSDK-UGU4PITWNI"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T18:30:00.000Z"); return t}()),
		// 			TimeGrain: to.Ptr("1.00:00:00"),
		// 			Unit: to.Ptr(armstorsimple1200series.MetricUnitBytes),
		// 			Values: []*armstorsimple1200series.MetricData{
		// 			},
		// 	}},
		// }
	}
}
