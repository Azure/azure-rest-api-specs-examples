package armdatafactory_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v10"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ActivityRuns_QueryByPipelineRun.json
func ExampleActivityRunsClient_QueryByPipelineRun() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewActivityRunsClient().QueryByPipelineRun(ctx, "exampleResourceGroup", "exampleFactoryName", "2f7fdb90-5df1-4b8e-ac2f-064cfa58202b", armdatafactory.RunFilterParameters{
		LastUpdatedAfter:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:36:44.334Z"); return t }()),
		LastUpdatedBefore: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:49:48.368Z"); return t }()),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ActivityRunsQueryResponse = armdatafactory.ActivityRunsQueryResponse{
	// 	Value: []*armdatafactory.ActivityRun{
	// 		{
	// 			AdditionalProperties: map[string]any{
	// 				"retryAttempt": nil,
	// 				"userProperties": map[string]any{
	// 				},
	// 			},
	// 			ActivityName: to.Ptr("ExampleForeachActivity"),
	// 			ActivityRunEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:38:11.544Z"); return t}()),
	// 			ActivityRunID: to.Ptr("f30c5514-fb85-43ed-9fa4-768d42e58680"),
	// 			ActivityRunStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:37:49.480Z"); return t}()),
	// 			ActivityType: to.Ptr("ForEach"),
	// 			DurationInMs: to.Ptr[int32](22064),
	// 			Error: map[string]any{
	// 				"errorCode": "",
	// 				"failureType": "",
	// 				"message": "",
	// 				"target": "ExampleForeachActivity",
	// 			},
	// 			Input: map[string]any{
	// 			},
	// 			LinkedServiceName: to.Ptr(""),
	// 			Output: map[string]any{
	// 			},
	// 			PipelineName: to.Ptr("examplePipeline"),
	// 			PipelineRunID: to.Ptr("2f7fdb90-5df1-4b8e-ac2f-064cfa58202b"),
	// 			Status: to.Ptr("Succeeded"),
	// 		},
	// 		{
	// 			AdditionalProperties: map[string]any{
	// 				"retryAttempt": nil,
	// 				"userProperties": map[string]any{
	// 				},
	// 			},
	// 			ActivityName: to.Ptr("ExampleCopyActivity"),
	// 			ActivityRunEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:38:07.418Z"); return t}()),
	// 			ActivityRunID: to.Ptr("a96678c8-7167-4f00-b629-afccfbad4e51"),
	// 			ActivityRunStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:37:50.246Z"); return t}()),
	// 			ActivityType: to.Ptr("Copy"),
	// 			DurationInMs: to.Ptr[int32](17172),
	// 			Error: map[string]any{
	// 				"errorCode": "",
	// 				"failureType": "",
	// 				"message": "",
	// 				"target": "ExampleCopyActivity",
	// 			},
	// 			Input: map[string]any{
	// 				"dataIntegrationUnits": float64(32),
	// 				"sink":map[string]any{
	// 					"type": "BlobSink",
	// 				},
	// 				"source":map[string]any{
	// 					"type": "BlobSource",
	// 				},
	// 			},
	// 			LinkedServiceName: to.Ptr(""),
	// 			Output: map[string]any{
	// 				"copyDuration": float64(6),
	// 				"dataRead": float64(142000),
	// 				"dataWritten": float64(142000),
	// 				"effectiveIntegrationRuntime": "DefaultIntegrationRuntime (East US)",
	// 				"errors":[]any{
	// 				},
	// 				"executionDetails":[]any{
	// 					map[string]any{
	// 						"detailedDurations":map[string]any{
	// 							"queuingDuration": float64(4),
	// 							"transferDuration": float64(2),
	// 						},
	// 						"duration": float64(6),
	// 						"sink":map[string]any{
	// 							"type": "AzureBlob",
	// 						},
	// 						"source":map[string]any{
	// 							"type": "AzureBlob",
	// 						},
	// 						"start": "2018-06-16T00:37:50.68834Z",
	// 						"status": "Succeeded",
	// 						"usedCloudDataMovementUnits": float64(4),
	// 						"usedParallelCopies": float64(1),
	// 					},
	// 				},
	// 				"filesRead": float64(1),
	// 				"filesWritten": float64(1),
	// 				"throughput": float64(23.112),
	// 				"usedCloudDataMovementUnits": float64(4),
	// 				"usedParallelCopies": float64(1),
	// 			},
	// 			PipelineName: to.Ptr("examplePipeline"),
	// 			PipelineRunID: to.Ptr("2f7fdb90-5df1-4b8e-ac2f-064cfa58202b"),
	// 			Status: to.Ptr("Succeeded"),
	// 	}},
	// }
}
