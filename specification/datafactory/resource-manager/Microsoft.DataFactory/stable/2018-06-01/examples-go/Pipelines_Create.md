Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatafactory%2Farmdatafactory%2Fv0.3.0/sdk/resourcemanager/datafactory/armdatafactory/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_Create.json
func ExamplePipelinesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewPipelinesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<pipeline-name>",
		armdatafactory.PipelineResource{
			Properties: &armdatafactory.Pipeline{
				Activities: []armdatafactory.ActivityClassification{
					&armdatafactory.ForEachActivity{
						Name: to.StringPtr("<name>"),
						Type: to.StringPtr("<type>"),
						TypeProperties: &armdatafactory.ForEachActivityTypeProperties{
							Activities: []armdatafactory.ActivityClassification{
								&armdatafactory.CopyActivity{
									Name: to.StringPtr("<name>"),
									Type: to.StringPtr("<type>"),
									Inputs: []*armdatafactory.DatasetReference{
										{
											Type: armdatafactory.DatasetReferenceType("DatasetReference").ToPtr(),
											Parameters: map[string]interface{}{
												"MyFileName":   "examplecontainer.csv",
												"MyFolderPath": "examplecontainer",
											},
											ReferenceName: to.StringPtr("<reference-name>"),
										}},
									Outputs: []*armdatafactory.DatasetReference{
										{
											Type: armdatafactory.DatasetReferenceType("DatasetReference").ToPtr(),
											Parameters: map[string]interface{}{
												"MyFileName": map[string]interface{}{
													"type":  "Expression",
													"value": "@item()",
												},
												"MyFolderPath": "examplecontainer",
											},
											ReferenceName: to.StringPtr("<reference-name>"),
										}},
									TypeProperties: &armdatafactory.CopyActivityTypeProperties{
										DataIntegrationUnits: float64(32),
										Sink: &armdatafactory.BlobSink{
											Type: to.StringPtr("<type>"),
										},
										Source: &armdatafactory.BlobSource{
											Type: to.StringPtr("<type>"),
										},
									},
								}},
							IsSequential: to.BoolPtr(true),
							Items: &armdatafactory.Expression{
								Type:  armdatafactory.ExpressionType("Expression").ToPtr(),
								Value: to.StringPtr("<value>"),
							},
						},
					}},
				Parameters: map[string]*armdatafactory.ParameterSpecification{
					"JobId": {
						Type: armdatafactory.ParameterType("String").ToPtr(),
					},
					"OutputBlobNameList": {
						Type: armdatafactory.ParameterType("Array").ToPtr(),
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
						Type: armdatafactory.VariableType("Array").ToPtr(),
					},
				},
			},
		},
		&armdatafactory.PipelinesClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PipelinesClientCreateOrUpdateResult)
}
```
