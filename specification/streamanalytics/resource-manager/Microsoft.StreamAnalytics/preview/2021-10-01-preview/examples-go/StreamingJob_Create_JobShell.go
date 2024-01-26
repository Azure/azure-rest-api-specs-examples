package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/StreamingJob_Create_JobShell.json
func ExampleStreamingJobsClient_BeginCreateOrReplace_createAStreamingJobShellAStreamingJobWithNoInputsOutputsTransformationOrFunctions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStreamingJobsClient().BeginCreateOrReplace(ctx, "sjrg6936", "sj59", armstreamanalytics.StreamingJob{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"key1":      to.Ptr("value1"),
			"key3":      to.Ptr("value3"),
			"randomKey": to.Ptr("randomValue"),
		},
		Properties: &armstreamanalytics.StreamingJobProperties{
			CompatibilityLevel:                 to.Ptr(armstreamanalytics.CompatibilityLevelOne0),
			DataLocale:                         to.Ptr("en-US"),
			EventsLateArrivalMaxDelayInSeconds: to.Ptr[int32](16),
			EventsOutOfOrderMaxDelayInSeconds:  to.Ptr[int32](5),
			EventsOutOfOrderPolicy:             to.Ptr(armstreamanalytics.EventsOutOfOrderPolicyDrop),
			Functions:                          []*armstreamanalytics.Function{},
			Inputs:                             []*armstreamanalytics.Input{},
			OutputErrorPolicy:                  to.Ptr(armstreamanalytics.OutputErrorPolicyDrop),
			Outputs:                            []*armstreamanalytics.Output{},
			SKU: &armstreamanalytics.SKU{
				Name: to.Ptr(armstreamanalytics.SKUNameStandard),
			},
		},
	}, &armstreamanalytics.StreamingJobsClientBeginCreateOrReplaceOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StreamingJob = armstreamanalytics.StreamingJob{
	// 	Name: to.Ptr("sj59"),
	// 	Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs"),
	// 	ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg6936/providers/Microsoft.StreamAnalytics/streamingjobs/sj59"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key3": to.Ptr("value3"),
	// 		"randomKey": to.Ptr("randomValue"),
	// 	},
	// 	Properties: &armstreamanalytics.StreamingJobProperties{
	// 		CompatibilityLevel: to.Ptr(armstreamanalytics.CompatibilityLevelOne0),
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-11T04:37:04.697Z"); return t}()),
	// 		DataLocale: to.Ptr("en-US"),
	// 		EventsLateArrivalMaxDelayInSeconds: to.Ptr[int32](16),
	// 		EventsOutOfOrderMaxDelayInSeconds: to.Ptr[int32](5),
	// 		EventsOutOfOrderPolicy: to.Ptr(armstreamanalytics.EventsOutOfOrderPolicyDrop),
	// 		Functions: []*armstreamanalytics.Function{
	// 		},
	// 		Inputs: []*armstreamanalytics.Input{
	// 		},
	// 		JobID: to.Ptr("d53ecc3c-fcb0-485d-9caf-25e20fcb2061"),
	// 		JobState: to.Ptr("Created"),
	// 		OutputErrorPolicy: to.Ptr(armstreamanalytics.OutputErrorPolicyDrop),
	// 		Outputs: []*armstreamanalytics.Output{
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SKU: &armstreamanalytics.SKU{
	// 			Name: to.Ptr(armstreamanalytics.SKUNameStandard),
	// 		},
	// 	},
	// }
}
