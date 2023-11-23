package armmonitor_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2021-05-01/examples/PostMultiResourceMetricBody.json
func ExampleMetricsClient_ListAtSubscriptionScopePost_postRequestForSubscriptionLevelMetricDataUsingBodyParams() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMetricsClient().ListAtSubscriptionScopePost(ctx, "westus2", &armmonitor.MetricsClientListAtSubscriptionScopePostOptions{Timespan: nil,
		Interval:            nil,
		Metricnames:         nil,
		Aggregation:         nil,
		Top:                 nil,
		Orderby:             nil,
		Filter:              nil,
		ResultType:          nil,
		Metricnamespace:     nil,
		AutoAdjustTimegrain: nil,
		ValidateDimensions:  nil,
		Body: &armmonitor.SubscriptionScopeMetricsRequestBodyParameters{
			Aggregation:         to.Ptr("count"),
			AutoAdjustTimegrain: to.Ptr(true),
			Filter:              to.Ptr("LUN eq '0' and Microsoft.ResourceId eq '*'"),
			Interval:            to.Ptr("PT6H"),
			MetricNames:         to.Ptr("Data Disk Max Burst IOPS"),
			MetricNamespace:     to.Ptr("microsoft.compute/virtualmachines"),
			OrderBy:             to.Ptr("count desc"),
			RollUpBy:            to.Ptr("LUN"),
			Timespan:            to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "NaN-NaN-NaNTNaN:NaN:NaN.NaNZ"); return t }()),
			Top:                 to.Ptr[int32](10),
			ValidateDimensions:  to.Ptr(false),
		},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SubscriptionScopeMetricResponse = armmonitor.SubscriptionScopeMetricResponse{
	// 	Cost: to.Ptr[int32](4679),
	// 	Interval: to.Ptr("PT6H"),
	// 	Namespace: to.Ptr("microsoft.compute/virtualmachines"),
	// 	Resourceregion: to.Ptr("westus2"),
	// 	Timespan: to.Ptr("2021-06-08T19:00:00Z/2021-06-12T01:00:00Z"),
	// 	Value: []*armmonitor.SubscriptionScopeMetric{
	// 		{
	// 			Name: &armmonitor.LocalizableString{
	// 				LocalizedValue: to.Ptr("Data Disk Max Burst IOPS"),
	// 				Value: to.Ptr("Data Disk Max Burst IOPS"),
	// 			},
	// 			Type: to.Ptr("Microsoft.Insights/metrics"),
	// 			DisplayDescription: to.Ptr("Maximum IOPS Data Disk can achieve with bursting"),
	// 			ErrorCode: to.Ptr("Success"),
	// 			ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/Microsoft.Insights/metrics/Data Disk Max Burst IOPS"),
	// 			Timeseries: []*armmonitor.TimeSeriesElement{
	// 				{
	// 					Data: []*armmonitor.MetricValue{
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-08T19:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T01:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T07:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T13:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T19:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T01:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T07:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](413),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T13:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T19:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-11T01:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-11T07:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-11T13:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-11T19:00:00.000Z"); return t}()),
	// 					}},
	// 					Metadatavalues: []*armmonitor.MetadataValue{
	// 						{
	// 							Name: &armmonitor.LocalizableString{
	// 								LocalizedValue: to.Ptr("microsoft.resourceid"),
	// 								Value: to.Ptr("microsoft.resourceid"),
	// 							},
	// 							Value: to.Ptr("/subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/resourceGroups/sas1/providers/Microsoft.Compute/virtualMachines/sas1-dev"),
	// 					}},
	// 				},
	// 				{
	// 					Data: []*armmonitor.MetricValue{
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-08T19:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](133),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T01:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T07:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T13:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T19:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T01:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T07:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T13:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T19:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-11T01:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-11T07:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-11T13:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-11T19:00:00.000Z"); return t}()),
	// 					}},
	// 					Metadatavalues: []*armmonitor.MetadataValue{
	// 						{
	// 							Name: &armmonitor.LocalizableString{
	// 								LocalizedValue: to.Ptr("microsoft.resourceid"),
	// 								Value: to.Ptr("microsoft.resourceid"),
	// 							},
	// 							Value: to.Ptr("/subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/resourceGroups/sas2/providers/Microsoft.Compute/virtualMachines/sas2-vm"),
	// 					}},
	// 				},
	// 				{
	// 					Data: []*armmonitor.MetricValue{
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-08T19:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T01:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T07:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T13:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T19:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T01:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T07:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T13:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T19:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](78),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-11T01:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-11T07:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-11T13:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-11T19:00:00.000Z"); return t}()),
	// 					}},
	// 					Metadatavalues: []*armmonitor.MetadataValue{
	// 						{
	// 							Name: &armmonitor.LocalizableString{
	// 								LocalizedValue: to.Ptr("microsoft.resourceid"),
	// 								Value: to.Ptr("microsoft.resourceid"),
	// 							},
	// 							Value: to.Ptr("/subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/resourceGroups/sas3/providers/Microsoft.Compute/virtualMachines/sas3-vm"),
	// 					}},
	// 				},
	// 				{
	// 					Data: []*armmonitor.MetricValue{
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-08T19:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T01:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T07:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T13:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T19:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T01:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T07:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T13:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T19:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-11T01:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-11T07:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-11T13:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-11T19:00:00.000Z"); return t}()),
	// 					}},
	// 					Metadatavalues: []*armmonitor.MetadataValue{
	// 						{
	// 							Name: &armmonitor.LocalizableString{
	// 								LocalizedValue: to.Ptr("microsoft.resourceid"),
	// 								Value: to.Ptr("microsoft.resourceid"),
	// 							},
	// 							Value: to.Ptr("/subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/resourceGroups/sas4/providers/Microsoft.Compute/virtualMachines/sas4-vm"),
	// 					}},
	// 				},
	// 				{
	// 					Data: []*armmonitor.MetricValue{
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-08T19:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T01:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T07:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T13:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T19:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T01:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T07:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T13:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T19:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-11T01:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-11T07:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-11T13:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Count: to.Ptr[float64](72),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-11T19:00:00.000Z"); return t}()),
	// 					}},
	// 					Metadatavalues: []*armmonitor.MetadataValue{
	// 						{
	// 							Name: &armmonitor.LocalizableString{
	// 								LocalizedValue: to.Ptr("microsoft.resourceid"),
	// 								Value: to.Ptr("microsoft.resourceid"),
	// 							},
	// 							Value: to.Ptr("/subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/resourceGroups/sas5/providers/Microsoft.Compute/virtualMachines/sas5-vm-asc"),
	// 					}},
	// 			}},
	// 			Unit: to.Ptr(armmonitor.MetricUnitCount),
	// 	}},
	// }
}
