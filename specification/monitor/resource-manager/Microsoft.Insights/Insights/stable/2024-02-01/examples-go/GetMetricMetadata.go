package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: 2024-02-01/GetMetricMetadata.json
func ExampleMetricsClient_List_getMetricForMetadata() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMetricsClient().List(ctx, "subscriptions/1f3fa6d2-851c-4a91-9087-1a050f3a9c38/resourceGroups/todking/providers/Microsoft.Storage/storageAccounts/tkfileserv/blobServices/default", &armmonitor.MetricsClientListOptions{
		Filter:          to.Ptr("Tier eq '*'"),
		Metricnamespace: to.Ptr("Microsoft.Storage/storageAccounts/blobServices"),
		Timespan:        to.Ptr("2017-04-14T02:20:00Z/2017-04-14T04:20:00Z")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmonitor.MetricsClientListResponse{
	// 	Response: armmonitor.Response{
	// 		Interval: to.Ptr("PT1H"),
	// 		Namespace: to.Ptr("microsoft.storage/storageaccounts/blobservices"),
	// 		Resourceregion: to.Ptr("westus2"),
	// 		Timespan: to.Ptr("2021-04-15T02:18:00Z/2021-04-22T02:18:00Z"),
	// 		Value: []*armmonitor.Metric{
	// 			{
	// 				Name: &armmonitor.LocalizableString{
	// 					LocalizedValue: to.Ptr("Blob Count"),
	// 					Value: to.Ptr("BlobCount"),
	// 				},
	// 				Type: to.Ptr("Microsoft.Insights/metrics"),
	// 				DisplayDescription: to.Ptr("The number of blob objects stored in the storage account."),
	// 				ID: to.Ptr("/subscriptions/1f3fa6d2-851c-4a91-9087-1a050f3a9c38/resourceGroups/todking/providers/Microsoft.Storage/storageAccounts/tkfileserv/blobServices/default/providers/Microsoft.Insights/metrics/BlobCount"),
	// 				Timeseries: []*armmonitor.TimeSeriesElement{
	// 					{
	// 						Metadatavalues: []*armmonitor.MetadataValue{
	// 							{
	// 								Name: &armmonitor.LocalizableString{
	// 									LocalizedValue: to.Ptr("tier"),
	// 									Value: to.Ptr("tier"),
	// 								},
	// 								Value: to.Ptr("Cool"),
	// 							},
	// 						},
	// 					},
	// 					{
	// 						Metadatavalues: []*armmonitor.MetadataValue{
	// 							{
	// 								Name: &armmonitor.LocalizableString{
	// 									LocalizedValue: to.Ptr("tier"),
	// 									Value: to.Ptr("tier"),
	// 								},
	// 								Value: to.Ptr("Archive"),
	// 							},
	// 						},
	// 					},
	// 					{
	// 						Metadatavalues: []*armmonitor.MetadataValue{
	// 							{
	// 								Name: &armmonitor.LocalizableString{
	// 									LocalizedValue: to.Ptr("tier"),
	// 									Value: to.Ptr("tier"),
	// 								},
	// 								Value: to.Ptr("Standard"),
	// 							},
	// 						},
	// 					},
	// 					{
	// 						Metadatavalues: []*armmonitor.MetadataValue{
	// 							{
	// 								Name: &armmonitor.LocalizableString{
	// 									LocalizedValue: to.Ptr("tier"),
	// 									Value: to.Ptr("tier"),
	// 								},
	// 								Value: to.Ptr("Untiered"),
	// 							},
	// 						},
	// 					},
	// 					{
	// 						Metadatavalues: []*armmonitor.MetadataValue{
	// 							{
	// 								Name: &armmonitor.LocalizableString{
	// 									LocalizedValue: to.Ptr("tier"),
	// 									Value: to.Ptr("tier"),
	// 								},
	// 								Value: to.Ptr("Hot"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 				Unit: to.Ptr(armmonitor.MetricUnitCount),
	// 			},
	// 		},
	// 	},
	// }
}
