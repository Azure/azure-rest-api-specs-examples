package armnotificationhubs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs/v2"
)

// Generated from example definition: 2023-10-01-preview/Namespaces/PrivateLinkResourceList.json
func ExamplePrivateEndpointConnectionsClient_NewListGroupIDsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnotificationhubs.NewClientFactory("29cfa613-cbbc-4512-b1d6-1b3a92c7fa40", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListGroupIDsPager("5ktrial", "nh-sdk-ns", nil)
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
		// page = armnotificationhubs.PrivateEndpointConnectionsClientListGroupIDsResponse{
		// 	PrivateLinkResourceListResult: armnotificationhubs.PrivateLinkResourceListResult{
		// 		Value: []*armnotificationhubs.PrivateLinkResource{
		// 			{
		// 				Name: to.Ptr("namespace"),
		// 				Type: to.Ptr("Microsoft.NotificationHubs/namespaces/privateLinkResources"),
		// 				ID: to.Ptr("/subscriptions/29cfa613-cbbc-4512-b1d6-1b3a92c7fa40/resourceGroups/5ktrial/providers/Microsoft.NotificationHubs/namespaces/nh-sdk-ns/privateLinkResources/namespace"),
		// 				Properties: &armnotificationhubs.PrivateLinkResourceProperties{
		// 					GroupID: to.Ptr("namespace"),
		// 					RequiredMembers: []*string{
		// 						to.Ptr("namespace"),
		// 					},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.servicebus.windows.net"),
		// 						to.Ptr("privatelink.notificationhub.windows.net"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
