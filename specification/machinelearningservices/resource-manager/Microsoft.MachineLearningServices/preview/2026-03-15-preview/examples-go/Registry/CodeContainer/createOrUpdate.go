package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Registry/CodeContainer/createOrUpdate.json
func ExampleRegistryCodeContainersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRegistryCodeContainersClient().BeginCreateOrUpdate(ctx, "testrg123", "testregistry", "testContainer", armmachinelearning.CodeContainer{
		Properties: &armmachinelearning.CodeContainerProperties{
			Description: to.Ptr("string"),
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
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
	// res = armmachinelearning.RegistryCodeContainersClientCreateOrUpdateResponse{
	// 	CodeContainer: armmachinelearning.CodeContainer{
	// 		Name: to.Ptr("testContainer"),
	// 		Type: to.Ptr("Microsoft.MachineLearningServices/registries/codes"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/registries/testregistry/codes/testContainer"),
	// 		Properties: &armmachinelearning.CodeContainerProperties{
	// 			Description: to.Ptr("string"),
	// 			Tags: map[string]*string{
	// 				"property1": to.Ptr("string"),
	// 				"property2": to.Ptr("string"),
	// 			},
	// 		},
	// 		SystemData: &armmachinelearning.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T12:00:00.000Z"); return t}()),
	// 			CreatedBy: to.Ptr("John Smith"),
	// 			CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T12:00:00.000Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("John Smith"),
	// 			LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
