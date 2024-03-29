package armhanaonazure_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2020-02-07-preview/examples/HanaOperations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhanaonazure.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationList = armhanaonazure.OperationList{
		// 	Value: []*armhanaonazure.Operation{
		// 		{
		// 			Name: to.Ptr("HanaOp1"),
		// 			Display: &armhanaonazure.Display{
		// 				Description: to.Ptr("HanaOp1Description"),
		// 				Operation: to.Ptr("HanaOp1OperationName"),
		// 				Origin: to.Ptr("HanaOp1Origin"),
		// 				Provider: to.Ptr("HanaOp1ProviderName"),
		// 				Resource: to.Ptr("HanaOp1ResourceName"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("HanaOp2"),
		// 			Display: &armhanaonazure.Display{
		// 				Description: to.Ptr("HanaOp2Description"),
		// 				Operation: to.Ptr("HanaOp2OperationName"),
		// 				Origin: to.Ptr("HanaOp2Origin"),
		// 				Provider: to.Ptr("HanaOp2ProviderName"),
		// 				Resource: to.Ptr("HanaOp2ResourceName"),
		// 			},
		// 	}},
		// }
	}
}
