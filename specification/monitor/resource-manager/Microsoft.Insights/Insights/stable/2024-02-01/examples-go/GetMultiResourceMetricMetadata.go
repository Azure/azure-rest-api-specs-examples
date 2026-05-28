package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: 2024-02-01/GetMultiResourceMetricMetadata.json
func ExampleMetricsClient_ListAtSubscriptionScope_getSubscriptionLevelMetricMetadata() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("92d2a2d8-b514-432d-8cc9-a5f9272630d5", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMetricsClient().ListAtSubscriptionScope(ctx, "westus2", &armmonitor.MetricsClientListAtSubscriptionScopeOptions{
		Filter:          to.Ptr("LUN eq '0'"),
		Metricnames:     to.Ptr("Data Disk Max Burst IOPS"),
		Metricnamespace: to.Ptr("microsoft.compute/virtualmachines"),
		Timespan:        to.Ptr("2021-06-10T02:23:16.129Z/2021-06-12T02:23:16.129Z")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmonitor.MetricsClientListAtSubscriptionScopeResponse{
	// 	Response: armmonitor.Response{
	// 		Interval: to.Ptr("PT1M"),
	// 		Namespace: to.Ptr("microsoft.compute/virtualmachines"),
	// 		Resourceregion: to.Ptr("westus2"),
	// 		Timespan: to.Ptr("2021-06-10T02:23:16Z/2021-06-12T02:23:16Z"),
	// 		Value: []*armmonitor.Metric{
	// 			{
	// 				Name: &armmonitor.LocalizableString{
	// 					LocalizedValue: to.Ptr("Data Disk Read Bytes/Sec"),
	// 					Value: to.Ptr("Data Disk Read Bytes/sec"),
	// 				},
	// 				Type: to.Ptr("Microsoft.Insights/metrics"),
	// 				DisplayDescription: to.Ptr("Bytes/Sec read from a single disk during monitoring period"),
	// 				ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/Microsoft.Insights/metrics/Data Disk Read Bytes/sec"),
	// 				Timeseries: []*armmonitor.TimeSeriesElement{
	// 					{
	// 						Metadatavalues: []*armmonitor.MetadataValue{
	// 							{
	// 								Name: &armmonitor.LocalizableString{
	// 									LocalizedValue: to.Ptr("lun"),
	// 									Value: to.Ptr("lun"),
	// 								},
	// 								Value: to.Ptr("0"),
	// 							},
	// 						},
	// 					},
	// 					{
	// 						Metadatavalues: []*armmonitor.MetadataValue{
	// 							{
	// 								Name: &armmonitor.LocalizableString{
	// 									LocalizedValue: to.Ptr("lun"),
	// 									Value: to.Ptr("lun"),
	// 								},
	// 								Value: to.Ptr("1"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 				Unit: to.Ptr(armmonitor.MetricUnitBytesPerSecond),
	// 			},
	// 		},
	// 	},
	// }
}
