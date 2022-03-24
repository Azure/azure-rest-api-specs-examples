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

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_Create.json
func ExampleDataFlowsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewDataFlowsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<data-flow-name>",
		armdatafactory.DataFlowResource{
			Properties: &armdatafactory.MappingDataFlow{
				Type:        to.StringPtr("<type>"),
				Description: to.StringPtr("<description>"),
				TypeProperties: &armdatafactory.MappingDataFlowTypeProperties{
					Script: to.StringPtr("<script>"),
					Sinks: []*armdatafactory.DataFlowSink{
						{
							Name: to.StringPtr("<name>"),
							Dataset: &armdatafactory.DatasetReference{
								Type:          armdatafactory.DatasetReferenceType("DatasetReference").ToPtr(),
								ReferenceName: to.StringPtr("<reference-name>"),
							},
						},
						{
							Name: to.StringPtr("<name>"),
							Dataset: &armdatafactory.DatasetReference{
								Type:          armdatafactory.DatasetReferenceType("DatasetReference").ToPtr(),
								ReferenceName: to.StringPtr("<reference-name>"),
							},
						}},
					Sources: []*armdatafactory.DataFlowSource{
						{
							Name: to.StringPtr("<name>"),
							Dataset: &armdatafactory.DatasetReference{
								Type:          armdatafactory.DatasetReferenceType("DatasetReference").ToPtr(),
								ReferenceName: to.StringPtr("<reference-name>"),
							},
						},
						{
							Name: to.StringPtr("<name>"),
							Dataset: &armdatafactory.DatasetReference{
								Type:          armdatafactory.DatasetReferenceType("DatasetReference").ToPtr(),
								ReferenceName: to.StringPtr("<reference-name>"),
							},
						}},
				},
			},
		},
		&armdatafactory.DataFlowsClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DataFlowsClientCreateOrUpdateResult)
}
```
