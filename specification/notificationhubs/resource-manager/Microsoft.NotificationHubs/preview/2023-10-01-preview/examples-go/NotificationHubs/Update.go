package armnotificationhubs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/85cfba195a19120f309bd292c4261aa53a586adb/specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/NotificationHubs/Update.json
func ExampleClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnotificationhubs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Update(ctx, "sdkresourceGroup", "nh-sdk-ns", "sdk-notificationHubs-8708", armnotificationhubs.NotificationHubPatchParameters{
		Properties: &armnotificationhubs.NotificationHubProperties{
			GCMCredential: &armnotificationhubs.GCMCredential{
				Properties: &armnotificationhubs.GCMCredentialProperties{
					GCMEndpoint:  to.Ptr("https://fcm.googleapis.com/fcm/send"),
					GoogleAPIKey: to.Ptr("###################################"),
				},
			},
			RegistrationTTL: to.Ptr("10675199.02:48:05.4775807"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NotificationHubResource = armnotificationhubs.NotificationHubResource{
	// 	Name: to.Ptr("test"),
	// 	Type: to.Ptr("Microsoft.NotificationHubs/namespaces/notificationHubs"),
	// 	ID: to.Ptr("/subscriptions/29cfa613-cbbc-4512-b1d6-1b3a92c7fa40/resourceGroups/5ktrial/providers/Microsoft.NotificationHubs/namespaces/nh-sdk-ns/notificationHubs/test"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"hubTag1": to.Ptr("hubTagValue1"),
	// 		"hubTag2": to.Ptr("hubTagValue2"),
	// 	},
	// 	Properties: &armnotificationhubs.NotificationHubProperties{
	// 		Name: to.Ptr("test"),
	// 		DailyMaxActiveDevices: to.Ptr[int64](0),
	// 		RegistrationTTL: to.Ptr("10675199.02:48:05.4775807"),
	// 	},
	// }
}
