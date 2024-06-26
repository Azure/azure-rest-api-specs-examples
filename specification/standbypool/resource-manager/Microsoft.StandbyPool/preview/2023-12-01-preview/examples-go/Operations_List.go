package armstandbypool_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/standbypool/armstandbypool"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/standbypool/resource-manager/Microsoft.StandbyPool/preview/2023-12-01-preview/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstandbypool.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armstandbypool.OperationListResult{
		// 	Value: []*armstandbypool.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.StandbyPool/standbyContainerGroupPools/write"),
		// 			ActionType: to.Ptr(armstandbypool.ActionTypeInternal),
		// 			Display: &armstandbypool.OperationDisplay{
		// 				Description: to.Ptr("Create a StandbyContainerGroupPools Resource"),
		// 				Operation: to.Ptr("StandbyContainerGroupPools_Create"),
		// 				Provider: to.Ptr("Microsoft.StandbyPool"),
		// 				Resource: to.Ptr("standbyContainerGroupPools"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 			Origin: to.Ptr(armstandbypool.OriginUser),
		// 	}},
		// }
	}
}
