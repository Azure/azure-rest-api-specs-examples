package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/20312e2b31df58f0ea7560e87062d62aa92f0a14/specification/redis/resource-manager/Microsoft.Cache/stable/2023-08-01/examples/RedisCacheListUpgradeNotifications.json
func ExampleClient_NewListUpgradeNotificationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListUpgradeNotificationsPager("rg1", "cache1", 5000, nil)
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
		// page.NotificationListResponse = armredis.NotificationListResponse{
		// 	Value: []*armredis.UpgradeNotification{
		// 		{
		// 			Name: to.Ptr("notification1"),
		// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-16T23:20:50.52Z"); return t}()),
		// 			UpsellNotification: map[string]*string{
		// 			},
		// 	}},
		// }
	}
}
