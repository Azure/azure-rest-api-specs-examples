package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/be0d1c383d683a8eb22ab5fd5b75e380ac3c2eea/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/GlobalParameters_Create.json
func ExampleGlobalParametersClient_CreateOrUpdate_globalParametersCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGlobalParametersClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "default", armdatafactory.GlobalParameterResource{
		Properties: map[string]*armdatafactory.GlobalParameterSpecification{
			"waitTime": {
				Type:  to.Ptr(armdatafactory.GlobalParameterTypeInt),
				Value: float64(5),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GlobalParameterResource = armdatafactory.GlobalParameterResource{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/globalParameters"),
	// 	Etag: to.Ptr("0a008ad4-0000-0000-0000-5b245c6e0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/globalParameters/default"),
	// 	Properties: map[string]*armdatafactory.GlobalParameterSpecification{
	// 		"waitTime": &armdatafactory.GlobalParameterSpecification{
	// 			Type: to.Ptr(armdatafactory.GlobalParameterTypeInt),
	// 			Value: float64(5),
	// 		},
	// 	},
	// }
}
