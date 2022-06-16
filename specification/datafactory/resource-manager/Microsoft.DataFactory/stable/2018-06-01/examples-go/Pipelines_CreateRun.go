package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_CreateRun.json
func ExamplePipelinesClient_CreateRun() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPipelinesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateRun(ctx,
		"exampleResourceGroup",
		"exampleFactoryName",
		"examplePipeline",
		&armdatafactory.PipelinesClientCreateRunOptions{ReferencePipelineRunID: nil,
			IsRecovery:        nil,
			StartActivityName: nil,
			StartFromFailure:  nil,
			Parameters: map[string]interface{}{
				"OutputBlobNameList": []interface{}{
					"exampleoutput.csv",
				},
			},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
