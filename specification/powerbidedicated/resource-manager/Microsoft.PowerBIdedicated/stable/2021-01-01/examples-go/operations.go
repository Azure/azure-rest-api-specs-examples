package armpowerbidedicated_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/operations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpowerbidedicated.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armpowerbidedicated.OperationListResult{
		// 	Value: []*armpowerbidedicated.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerBIDedicated/checkNameAvailability/action"),
		// 			Display: &armpowerbidedicated.OperationDisplay{
		// 				Description: to.Ptr("Checks that given Power BI Dedicated Capacity name is valid and not in use."),
		// 				Operation: to.Ptr("Capacities_CheckNameAvailability"),
		// 				Provider: to.Ptr("Microsoft.PowerBIDedicated"),
		// 				Resource: to.Ptr("checkNameAvailability"),
		// 			},
		// 	}},
		// }
	}
}
