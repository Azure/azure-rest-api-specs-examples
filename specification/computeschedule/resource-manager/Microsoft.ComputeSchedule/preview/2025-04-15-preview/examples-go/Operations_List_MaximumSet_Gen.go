package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule"
)

// Generated from example definition: 2025-04-15-preview/Operations_List_MaximumSet_Gen.json
func ExampleOperationsClient_NewListPager_operationsListMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armcomputeschedule.OperationsClientListResponse{
		// 	OperationListResult: armcomputeschedule.OperationListResult{
		// 		Value: []*armcomputeschedule.Operation{
		// 			{
		// 				Name: to.Ptr("tew"),
		// 				IsDataAction: to.Ptr(true),
		// 				Display: &armcomputeschedule.OperationDisplay{
		// 					Provider: to.Ptr("kphlo"),
		// 					Resource: to.Ptr("vjhxmakxdtbnqhyjhclbargzt"),
		// 					Operation: to.Ptr("mmfsfkofhhfvgeryxbpdu"),
		// 					Description: to.Ptr("trdrldwroctmfwumqodfdxiuts"),
		// 				},
		// 				Origin: to.Ptr(armcomputeschedule.OriginUser),
		// 				ActionType: to.Ptr(armcomputeschedule.ActionTypeInternal),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
