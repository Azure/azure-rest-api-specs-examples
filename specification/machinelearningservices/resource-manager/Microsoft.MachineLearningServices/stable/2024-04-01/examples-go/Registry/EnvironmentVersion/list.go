package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Registry/EnvironmentVersion/list.json
func ExampleRegistryEnvironmentVersionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRegistryEnvironmentVersionsClient().NewListPager("test-rg", "my-aml-regsitry", "string", &armmachinelearning.RegistryEnvironmentVersionsClientListOptions{OrderBy: to.Ptr("string"),
		Top:          to.Ptr[int32](1),
		Skip:         nil,
		ListViewType: nil,
	})
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
		// page.EnvironmentVersionResourceArmPaginatedResult = armmachinelearning.EnvironmentVersionResourceArmPaginatedResult{
		// 	Value: []*armmachinelearning.EnvironmentVersion{
		// 		{
		// 			Name: to.Ptr("string"),
		// 			Type: to.Ptr("string"),
		// 			ID: to.Ptr("string"),
		// 			SystemData: &armmachinelearning.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
		// 			},
		// 			Properties: &armmachinelearning.EnvironmentVersionProperties{
		// 				Description: to.Ptr("string"),
		// 				Properties: map[string]*string{
		// 					"string": to.Ptr("string"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"string": to.Ptr("string"),
		// 				},
		// 				IsAnonymous: to.Ptr(false),
		// 				Build: &armmachinelearning.BuildContext{
		// 					ContextURI: to.Ptr("https://storage-account.blob.core.windows.net/azureml/DockerBuildContext/95ddede6b9b8c4e90472db3acd0a8d28/"),
		// 					DockerfilePath: to.Ptr("prod/Dockerfile"),
		// 				},
		// 				CondaFile: to.Ptr("string"),
		// 				EnvironmentType: to.Ptr(armmachinelearning.EnvironmentTypeCurated),
		// 				Image: to.Ptr("docker.io/tensorflow/serving:latest"),
		// 				InferenceConfig: &armmachinelearning.InferenceContainerProperties{
		// 					LivenessRoute: &armmachinelearning.Route{
		// 						Path: to.Ptr("string"),
		// 						Port: to.Ptr[int32](1),
		// 					},
		// 					ReadinessRoute: &armmachinelearning.Route{
		// 						Path: to.Ptr("string"),
		// 						Port: to.Ptr[int32](1),
		// 					},
		// 					ScoringRoute: &armmachinelearning.Route{
		// 						Path: to.Ptr("string"),
		// 						Port: to.Ptr[int32](1),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
