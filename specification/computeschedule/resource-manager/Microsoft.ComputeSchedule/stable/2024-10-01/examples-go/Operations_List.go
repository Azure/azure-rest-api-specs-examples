package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a4eca6b060cf70da696963245656fdc440b666b/specification/computeschedule/resource-manager/Microsoft.ComputeSchedule/stable/2024-10-01/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armcomputeschedule.OperationListResult{
		// 	Value: []*armcomputeschedule.Operation{
		// 		{
		// 			Name: to.Ptr("mtiwosbky"),
		// 			ActionType: to.Ptr(armcomputeschedule.ActionTypeInternal),
		// 			Display: &armcomputeschedule.OperationDisplay{
		// 				Description: to.Ptr("moyje"),
		// 				Operation: to.Ptr("tuneyqwanedwnnbztrmq"),
		// 				Provider: to.Ptr("vtlhmqtfhlyllnplzpdpq"),
		// 				Resource: to.Ptr("epj"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 			Origin: to.Ptr(armcomputeschedule.OriginUser),
		// 	}},
		// }
	}
}
