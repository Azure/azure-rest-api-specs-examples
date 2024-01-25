package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/StreamingJob_List_ByResourceGroup_NoExpand.json
func ExampleStreamingJobsClient_NewListByResourceGroupPager_listAllStreamingJobsInAResourceGroupAndDoNotUseTheExpandODataQueryParameter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStreamingJobsClient().NewListByResourceGroupPager("sjrg6936", &armstreamanalytics.StreamingJobsClientListByResourceGroupOptions{Expand: nil})
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
		// page.StreamingJobListResult = armstreamanalytics.StreamingJobListResult{
		// 	Value: []*armstreamanalytics.StreamingJob{
		// 		{
		// 			Name: to.Ptr("sj59"),
		// 			Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs"),
		// 			ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg6936/providers/Microsoft.StreamAnalytics/streamingjobs/sj59"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key3": to.Ptr("value3"),
		// 				"randomKey": to.Ptr("randomValue"),
		// 			},
		// 			Properties: &armstreamanalytics.StreamingJobProperties{
		// 				CompatibilityLevel: to.Ptr(armstreamanalytics.CompatibilityLevelOne0),
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-11T04:37:04.697Z"); return t}()),
		// 				DataLocale: to.Ptr("en-US"),
		// 				Etag: to.Ptr("3e6872bc-c9d0-45b6-91b6-da66f1773056"),
		// 				EventsLateArrivalMaxDelayInSeconds: to.Ptr[int32](13),
		// 				EventsOutOfOrderMaxDelayInSeconds: to.Ptr[int32](21),
		// 				EventsOutOfOrderPolicy: to.Ptr(armstreamanalytics.EventsOutOfOrderPolicyDrop),
		// 				JobID: to.Ptr("d53ecc3c-fcb0-485d-9caf-25e20fcb2061"),
		// 				JobState: to.Ptr("Created"),
		// 				OutputErrorPolicy: to.Ptr(armstreamanalytics.OutputErrorPolicyDrop),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SKU: &armstreamanalytics.SKU{
		// 					Name: to.Ptr(armstreamanalytics.SKUNameStandard),
		// 				},
		// 			},
		// 			SKU: &armstreamanalytics.SKU{
		// 				Name: to.Ptr(armstreamanalytics.SKUNameStandard),
		// 				Capacity: to.Ptr[int32](36),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sj69"),
		// 			Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs"),
		// 			ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg6936/providers/Microsoft.StreamAnalytics/streamingjobs/sj69"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key3": to.Ptr("value3"),
		// 				"randomKey": to.Ptr("randomValue"),
		// 			},
		// 			Properties: &armstreamanalytics.StreamingJobProperties{
		// 				CompatibilityLevel: to.Ptr(armstreamanalytics.CompatibilityLevelOne0),
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-11T04:38:04.697Z"); return t}()),
		// 				DataLocale: to.Ptr("en-US"),
		// 				Etag: to.Ptr("99538949-a164-4e2f-a991-40303e86024f"),
		// 				EventsLateArrivalMaxDelayInSeconds: to.Ptr[int32](6),
		// 				EventsOutOfOrderMaxDelayInSeconds: to.Ptr[int32](5),
		// 				EventsOutOfOrderPolicy: to.Ptr(armstreamanalytics.EventsOutOfOrderPolicyDrop),
		// 				JobID: to.Ptr("817b36cf-a161-4a9e-86f2-eb00b3566d88"),
		// 				JobState: to.Ptr("Created"),
		// 				OutputErrorPolicy: to.Ptr(armstreamanalytics.OutputErrorPolicyDrop),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SKU: &armstreamanalytics.SKU{
		// 					Name: to.Ptr(armstreamanalytics.SKUNameStandard),
		// 				},
		// 			},
		// 			SKU: &armstreamanalytics.SKU{
		// 				Name: to.Ptr(armstreamanalytics.SKUNameStandard),
		// 				Capacity: to.Ptr[int32](36),
		// 			},
		// 	}},
		// }
	}
}
