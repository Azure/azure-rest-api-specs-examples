package armdatatransfer_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datatransfer/armdatatransfer"
)

// Generated from example definition: 2026-02-06-preview/listConnectionsBySubscription.json
func ExampleConnectionsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatatransfer.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectionsClient().NewListBySubscriptionPager(nil)
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
		// page = armdatatransfer.ConnectionsClientListBySubscriptionResponse{
		// 	ConnectionListResult: armdatatransfer.ConnectionListResult{
		// 		Value: []*armdatatransfer.Connection{
		// 			{
		// 				Name: to.Ptr("testConnection"),
		// 				Type: to.Ptr("Microsoft.AzureDataTransfer/connections"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.AzureDataTransfer/connections/testConnection"),
		// 				Location: to.Ptr("East US"),
		// 				Properties: &armdatatransfer.ConnectionProperties{
		// 					Justification: to.Ptr("justification"),
		// 					Pipeline: to.Ptr("testdc"),
		// 					RequirementID: to.Ptr("id"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("testConnection2"),
		// 				Type: to.Ptr("Microsoft.AzureDataTransfer/connections"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.AzureDataTransfer/connections/testConnection2"),
		// 				Location: to.Ptr("East US"),
		// 				Properties: &armdatatransfer.ConnectionProperties{
		// 					Justification: to.Ptr("justification"),
		// 					Pipeline: to.Ptr("testdc"),
		// 					RequirementID: to.Ptr("id"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
