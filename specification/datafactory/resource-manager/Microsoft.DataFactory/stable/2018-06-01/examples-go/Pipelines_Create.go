package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_Create.json
func ExamplePipelinesClient_CreateOrUpdate_pipelinesCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPipelinesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "examplePipeline", armdatafactory.PipelineResource{
		Properties: &armdatafactory.Pipeline{
			Activities: []armdatafactory.ActivityClassification{
				&armdatafactory.ForEachActivity{
					Name: to.Ptr("ExampleForeachActivity"),
					Type: to.Ptr("ForEach"),
					TypeProperties: &armdatafactory.ForEachActivityTypeProperties{
						Activities: []armdatafactory.ActivityClassification{
							&armdatafactory.CopyActivity{
								Name: to.Ptr("ExampleCopyActivity"),
								Type: to.Ptr("Copy"),
								Inputs: []*armdatafactory.DatasetReference{
									{
										Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
										Parameters: map[string]interface{}{
											"MyFileName":   "examplecontainer.csv",
											"MyFolderPath": "examplecontainer",
										},
										ReferenceName: to.Ptr("exampleDataset"),
									}},
								Outputs: []*armdatafactory.DatasetReference{
									{
										Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
										Parameters: map[string]interface{}{
											"MyFileName": map[string]interface{}{
												"type":  "Expression",
												"value": "@item()",
											},
											"MyFolderPath": "examplecontainer",
										},
										ReferenceName: to.Ptr("exampleDataset"),
									}},
								TypeProperties: &armdatafactory.CopyActivityTypeProperties{
									DataIntegrationUnits: float64(32),
									Sink: &armdatafactory.BlobSink{
										Type: to.Ptr("BlobSink"),
									},
									Source: &armdatafactory.BlobSource{
										Type: to.Ptr("BlobSource"),
									},
								},
							}},
						IsSequential: to.Ptr(true),
						Items: &armdatafactory.Expression{
							Type:  to.Ptr(armdatafactory.ExpressionTypeExpression),
							Value: to.Ptr("@pipeline().parameters.OutputBlobNameList"),
						},
					},
				}},
			Parameters: map[string]*armdatafactory.ParameterSpecification{
				"JobId": {
					Type: to.Ptr(armdatafactory.ParameterTypeString),
				},
				"OutputBlobNameList": {
					Type: to.Ptr(armdatafactory.ParameterTypeArray),
				},
			},
			Policy: &armdatafactory.PipelinePolicy{
				ElapsedTimeMetric: &armdatafactory.PipelineElapsedTimeMetricPolicy{
					Duration: "0.00:10:00",
				},
			},
			RunDimensions: map[string]interface{}{
				"JobId": map[string]interface{}{
					"type":  "Expression",
					"value": "@pipeline().parameters.JobId",
				},
			},
			Variables: map[string]*armdatafactory.VariableSpecification{
				"TestVariableArray": {
					Type: to.Ptr(armdatafactory.VariableTypeArray),
				},
			},
		},
	}, &armdatafactory.PipelinesClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
