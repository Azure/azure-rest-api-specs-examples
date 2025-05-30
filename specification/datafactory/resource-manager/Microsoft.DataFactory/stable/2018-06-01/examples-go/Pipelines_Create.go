package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v10"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_Create.json
func ExamplePipelinesClient_CreateOrUpdate_pipelinesCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPipelinesClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "examplePipeline", armdatafactory.PipelineResource{
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
										Parameters: map[string]any{
											"MyFileName":   "examplecontainer.csv",
											"MyFolderPath": "examplecontainer",
										},
										ReferenceName: to.Ptr("exampleDataset"),
									}},
								Outputs: []*armdatafactory.DatasetReference{
									{
										Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
										Parameters: map[string]any{
											"MyFileName": map[string]any{
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
			RunDimensions: map[string]any{
				"JobId": map[string]any{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PipelineResource = armdatafactory.PipelineResource{
	// 	Name: to.Ptr("examplePipeline"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/pipelines"),
	// 	Etag: to.Ptr("0a0069d4-0000-0000-0000-5b245bd50000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/pipelines/examplePipeline"),
	// 	Properties: &armdatafactory.Pipeline{
	// 		Activities: []armdatafactory.ActivityClassification{
	// 			&armdatafactory.ForEachActivity{
	// 				Name: to.Ptr("ExampleForeachActivity"),
	// 				Type: to.Ptr("ForEach"),
	// 				TypeProperties: &armdatafactory.ForEachActivityTypeProperties{
	// 					Activities: []armdatafactory.ActivityClassification{
	// 						&armdatafactory.CopyActivity{
	// 							Name: to.Ptr("ExampleCopyActivity"),
	// 							Type: to.Ptr("Copy"),
	// 							Inputs: []*armdatafactory.DatasetReference{
	// 								{
	// 									Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 									Parameters: map[string]any{
	// 										"MyFileName": "examplecontainer.csv",
	// 										"MyFolderPath": "examplecontainer",
	// 									},
	// 									ReferenceName: to.Ptr("exampleDataset"),
	// 							}},
	// 							Outputs: []*armdatafactory.DatasetReference{
	// 								{
	// 									Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 									Parameters: map[string]any{
	// 										"MyFileName": map[string]any{
	// 											"type": "Expression",
	// 											"value": "@item()",
	// 										},
	// 										"MyFolderPath": "examplecontainer",
	// 									},
	// 									ReferenceName: to.Ptr("exampleDataset"),
	// 							}},
	// 							TypeProperties: &armdatafactory.CopyActivityTypeProperties{
	// 								DataIntegrationUnits: float64(32),
	// 								Sink: &armdatafactory.BlobSink{
	// 									Type: to.Ptr("BlobSink"),
	// 								},
	// 								Source: &armdatafactory.BlobSource{
	// 									Type: to.Ptr("BlobSource"),
	// 								},
	// 							},
	// 					}},
	// 					IsSequential: to.Ptr(true),
	// 					Items: &armdatafactory.Expression{
	// 						Type: to.Ptr(armdatafactory.ExpressionTypeExpression),
	// 						Value: to.Ptr("@pipeline().parameters.OutputBlobNameList"),
	// 					},
	// 				},
	// 		}},
	// 		Parameters: map[string]*armdatafactory.ParameterSpecification{
	// 			"JobId": &armdatafactory.ParameterSpecification{
	// 				Type: to.Ptr(armdatafactory.ParameterTypeString),
	// 			},
	// 			"OutputBlobNameList": &armdatafactory.ParameterSpecification{
	// 				Type: to.Ptr(armdatafactory.ParameterTypeArray),
	// 			},
	// 		},
	// 		RunDimensions: map[string]any{
	// 			"JobId": map[string]any{
	// 				"type": "Expression",
	// 				"value": "@pipeline().parameters.JobId",
	// 			},
	// 		},
	// 		Variables: map[string]*armdatafactory.VariableSpecification{
	// 			"TestVariableArray": &armdatafactory.VariableSpecification{
	// 				Type: to.Ptr(armdatafactory.VariableTypeArray),
	// 			},
	// 		},
	// 	},
	// }
}
