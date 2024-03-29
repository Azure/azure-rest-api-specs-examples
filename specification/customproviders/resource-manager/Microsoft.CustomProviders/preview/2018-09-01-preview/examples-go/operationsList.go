package armcustomproviders_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customproviders/armcustomproviders"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/customproviders/resource-manager/Microsoft.CustomProviders/preview/2018-09-01-preview/examples/operationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomproviders.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.ResourceProviderOperationList = armcustomproviders.ResourceProviderOperationList{
		// 	Value: []*armcustomproviders.ResourceProviderOperation{
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomProviders/resourceProviders/read"),
		// 			Display: &armcustomproviders.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("List or get the Custom Providers"),
		// 				Operation: to.Ptr("List Custom Providers"),
		// 				Provider: to.Ptr("Azure Custom Providers"),
		// 				Resource: to.Ptr("resourceProviders"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomProviders/resourceProviders/delete"),
		// 			Display: &armcustomproviders.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("List or get the Custom Providers"),
		// 				Operation: to.Ptr("List Custom Providers"),
		// 				Provider: to.Ptr("Azure Custom Providers"),
		// 				Resource: to.Ptr("resourceProviders"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomProviders/resourceProviders/wrte"),
		// 			Display: &armcustomproviders.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("List or get the Custom Providers"),
		// 				Operation: to.Ptr("List Custom Providers"),
		// 				Provider: to.Ptr("Azure Custom Providers"),
		// 				Resource: to.Ptr("resourceProviders"),
		// 			},
		// 	}},
		// }
	}
}
