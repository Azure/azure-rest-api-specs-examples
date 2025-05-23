package armstandbypool_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/standbypool/armstandbypool/v2"
)

// Generated from example definition: 2025-03-01/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstandbypool.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armstandbypool.OperationsClientListResponse{
		// 	OperationListResult: armstandbypool.OperationListResult{
		// 		Value: []*armstandbypool.Operation{
		// 			{
		// 				Name: to.Ptr("Microsoft.StandbyPool/standbyContainerGroupPools/write"),
		// 				IsDataAction: to.Ptr(true),
		// 				Display: &armstandbypool.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.StandbyPool"),
		// 					Resource: to.Ptr("standbyContainerGroupPools"),
		// 					Operation: to.Ptr("StandbyContainerGroupPools_Create"),
		// 					Description: to.Ptr("Create a StandbyContainerGroupPools Resource"),
		// 				},
		// 				Origin: to.Ptr(armstandbypool.OriginUser),
		// 				ActionType: to.Ptr(armstandbypool.ActionTypeInternal),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://example.com/providers/Microsoft.StandbyPool/operations"),
		// 	},
		// }
	}
}
