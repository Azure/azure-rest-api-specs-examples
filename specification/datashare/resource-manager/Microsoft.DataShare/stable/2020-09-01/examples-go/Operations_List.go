package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationList = armdatashare.OperationList{
		// 	Value: []*armdatashare.OperationModel{
		// 		{
		// 			Name: to.Ptr("Microsoft.DataShare/operations/read"),
		// 			Display: &armdatashare.OperationModelProperties{
		// 				Description: to.Ptr("Reads all available operations in Data Share Resource Provider."),
		// 				Operation: to.Ptr("Read all operations"),
		// 				Provider: to.Ptr("Microsoft Data Share"),
		// 				Resource: to.Ptr("Data Share Resource Provider"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataShare/register/action"),
		// 			Display: &armdatashare.OperationModelProperties{
		// 				Description: to.Ptr("Register the subscription for the Data Share Resource Provider."),
		// 				Operation: to.Ptr("Register Data Share Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Data Share"),
		// 				Resource: to.Ptr("Data Share Resource Provider"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataShare/unregister/action"),
		// 			Display: &armdatashare.OperationModelProperties{
		// 				Description: to.Ptr("Unregister the subscription for the Data Share Resource Provider."),
		// 				Operation: to.Ptr("Unregister Data Share Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Data Share"),
		// 				Resource: to.Ptr("Data Share Resource Provider"),
		// 			},
		// 	}},
		// }
	}
}
