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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_Create.json
func ExampleDataFlowsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdatafactory.NewDataFlowsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<data-flow-name>",
		armdatafactory.DataFlowResource{
			Properties: &armdatafactory.MappingDataFlow{
				Type:        to.Ptr("<type>"),
				Description: to.Ptr("<description>"),
				TypeProperties: &armdatafactory.MappingDataFlowTypeProperties{
					Script: to.Ptr("<script>"),
					Sinks: []*armdatafactory.DataFlowSink{
						{
							Name: to.Ptr("<name>"),
							Dataset: &armdatafactory.DatasetReference{
								Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
								ReferenceName: to.Ptr("<reference-name>"),
							},
						},
						{
							Name: to.Ptr("<name>"),
							Dataset: &armdatafactory.DatasetReference{
								Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
								ReferenceName: to.Ptr("<reference-name>"),
							},
						}},
					Sources: []*armdatafactory.DataFlowSource{
						{
							Name: to.Ptr("<name>"),
							Dataset: &armdatafactory.DatasetReference{
								Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
								ReferenceName: to.Ptr("<reference-name>"),
							},
						},
						{
							Name: to.Ptr("<name>"),
							Dataset: &armdatafactory.DatasetReference{
								Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
								ReferenceName: to.Ptr("<reference-name>"),
							},
						}},
				},
			},
		},
		&armdatafactory.DataFlowsClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
