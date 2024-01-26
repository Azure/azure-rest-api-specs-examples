package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/StreamingJob_Create_UserAssignedIdentity.json
func ExampleStreamingJobsClient_BeginCreateOrReplace_createAStreamingJobShellAStreamingJobWithNoInputsOutputsTransformationOrFunctionsWithUserAssignedIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStreamingJobsClient().BeginCreateOrReplace(ctx, "sjrg", "sjName", armstreamanalytics.StreamingJob{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"key1":      to.Ptr("value1"),
			"key3":      to.Ptr("value3"),
			"randomKey": to.Ptr("randomValue"),
		},
		Identity: &armstreamanalytics.Identity{
			Type: to.Ptr("UserAssigned"),
			UserAssignedIdentities: map[string]any{
				"/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.ManagedIdentity/userAssignedIdentities/sdkIdentity": map[string]any{},
			},
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
	// 	Name: to.Ptr("sjName"),
	// 	Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs"),
	// 	ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg/providers/Microsoft.StreamAnalytics/streamingjobs/sjName"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key3": to.Ptr("value3"),
	// 		"randomKey": to.Ptr("randomValue"),
	// 	},
	// 	Identity: &armstreamanalytics.Identity{
	// 		Type: to.Ptr("UserAssigned"),
	// 		PrincipalID: to.Ptr("c10a9ec7-7136-441f-9e90-d17cd1a51b94"),
	// 		TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 		UserAssignedIdentities: map[string]any{
	// 			"/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.ManagedIdentity/userAssignedIdentities/sdkIdentity": map[string]any{
	// 			},
	// 		},
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
