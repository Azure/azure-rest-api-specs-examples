package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/InferenceGroup/modifyDeltaModelsAsync.json
func ExampleInferenceGroupsClient_BeginModifyDeltaModelsAsync() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInferenceGroupsClient().BeginModifyDeltaModelsAsync(ctx, "test-rg", "my-aml-workspace", "string", "string", armmachinelearning.DeltaModelModifyRequest{
		AddDeltaModels: []*string{
			to.Ptr("string"),
		},
		RemoveDeltaModels: []*string{
			to.Ptr("string"),
		},
		TargetBaseModel: to.Ptr("azureml://registries/test-registry/models/modelabc/versions/1"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}
