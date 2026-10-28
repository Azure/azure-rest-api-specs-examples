package armdatatransfer_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datatransfer/armdatatransfer"
)

// Generated from example definition: 2026-02-06-preview/listPendingConnections.json
func ExampleConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatatransfer.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectionsClient().NewListPager("testRG", "testConnection", nil)
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
		// page = armdatatransfer.ConnectionsClientListResponse{
		// 	PendingConnectionListResult: armdatatransfer.PendingConnectionListResult{
		// 		Value: []*armdatatransfer.PendingConnection{
		// 			{
		// 				Name: to.Ptr("testConnection"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.AzureDataTransfer/connections/testConnection"),
		// 				Justification: to.Ptr("justification"),
		// 				Location: to.Ptr("East US"),
		// 				Pipeline: to.Ptr("testdc"),
		// 				RequirementID: to.Ptr("id"),
		// 				SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			},
		// 			{
		// 				Name: to.Ptr("testConnection"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.AzureDataTransfer/connections/testConnection"),
		// 				Justification: to.Ptr("justification"),
		// 				Location: to.Ptr("East US"),
		// 				Pipeline: to.Ptr("testdc"),
		// 				RequirementID: to.Ptr("id"),
		// 				SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
