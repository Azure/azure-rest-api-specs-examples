package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01a71545e82bb98b8137d3038150c436d46a59ed/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/PipelineRuns_Get.json
func ExamplePipelineRunsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPipelineRunsClient().Get(ctx, "exampleResourceGroup", "exampleFactoryName", "2f7fdb90-5df1-4b8e-ac2f-064cfa58202b", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PipelineRun = armdatafactory.PipelineRun{
	// 	AdditionalProperties: map[string]any{
	// 		"annotations": []any{
	// 		},
	// 	},
	// 	DurationInMs: to.Ptr[int32](28105),
	// 	InvokedBy: &armdatafactory.PipelineRunInvokedBy{
	// 		Name: to.Ptr("Manual"),
	// 		ID: to.Ptr("80a01654a9d34ad18b3fcac5d5d76b67"),
	// 	},
	// 	LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:38:12.731Z"); return t}()),
	// 	Message: to.Ptr(""),
	// 	Parameters: map[string]*string{
	// 		"OutputBlobNameList": to.Ptr("[\"exampleoutput.csv\"]"),
	// 	},
	// 	PipelineName: to.Ptr("examplePipeline"),
	// 	RunEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:38:12.731Z"); return t}()),
	// 	RunID: to.Ptr("2f7fdb90-5df1-4b8e-ac2f-064cfa58202b"),
	// 	RunStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:37:44.625Z"); return t}()),
	// 	Status: to.Ptr("Succeeded"),
	// }
}
