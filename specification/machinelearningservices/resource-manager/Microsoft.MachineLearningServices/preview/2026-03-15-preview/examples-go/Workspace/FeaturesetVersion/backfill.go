package armmachinelearning_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Workspace/FeaturesetVersion/backfill.json
func ExampleFeaturesetVersionsClient_BeginBackfill() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFeaturesetVersionsClient().BeginBackfill(ctx, "test-rg", "my-aml-workspace", "string", "string", armmachinelearning.FeaturesetVersionBackfillRequest{
		Description: to.Ptr("string"),
		DataAvailabilityStatus: []*armmachinelearning.DataAvailabilityStatus{
			to.Ptr(armmachinelearning.DataAvailabilityStatusNone),
		},
		DisplayName: to.Ptr("string"),
		FeatureWindow: &armmachinelearning.FeatureWindow{
			FeatureWindowEnd:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999+00:51"); return t }()),
			FeatureWindowStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999+00:51"); return t }()),
		},
		JobID: to.Ptr("string"),
		Resource: &armmachinelearning.MaterializationComputeResource{
			InstanceType: to.Ptr("string"),
		},
		SparkConfiguration: map[string]*string{
			"string": to.Ptr("string"),
		},
		Tags: map[string]*string{
			"string": to.Ptr("string"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmachinelearning.FeaturesetVersionsClientBackfillResponse{
	// 	FeaturesetVersionBackfillResponse: armmachinelearning.FeaturesetVersionBackfillResponse{
	// 		JobIDs: []*string{
	// 			to.Ptr("string"),
	// 			to.Ptr("string"),
	// 		},
	// 	},
	// }
}
