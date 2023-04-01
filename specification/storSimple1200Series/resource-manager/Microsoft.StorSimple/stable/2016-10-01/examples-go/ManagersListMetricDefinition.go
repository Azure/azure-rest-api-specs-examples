package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/ManagersListMetricDefinition.json
func ExampleManagersClient_NewListMetricDefinitionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagersClient().NewListMetricDefinitionPager("ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
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
		// page.MetricDefinitionList = armstorsimple1200series.MetricDefinitionList{
		// 	Value: []*armstorsimple1200series.MetricDefinition{
		// 		{
		// 			Name: &armstorsimple1200series.MetricName{
		// 				LocalizedValue: to.Ptr("Primary Storage Used"),
		// 				Value: to.Ptr("HostUsedStorage"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/metricsDefinitions"),
		// 			Dimensions: []*armstorsimple1200series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Manager"),
		// 					Value: to.Ptr("hAzureSDKOperations"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple1200series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1M"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P3M"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1Y"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armstorsimple1200series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("https://pod01-cis2.sea.storsimple.windowsazure.com/managers/4239154091695873374"),
		// 			Unit: to.Ptr(armstorsimple1200series.MetricUnitBytes),
		// 		},
		// 		{
		// 			Name: &armstorsimple1200series.MetricName{
		// 				LocalizedValue: to.Ptr("Cloud Storage Used"),
		// 				Value: to.Ptr("CloudConsumedStorage"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/metricsDefinitions"),
		// 			Dimensions: []*armstorsimple1200series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Manager"),
		// 					Value: to.Ptr("hAzureSDKOperations"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple1200series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1M"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P3M"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1Y"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armstorsimple1200series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("https://pod01-cis2.sea.storsimple.windowsazure.com/managers/4239154091695873374"),
		// 			Unit: to.Ptr(armstorsimple1200series.MetricUnitBytes),
		// 	}},
		// }
	}
}
