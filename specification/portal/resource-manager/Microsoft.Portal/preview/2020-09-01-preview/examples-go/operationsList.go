package armportal_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/portal/resource-manager/Microsoft.Portal/preview/2020-09-01-preview/examples/operationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armportal.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.ResourceProviderOperationList = armportal.ResourceProviderOperationList{
		// 	Value: []*armportal.ResourceProviderOperation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Portal/dashboards/read"),
		// 			Display: &armportal.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Reads the dashboards for the subscription."),
		// 				Operation: to.Ptr("Get Dashboard"),
		// 				Provider: to.Ptr("Microsoft Portal"),
		// 				Resource: to.Ptr("Dashboards"),
		// 			},
		// 			IsDataAction: to.Ptr("false"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Portal/dashboards/write"),
		// 			Display: &armportal.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Add or modify dashboard to a subscription."),
		// 				Operation: to.Ptr("Set Dashboard"),
		// 				Provider: to.Ptr("Microsoft Portal"),
		// 				Resource: to.Ptr("Dashboards"),
		// 			},
		// 			IsDataAction: to.Ptr("false"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Portal/dashboards/delete"),
		// 			Display: &armportal.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Removes the dashboard from the subscription."),
		// 				Operation: to.Ptr("Delete Dashboard"),
		// 				Provider: to.Ptr("Microsoft Portal"),
		// 				Resource: to.Ptr("Dashboards"),
		// 			},
		// 			IsDataAction: to.Ptr("false"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Portal/register/action"),
		// 			Display: &armportal.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Registers the subscription for the portal resource provider and enables shared dashboards."),
		// 				Operation: to.Ptr("Registers the Portal Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Portal"),
		// 				Resource: to.Ptr("Portal Resource Provider"),
		// 			},
		// 			IsDataAction: to.Ptr("false"),
		// 	}},
		// }
	}
}
