package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/GlobalParameters_ListByFactory.json
func ExampleGlobalParametersClient_NewListByFactoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGlobalParametersClient().NewListByFactoryPager("exampleResourceGroup", "exampleFactoryName", nil)
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
		// page.GlobalParameterListResponse = armdatafactory.GlobalParameterListResponse{
		// 	Value: []*armdatafactory.GlobalParameterResource{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories/globalParameters"),
		// 			Etag: to.Ptr("da00a1c3-0000-0400-0000-6241f3290000"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/globalParameters/default"),
		// 			Properties: map[string]*armdatafactory.GlobalParameterSpecification{
		// 				"copyPipelineTest": &armdatafactory.GlobalParameterSpecification{
		// 					Type: to.Ptr(armdatafactory.GlobalParameterTypeObject),
		// 					Value: map[string]any{
		// 						"mySinkDatasetFolderPath": "exampleOutput",
		// 						"mySourceDatasetFolderPath": "exampleInput/",
		// 						"testingEmptyContextParams": "",
		// 					},
		// 				},
		// 				"mySourceDatasetFolderPath": &armdatafactory.GlobalParameterSpecification{
		// 					Type: to.Ptr(armdatafactory.GlobalParameterTypeString),
		// 					Value: "input",
		// 				},
		// 				"url": &armdatafactory.GlobalParameterSpecification{
		// 					Type: to.Ptr(armdatafactory.GlobalParameterTypeString),
		// 					Value: "https://testuri.test",
		// 				},
		// 				"validADFOffice365Uris": &armdatafactory.GlobalParameterSpecification{
		// 					Type: to.Ptr(armdatafactory.GlobalParameterTypeArray),
		// 					Value: []any{
		// 						"https://testuri.test",
		// 						"https://testuri.test",
		// 					},
		// 				},
		// 				"waitTime": &armdatafactory.GlobalParameterSpecification{
		// 					Type: to.Ptr(armdatafactory.GlobalParameterTypeInt),
		// 					Value: float64(5),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
