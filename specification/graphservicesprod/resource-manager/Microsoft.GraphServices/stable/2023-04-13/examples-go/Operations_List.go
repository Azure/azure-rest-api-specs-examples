package armgraphservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/graphservices/armgraphservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0d41e635294dce73dfa99b07f3da4b68a9c9e29c/specification/graphservicesprod/resource-manager/Microsoft.GraphServices/stable/2023-04-13/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armgraphservices.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armgraphservices.OperationListResult{
		// 	Value: []*armgraphservices.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.GraphServices/accounts/read"),
		// 			Display: &armgraphservices.OperationDisplay{
		// 				Description: to.Ptr("Get Microsoft Graph Services Account resources"),
		// 				Operation: to.Ptr("Get Microsoft Graph Services Account resources"),
		// 				Provider: to.Ptr("Microsoft Graph Services"),
		// 				Resource: to.Ptr("Microsoft Graph Services Account"),
		// 			},
		// 	}},
		// }
	}
}
