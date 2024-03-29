package armpurview_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurview.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationList = armpurview.OperationList{
		// 	Value: []*armpurview.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Purview/operations/read"),
		// 			Display: &armpurview.OperationDisplay{
		// 				Description: to.Ptr("Reads all available operations in Purview Resource Provider."),
		// 				Operation: to.Ptr("Read all operations"),
		// 				Provider: to.Ptr("Microsoft Purview"),
		// 				Resource: to.Ptr("Operations"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Purview/register/action"),
		// 			Display: &armpurview.OperationDisplay{
		// 				Description: to.Ptr("Register the subscription for Purview Resource Provider"),
		// 				Operation: to.Ptr("Register Purview Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Purview"),
		// 				Resource: to.Ptr("Purview Resource Provider"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Purview/unregister/action"),
		// 			Display: &armpurview.OperationDisplay{
		// 				Description: to.Ptr("Unregister  the subscription for Purview Resource Provider"),
		// 				Operation: to.Ptr("Unregister Purview Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Purview"),
		// 				Resource: to.Ptr("Purview Resource Provider"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Purview/accounts/read"),
		// 			Display: &armpurview.OperationDisplay{
		// 				Description: to.Ptr("Read account resource for Purview Resource Provider."),
		// 				Operation: to.Ptr("Read account resource"),
		// 				Provider: to.Ptr("Microsoft Purview"),
		// 				Resource: to.Ptr("Account"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Purview/accounts/write"),
		// 			Display: &armpurview.OperationDisplay{
		// 				Description: to.Ptr("Write account resource for Purview Resource Provider."),
		// 				Operation: to.Ptr("Write account resource"),
		// 				Provider: to.Ptr("Microsoft Purview"),
		// 				Resource: to.Ptr("Account"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Purview/accounts/delete"),
		// 			Display: &armpurview.OperationDisplay{
		// 				Description: to.Ptr("Delete account resource for Purview Resource Provider."),
		// 				Operation: to.Ptr("Delete account resource"),
		// 				Provider: to.Ptr("Microsoft Purview"),
		// 				Resource: to.Ptr("Account"),
		// 			},
		// 	}},
		// }
	}
}
