package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/be0d1c383d683a8eb22ab5fd5b75e380ac3c2eea/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Datasets_ListByFactory.json
func ExampleDatasetsClient_NewListByFactoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatasetsClient().NewListByFactoryPager("exampleResourceGroup", "exampleFactoryName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.DatasetListResponse = armdatafactory.DatasetListResponse{
		// 	Value: []*armdatafactory.DatasetResource{
		// 		{
		// 			Name: to.Ptr("exampleDataset"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories/datasets"),
		// 			Etag: to.Ptr("0a0068d4-0000-0000-0000-5b245bd30000"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/datasets/exampleDataset"),
		// 			Properties: &armdatafactory.AzureBlobDataset{
		// 				Type: to.Ptr("AzureBlob"),
		// 				Description: to.Ptr("Example description"),
		// 				LinkedServiceName: &armdatafactory.LinkedServiceReference{
		// 					Type: to.Ptr(armdatafactory.LinkedServiceReferenceTypeLinkedServiceReference),
		// 					ReferenceName: to.Ptr("exampleLinkedService"),
		// 				},
		// 				Parameters: map[string]*armdatafactory.ParameterSpecification{
		// 					"MyFileName": &armdatafactory.ParameterSpecification{
		// 						Type: to.Ptr(armdatafactory.ParameterTypeString),
		// 					},
		// 					"MyFolderPath": &armdatafactory.ParameterSpecification{
		// 						Type: to.Ptr(armdatafactory.ParameterTypeString),
		// 					},
		// 				},
		// 				TypeProperties: &armdatafactory.AzureBlobDatasetTypeProperties{
		// 					Format: &armdatafactory.TextFormat{
		// 						Type: to.Ptr("TextFormat"),
		// 					},
		// 					FileName: map[string]any{
		// 						"type": "Expression",
		// 						"value": "@dataset().MyFileName",
		// 					},
		// 					FolderPath: map[string]any{
		// 						"type": "Expression",
		// 						"value": "@dataset().MyFolderPath",
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
