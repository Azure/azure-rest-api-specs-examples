package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/EnvironmentVersion/createOrUpdate.json
func ExampleEnvironmentVersionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewEnvironmentVersionsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"test-rg",
		"my-aml-workspace",
		"string",
		"string",
		armmachinelearning.EnvironmentVersionData{
			Properties: &armmachinelearning.EnvironmentVersionDetails{
				Description: to.Ptr("string"),
				Properties: map[string]*string{
					"string": to.Ptr("string"),
				},
				Tags: map[string]*string{
					"string": to.Ptr("string"),
				},
				IsAnonymous: to.Ptr(false),
				Build: &armmachinelearning.BuildContext{
					ContextURI:     to.Ptr("https://storage-account.blob.core.windows.net/azureml/DockerBuildContext/95ddede6b9b8c4e90472db3acd0a8d28/"),
					DockerfilePath: to.Ptr("prod/Dockerfile"),
				},
				CondaFile: to.Ptr("string"),
				Image:     to.Ptr("docker.io/tensorflow/serving:latest"),
				InferenceConfig: &armmachinelearning.InferenceContainerProperties{
					LivenessRoute: &armmachinelearning.Route{
						Path: to.Ptr("string"),
						Port: to.Ptr[int32](1),
					},
					ReadinessRoute: &armmachinelearning.Route{
						Path: to.Ptr("string"),
						Port: to.Ptr[int32](1),
					},
					ScoringRoute: &armmachinelearning.Route{
						Path: to.Ptr("string"),
						Port: to.Ptr[int32](1),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
