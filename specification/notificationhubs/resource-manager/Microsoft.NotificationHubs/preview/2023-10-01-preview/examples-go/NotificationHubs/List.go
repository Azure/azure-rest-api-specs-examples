package armnotificationhubs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/85cfba195a19120f309bd292c4261aa53a586adb/specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/NotificationHubs/List.json
func ExampleClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnotificationhubs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListPager("5ktrial", "nh-sdk-ns", &armnotificationhubs.ClientListOptions{SkipToken: nil,
		Top: nil,
	})
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
		// page.NotificationHubListResult = armnotificationhubs.NotificationHubListResult{
		// 	Value: []*armnotificationhubs.NotificationHubResource{
		// 		{
		// 			Name: to.Ptr("test"),
		// 			Type: to.Ptr("Microsoft.NotificationHubs/namespaces/notificationHubs"),
		// 			ID: to.Ptr("/subscriptions/29cfa613-cbbc-4512-b1d6-1b3a92c7fa40/resourceGroups/5ktrial/providers/Microsoft.NotificationHubs/namespaces/nh-sdk-ns/notificationHubs/test"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 				"hubTag1": to.Ptr("hubTagValue1"),
		// 				"hubTag2": to.Ptr("hubTagValue2"),
		// 			},
		// 			Properties: &armnotificationhubs.NotificationHubProperties{
		// 				Name: to.Ptr("test"),
		// 				DailyMaxActiveDevices: to.Ptr[int64](0),
		// 				RegistrationTTL: to.Ptr("10675199.02:48:05.4775807"),
		// 			},
		// 	}},
		// }
	}
}
