package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Registry/EnvironmentContainer/createOrUpdate.json
func ExampleRegistryEnvironmentContainersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRegistryEnvironmentContainersClient().BeginCreateOrUpdate(ctx, "testrg123", "testregistry", "testEnvironment", armmachinelearning.EnvironmentContainer{
		Properties: &armmachinelearning.EnvironmentContainerProperties{
			Description: to.Ptr("string"),
			Properties: map[string]*string{
				"additionalProp1": to.Ptr("string"),
				"additionalProp2": to.Ptr("string"),
				"additionalProp3": to.Ptr("string"),
			},
			Tags: map[string]*string{
				"additionalProp1": to.Ptr("string"),
				"additionalProp2": to.Ptr("string"),
				"additionalProp3": to.Ptr("string"),
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
	// res = armmachinelearning.RegistryEnvironmentContainersClientCreateOrUpdateResponse{
	// 	EnvironmentContainer: armmachinelearning.EnvironmentContainer{
	// 		Name: to.Ptr("testEnvironment"),
	// 		Type: to.Ptr("Microsoft.MachineLearningServices/registries/environments"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/registries/testregistry/environments/testEnvironment"),
	// 		Properties: &armmachinelearning.EnvironmentContainerProperties{
	// 			Description: to.Ptr("string"),
	// 			Properties: map[string]*string{
	// 				"additionalProp1": to.Ptr("string"),
	// 				"additionalProp2": to.Ptr("string"),
	// 				"additionalProp3": to.Ptr("string"),
	// 			},
	// 			Tags: map[string]*string{
	// 				"additionalProp1": to.Ptr("string"),
	// 				"additionalProp2": to.Ptr("string"),
	// 				"additionalProp3": to.Ptr("string"),
	// 			},
	// 		},
	// 		SystemData: &armmachinelearning.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-04T03:39:11.300Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-04T03:39:11.300Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
