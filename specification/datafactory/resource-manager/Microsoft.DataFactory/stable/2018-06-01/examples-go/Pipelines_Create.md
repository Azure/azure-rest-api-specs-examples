Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatafactory%2Farmdatafactory%2Fv0.5.0/sdk/resourcemanager/datafactory/armdatafactory/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_Create.json
func ExamplePipelinesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPipelinesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<pipeline-name>",
		armdatafactory.PipelineResource{
			Properties: &armdatafactory.Pipeline{
				Activities: []armdatafactory.ActivityClassification{
					&armdatafactory.ForEachActivity{
						Name: to.Ptr("<name>"),
						Type: to.Ptr("<type>"),
						TypeProperties: &armdatafactory.ForEachActivityTypeProperties{
							Activities: []armdatafactory.ActivityClassification{
								&armdatafactory.CopyActivity{
									Name: to.Ptr("<name>"),
									Type: to.Ptr("<type>"),
									Inputs: []*armdatafactory.DatasetReference{
										{
											Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
											Parameters: map[string]interface{}{
												"MyFileName":   "examplecontainer.csv",
												"MyFolderPath": "examplecontainer",
											},
											ReferenceName: to.Ptr("<reference-name>"),
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
											ReferenceName: to.Ptr("<reference-name>"),
										}},
									TypeProperties: &armdatafactory.CopyActivityTypeProperties{
										DataIntegrationUnits: float64(32),
										Sink: &armdatafactory.BlobSink{
											Type: to.Ptr("<type>"),
										},
										Source: &armdatafactory.BlobSource{
											Type: to.Ptr("<type>"),
										},
									},
								}},
							IsSequential: to.Ptr(true),
							Items: &armdatafactory.Expression{
								Type:  to.Ptr(armdatafactory.ExpressionTypeExpression),
								Value: to.Ptr("<value>"),
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
		},
		&armdatafactory.PipelinesClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
