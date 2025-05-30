package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// Generated from example definition: 2024-12-01/OperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page = armhybridconnectivity.OperationsClientListResponse{
		// 	OperationListResult: armhybridconnectivity.OperationListResult{
		// 		Value: []*armhybridconnectivity.Operation{
		// 			{
		// 				Name: to.Ptr("Microsoft.HybridConnectivity/operations/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armhybridconnectivity.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.HybridConnectivity"),
		// 					Resource: to.Ptr("Operations"),
		// 					Operation: to.Ptr("Get operations"),
		// 					Description: to.Ptr("Get the list of Operations"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.HybridConnectivity/endpoints/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armhybridconnectivity.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.HybridConnectivity"),
		// 					Resource: to.Ptr("endpoints"),
		// 					Operation: to.Ptr("Get/List endpoints"),
		// 					Description: to.Ptr("Get or list of endpoints to the target resource."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.HybridConnectivity/endpoints/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armhybridconnectivity.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.HybridConnectivity"),
		// 					Resource: to.Ptr("endpoints"),
		// 					Operation: to.Ptr("Create/Update endpoint"),
		// 					Description: to.Ptr("Create or update the endpoint to the target resource."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.HybridConnectivity/endpoints/delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armhybridconnectivity.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.HybridConnectivity"),
		// 					Resource: to.Ptr("endpoints"),
		// 					Operation: to.Ptr("Delete endpoint"),
		// 					Description: to.Ptr("Deletes the endpoint access to the target resource."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.HybridConnectivity/endpoints/listCredentials/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armhybridconnectivity.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.HybridConnectivity"),
		// 					Resource: to.Ptr("endpoints"),
		// 					Operation: to.Ptr("List credentials for endpoint access"),
		// 					Description: to.Ptr("List the endpoint access credentials to the resource."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.HybridConnectivity/endpoints/listIngressGatewayCredentials/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armhybridconnectivity.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.HybridConnectivity"),
		// 					Resource: to.Ptr("endpoints"),
		// 					Operation: to.Ptr("List credentials for ingress gateway"),
		// 					Description: to.Ptr("List the ingress gateway credentials to the resource."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.HybridConnectivity/register/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armhybridconnectivity.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.HybridConnectivity"),
		// 					Resource: to.Ptr("Microsoft.HybridConnectivity"),
		// 					Operation: to.Ptr("Register the Microsoft.HybridConnectivity"),
		// 					Description: to.Ptr("Register the subscription for Microsoft.HybridConnectivity"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.HybridConnectivity/unregister/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armhybridconnectivity.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.HybridConnectivity"),
		// 					Resource: to.Ptr("Microsoft.HybridConnectivity"),
		// 					Operation: to.Ptr("Unregister the Microsoft.HybridConnectivity"),
		// 					Description: to.Ptr("Unregister the subscription for Microsoft.HybridConnectivity"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
