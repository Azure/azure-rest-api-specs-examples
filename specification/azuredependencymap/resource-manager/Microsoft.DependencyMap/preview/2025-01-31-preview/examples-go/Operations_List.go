package armdependencymap_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dependencymap/armdependencymap"
)

// Generated from example definition: 2025-01-31-preview/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdependencymap.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armdependencymap.OperationsClientListResponse{
		// 	OperationListResult: armdependencymap.OperationListResult{
		// 		Value: []*armdependencymap.Operation{
		// 			{
		// 				Name: to.Ptr("qzkexod"),
		// 				IsDataAction: to.Ptr(true),
		// 				Display: &armdependencymap.OperationDisplay{
		// 					Provider: to.Ptr("edcbcbkhk"),
		// 					Resource: to.Ptr("gpuxsogqqoaczsvhznieqtlhqgiuln"),
		// 					Operation: to.Ptr("nlu"),
		// 					Description: to.Ptr("cxjpanznyqeoxxjywlfcjkwrl"),
		// 				},
		// 				Origin: to.Ptr(armdependencymap.OriginUser),
		// 				ActionType: to.Ptr(armdependencymap.ActionTypeInternal),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("ql"),
		// 	},
		// }
	}
}
