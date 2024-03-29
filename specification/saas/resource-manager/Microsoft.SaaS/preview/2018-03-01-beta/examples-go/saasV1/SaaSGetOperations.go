package armsaas_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/saas/resource-manager/Microsoft.SaaS/preview/2018-03-01-beta/examples/saasV1/SaaSGetOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsaas.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.AppOperationsResponseWithContinuation = armsaas.AppOperationsResponseWithContinuation{
		// 	Value: []*armsaas.AppOperation{
		// 		{
		// 			Name: to.Ptr("Microsoft.SaaS/register/action"),
		// 			Display: &armsaas.AppOperationDisplay{
		// 				Description: to.Ptr("Register Saas resource provider in the tenant"),
		// 				Operation: to.Ptr("Register SaaS resource provider"),
		// 				Provider: to.Ptr("Microsoft"),
		// 				Resource: to.Ptr("SaaS Resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SaaS/saasresources/read"),
		// 			Display: &armsaas.AppOperationDisplay{
		// 				Description: to.Ptr("Get all SaaS resources or one resource"),
		// 				Operation: to.Ptr("Get SaaS Resources"),
		// 				Provider: to.Ptr("Microsoft"),
		// 				Resource: to.Ptr("SaaS Resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SaaS/saasresources/write"),
		// 			Display: &armsaas.AppOperationDisplay{
		// 				Description: to.Ptr("Create or Update a SaaS resource"),
		// 				Operation: to.Ptr("Create or Update a SaaS resource"),
		// 				Provider: to.Ptr("Microsoft"),
		// 				Resource: to.Ptr("SaaS Resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SaaS/saasresources/delete"),
		// 			Display: &armsaas.AppOperationDisplay{
		// 				Description: to.Ptr("Delete a SaaS Resource"),
		// 				Operation: to.Ptr("Delete a SaaS Resource"),
		// 				Provider: to.Ptr("Microsoft"),
		// 				Resource: to.Ptr("SaaS Resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SaaS/resources/read"),
		// 			Display: &armsaas.AppOperationDisplay{
		// 				Description: to.Ptr("Get one resource"),
		// 				Operation: to.Ptr("Get SaaS Subscription Level Resources"),
		// 				Provider: to.Ptr("Microsoft"),
		// 				Resource: to.Ptr("SaaS Subscription Level Resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SaaS/resources/write"),
		// 			Display: &armsaas.AppOperationDisplay{
		// 				Description: to.Ptr("Create or Update a SaaS Subscription Level resource"),
		// 				Operation: to.Ptr("Create or Update a SaaS Subscription Level resource"),
		// 				Provider: to.Ptr("Microsoft"),
		// 				Resource: to.Ptr("SaaS Subscription Level Resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SaaS/resources/delete"),
		// 			Display: &armsaas.AppOperationDisplay{
		// 				Description: to.Ptr("Delete a SaaS Subscription Level Resource"),
		// 				Operation: to.Ptr("Delete a SaaS Subscription Level Resource"),
		// 				Provider: to.Ptr("Microsoft"),
		// 				Resource: to.Ptr("SaaS Subscription Level Resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SaaS/saasresources/listaccesstoken/action"),
		// 			Display: &armsaas.AppOperationDisplay{
		// 				Description: to.Ptr("Get the access Token for a SaaS resource"),
		// 				Operation: to.Ptr("Get the access Token"),
		// 				Provider: to.Ptr("Microsoft"),
		// 				Resource: to.Ptr("SaaS Resources"),
		// 			},
		// 	}},
		// }
	}
}
