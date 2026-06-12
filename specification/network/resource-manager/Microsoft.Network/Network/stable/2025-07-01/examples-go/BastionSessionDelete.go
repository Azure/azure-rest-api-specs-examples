package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/BastionSessionDelete.json
func ExampleManagementClient_NewDisconnectActiveSessionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagementClient().NewDisconnectActiveSessionsPager("rg1", "bastionhosttenant", armnetwork.SessionIDs{
		SessionIDs: []*string{
			to.Ptr("session1"),
			to.Ptr("session2"),
			to.Ptr("session3"),
		},
	}, nil)
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
		// page = armnetwork.ManagementClientDisconnectActiveSessionsResponse{
		// 	BastionSessionDeleteResult: armnetwork.BastionSessionDeleteResult{
		// 		Value: []*armnetwork.BastionSessionState{
		// 			{
		// 				Message: to.Ptr("session session1 invalidated!"),
		// 				SessionID: to.Ptr("session1"),
		// 				State: to.Ptr("Disconnected"),
		// 			},
		// 			{
		// 				Message: to.Ptr("session session2 could not be disconnected!"),
		// 				SessionID: to.Ptr("session2"),
		// 				State: to.Ptr("Failed"),
		// 			},
		// 			{
		// 				Message: to.Ptr("session session3 not found!"),
		// 				SessionID: to.Ptr("session3"),
		// 				State: to.Ptr("NotFound"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
