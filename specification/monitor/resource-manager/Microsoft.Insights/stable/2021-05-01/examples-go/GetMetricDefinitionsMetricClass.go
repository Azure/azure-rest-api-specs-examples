package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2021-05-01/examples/GetMetricDefinitionsMetricClass.json
func ExampleMetricDefinitionsClient_NewListPager_getStorageCacheMetricDefinitionsWithMetricClass() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMetricDefinitionsClient().NewListPager("subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache", &armmonitor.MetricDefinitionsClientListOptions{Metricnamespace: to.Ptr("microsoft.storagecache/caches")})
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
		// page.MetricDefinitionCollection = armmonitor.MetricDefinitionCollection{
		// 	Value: []*armmonitor.MetricDefinition{
		// 		{
		// 			Name: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("Total Client IOPS"),
		// 				Value: to.Ptr("ClientIOPS"),
		// 			},
		// 			DisplayDescription: to.Ptr("The rate of client file operations processed by the Cache."),
		// 			ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/ClientIOPS"),
		// 			IsDimensionRequired: to.Ptr(false),
		// 			MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 				{
		// 					Retention: to.Ptr("P93D"),
		// 					TimeGrain: to.Ptr("PT1M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P93D"),
		// 					TimeGrain: to.Ptr("PT5M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P93D"),
		// 					TimeGrain: to.Ptr("PT15M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P93D"),
		// 					TimeGrain: to.Ptr("PT30M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P93D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P93D"),
		// 					TimeGrain: to.Ptr("PT6H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P93D"),
		// 					TimeGrain: to.Ptr("PT12H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P93D"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			MetricClass: to.Ptr(armmonitor.MetricClassTransactions),
		// 			Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 			PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 			SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 				to.Ptr(armmonitor.AggregationTypeMinimum),
		// 				to.Ptr(armmonitor.AggregationTypeMaximum),
		// 				to.Ptr(armmonitor.AggregationTypeAverage)},
		// 				Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 			},
		// 			{
		// 				Name: &armmonitor.LocalizableString{
		// 					LocalizedValue: to.Ptr("Average Client Latency"),
		// 					Value: to.Ptr("ClientLatency"),
		// 				},
		// 				DisplayDescription: to.Ptr("Average latency of client file operations to the Cache."),
		// 				ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/ClientLatency"),
		// 				IsDimensionRequired: to.Ptr(false),
		// 				MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 					{
		// 						Retention: to.Ptr("P93D"),
		// 						TimeGrain: to.Ptr("PT1M"),
		// 					},
		// 					{
		// 						Retention: to.Ptr("P93D"),
		// 						TimeGrain: to.Ptr("PT5M"),
		// 					},
		// 					{
		// 						Retention: to.Ptr("P93D"),
		// 						TimeGrain: to.Ptr("PT15M"),
		// 					},
		// 					{
		// 						Retention: to.Ptr("P93D"),
		// 						TimeGrain: to.Ptr("PT30M"),
		// 					},
		// 					{
		// 						Retention: to.Ptr("P93D"),
		// 						TimeGrain: to.Ptr("PT1H"),
		// 					},
		// 					{
		// 						Retention: to.Ptr("P93D"),
		// 						TimeGrain: to.Ptr("PT6H"),
		// 					},
		// 					{
		// 						Retention: to.Ptr("P93D"),
		// 						TimeGrain: to.Ptr("PT12H"),
		// 					},
		// 					{
		// 						Retention: to.Ptr("P93D"),
		// 						TimeGrain: to.Ptr("P1D"),
		// 				}},
		// 				MetricClass: to.Ptr(armmonitor.MetricClassLatency),
		// 				Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 				PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 				ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 				SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 					to.Ptr(armmonitor.AggregationTypeMinimum),
		// 					to.Ptr(armmonitor.AggregationTypeMaximum),
		// 					to.Ptr(armmonitor.AggregationTypeAverage)},
		// 					Unit: to.Ptr(armmonitor.MetricUnitMilliSeconds),
		// 				},
		// 				{
		// 					Name: &armmonitor.LocalizableString{
		// 						LocalizedValue: to.Ptr("Client Read IOPS"),
		// 						Value: to.Ptr("ClientReadIOPS"),
		// 					},
		// 					DisplayDescription: to.Ptr("Client read operations per second."),
		// 					ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/ClientReadIOPS"),
		// 					IsDimensionRequired: to.Ptr(false),
		// 					MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 						{
		// 							Retention: to.Ptr("P93D"),
		// 							TimeGrain: to.Ptr("PT1M"),
		// 						},
		// 						{
		// 							Retention: to.Ptr("P93D"),
		// 							TimeGrain: to.Ptr("PT5M"),
		// 						},
		// 						{
		// 							Retention: to.Ptr("P93D"),
		// 							TimeGrain: to.Ptr("PT15M"),
		// 						},
		// 						{
		// 							Retention: to.Ptr("P93D"),
		// 							TimeGrain: to.Ptr("PT30M"),
		// 						},
		// 						{
		// 							Retention: to.Ptr("P93D"),
		// 							TimeGrain: to.Ptr("PT1H"),
		// 						},
		// 						{
		// 							Retention: to.Ptr("P93D"),
		// 							TimeGrain: to.Ptr("PT6H"),
		// 						},
		// 						{
		// 							Retention: to.Ptr("P93D"),
		// 							TimeGrain: to.Ptr("PT12H"),
		// 						},
		// 						{
		// 							Retention: to.Ptr("P93D"),
		// 							TimeGrain: to.Ptr("P1D"),
		// 					}},
		// 					MetricClass: to.Ptr(armmonitor.MetricClassTransactions),
		// 					Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 					PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 					ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 					SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 						to.Ptr(armmonitor.AggregationTypeMinimum),
		// 						to.Ptr(armmonitor.AggregationTypeMaximum),
		// 						to.Ptr(armmonitor.AggregationTypeAverage)},
		// 						Unit: to.Ptr(armmonitor.MetricUnitCountPerSecond),
		// 					},
		// 					{
		// 						Name: &armmonitor.LocalizableString{
		// 							LocalizedValue: to.Ptr("Average Cache Read Throughput"),
		// 							Value: to.Ptr("ClientReadThroughput"),
		// 						},
		// 						DisplayDescription: to.Ptr("Client read data transfer rate."),
		// 						ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/ClientReadThroughput"),
		// 						IsDimensionRequired: to.Ptr(false),
		// 						MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 							{
		// 								Retention: to.Ptr("P93D"),
		// 								TimeGrain: to.Ptr("PT1M"),
		// 							},
		// 							{
		// 								Retention: to.Ptr("P93D"),
		// 								TimeGrain: to.Ptr("PT5M"),
		// 							},
		// 							{
		// 								Retention: to.Ptr("P93D"),
		// 								TimeGrain: to.Ptr("PT15M"),
		// 							},
		// 							{
		// 								Retention: to.Ptr("P93D"),
		// 								TimeGrain: to.Ptr("PT30M"),
		// 							},
		// 							{
		// 								Retention: to.Ptr("P93D"),
		// 								TimeGrain: to.Ptr("PT1H"),
		// 							},
		// 							{
		// 								Retention: to.Ptr("P93D"),
		// 								TimeGrain: to.Ptr("PT6H"),
		// 							},
		// 							{
		// 								Retention: to.Ptr("P93D"),
		// 								TimeGrain: to.Ptr("PT12H"),
		// 							},
		// 							{
		// 								Retention: to.Ptr("P93D"),
		// 								TimeGrain: to.Ptr("P1D"),
		// 						}},
		// 						MetricClass: to.Ptr(armmonitor.MetricClassTransactions),
		// 						Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 						PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 						ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 						SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 							to.Ptr(armmonitor.AggregationTypeMinimum),
		// 							to.Ptr(armmonitor.AggregationTypeMaximum),
		// 							to.Ptr(armmonitor.AggregationTypeAverage)},
		// 							Unit: to.Ptr(armmonitor.MetricUnitBytesPerSecond),
		// 						},
		// 						{
		// 							Name: &armmonitor.LocalizableString{
		// 								LocalizedValue: to.Ptr("Client Write IOPS"),
		// 								Value: to.Ptr("ClientWriteIOPS"),
		// 							},
		// 							DisplayDescription: to.Ptr("Client write operations per second."),
		// 							ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/ClientWriteIOPS"),
		// 							IsDimensionRequired: to.Ptr(false),
		// 							MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 								{
		// 									Retention: to.Ptr("P93D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 								},
		// 								{
		// 									Retention: to.Ptr("P93D"),
		// 									TimeGrain: to.Ptr("PT5M"),
		// 								},
		// 								{
		// 									Retention: to.Ptr("P93D"),
		// 									TimeGrain: to.Ptr("PT15M"),
		// 								},
		// 								{
		// 									Retention: to.Ptr("P93D"),
		// 									TimeGrain: to.Ptr("PT30M"),
		// 								},
		// 								{
		// 									Retention: to.Ptr("P93D"),
		// 									TimeGrain: to.Ptr("PT1H"),
		// 								},
		// 								{
		// 									Retention: to.Ptr("P93D"),
		// 									TimeGrain: to.Ptr("PT6H"),
		// 								},
		// 								{
		// 									Retention: to.Ptr("P93D"),
		// 									TimeGrain: to.Ptr("PT12H"),
		// 								},
		// 								{
		// 									Retention: to.Ptr("P93D"),
		// 									TimeGrain: to.Ptr("P1D"),
		// 							}},
		// 							MetricClass: to.Ptr(armmonitor.MetricClassTransactions),
		// 							Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 							PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 							ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 							SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 								to.Ptr(armmonitor.AggregationTypeMinimum),
		// 								to.Ptr(armmonitor.AggregationTypeMaximum),
		// 								to.Ptr(armmonitor.AggregationTypeAverage)},
		// 								Unit: to.Ptr(armmonitor.MetricUnitCountPerSecond),
		// 							},
		// 							{
		// 								Name: &armmonitor.LocalizableString{
		// 									LocalizedValue: to.Ptr("Average Cache Write Throughput"),
		// 									Value: to.Ptr("ClientWriteThroughput"),
		// 								},
		// 								DisplayDescription: to.Ptr("Client write data transfer rate."),
		// 								ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/ClientWriteThroughput"),
		// 								IsDimensionRequired: to.Ptr(false),
		// 								MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 									{
		// 										Retention: to.Ptr("P93D"),
		// 										TimeGrain: to.Ptr("PT1M"),
		// 									},
		// 									{
		// 										Retention: to.Ptr("P93D"),
		// 										TimeGrain: to.Ptr("PT5M"),
		// 									},
		// 									{
		// 										Retention: to.Ptr("P93D"),
		// 										TimeGrain: to.Ptr("PT15M"),
		// 									},
		// 									{
		// 										Retention: to.Ptr("P93D"),
		// 										TimeGrain: to.Ptr("PT30M"),
		// 									},
		// 									{
		// 										Retention: to.Ptr("P93D"),
		// 										TimeGrain: to.Ptr("PT1H"),
		// 									},
		// 									{
		// 										Retention: to.Ptr("P93D"),
		// 										TimeGrain: to.Ptr("PT6H"),
		// 									},
		// 									{
		// 										Retention: to.Ptr("P93D"),
		// 										TimeGrain: to.Ptr("PT12H"),
		// 									},
		// 									{
		// 										Retention: to.Ptr("P93D"),
		// 										TimeGrain: to.Ptr("P1D"),
		// 								}},
		// 								MetricClass: to.Ptr(armmonitor.MetricClassTransactions),
		// 								Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 								PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 								ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 								SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 									to.Ptr(armmonitor.AggregationTypeMinimum),
		// 									to.Ptr(armmonitor.AggregationTypeMaximum),
		// 									to.Ptr(armmonitor.AggregationTypeAverage)},
		// 									Unit: to.Ptr(armmonitor.MetricUnitBytesPerSecond),
		// 								},
		// 								{
		// 									Name: &armmonitor.LocalizableString{
		// 										LocalizedValue: to.Ptr("Client Metadata Read IOPS"),
		// 										Value: to.Ptr("ClientMetadataReadIOPS"),
		// 									},
		// 									DisplayDescription: to.Ptr("The rate of client file operations sent to the Cache, excluding data reads, that do not modify persistent state."),
		// 									ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/ClientMetadataReadIOPS"),
		// 									IsDimensionRequired: to.Ptr(false),
		// 									MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 										{
		// 											Retention: to.Ptr("P93D"),
		// 											TimeGrain: to.Ptr("PT1M"),
		// 										},
		// 										{
		// 											Retention: to.Ptr("P93D"),
		// 											TimeGrain: to.Ptr("PT5M"),
		// 										},
		// 										{
		// 											Retention: to.Ptr("P93D"),
		// 											TimeGrain: to.Ptr("PT15M"),
		// 										},
		// 										{
		// 											Retention: to.Ptr("P93D"),
		// 											TimeGrain: to.Ptr("PT30M"),
		// 										},
		// 										{
		// 											Retention: to.Ptr("P93D"),
		// 											TimeGrain: to.Ptr("PT1H"),
		// 										},
		// 										{
		// 											Retention: to.Ptr("P93D"),
		// 											TimeGrain: to.Ptr("PT6H"),
		// 										},
		// 										{
		// 											Retention: to.Ptr("P93D"),
		// 											TimeGrain: to.Ptr("PT12H"),
		// 										},
		// 										{
		// 											Retention: to.Ptr("P93D"),
		// 											TimeGrain: to.Ptr("P1D"),
		// 									}},
		// 									MetricClass: to.Ptr(armmonitor.MetricClassTransactions),
		// 									Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 									PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 									ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 									SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 										to.Ptr(armmonitor.AggregationTypeMinimum),
		// 										to.Ptr(armmonitor.AggregationTypeMaximum),
		// 										to.Ptr(armmonitor.AggregationTypeAverage)},
		// 										Unit: to.Ptr(armmonitor.MetricUnitCountPerSecond),
		// 									},
		// 									{
		// 										Name: &armmonitor.LocalizableString{
		// 											LocalizedValue: to.Ptr("Client Metadata Write IOPS"),
		// 											Value: to.Ptr("ClientMetadataWriteIOPS"),
		// 										},
		// 										DisplayDescription: to.Ptr("The rate of client file operations sent to the Cache, excluding data writes, that modify persistent state."),
		// 										ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/ClientMetadataWriteIOPS"),
		// 										IsDimensionRequired: to.Ptr(false),
		// 										MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 											{
		// 												Retention: to.Ptr("P93D"),
		// 												TimeGrain: to.Ptr("PT1M"),
		// 											},
		// 											{
		// 												Retention: to.Ptr("P93D"),
		// 												TimeGrain: to.Ptr("PT5M"),
		// 											},
		// 											{
		// 												Retention: to.Ptr("P93D"),
		// 												TimeGrain: to.Ptr("PT15M"),
		// 											},
		// 											{
		// 												Retention: to.Ptr("P93D"),
		// 												TimeGrain: to.Ptr("PT30M"),
		// 											},
		// 											{
		// 												Retention: to.Ptr("P93D"),
		// 												TimeGrain: to.Ptr("PT1H"),
		// 											},
		// 											{
		// 												Retention: to.Ptr("P93D"),
		// 												TimeGrain: to.Ptr("PT6H"),
		// 											},
		// 											{
		// 												Retention: to.Ptr("P93D"),
		// 												TimeGrain: to.Ptr("PT12H"),
		// 											},
		// 											{
		// 												Retention: to.Ptr("P93D"),
		// 												TimeGrain: to.Ptr("P1D"),
		// 										}},
		// 										MetricClass: to.Ptr(armmonitor.MetricClassTransactions),
		// 										Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 										PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 										ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 										SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 											to.Ptr(armmonitor.AggregationTypeMinimum),
		// 											to.Ptr(armmonitor.AggregationTypeMaximum),
		// 											to.Ptr(armmonitor.AggregationTypeAverage)},
		// 											Unit: to.Ptr(armmonitor.MetricUnitCountPerSecond),
		// 										},
		// 										{
		// 											Name: &armmonitor.LocalizableString{
		// 												LocalizedValue: to.Ptr("Client Lock IOPS"),
		// 												Value: to.Ptr("ClientLockIOPS"),
		// 											},
		// 											DisplayDescription: to.Ptr("Client file locking operations per second."),
		// 											ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/ClientLockIOPS"),
		// 											IsDimensionRequired: to.Ptr(false),
		// 											MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 												{
		// 													Retention: to.Ptr("P93D"),
		// 													TimeGrain: to.Ptr("PT1M"),
		// 												},
		// 												{
		// 													Retention: to.Ptr("P93D"),
		// 													TimeGrain: to.Ptr("PT5M"),
		// 												},
		// 												{
		// 													Retention: to.Ptr("P93D"),
		// 													TimeGrain: to.Ptr("PT15M"),
		// 												},
		// 												{
		// 													Retention: to.Ptr("P93D"),
		// 													TimeGrain: to.Ptr("PT30M"),
		// 												},
		// 												{
		// 													Retention: to.Ptr("P93D"),
		// 													TimeGrain: to.Ptr("PT1H"),
		// 												},
		// 												{
		// 													Retention: to.Ptr("P93D"),
		// 													TimeGrain: to.Ptr("PT6H"),
		// 												},
		// 												{
		// 													Retention: to.Ptr("P93D"),
		// 													TimeGrain: to.Ptr("PT12H"),
		// 												},
		// 												{
		// 													Retention: to.Ptr("P93D"),
		// 													TimeGrain: to.Ptr("P1D"),
		// 											}},
		// 											MetricClass: to.Ptr(armmonitor.MetricClassTransactions),
		// 											Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 											PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 											ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 											SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 												to.Ptr(armmonitor.AggregationTypeMinimum),
		// 												to.Ptr(armmonitor.AggregationTypeMaximum),
		// 												to.Ptr(armmonitor.AggregationTypeAverage)},
		// 												Unit: to.Ptr(armmonitor.MetricUnitCountPerSecond),
		// 											},
		// 											{
		// 												Name: &armmonitor.LocalizableString{
		// 													LocalizedValue: to.Ptr("Storage Target Health"),
		// 													Value: to.Ptr("StorageTargetHealth"),
		// 												},
		// 												DisplayDescription: to.Ptr("Boolean results of connectivity test between the Cache and Storage Targets."),
		// 												ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/StorageTargetHealth"),
		// 												IsDimensionRequired: to.Ptr(false),
		// 												MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 													{
		// 														Retention: to.Ptr("P93D"),
		// 														TimeGrain: to.Ptr("PT1M"),
		// 													},
		// 													{
		// 														Retention: to.Ptr("P93D"),
		// 														TimeGrain: to.Ptr("PT5M"),
		// 													},
		// 													{
		// 														Retention: to.Ptr("P93D"),
		// 														TimeGrain: to.Ptr("PT15M"),
		// 													},
		// 													{
		// 														Retention: to.Ptr("P93D"),
		// 														TimeGrain: to.Ptr("PT30M"),
		// 													},
		// 													{
		// 														Retention: to.Ptr("P93D"),
		// 														TimeGrain: to.Ptr("PT1H"),
		// 													},
		// 													{
		// 														Retention: to.Ptr("P93D"),
		// 														TimeGrain: to.Ptr("PT6H"),
		// 													},
		// 													{
		// 														Retention: to.Ptr("P93D"),
		// 														TimeGrain: to.Ptr("PT12H"),
		// 													},
		// 													{
		// 														Retention: to.Ptr("P93D"),
		// 														TimeGrain: to.Ptr("P1D"),
		// 												}},
		// 												MetricClass: to.Ptr(armmonitor.MetricClassErrors),
		// 												Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 												PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 												ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 												SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 													to.Ptr(armmonitor.AggregationTypeMinimum),
		// 													to.Ptr(armmonitor.AggregationTypeMaximum),
		// 													to.Ptr(armmonitor.AggregationTypeAverage)},
		// 													Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 												},
		// 												{
		// 													Name: &armmonitor.LocalizableString{
		// 														LocalizedValue: to.Ptr("Uptime"),
		// 														Value: to.Ptr("Uptime"),
		// 													},
		// 													DisplayDescription: to.Ptr("Boolean results of connectivity test between the Cache and monitoring system."),
		// 													ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/Uptime"),
		// 													IsDimensionRequired: to.Ptr(false),
		// 													MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 														{
		// 															Retention: to.Ptr("P93D"),
		// 															TimeGrain: to.Ptr("PT1M"),
		// 														},
		// 														{
		// 															Retention: to.Ptr("P93D"),
		// 															TimeGrain: to.Ptr("PT5M"),
		// 														},
		// 														{
		// 															Retention: to.Ptr("P93D"),
		// 															TimeGrain: to.Ptr("PT15M"),
		// 														},
		// 														{
		// 															Retention: to.Ptr("P93D"),
		// 															TimeGrain: to.Ptr("PT30M"),
		// 														},
		// 														{
		// 															Retention: to.Ptr("P93D"),
		// 															TimeGrain: to.Ptr("PT1H"),
		// 														},
		// 														{
		// 															Retention: to.Ptr("P93D"),
		// 															TimeGrain: to.Ptr("PT6H"),
		// 														},
		// 														{
		// 															Retention: to.Ptr("P93D"),
		// 															TimeGrain: to.Ptr("PT12H"),
		// 														},
		// 														{
		// 															Retention: to.Ptr("P93D"),
		// 															TimeGrain: to.Ptr("P1D"),
		// 													}},
		// 													MetricClass: to.Ptr(armmonitor.MetricClassAvailability),
		// 													Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 													PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 													ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 													SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 														to.Ptr(armmonitor.AggregationTypeMinimum),
		// 														to.Ptr(armmonitor.AggregationTypeMaximum),
		// 														to.Ptr(armmonitor.AggregationTypeAverage)},
		// 														Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 													},
		// 													{
		// 														Name: &armmonitor.LocalizableString{
		// 															LocalizedValue: to.Ptr("Total StorageTarget IOPS"),
		// 															Value: to.Ptr("StorageTargetIOPS"),
		// 														},
		// 														Dimensions: []*armmonitor.LocalizableString{
		// 															{
		// 																LocalizedValue: to.Ptr("StorageTarget"),
		// 																Value: to.Ptr("StorageTarget"),
		// 														}},
		// 														DisplayDescription: to.Ptr("The rate of all file operations the Cache sends to a particular StorageTarget."),
		// 														ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/StorageTargetIOPS"),
		// 														IsDimensionRequired: to.Ptr(false),
		// 														MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 															{
		// 																Retention: to.Ptr("P93D"),
		// 																TimeGrain: to.Ptr("PT1M"),
		// 															},
		// 															{
		// 																Retention: to.Ptr("P93D"),
		// 																TimeGrain: to.Ptr("PT5M"),
		// 															},
		// 															{
		// 																Retention: to.Ptr("P93D"),
		// 																TimeGrain: to.Ptr("PT15M"),
		// 															},
		// 															{
		// 																Retention: to.Ptr("P93D"),
		// 																TimeGrain: to.Ptr("PT30M"),
		// 															},
		// 															{
		// 																Retention: to.Ptr("P93D"),
		// 																TimeGrain: to.Ptr("PT1H"),
		// 															},
		// 															{
		// 																Retention: to.Ptr("P93D"),
		// 																TimeGrain: to.Ptr("PT6H"),
		// 															},
		// 															{
		// 																Retention: to.Ptr("P93D"),
		// 																TimeGrain: to.Ptr("PT12H"),
		// 															},
		// 															{
		// 																Retention: to.Ptr("P93D"),
		// 																TimeGrain: to.Ptr("P1D"),
		// 														}},
		// 														MetricClass: to.Ptr(armmonitor.MetricClassTransactions),
		// 														Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 														PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 														ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 														SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 															to.Ptr(armmonitor.AggregationTypeMinimum),
		// 															to.Ptr(armmonitor.AggregationTypeMaximum),
		// 															to.Ptr(armmonitor.AggregationTypeAverage)},
		// 															Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 														},
		// 														{
		// 															Name: &armmonitor.LocalizableString{
		// 																LocalizedValue: to.Ptr("StorageTarget Write IOPS"),
		// 																Value: to.Ptr("StorageTargetWriteIOPS"),
		// 															},
		// 															Dimensions: []*armmonitor.LocalizableString{
		// 																{
		// 																	LocalizedValue: to.Ptr("StorageTarget"),
		// 																	Value: to.Ptr("StorageTarget"),
		// 															}},
		// 															DisplayDescription: to.Ptr("The rate of the file write operations the Cache sends to a particular StorageTarget."),
		// 															ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/StorageTargetWriteIOPS"),
		// 															IsDimensionRequired: to.Ptr(false),
		// 															MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																{
		// 																	Retention: to.Ptr("P93D"),
		// 																	TimeGrain: to.Ptr("PT1M"),
		// 																},
		// 																{
		// 																	Retention: to.Ptr("P93D"),
		// 																	TimeGrain: to.Ptr("PT5M"),
		// 																},
		// 																{
		// 																	Retention: to.Ptr("P93D"),
		// 																	TimeGrain: to.Ptr("PT15M"),
		// 																},
		// 																{
		// 																	Retention: to.Ptr("P93D"),
		// 																	TimeGrain: to.Ptr("PT30M"),
		// 																},
		// 																{
		// 																	Retention: to.Ptr("P93D"),
		// 																	TimeGrain: to.Ptr("PT1H"),
		// 																},
		// 																{
		// 																	Retention: to.Ptr("P93D"),
		// 																	TimeGrain: to.Ptr("PT6H"),
		// 																},
		// 																{
		// 																	Retention: to.Ptr("P93D"),
		// 																	TimeGrain: to.Ptr("PT12H"),
		// 																},
		// 																{
		// 																	Retention: to.Ptr("P93D"),
		// 																	TimeGrain: to.Ptr("P1D"),
		// 															}},
		// 															MetricClass: to.Ptr(armmonitor.MetricClassTransactions),
		// 															Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 															PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 															ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 															SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																to.Ptr(armmonitor.AggregationTypeMinimum),
		// 																to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																to.Ptr(armmonitor.AggregationTypeAverage)},
		// 																Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 															},
		// 															{
		// 																Name: &armmonitor.LocalizableString{
		// 																	LocalizedValue: to.Ptr("StorageTarget Asynchronous Write Throughput"),
		// 																	Value: to.Ptr("StorageTargetAsyncWriteThroughput"),
		// 																},
		// 																Dimensions: []*armmonitor.LocalizableString{
		// 																	{
		// 																		LocalizedValue: to.Ptr("StorageTarget"),
		// 																		Value: to.Ptr("StorageTarget"),
		// 																}},
		// 																DisplayDescription: to.Ptr("The rate the Cache asynchronously writes data to a particular StorageTarget. These are opportunistic writes that do not cause clients to block."),
		// 																ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/StorageTargetAsyncWriteThroughput"),
		// 																IsDimensionRequired: to.Ptr(false),
		// 																MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																	{
		// 																		Retention: to.Ptr("P93D"),
		// 																		TimeGrain: to.Ptr("PT1M"),
		// 																	},
		// 																	{
		// 																		Retention: to.Ptr("P93D"),
		// 																		TimeGrain: to.Ptr("PT5M"),
		// 																	},
		// 																	{
		// 																		Retention: to.Ptr("P93D"),
		// 																		TimeGrain: to.Ptr("PT15M"),
		// 																	},
		// 																	{
		// 																		Retention: to.Ptr("P93D"),
		// 																		TimeGrain: to.Ptr("PT30M"),
		// 																	},
		// 																	{
		// 																		Retention: to.Ptr("P93D"),
		// 																		TimeGrain: to.Ptr("PT1H"),
		// 																	},
		// 																	{
		// 																		Retention: to.Ptr("P93D"),
		// 																		TimeGrain: to.Ptr("PT6H"),
		// 																	},
		// 																	{
		// 																		Retention: to.Ptr("P93D"),
		// 																		TimeGrain: to.Ptr("PT12H"),
		// 																	},
		// 																	{
		// 																		Retention: to.Ptr("P93D"),
		// 																		TimeGrain: to.Ptr("P1D"),
		// 																}},
		// 																MetricClass: to.Ptr(armmonitor.MetricClassTransactions),
		// 																Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 																PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 																SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																	to.Ptr(armmonitor.AggregationTypeMinimum),
		// 																	to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																	to.Ptr(armmonitor.AggregationTypeAverage)},
		// 																	Unit: to.Ptr(armmonitor.MetricUnitBytesPerSecond),
		// 																},
		// 																{
		// 																	Name: &armmonitor.LocalizableString{
		// 																		LocalizedValue: to.Ptr("StorageTarget Synchronous Write Throughput"),
		// 																		Value: to.Ptr("StorageTargetSyncWriteThroughput"),
		// 																	},
		// 																	Dimensions: []*armmonitor.LocalizableString{
		// 																		{
		// 																			LocalizedValue: to.Ptr("StorageTarget"),
		// 																			Value: to.Ptr("StorageTarget"),
		// 																	}},
		// 																	DisplayDescription: to.Ptr("The rate the Cache synchronously writes data to a particular StorageTarget. These are writes that do cause clients to block."),
		// 																	ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/StorageTargetSyncWriteThroughput"),
		// 																	IsDimensionRequired: to.Ptr(false),
		// 																	MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																		{
		// 																			Retention: to.Ptr("P93D"),
		// 																			TimeGrain: to.Ptr("PT1M"),
		// 																		},
		// 																		{
		// 																			Retention: to.Ptr("P93D"),
		// 																			TimeGrain: to.Ptr("PT5M"),
		// 																		},
		// 																		{
		// 																			Retention: to.Ptr("P93D"),
		// 																			TimeGrain: to.Ptr("PT15M"),
		// 																		},
		// 																		{
		// 																			Retention: to.Ptr("P93D"),
		// 																			TimeGrain: to.Ptr("PT30M"),
		// 																		},
		// 																		{
		// 																			Retention: to.Ptr("P93D"),
		// 																			TimeGrain: to.Ptr("PT1H"),
		// 																		},
		// 																		{
		// 																			Retention: to.Ptr("P93D"),
		// 																			TimeGrain: to.Ptr("PT6H"),
		// 																		},
		// 																		{
		// 																			Retention: to.Ptr("P93D"),
		// 																			TimeGrain: to.Ptr("PT12H"),
		// 																		},
		// 																		{
		// 																			Retention: to.Ptr("P93D"),
		// 																			TimeGrain: to.Ptr("P1D"),
		// 																	}},
		// 																	MetricClass: to.Ptr(armmonitor.MetricClassTransactions),
		// 																	Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 																	PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																	ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 																	SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																		to.Ptr(armmonitor.AggregationTypeMinimum),
		// 																		to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																		to.Ptr(armmonitor.AggregationTypeAverage)},
		// 																		Unit: to.Ptr(armmonitor.MetricUnitBytesPerSecond),
		// 																	},
		// 																	{
		// 																		Name: &armmonitor.LocalizableString{
		// 																			LocalizedValue: to.Ptr("StorageTarget Total Write Throughput"),
		// 																			Value: to.Ptr("StorageTargetTotalWriteThroughput"),
		// 																		},
		// 																		Dimensions: []*armmonitor.LocalizableString{
		// 																			{
		// 																				LocalizedValue: to.Ptr("StorageTarget"),
		// 																				Value: to.Ptr("StorageTarget"),
		// 																		}},
		// 																		DisplayDescription: to.Ptr("The total rate that the Cache writes data to a particular StorageTarget."),
		// 																		ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/StorageTargetTotalWriteThroughput"),
		// 																		IsDimensionRequired: to.Ptr(false),
		// 																		MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																			{
		// 																				Retention: to.Ptr("P93D"),
		// 																				TimeGrain: to.Ptr("PT1M"),
		// 																			},
		// 																			{
		// 																				Retention: to.Ptr("P93D"),
		// 																				TimeGrain: to.Ptr("PT5M"),
		// 																			},
		// 																			{
		// 																				Retention: to.Ptr("P93D"),
		// 																				TimeGrain: to.Ptr("PT15M"),
		// 																			},
		// 																			{
		// 																				Retention: to.Ptr("P93D"),
		// 																				TimeGrain: to.Ptr("PT30M"),
		// 																			},
		// 																			{
		// 																				Retention: to.Ptr("P93D"),
		// 																				TimeGrain: to.Ptr("PT1H"),
		// 																			},
		// 																			{
		// 																				Retention: to.Ptr("P93D"),
		// 																				TimeGrain: to.Ptr("PT6H"),
		// 																			},
		// 																			{
		// 																				Retention: to.Ptr("P93D"),
		// 																				TimeGrain: to.Ptr("PT12H"),
		// 																			},
		// 																			{
		// 																				Retention: to.Ptr("P93D"),
		// 																				TimeGrain: to.Ptr("P1D"),
		// 																		}},
		// 																		MetricClass: to.Ptr(armmonitor.MetricClassTransactions),
		// 																		Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 																		PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																		ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 																		SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																			to.Ptr(armmonitor.AggregationTypeMinimum),
		// 																			to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																			to.Ptr(armmonitor.AggregationTypeAverage)},
		// 																			Unit: to.Ptr(armmonitor.MetricUnitBytesPerSecond),
		// 																		},
		// 																		{
		// 																			Name: &armmonitor.LocalizableString{
		// 																				LocalizedValue: to.Ptr("StorageTarget Latency"),
		// 																				Value: to.Ptr("StorageTargetLatency"),
		// 																			},
		// 																			Dimensions: []*armmonitor.LocalizableString{
		// 																				{
		// 																					LocalizedValue: to.Ptr("StorageTarget"),
		// 																					Value: to.Ptr("StorageTarget"),
		// 																			}},
		// 																			DisplayDescription: to.Ptr("The average round trip latency of all the file operations the Cache sends to a partricular StorageTarget."),
		// 																			ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/StorageTargetLatency"),
		// 																			IsDimensionRequired: to.Ptr(false),
		// 																			MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																				{
		// 																					Retention: to.Ptr("P93D"),
		// 																					TimeGrain: to.Ptr("PT1M"),
		// 																				},
		// 																				{
		// 																					Retention: to.Ptr("P93D"),
		// 																					TimeGrain: to.Ptr("PT5M"),
		// 																				},
		// 																				{
		// 																					Retention: to.Ptr("P93D"),
		// 																					TimeGrain: to.Ptr("PT15M"),
		// 																				},
		// 																				{
		// 																					Retention: to.Ptr("P93D"),
		// 																					TimeGrain: to.Ptr("PT30M"),
		// 																				},
		// 																				{
		// 																					Retention: to.Ptr("P93D"),
		// 																					TimeGrain: to.Ptr("PT1H"),
		// 																				},
		// 																				{
		// 																					Retention: to.Ptr("P93D"),
		// 																					TimeGrain: to.Ptr("PT6H"),
		// 																				},
		// 																				{
		// 																					Retention: to.Ptr("P93D"),
		// 																					TimeGrain: to.Ptr("PT12H"),
		// 																				},
		// 																				{
		// 																					Retention: to.Ptr("P93D"),
		// 																					TimeGrain: to.Ptr("P1D"),
		// 																			}},
		// 																			MetricClass: to.Ptr(armmonitor.MetricClassTransactions),
		// 																			Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 																			PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																			ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 																			SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																				to.Ptr(armmonitor.AggregationTypeMinimum),
		// 																				to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																				to.Ptr(armmonitor.AggregationTypeAverage)},
		// 																				Unit: to.Ptr(armmonitor.MetricUnitMilliSeconds),
		// 																			},
		// 																			{
		// 																				Name: &armmonitor.LocalizableString{
		// 																					LocalizedValue: to.Ptr("StorageTarget Metadata Read IOPS"),
		// 																					Value: to.Ptr("StorageTargetMetadataReadIOPS"),
		// 																				},
		// 																				Dimensions: []*armmonitor.LocalizableString{
		// 																					{
		// 																						LocalizedValue: to.Ptr("StorageTarget"),
		// 																						Value: to.Ptr("StorageTarget"),
		// 																				}},
		// 																				DisplayDescription: to.Ptr("The rate of file operations that do not modify persistent state, and excluding the read operation, that the Cache sends to a particular StorageTarget."),
		// 																				ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/StorageTargetMetadataReadIOPS"),
		// 																				IsDimensionRequired: to.Ptr(false),
		// 																				MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																					{
		// 																						Retention: to.Ptr("P93D"),
		// 																						TimeGrain: to.Ptr("PT1M"),
		// 																					},
		// 																					{
		// 																						Retention: to.Ptr("P93D"),
		// 																						TimeGrain: to.Ptr("PT5M"),
		// 																					},
		// 																					{
		// 																						Retention: to.Ptr("P93D"),
		// 																						TimeGrain: to.Ptr("PT15M"),
		// 																					},
		// 																					{
		// 																						Retention: to.Ptr("P93D"),
		// 																						TimeGrain: to.Ptr("PT30M"),
		// 																					},
		// 																					{
		// 																						Retention: to.Ptr("P93D"),
		// 																						TimeGrain: to.Ptr("PT1H"),
		// 																					},
		// 																					{
		// 																						Retention: to.Ptr("P93D"),
		// 																						TimeGrain: to.Ptr("PT6H"),
		// 																					},
		// 																					{
		// 																						Retention: to.Ptr("P93D"),
		// 																						TimeGrain: to.Ptr("PT12H"),
		// 																					},
		// 																					{
		// 																						Retention: to.Ptr("P93D"),
		// 																						TimeGrain: to.Ptr("P1D"),
		// 																				}},
		// 																				MetricClass: to.Ptr(armmonitor.MetricClassTransactions),
		// 																				Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 																				PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																				ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 																				SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																					to.Ptr(armmonitor.AggregationTypeMinimum),
		// 																					to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																					to.Ptr(armmonitor.AggregationTypeAverage)},
		// 																					Unit: to.Ptr(armmonitor.MetricUnitCountPerSecond),
		// 																				},
		// 																				{
		// 																					Name: &armmonitor.LocalizableString{
		// 																						LocalizedValue: to.Ptr("StorageTarget Metadata Write IOPS"),
		// 																						Value: to.Ptr("StorageTargetMetadataWriteIOPS"),
		// 																					},
		// 																					Dimensions: []*armmonitor.LocalizableString{
		// 																						{
		// 																							LocalizedValue: to.Ptr("StorageTarget"),
		// 																							Value: to.Ptr("StorageTarget"),
		// 																					}},
		// 																					DisplayDescription: to.Ptr("The rate of file operations that do modify persistent state and excluding the write operation, that the Cache sends to a particular StorageTarget."),
		// 																					ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/StorageTargetMetadataWriteIOPS"),
		// 																					IsDimensionRequired: to.Ptr(false),
		// 																					MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																						{
		// 																							Retention: to.Ptr("P93D"),
		// 																							TimeGrain: to.Ptr("PT1M"),
		// 																						},
		// 																						{
		// 																							Retention: to.Ptr("P93D"),
		// 																							TimeGrain: to.Ptr("PT5M"),
		// 																						},
		// 																						{
		// 																							Retention: to.Ptr("P93D"),
		// 																							TimeGrain: to.Ptr("PT15M"),
		// 																						},
		// 																						{
		// 																							Retention: to.Ptr("P93D"),
		// 																							TimeGrain: to.Ptr("PT30M"),
		// 																						},
		// 																						{
		// 																							Retention: to.Ptr("P93D"),
		// 																							TimeGrain: to.Ptr("PT1H"),
		// 																						},
		// 																						{
		// 																							Retention: to.Ptr("P93D"),
		// 																							TimeGrain: to.Ptr("PT6H"),
		// 																						},
		// 																						{
		// 																							Retention: to.Ptr("P93D"),
		// 																							TimeGrain: to.Ptr("PT12H"),
		// 																						},
		// 																						{
		// 																							Retention: to.Ptr("P93D"),
		// 																							TimeGrain: to.Ptr("P1D"),
		// 																					}},
		// 																					MetricClass: to.Ptr(armmonitor.MetricClassTransactions),
		// 																					Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 																					PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																					ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 																					SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																						to.Ptr(armmonitor.AggregationTypeMinimum),
		// 																						to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																						to.Ptr(armmonitor.AggregationTypeAverage)},
		// 																						Unit: to.Ptr(armmonitor.MetricUnitCountPerSecond),
		// 																					},
		// 																					{
		// 																						Name: &armmonitor.LocalizableString{
		// 																							LocalizedValue: to.Ptr("StorageTarget Read IOPS"),
		// 																							Value: to.Ptr("StorageTargetReadIOPS"),
		// 																						},
		// 																						Dimensions: []*armmonitor.LocalizableString{
		// 																							{
		// 																								LocalizedValue: to.Ptr("StorageTarget"),
		// 																								Value: to.Ptr("StorageTarget"),
		// 																						}},
		// 																						DisplayDescription: to.Ptr("The rate of file read operations the Cache sends to a particular StorageTarget."),
		// 																						ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/StorageTargetReadIOPS"),
		// 																						IsDimensionRequired: to.Ptr(false),
		// 																						MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																							{
		// 																								Retention: to.Ptr("P93D"),
		// 																								TimeGrain: to.Ptr("PT1M"),
		// 																							},
		// 																							{
		// 																								Retention: to.Ptr("P93D"),
		// 																								TimeGrain: to.Ptr("PT5M"),
		// 																							},
		// 																							{
		// 																								Retention: to.Ptr("P93D"),
		// 																								TimeGrain: to.Ptr("PT15M"),
		// 																							},
		// 																							{
		// 																								Retention: to.Ptr("P93D"),
		// 																								TimeGrain: to.Ptr("PT30M"),
		// 																							},
		// 																							{
		// 																								Retention: to.Ptr("P93D"),
		// 																								TimeGrain: to.Ptr("PT1H"),
		// 																							},
		// 																							{
		// 																								Retention: to.Ptr("P93D"),
		// 																								TimeGrain: to.Ptr("PT6H"),
		// 																							},
		// 																							{
		// 																								Retention: to.Ptr("P93D"),
		// 																								TimeGrain: to.Ptr("PT12H"),
		// 																							},
		// 																							{
		// 																								Retention: to.Ptr("P93D"),
		// 																								TimeGrain: to.Ptr("P1D"),
		// 																						}},
		// 																						MetricClass: to.Ptr(armmonitor.MetricClassTransactions),
		// 																						Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 																						PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																						ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 																						SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																							to.Ptr(armmonitor.AggregationTypeMinimum),
		// 																							to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																							to.Ptr(armmonitor.AggregationTypeAverage)},
		// 																							Unit: to.Ptr(armmonitor.MetricUnitCountPerSecond),
		// 																						},
		// 																						{
		// 																							Name: &armmonitor.LocalizableString{
		// 																								LocalizedValue: to.Ptr("StorageTarget Read Ahead Throughput"),
		// 																								Value: to.Ptr("StorageTargetReadAheadThroughput"),
		// 																							},
		// 																							Dimensions: []*armmonitor.LocalizableString{
		// 																								{
		// 																									LocalizedValue: to.Ptr("StorageTarget"),
		// 																									Value: to.Ptr("StorageTarget"),
		// 																							}},
		// 																							DisplayDescription: to.Ptr("The rate the Cache opportunisticly reads data from the StorageTarget."),
		// 																							ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/StorageTargetReadAheadThroughput"),
		// 																							IsDimensionRequired: to.Ptr(false),
		// 																							MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																								{
		// 																									Retention: to.Ptr("P93D"),
		// 																									TimeGrain: to.Ptr("PT1M"),
		// 																								},
		// 																								{
		// 																									Retention: to.Ptr("P93D"),
		// 																									TimeGrain: to.Ptr("PT5M"),
		// 																								},
		// 																								{
		// 																									Retention: to.Ptr("P93D"),
		// 																									TimeGrain: to.Ptr("PT15M"),
		// 																								},
		// 																								{
		// 																									Retention: to.Ptr("P93D"),
		// 																									TimeGrain: to.Ptr("PT30M"),
		// 																								},
		// 																								{
		// 																									Retention: to.Ptr("P93D"),
		// 																									TimeGrain: to.Ptr("PT1H"),
		// 																								},
		// 																								{
		// 																									Retention: to.Ptr("P93D"),
		// 																									TimeGrain: to.Ptr("PT6H"),
		// 																								},
		// 																								{
		// 																									Retention: to.Ptr("P93D"),
		// 																									TimeGrain: to.Ptr("PT12H"),
		// 																								},
		// 																								{
		// 																									Retention: to.Ptr("P93D"),
		// 																									TimeGrain: to.Ptr("P1D"),
		// 																							}},
		// 																							MetricClass: to.Ptr(armmonitor.MetricClassTransactions),
		// 																							Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 																							PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																							ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 																							SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																								to.Ptr(armmonitor.AggregationTypeMinimum),
		// 																								to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																								to.Ptr(armmonitor.AggregationTypeAverage)},
		// 																								Unit: to.Ptr(armmonitor.MetricUnitBytesPerSecond),
		// 																							},
		// 																							{
		// 																								Name: &armmonitor.LocalizableString{
		// 																									LocalizedValue: to.Ptr("StorageTarget Fill Throughput"),
		// 																									Value: to.Ptr("StorageTargetFillThroughput"),
		// 																								},
		// 																								Dimensions: []*armmonitor.LocalizableString{
		// 																									{
		// 																										LocalizedValue: to.Ptr("StorageTarget"),
		// 																										Value: to.Ptr("StorageTarget"),
		// 																								}},
		// 																								DisplayDescription: to.Ptr("The rate the Cache reads data from the StorageTarget to handle a cache miss."),
		// 																								ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/StorageTargetFillThroughput"),
		// 																								IsDimensionRequired: to.Ptr(false),
		// 																								MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																									{
		// 																										Retention: to.Ptr("P93D"),
		// 																										TimeGrain: to.Ptr("PT1M"),
		// 																									},
		// 																									{
		// 																										Retention: to.Ptr("P93D"),
		// 																										TimeGrain: to.Ptr("PT5M"),
		// 																									},
		// 																									{
		// 																										Retention: to.Ptr("P93D"),
		// 																										TimeGrain: to.Ptr("PT15M"),
		// 																									},
		// 																									{
		// 																										Retention: to.Ptr("P93D"),
		// 																										TimeGrain: to.Ptr("PT30M"),
		// 																									},
		// 																									{
		// 																										Retention: to.Ptr("P93D"),
		// 																										TimeGrain: to.Ptr("PT1H"),
		// 																									},
		// 																									{
		// 																										Retention: to.Ptr("P93D"),
		// 																										TimeGrain: to.Ptr("PT6H"),
		// 																									},
		// 																									{
		// 																										Retention: to.Ptr("P93D"),
		// 																										TimeGrain: to.Ptr("PT12H"),
		// 																									},
		// 																									{
		// 																										Retention: to.Ptr("P93D"),
		// 																										TimeGrain: to.Ptr("P1D"),
		// 																								}},
		// 																								MetricClass: to.Ptr(armmonitor.MetricClassTransactions),
		// 																								Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 																								PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																								ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 																								SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																									to.Ptr(armmonitor.AggregationTypeMinimum),
		// 																									to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																									to.Ptr(armmonitor.AggregationTypeAverage)},
		// 																									Unit: to.Ptr(armmonitor.MetricUnitBytesPerSecond),
		// 																								},
		// 																								{
		// 																									Name: &armmonitor.LocalizableString{
		// 																										LocalizedValue: to.Ptr("StorageTarget Total Read Throughput"),
		// 																										Value: to.Ptr("StorageTargetTotalReadThroughput"),
		// 																									},
		// 																									Dimensions: []*armmonitor.LocalizableString{
		// 																										{
		// 																											LocalizedValue: to.Ptr("StorageTarget"),
		// 																											Value: to.Ptr("StorageTarget"),
		// 																									}},
		// 																									DisplayDescription: to.Ptr("The total rate that the Cache reads data from a particular StorageTarget."),
		// 																									ID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache/providers/microsoft.insights/metricdefinitions/StorageTargetTotalReadThroughput"),
		// 																									IsDimensionRequired: to.Ptr(false),
		// 																									MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																										{
		// 																											Retention: to.Ptr("P93D"),
		// 																											TimeGrain: to.Ptr("PT1M"),
		// 																										},
		// 																										{
		// 																											Retention: to.Ptr("P93D"),
		// 																											TimeGrain: to.Ptr("PT5M"),
		// 																										},
		// 																										{
		// 																											Retention: to.Ptr("P93D"),
		// 																											TimeGrain: to.Ptr("PT15M"),
		// 																										},
		// 																										{
		// 																											Retention: to.Ptr("P93D"),
		// 																											TimeGrain: to.Ptr("PT30M"),
		// 																										},
		// 																										{
		// 																											Retention: to.Ptr("P93D"),
		// 																											TimeGrain: to.Ptr("PT1H"),
		// 																										},
		// 																										{
		// 																											Retention: to.Ptr("P93D"),
		// 																											TimeGrain: to.Ptr("PT6H"),
		// 																										},
		// 																										{
		// 																											Retention: to.Ptr("P93D"),
		// 																											TimeGrain: to.Ptr("PT12H"),
		// 																										},
		// 																										{
		// 																											Retention: to.Ptr("P93D"),
		// 																											TimeGrain: to.Ptr("P1D"),
		// 																									}},
		// 																									MetricClass: to.Ptr(armmonitor.MetricClassTransactions),
		// 																									Namespace: to.Ptr("microsoft.storagecache/caches"),
		// 																									PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																									ResourceID: to.Ptr("/subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache"),
		// 																									SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																										to.Ptr(armmonitor.AggregationTypeMinimum),
		// 																										to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																										to.Ptr(armmonitor.AggregationTypeAverage)},
		// 																										Unit: to.Ptr(armmonitor.MetricUnitBytesPerSecond),
		// 																								}},
		// 																							}
	}
}
