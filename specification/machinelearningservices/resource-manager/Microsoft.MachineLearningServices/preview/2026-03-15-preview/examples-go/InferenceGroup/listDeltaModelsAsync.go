package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/InferenceGroup/listDeltaModelsAsync.json
func ExampleInferenceGroupsClient_NewListDeltaModelsAsyncPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInferenceGroupsClient().NewListDeltaModelsAsyncPager("test-rg", "my-aml-workspace", "string", "string", armmachinelearning.DeltaModelListRequest{
		Count:           to.Ptr[int32](-1),
		SkipToken:       to.Ptr("string"),
		TargetBaseModel: to.Ptr("azureml://registries/test-registry/models/modelabc/versions/1"),
	}, nil)
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
		// page = armmachinelearning.InferenceGroupsClientListDeltaModelsAsyncResponse{
		// 	StringArmPaginatedResult: armmachinelearning.StringArmPaginatedResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/my-aml-workspace/inferenceGroups/testInferenceGroupName/listDeltaModels?api-version=2025-07-01-preview&$skip=2"),
		// 		Value: []*string{
		// 			to.Ptr("string"),
		// 		},
		// 	},
		// }
	}
}
