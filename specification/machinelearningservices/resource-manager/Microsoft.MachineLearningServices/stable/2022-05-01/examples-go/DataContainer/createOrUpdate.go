package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-05-01/examples/DataContainer/createOrUpdate.json
func ExampleDataContainersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewDataContainersClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"testrg123",
		"workspace123",
		"datacontainer123",
		armmachinelearning.DataContainer{
			Properties: &armmachinelearning.DataContainerProperties{
				Description: to.Ptr("string"),
				Properties: map[string]*string{
					"properties1": to.Ptr("value1"),
					"properties2": to.Ptr("value2"),
				},
				Tags: map[string]*string{
					"tag1": to.Ptr("value1"),
					"tag2": to.Ptr("value2"),
				},
				DataType: to.Ptr(armmachinelearning.DataType("UriFile")),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
