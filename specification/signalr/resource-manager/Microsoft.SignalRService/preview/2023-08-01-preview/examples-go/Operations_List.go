package armsignalr_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsignalr.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationList = armsignalr.OperationList{
		// 	Value: []*armsignalr.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.SignalRService/SignalR/read"),
		// 			Display: &armsignalr.OperationDisplay{
		// 				Description: to.Ptr("View the resource settings and configurations in the management portal or through API"),
		// 				Operation: to.Ptr("Manage SignalR (read-only)"),
		// 				Provider: to.Ptr("Microsoft.SignalRService"),
		// 				Resource: to.Ptr("SignalR"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Properties: &armsignalr.OperationProperties{
		// 			},
		// 	}},
		// }
	}
}
