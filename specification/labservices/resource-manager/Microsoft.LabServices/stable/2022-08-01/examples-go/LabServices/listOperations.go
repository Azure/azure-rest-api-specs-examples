package armlabservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/labservices/armlabservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4c2cdccf6ca3281dd50ed8788ce1de2e0d480973/specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/LabServices/listOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlabservices.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armlabservices.OperationListResult{
		// 	Value: []*armlabservices.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.LabServices/labPlans/write"),
		// 			Display: &armlabservices.OperationDisplay{
		// 				Description: to.Ptr("Create new or update an existing lab plan."),
		// 				Operation: to.Ptr("Create or Update Lab Plan"),
		// 				Provider: to.Ptr("Microsoft Azure Lab Services"),
		// 				Resource: to.Ptr("Microsoft.LabServices/labPlans"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armlabservices.OriginUserSystem),
		// 	}},
		// }
	}
}
