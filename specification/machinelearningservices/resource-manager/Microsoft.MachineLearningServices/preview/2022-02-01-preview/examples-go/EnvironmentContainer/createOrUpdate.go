package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/EnvironmentContainer/createOrUpdate.json
func ExampleEnvironmentContainersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewEnvironmentContainersClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"testrg123",
		"testworkspace",
		"testEnvironment",
		armmachinelearning.EnvironmentContainerData{
			Properties: &armmachinelearning.EnvironmentContainerDetails{
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
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
