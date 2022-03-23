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

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Datasets_Create.json
func ExampleDatasetsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewDatasetsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<dataset-name>",
		armdatafactory.DatasetResource{
			Properties: &armdatafactory.AzureBlobDataset{
				Type: to.StringPtr("<type>"),
				LinkedServiceName: &armdatafactory.LinkedServiceReference{
					Type:          armdatafactory.LinkedServiceReferenceType("LinkedServiceReference").ToPtr(),
					ReferenceName: to.StringPtr("<reference-name>"),
				},
				Parameters: map[string]*armdatafactory.ParameterSpecification{
					"MyFileName": {
						Type: armdatafactory.ParameterType("String").ToPtr(),
					},
					"MyFolderPath": {
						Type: armdatafactory.ParameterType("String").ToPtr(),
					},
				},
				TypeProperties: &armdatafactory.AzureBlobDatasetTypeProperties{
					Format: &armdatafactory.TextFormat{
						Type: to.StringPtr("<type>"),
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
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DatasetsClientCreateOrUpdateResult)
}
```
