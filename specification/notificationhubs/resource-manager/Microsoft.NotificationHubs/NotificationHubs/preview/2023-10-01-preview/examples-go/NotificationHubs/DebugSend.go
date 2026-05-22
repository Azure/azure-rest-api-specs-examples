package armnotificationhubs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs/v2"
)

// Generated from example definition: 2023-10-01-preview/NotificationHubs/DebugSend.json
func ExampleClient_DebugSend() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnotificationhubs.NewClientFactory("29cfa613-cbbc-4512-b1d6-1b3a92c7fa40", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().DebugSend(ctx, "5ktrial", "nh-sdk-ns", "nh-sdk-hub", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnotificationhubs.ClientDebugSendResponse{
	// 	DebugSendResponse: &armnotificationhubs.DebugSendResponse{
	// 		Name: to.Ptr("test"),
	// 		Type: to.Ptr("Microsoft.NotificationHubs/namespaces/notificationHubs/debugSend"),
	// 		ID: to.Ptr("/subscriptions/29cfa613-cbbc-4512-b1d6-1b3a92c7fa40/resourceGroups/5ktrial/providers/Microsoft.NotificationHubs/namespaces/nh-sdk-ns/notificationHubs/test"),
	// 		Properties: &armnotificationhubs.DebugSendResult{
	// 			Failure: to.Ptr[int64](0),
	// 			Results: []*armnotificationhubs.RegistrationResult{
	// 			},
	// 			Success: to.Ptr[int64](0),
	// 		},
	// 	},
	// }
}
