package armdurabletask_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/durabletask/armdurabletask"
)

// Generated from example definition: 2026-02-01/PrivateEndpointConnections_List_MaximumSet_Gen.json
func ExampleSchedulersClient_NewListPrivateEndpointConnectionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdurabletask.NewClientFactory("851A7597-D699-45CC-899B-7487A5B3B775", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSchedulersClient().NewListPrivateEndpointConnectionsPager("rgdurabletask", "testscheduler", nil)
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
		// page = armdurabletask.SchedulersClientListPrivateEndpointConnectionsResponse{
		// 	PrivateEndpointConnectionListResult: armdurabletask.PrivateEndpointConnectionListResult{
		// 		Value: []*armdurabletask.PrivateEndpointConnection{
		// 			{
		// 				Properties: &armdurabletask.PrivateEndpointConnectionProperties{
		// 					PrivateEndpoint: &armdurabletask.PrivateEndpoint{
		// 						ID: to.Ptr("vjjxatyilmgjaervqztrmlpfodvbo"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armdurabletask.PrivateLinkServiceConnectionState{
		// 						Status: to.Ptr(armdurabletask.PrivateEndpointServiceConnectionStatusPending),
		// 						ActionsRequired: to.Ptr("mxymqfbbmpwjxsroldlsd"),
		// 						Description: to.Ptr("ujdcsoyxljivwsgfkexhotaxcmzq"),
		// 					},
		// 					GroupIDs: []*string{
		// 						to.Ptr("xnnrzmowptxnijdojrntrbm"),
		// 					},
		// 					ProvisioningState: to.Ptr(armdurabletask.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 				},
		// 				ID: to.Ptr("/subscriptions/851A7597-D699-45CC-899B-7487A5B3B775/resourceGroups/rgdurabletask/providers/Microsoft.DurableTask/schedulers/testscheduler/privateEndpointConnections/spzckqrbhfnabu"),
		// 				Name: to.Ptr("spzckqrbhfnabu"),
		// 				Type: to.Ptr("Microsoft.DurableTask/schedulers/privateEndpointConnections"),
		// 				SystemData: &armdurabletask.SystemData{
		// 					CreatedBy: to.Ptr("tenmbevaunjzikxowqexrsx"),
		// 					CreatedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-17T15:34:17.365Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("xfvdcegtj"),
		// 					LastModifiedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-17T15:34:17.366Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/anend"),
		// 	},
		// }
	}
}
