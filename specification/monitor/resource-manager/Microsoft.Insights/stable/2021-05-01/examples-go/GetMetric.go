package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2021-05-01/examples/GetMetric.json
func ExampleMetricsClient_List_getMetricForData() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMetricsClient().List(ctx, "subscriptions/1f3fa6d2-851c-4a91-9087-1a050f3a9c38/resourceGroups/todking/providers/Microsoft.Storage/storageAccounts/tkfileserv/blobServices/default", &armmonitor.MetricsClientListOptions{Timespan: to.Ptr("2021-04-20T09:00:00.000Z/2021-04-20T14:00:00.000Z"),
		Interval:            to.Ptr("PT6H"),
		Metricnames:         to.Ptr("BlobCount,BlobCapacity"),
		Aggregation:         to.Ptr("average,minimum,maximum"),
		Top:                 to.Ptr[int32](5),
		Orderby:             to.Ptr("average asc"),
		Filter:              to.Ptr("Tier eq '*'"),
		ResultType:          nil,
		Metricnamespace:     to.Ptr("Microsoft.Storage/storageAccounts/blobServices"),
		AutoAdjustTimegrain: to.Ptr(true),
		ValidateDimensions:  to.Ptr(false),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Response = armmonitor.Response{
	// 	Cost: to.Ptr[int32](598),
	// 	Interval: to.Ptr("PT1H"),
	// 	Namespace: to.Ptr("microsoft.storage/storageaccounts/blobservices"),
	// 	Resourceregion: to.Ptr("westus2"),
	// 	Timespan: to.Ptr("2021-04-20T09:00:00Z/2021-04-20T14:00:00Z"),
	// 	Value: []*armmonitor.Metric{
	// 		{
	// 			Name: &armmonitor.LocalizableString{
	// 				LocalizedValue: to.Ptr("Blob Count"),
	// 				Value: to.Ptr("BlobCount"),
	// 			},
	// 			Type: to.Ptr("Microsoft.Insights/metrics"),
	// 			DisplayDescription: to.Ptr("The number of blob objects stored in the storage account."),
	// 			ErrorCode: to.Ptr("Success"),
	// 			ID: to.Ptr("/subscriptions/1f3fa6d2-851c-4a91-9087-1a050f3a9c38/resourceGroups/todking/providers/Microsoft.Storage/storageAccounts/tkfileserv/blobServices/default/providers/Microsoft.Insights/metrics/BlobCount"),
	// 			Timeseries: []*armmonitor.TimeSeriesElement{
	// 				{
	// 					Data: []*armmonitor.MetricValue{
	// 						{
	// 							Average: to.Ptr[float64](7),
	// 							Maximum: to.Ptr[float64](7),
	// 							Minimum: to.Ptr[float64](7),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T09:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](7),
	// 							Maximum: to.Ptr[float64](7),
	// 							Minimum: to.Ptr[float64](7),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T10:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](7),
	// 							Maximum: to.Ptr[float64](7),
	// 							Minimum: to.Ptr[float64](7),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T11:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](7),
	// 							Maximum: to.Ptr[float64](7),
	// 							Minimum: to.Ptr[float64](7),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T12:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](7),
	// 							Maximum: to.Ptr[float64](7),
	// 							Minimum: to.Ptr[float64](7),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T13:00:00.000Z"); return t}()),
	// 					}},
	// 					Metadatavalues: []*armmonitor.MetadataValue{
	// 						{
	// 							Name: &armmonitor.LocalizableString{
	// 								LocalizedValue: to.Ptr("tier"),
	// 								Value: to.Ptr("tier"),
	// 							},
	// 							Value: to.Ptr("Hot"),
	// 					}},
	// 				},
	// 				{
	// 					Data: []*armmonitor.MetricValue{
	// 						{
	// 							Average: to.Ptr[float64](2),
	// 							Maximum: to.Ptr[float64](2),
	// 							Minimum: to.Ptr[float64](2),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T09:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](2),
	// 							Maximum: to.Ptr[float64](2),
	// 							Minimum: to.Ptr[float64](2),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T10:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](2),
	// 							Maximum: to.Ptr[float64](2),
	// 							Minimum: to.Ptr[float64](2),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T11:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](2),
	// 							Maximum: to.Ptr[float64](2),
	// 							Minimum: to.Ptr[float64](2),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T12:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](2),
	// 							Maximum: to.Ptr[float64](2),
	// 							Minimum: to.Ptr[float64](2),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T13:00:00.000Z"); return t}()),
	// 					}},
	// 					Metadatavalues: []*armmonitor.MetadataValue{
	// 						{
	// 							Name: &armmonitor.LocalizableString{
	// 								LocalizedValue: to.Ptr("tier"),
	// 								Value: to.Ptr("tier"),
	// 							},
	// 							Value: to.Ptr("Standard"),
	// 					}},
	// 				},
	// 				{
	// 					Data: []*armmonitor.MetricValue{
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T09:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T10:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T11:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T12:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T13:00:00.000Z"); return t}()),
	// 					}},
	// 					Metadatavalues: []*armmonitor.MetadataValue{
	// 						{
	// 							Name: &armmonitor.LocalizableString{
	// 								LocalizedValue: to.Ptr("tier"),
	// 								Value: to.Ptr("tier"),
	// 							},
	// 							Value: to.Ptr("Cool"),
	// 					}},
	// 				},
	// 				{
	// 					Data: []*armmonitor.MetricValue{
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T09:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T10:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T11:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T12:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T13:00:00.000Z"); return t}()),
	// 					}},
	// 					Metadatavalues: []*armmonitor.MetadataValue{
	// 						{
	// 							Name: &armmonitor.LocalizableString{
	// 								LocalizedValue: to.Ptr("tier"),
	// 								Value: to.Ptr("tier"),
	// 							},
	// 							Value: to.Ptr("Archive"),
	// 					}},
	// 				},
	// 				{
	// 					Data: []*armmonitor.MetricValue{
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T09:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T10:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T11:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T12:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T13:00:00.000Z"); return t}()),
	// 					}},
	// 					Metadatavalues: []*armmonitor.MetadataValue{
	// 						{
	// 							Name: &armmonitor.LocalizableString{
	// 								LocalizedValue: to.Ptr("tier"),
	// 								Value: to.Ptr("tier"),
	// 							},
	// 							Value: to.Ptr("Untiered"),
	// 					}},
	// 			}},
	// 			Unit: to.Ptr(armmonitor.UnitCount),
	// 		},
	// 		{
	// 			Name: &armmonitor.LocalizableString{
	// 				LocalizedValue: to.Ptr("Blob Capacity"),
	// 				Value: to.Ptr("BlobCapacity"),
	// 			},
	// 			Type: to.Ptr("Microsoft.Insights/metrics"),
	// 			DisplayDescription: to.Ptr("The amount of storage used by the storage account’s Blob service in bytes."),
	// 			ErrorCode: to.Ptr("Success"),
	// 			ID: to.Ptr("/subscriptions/1f3fa6d2-851c-4a91-9087-1a050f3a9c38/resourceGroups/todking/providers/Microsoft.Storage/storageAccounts/tkfileserv/blobServices/default/providers/Microsoft.Insights/metrics/BlobCapacity"),
	// 			Timeseries: []*armmonitor.TimeSeriesElement{
	// 				{
	// 					Data: []*armmonitor.MetricValue{
	// 						{
	// 							Average: to.Ptr[float64](621492),
	// 							Maximum: to.Ptr[float64](621492),
	// 							Minimum: to.Ptr[float64](621492),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T09:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](621492),
	// 							Maximum: to.Ptr[float64](621492),
	// 							Minimum: to.Ptr[float64](621492),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T10:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](621492),
	// 							Maximum: to.Ptr[float64](621492),
	// 							Minimum: to.Ptr[float64](621492),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T11:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](621492),
	// 							Maximum: to.Ptr[float64](621492),
	// 							Minimum: to.Ptr[float64](621492),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T12:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](621492),
	// 							Maximum: to.Ptr[float64](621492),
	// 							Minimum: to.Ptr[float64](621492),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T13:00:00.000Z"); return t}()),
	// 					}},
	// 					Metadatavalues: []*armmonitor.MetadataValue{
	// 						{
	// 							Name: &armmonitor.LocalizableString{
	// 								LocalizedValue: to.Ptr("tier"),
	// 								Value: to.Ptr("tier"),
	// 							},
	// 							Value: to.Ptr("Standard"),
	// 					}},
	// 				},
	// 				{
	// 					Data: []*armmonitor.MetricValue{
	// 						{
	// 							Average: to.Ptr[float64](4733),
	// 							Maximum: to.Ptr[float64](4733),
	// 							Minimum: to.Ptr[float64](4733),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T09:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](4733),
	// 							Maximum: to.Ptr[float64](4733),
	// 							Minimum: to.Ptr[float64](4733),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T10:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](4733),
	// 							Maximum: to.Ptr[float64](4733),
	// 							Minimum: to.Ptr[float64](4733),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T11:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](4733),
	// 							Maximum: to.Ptr[float64](4733),
	// 							Minimum: to.Ptr[float64](4733),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T12:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](4733),
	// 							Maximum: to.Ptr[float64](4733),
	// 							Minimum: to.Ptr[float64](4733),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T13:00:00.000Z"); return t}()),
	// 					}},
	// 					Metadatavalues: []*armmonitor.MetadataValue{
	// 						{
	// 							Name: &armmonitor.LocalizableString{
	// 								LocalizedValue: to.Ptr("tier"),
	// 								Value: to.Ptr("tier"),
	// 							},
	// 							Value: to.Ptr("Hot"),
	// 					}},
	// 				},
	// 				{
	// 					Data: []*armmonitor.MetricValue{
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T09:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T10:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T11:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T12:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T13:00:00.000Z"); return t}()),
	// 					}},
	// 					Metadatavalues: []*armmonitor.MetadataValue{
	// 						{
	// 							Name: &armmonitor.LocalizableString{
	// 								LocalizedValue: to.Ptr("tier"),
	// 								Value: to.Ptr("tier"),
	// 							},
	// 							Value: to.Ptr("Archive"),
	// 					}},
	// 				},
	// 				{
	// 					Data: []*armmonitor.MetricValue{
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T09:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T10:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T11:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T12:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T13:00:00.000Z"); return t}()),
	// 					}},
	// 					Metadatavalues: []*armmonitor.MetadataValue{
	// 						{
	// 							Name: &armmonitor.LocalizableString{
	// 								LocalizedValue: to.Ptr("tier"),
	// 								Value: to.Ptr("tier"),
	// 							},
	// 							Value: to.Ptr("Untiered"),
	// 					}},
	// 				},
	// 				{
	// 					Data: []*armmonitor.MetricValue{
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T09:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T10:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T11:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T12:00:00.000Z"); return t}()),
	// 						},
	// 						{
	// 							Average: to.Ptr[float64](0),
	// 							Maximum: to.Ptr[float64](0),
	// 							Minimum: to.Ptr[float64](0),
	// 							TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-20T13:00:00.000Z"); return t}()),
	// 					}},
	// 					Metadatavalues: []*armmonitor.MetadataValue{
	// 						{
	// 							Name: &armmonitor.LocalizableString{
	// 								LocalizedValue: to.Ptr("tier"),
	// 								Value: to.Ptr("tier"),
	// 							},
	// 							Value: to.Ptr("Cool"),
	// 					}},
	// 			}},
	// 			Unit: to.Ptr(armmonitor.UnitBytes),
	// 	}},
	// }
}
