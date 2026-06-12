package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/InferenceGroup/getDeltaModelsStatusAsync.json
func ExampleInferenceGroupsClient_GetDeltaModelsStatusAsync() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInferenceGroupsClient().GetDeltaModelsStatusAsync(ctx, "test-rg", "my-aml-workspace", "string", "string", armmachinelearning.DeltaModelStatusRequest{
		DeltaModels: []*string{
			to.Ptr("string"),
		},
		TargetBaseModel: to.Ptr("azureml://registries/test-registry/models/modelabc/versions/1"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmachinelearning.InferenceGroupsClientGetDeltaModelsStatusAsyncResponse{
	// 	DeltaModelStatusResponse: armmachinelearning.DeltaModelStatusResponse{
	// 		ActualInstanceCount: to.Ptr[int32](0),
	// 		DeltaModels: map[string][]*armmachinelearning.DeltaModelCurrentState{
	// 			"string": []*armmachinelearning.DeltaModelCurrentState{
	// 				{
	// 					Count: to.Ptr[int32](0),
	// 					SampleInstanceID: to.Ptr("string"),
	// 					Status: to.Ptr("string"),
	// 				},
	// 			},
	// 		},
	// 		ExpectedInstanceCount: to.Ptr[int32](0),
	// 		RevisionID: to.Ptr("string"),
	// 		TargetBaseModel: to.Ptr("azureml://registries/test-registry/models/modelabc/versions/1"),
	// 	},
	// }
}
