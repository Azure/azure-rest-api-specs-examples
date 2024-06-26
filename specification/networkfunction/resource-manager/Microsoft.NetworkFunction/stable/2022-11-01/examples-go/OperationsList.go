package armnetworkfunction_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkfunction/armnetworkfunction/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/527f6d35fb0d85c48210ca0f6f6f42814d63bd33/specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-11-01/examples/OperationsList.json
func ExampleClient_NewListOperationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkfunction.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListOperationsPager(nil)
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
		// page.OperationListResult = armnetworkfunction.OperationListResult{
		// 	Value: []*armnetworkfunction.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.NetworkFunction/azureTrafficCollectors/write"),
		// 			Display: &armnetworkfunction.OperationDisplay{
		// 				Description: to.Ptr("Creates or Update a azure traffic collector"),
		// 				Operation: to.Ptr("Create/Update a azure traffic collector"),
		// 				Provider: to.Ptr("Microsoft NetworkFunction"),
		// 				Resource: to.Ptr("AzureTrafficCollector"),
		// 			},
		// 	}},
		// }
	}
}
