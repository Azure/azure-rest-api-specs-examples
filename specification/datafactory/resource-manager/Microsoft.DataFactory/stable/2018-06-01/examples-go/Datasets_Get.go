package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01a71545e82bb98b8137d3038150c436d46a59ed/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Datasets_Get.json
func ExampleDatasetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatasetsClient().Get(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleDataset", &armdatafactory.DatasetsClientGetOptions{IfNoneMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DatasetResource = armdatafactory.DatasetResource{
	// 	Name: to.Ptr("exampleDataset"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/datasets"),
	// 	Etag: to.Ptr("15004c4f-0000-0200-0000-5cbe090e0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/datasets/exampleDataset"),
	// 	Properties: &armdatafactory.AzureBlobDataset{
	// 		Type: to.Ptr("AzureBlob"),
	// 		Description: to.Ptr("Example description"),
	// 		LinkedServiceName: &armdatafactory.LinkedServiceReference{
	// 			Type: to.Ptr(armdatafactory.LinkedServiceReferenceTypeLinkedServiceReference),
	// 			ReferenceName: to.Ptr("exampleLinkedService"),
	// 		},
	// 		Parameters: map[string]*armdatafactory.ParameterSpecification{
	// 			"MyFileName": &armdatafactory.ParameterSpecification{
	// 				Type: to.Ptr(armdatafactory.ParameterTypeString),
	// 			},
	// 			"MyFolderPath": &armdatafactory.ParameterSpecification{
	// 				Type: to.Ptr(armdatafactory.ParameterTypeString),
	// 			},
	// 		},
	// 		TypeProperties: &armdatafactory.AzureBlobDatasetTypeProperties{
	// 			Format: &armdatafactory.TextFormat{
	// 				Type: to.Ptr("TextFormat"),
	// 			},
	// 			FileName: map[string]any{
	// 				"type": "Expression",
	// 				"value": "@dataset().MyFileName",
	// 			},
	// 			FolderPath: map[string]any{
	// 				"type": "Expression",
	// 				"value": "@dataset().MyFolderPath",
	// 			},
	// 		},
	// 	},
	// }
}
