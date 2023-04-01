package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DevicesListMetricDefinition.json
func ExampleDevicesClient_NewListMetricDefinitionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDevicesClient().NewListMetricDefinitionPager("Device05ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
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
		// page.MetricDefinitionList = armstorsimple8000series.MetricDefinitionList{
		// 	Value: []*armstorsimple8000series.MetricDefinition{
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("Primary Tiered Storage Used"),
		// 				Value: to.Ptr("PrimaryStorageTieredUsed"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions"),
		// 			Category: to.Ptr("CapacityUtilization"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Device"),
		// 					Value: to.Ptr("Device05ForSDKTest"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple8000series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("P1D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
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
		// 			PrimaryAggregationType: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitBytes),
		// 		},
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("Primary Locally Pinned Storage Used"),
		// 				Value: to.Ptr("PrimaryStorageLocallyPinnedUsed"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions"),
		// 			Category: to.Ptr("CapacityUtilization"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Device"),
		// 					Value: to.Ptr("Device05ForSDKTest"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple8000series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("P1D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
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
		// 			PrimaryAggregationType: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitBytes),
		// 		},
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("Cloud Storage Used"),
		// 				Value: to.Ptr("CloudStorageUsed"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions"),
		// 			Category: to.Ptr("CapacityUtilization"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Device"),
		// 					Value: to.Ptr("Device05ForSDKTest"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple8000series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("P1D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
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
		// 			PrimaryAggregationType: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitBytes),
		// 		},
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("Local Storage Used"),
		// 				Value: to.Ptr("LocalStorageUsed"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions"),
		// 			Category: to.Ptr("CapacityUtilization"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Device"),
		// 					Value: to.Ptr("Device05ForSDKTest"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple8000series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("P1D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
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
		// 			PrimaryAggregationType: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitBytes),
		// 		},
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("Read I/O Operations/s"),
		// 				Value: to.Ptr("InitiatorToDeviceReadOperations"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions"),
		// 			Category: to.Ptr("InitiatorToDeviceIOPerformance"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Device"),
		// 					Value: to.Ptr("Device05ForSDKTest"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple8000series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("PT6H"),
		// 					TimeGrain: to.Ptr("PT1M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
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
		// 			PrimaryAggregationType: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitCount),
		// 		},
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("Write I/O Operations/s"),
		// 				Value: to.Ptr("InitiatorToDeviceWriteOperations"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions"),
		// 			Category: to.Ptr("InitiatorToDeviceIOPerformance"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Device"),
		// 					Value: to.Ptr("Device05ForSDKTest"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple8000series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("PT6H"),
		// 					TimeGrain: to.Ptr("PT1M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
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
		// 			PrimaryAggregationType: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitCount),
		// 		},
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("Read Bytes/s"),
		// 				Value: to.Ptr("InitiatorToDeviceReadBytes"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions"),
		// 			Category: to.Ptr("InitiatorToDeviceIOPerformance"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Device"),
		// 					Value: to.Ptr("Device05ForSDKTest"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple8000series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("PT6H"),
		// 					TimeGrain: to.Ptr("PT1M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
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
		// 			PrimaryAggregationType: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitBytesPerSecond),
		// 		},
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("Write Bytes/s"),
		// 				Value: to.Ptr("InitiatorToDeviceWriteBytes"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions"),
		// 			Category: to.Ptr("InitiatorToDeviceIOPerformance"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Device"),
		// 					Value: to.Ptr("Device05ForSDKTest"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple8000series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("PT6H"),
		// 					TimeGrain: to.Ptr("PT1M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
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
		// 			PrimaryAggregationType: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitBytesPerSecond),
		// 		},
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("Read Latency"),
		// 				Value: to.Ptr("InitiatorToDeviceReadLatency"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions"),
		// 			Category: to.Ptr("InitiatorToDeviceIOPerformance"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Device"),
		// 					Value: to.Ptr("Device05ForSDKTest"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple8000series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("PT6H"),
		// 					TimeGrain: to.Ptr("PT1M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
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
		// 			PrimaryAggregationType: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitSeconds),
		// 		},
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("Write Latency"),
		// 				Value: to.Ptr("InitiatorToDeviceWriteLatency"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions"),
		// 			Category: to.Ptr("InitiatorToDeviceIOPerformance"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Device"),
		// 					Value: to.Ptr("Device05ForSDKTest"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple8000series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("PT6H"),
		// 					TimeGrain: to.Ptr("PT1M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
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
		// 			PrimaryAggregationType: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitSeconds),
		// 		},
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("Outstanding I/O"),
		// 				Value: to.Ptr("InitiatorToDeviceOutStandingIO"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions"),
		// 			Category: to.Ptr("InitiatorToDeviceIOPerformance"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Device"),
		// 					Value: to.Ptr("Device05ForSDKTest"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple8000series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("PT6H"),
		// 					TimeGrain: to.Ptr("PT1M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
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
		// 			PrimaryAggregationType: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitCount),
		// 		},
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("Read I/O Operations/s"),
		// 				Value: to.Ptr("DeviceToCloudReadOperations"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions"),
		// 			Category: to.Ptr("DeviceToCloudIOPerformance"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Device"),
		// 					Value: to.Ptr("Device05ForSDKTest"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple8000series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("PT6H"),
		// 					TimeGrain: to.Ptr("PT1M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
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
		// 			PrimaryAggregationType: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitCount),
		// 		},
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("Write I/O Operations/s"),
		// 				Value: to.Ptr("DeviceToCloudWriteOperations"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions"),
		// 			Category: to.Ptr("DeviceToCloudIOPerformance"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Device"),
		// 					Value: to.Ptr("Device05ForSDKTest"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple8000series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("PT6H"),
		// 					TimeGrain: to.Ptr("PT1M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
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
		// 			PrimaryAggregationType: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitCount),
		// 		},
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("Read Bytes/s"),
		// 				Value: to.Ptr("DeviceToCloudReadBytes"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions"),
		// 			Category: to.Ptr("DeviceToCloudIOPerformance"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Device"),
		// 					Value: to.Ptr("Device05ForSDKTest"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple8000series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("PT6H"),
		// 					TimeGrain: to.Ptr("PT1M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
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
		// 			PrimaryAggregationType: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitBytesPerSecond),
		// 		},
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("Write Bytes/s"),
		// 				Value: to.Ptr("DeviceToCloudWriteBytes"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions"),
		// 			Category: to.Ptr("DeviceToCloudIOPerformance"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Device"),
		// 					Value: to.Ptr("Device05ForSDKTest"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple8000series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("PT6H"),
		// 					TimeGrain: to.Ptr("PT1M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
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
		// 			PrimaryAggregationType: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitBytesPerSecond),
		// 		},
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("Read Latency"),
		// 				Value: to.Ptr("DeviceToCloudReadLatency"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions"),
		// 			Category: to.Ptr("DeviceToCloudIOPerformance"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Device"),
		// 					Value: to.Ptr("Device05ForSDKTest"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple8000series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("PT6H"),
		// 					TimeGrain: to.Ptr("PT1M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
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
		// 			PrimaryAggregationType: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitSeconds),
		// 		},
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("Write Latency"),
		// 				Value: to.Ptr("DeviceToCloudWriteLatency"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions"),
		// 			Category: to.Ptr("DeviceToCloudIOPerformance"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Device"),
		// 					Value: to.Ptr("Device05ForSDKTest"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple8000series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("PT6H"),
		// 					TimeGrain: to.Ptr("PT1M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
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
		// 			PrimaryAggregationType: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitSeconds),
		// 		},
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("Outstanding I/O"),
		// 				Value: to.Ptr("DeviceToCloudOutStandingIO"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions"),
		// 			Category: to.Ptr("DeviceToCloudIOPerformance"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Device"),
		// 					Value: to.Ptr("Device05ForSDKTest"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple8000series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("PT6H"),
		// 					TimeGrain: to.Ptr("PT1M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
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
		// 			PrimaryAggregationType: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitCount),
		// 		},
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("CPU Utilization"),
		// 				Value: to.Ptr("CPUUsage"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions"),
		// 			Category: to.Ptr("DevicePerformance"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Device"),
		// 					Value: to.Ptr("Device05ForSDKTest"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple8000series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("PT6H"),
		// 					TimeGrain: to.Ptr("PT1M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
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
		// 			PrimaryAggregationType: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitPercent),
		// 		},
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("Network Out"),
		// 				Value: to.Ptr("NetworkOut"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions"),
		// 			Category: to.Ptr("NetworkThroughput"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("NetworkInterface"),
		// 					Value: to.Ptr("data0"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple8000series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("P1D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
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
		// 			PrimaryAggregationType: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitBytesPerSecond),
		// 		},
		// 		{
		// 			Name: &armstorsimple8000series.MetricName{
		// 				LocalizedValue: to.Ptr("Network In"),
		// 				Value: to.Ptr("NetworkIn"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions"),
		// 			Category: to.Ptr("NetworkThroughput"),
		// 			Dimensions: []*armstorsimple8000series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("NetworkInterface"),
		// 					Value: to.Ptr("data0"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple8000series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("P1D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
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
		// 			PrimaryAggregationType: to.Ptr(armstorsimple8000series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple8000series.MetricUnitBytesPerSecond),
		// 	}},
		// }
	}
}
