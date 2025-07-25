package armpurestorageblock_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purestorageblock/armpurestorageblock"
)

// Generated from example definition: 2024-11-01/Operations_List_MaximumSet_Gen.json
func ExampleOperationsClient_NewListPager_operationsList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurestorageblock.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armpurestorageblock.OperationsClientListResponse{
		// 	OperationListResult: armpurestorageblock.OperationListResult{
		// 		Value: []*armpurestorageblock.Operation{
		// 			{
		// 				Name: to.Ptr("wikaab"),
		// 				IsDataAction: to.Ptr(true),
		// 				Display: &armpurestorageblock.OperationDisplay{
		// 					Provider: to.Ptr("oqlfyrrvbitoadnqagqkgxylxkow"),
		// 					Resource: to.Ptr("dxeqyes"),
		// 					Operation: to.Ptr("mk"),
		// 					Description: to.Ptr("zcgnhznothktczdlyovbbjwnxep"),
		// 				},
		// 				Origin: to.Ptr(armpurestorageblock.OriginUser),
		// 				ActionType: to.Ptr(armpurestorageblock.ActionTypeInternal),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
