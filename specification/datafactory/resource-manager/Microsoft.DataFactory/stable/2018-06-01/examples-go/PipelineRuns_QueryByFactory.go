package armdatafactory_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4afa6837cfb404d8e5ffa8a604a5e09996d6f79e/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/PipelineRuns_QueryByFactory.json
func ExamplePipelineRunsClient_QueryByFactory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPipelineRunsClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.QueryByFactory(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.RunFilterParameters{
		Filters: []*armdatafactory.RunQueryFilter{
			{
				Operand:  to.Ptr(armdatafactory.RunQueryFilterOperandPipelineName),
				Operator: to.Ptr(armdatafactory.RunQueryFilterOperatorEquals),
				Values: []*string{
					to.Ptr("examplePipeline")},
			}},
		LastUpdatedAfter:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:36:44.3345758Z"); return t }()),
		LastUpdatedBefore: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:49:48.3686473Z"); return t }()),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PipelineRunsQueryResponse = armdatafactory.PipelineRunsQueryResponse{
	// 	Value: []*armdatafactory.PipelineRun{
	// 		{
	// 			AdditionalProperties: map[string]any{
	// 				"annotations": []any{
	// 				},
	// 				"runDimension": map[string]any{
	// 					"JobId": "79c1cc52-265f-41a5-9553-be65e736fbd3",
	// 				},
	// 			},
	// 			DurationInMs: to.Ptr[int32](28105),
	// 			InvokedBy: &armdatafactory.PipelineRunInvokedBy{
	// 				Name: to.Ptr("Manual"),
	// 				ID: to.Ptr("80a01654a9d34ad18b3fcac5d5d76b67"),
	// 			},
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:38:12.7314495Z"); return t}()),
	// 			Message: to.Ptr(""),
	// 			Parameters: map[string]*string{
	// 				"OutputBlobNameList": to.Ptr("[\"exampleoutput.csv\"]"),
	// 			},
	// 			PipelineName: to.Ptr("examplePipeline"),
	// 			RunEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:38:12.7314495Z"); return t}()),
	// 			RunID: to.Ptr("2f7fdb90-5df1-4b8e-ac2f-064cfa58202b"),
	// 			RunStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:37:44.6257014Z"); return t}()),
	// 			Status: to.Ptr("Succeeded"),
	// 		},
	// 		{
	// 			AdditionalProperties: map[string]any{
	// 				"annotations": []any{
	// 				},
	// 				"runDimension": map[string]any{
	// 					"JobId": "84a3c493-0628-4b44-852f-ef5b3a11bdab",
	// 				},
	// 			},
	// 			InvokedBy: &armdatafactory.PipelineRunInvokedBy{
	// 				Name: to.Ptr("Manual"),
	// 				ID: to.Ptr("7c5fd7ef7e8a464b98b931cf15fcac66"),
	// 			},
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:39:51.216097Z"); return t}()),
	// 			Message: to.Ptr(""),
	// 			Parameters: map[string]*string{
	// 				"OutputBlobNameList": to.Ptr("[\"exampleoutput.csv\"]"),
	// 			},
	// 			PipelineName: to.Ptr("examplePipeline"),
	// 			RunID: to.Ptr("16ac5348-ff82-4f95-a80d-638c1d47b721"),
	// 			RunStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:39:49.2745128Z"); return t}()),
	// 			Status: to.Ptr("Cancelled"),
	// 	}},
	// }
}
