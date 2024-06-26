package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/stable/2023-03-15/examples/OperationsList.json
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
		// page.OperationListResult = armhybridconnectivity.OperationListResult{
		// 	Value: []*armhybridconnectivity.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridConnectivity/operations/read"),
		// 			Display: &armhybridconnectivity.OperationDisplay{
		// 				Description: to.Ptr("Get the list of Operations"),
		// 				Operation: to.Ptr("Get operations"),
		// 				Provider: to.Ptr("Microsoft.HybridConnectivity"),
		// 				Resource: to.Ptr("Operations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridConnectivity/endpoints/read"),
		// 			Display: &armhybridconnectivity.OperationDisplay{
		// 				Description: to.Ptr("Get or list of endpoints to the target resource."),
		// 				Operation: to.Ptr("Get/List endpoints"),
		// 				Provider: to.Ptr("Microsoft.HybridConnectivity"),
		// 				Resource: to.Ptr("endpoints"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridConnectivity/endpoints/write"),
		// 			Display: &armhybridconnectivity.OperationDisplay{
		// 				Description: to.Ptr("Create or update the endpoint to the target resource."),
		// 				Operation: to.Ptr("Create/Update endpoint"),
		// 				Provider: to.Ptr("Microsoft.HybridConnectivity"),
		// 				Resource: to.Ptr("endpoints"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridConnectivity/endpoints/delete"),
		// 			Display: &armhybridconnectivity.OperationDisplay{
		// 				Description: to.Ptr("Deletes the endpoint access to the target resource."),
		// 				Operation: to.Ptr("Delete endpoint"),
		// 				Provider: to.Ptr("Microsoft.HybridConnectivity"),
		// 				Resource: to.Ptr("endpoints"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridConnectivity/endpoints/listCredentials/action"),
		// 			Display: &armhybridconnectivity.OperationDisplay{
		// 				Description: to.Ptr("List the endpoint access credentials to the resource."),
		// 				Operation: to.Ptr("List credentials for endpoint access"),
		// 				Provider: to.Ptr("Microsoft.HybridConnectivity"),
		// 				Resource: to.Ptr("endpoints"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridConnectivity/endpoints/listIngressGatewayCredentials/action"),
		// 			Display: &armhybridconnectivity.OperationDisplay{
		// 				Description: to.Ptr("List the ingress gateway credentials to the resource."),
		// 				Operation: to.Ptr("List credentials for ingress gateway"),
		// 				Provider: to.Ptr("Microsoft.HybridConnectivity"),
		// 				Resource: to.Ptr("endpoints"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridConnectivity/register/action"),
		// 			Display: &armhybridconnectivity.OperationDisplay{
		// 				Description: to.Ptr("Register the subscription for Microsoft.HybridConnectivity"),
		// 				Operation: to.Ptr("Register the Microsoft.HybridConnectivity"),
		// 				Provider: to.Ptr("Microsoft.HybridConnectivity"),
		// 				Resource: to.Ptr("Microsoft.HybridConnectivity"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridConnectivity/unregister/action"),
		// 			Display: &armhybridconnectivity.OperationDisplay{
		// 				Description: to.Ptr("Unregister the subscription for Microsoft.HybridConnectivity"),
		// 				Operation: to.Ptr("Unregister the Microsoft.HybridConnectivity"),
		// 				Provider: to.Ptr("Microsoft.HybridConnectivity"),
		// 				Resource: to.Ptr("Microsoft.HybridConnectivity"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}
