package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v10"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/GetPrivateLinkResources.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().Get(ctx, "exampleResourceGroup", "exampleFactoryName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResourcesWrapper = armdatafactory.PrivateLinkResourcesWrapper{
	// 	Value: []*armdatafactory.PrivateLinkResource{
	// 		{
	// 			Name: to.Ptr("exampleFactoryName"),
	// 			Type: to.Ptr("Microsoft.DataFactory/factories/privateLinkResources"),
	// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName"),
	// 			Properties: &armdatafactory.PrivateLinkResourceProperties{
	// 				GroupID: to.Ptr("dataFactory"),
	// 				RequiredMembers: []*string{
	// 					to.Ptr("dataFactory")},
	// 					RequiredZoneNames: []*string{
	// 						to.Ptr("privatelink.datafactory.azure.net")},
	// 					},
	// 			}},
	// 		}
}
