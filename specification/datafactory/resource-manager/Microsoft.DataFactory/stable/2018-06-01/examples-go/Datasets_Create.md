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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Datasets_Create.json
func ExampleDatasetsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdatafactory.NewDatasetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<dataset-name>",
		armdatafactory.DatasetResource{
			Properties: &armdatafactory.AzureBlobDataset{
				Type: to.Ptr("<type>"),
				LinkedServiceName: &armdatafactory.LinkedServiceReference{
					Type:          to.Ptr(armdatafactory.LinkedServiceReferenceTypeLinkedServiceReference),
					ReferenceName: to.Ptr("<reference-name>"),
				},
				Parameters: map[string]*armdatafactory.ParameterSpecification{
					"MyFileName": {
						Type: to.Ptr(armdatafactory.ParameterTypeString),
					},
					"MyFolderPath": {
						Type: to.Ptr(armdatafactory.ParameterTypeString),
					},
				},
				TypeProperties: &armdatafactory.AzureBlobDatasetTypeProperties{
					Format: &armdatafactory.TextFormat{
						Type: to.Ptr("<type>"),
					},
					FileName: map[string]interface{}{
						"type":  "Expression",
						"value": "@dataset().MyFileName",
					},
					FolderPath: map[string]interface{}{
						"type":  "Expression",
						"value": "@dataset().MyFolderPath",
					},
				},
			},
		},
		&armdatafactory.DatasetsClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
