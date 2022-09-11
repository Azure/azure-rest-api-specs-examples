package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Datasets_Update.json
func ExampleDatasetsClient_CreateOrUpdate_datasetsUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewDatasetsClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleDataset", armdatafactory.DatasetResource{
		Properties: &armdatafactory.AzureBlobDataset{
			Type:        to.Ptr("AzureBlob"),
			Description: to.Ptr("Example description"),
			LinkedServiceName: &armdatafactory.LinkedServiceReference{
				Type:          to.Ptr(armdatafactory.LinkedServiceReferenceTypeLinkedServiceReference),
				ReferenceName: to.Ptr("exampleLinkedService"),
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
					Type: to.Ptr("TextFormat"),
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
	}, &armdatafactory.DatasetsClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
