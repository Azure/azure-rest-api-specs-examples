package armsubscription_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription/v2"
)

// Generated from example definition: 2025-11-01-preview/listTargetDirectory.json
func ExampleSubscriptionsClient_NewListTargetDirectoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsubscription.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSubscriptionsClient().NewListTargetDirectoryPager("ebe4f8fd-d8b3-4867-bcf4-b2407edd196d", nil)
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
		// page = armsubscription.SubscriptionsClientListTargetDirectoryResponse{
		// 	TargetDirectoryListResult: armsubscription.TargetDirectoryListResult{
		// 		Value: []*armsubscription.TargetDirectoryResult{
		// 			{
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.Subscription/changeTenantRequest"),
		// 				ID: to.Ptr("/subscriptions/ebe4f8fd-d8b3-4867-bcf4-b2407edd196d/providers/Microsoft.Subscription/changeTenantRequest/default"),
		// 				Properties: &armsubscription.TargetDirectoryResultProperties{
		// 					AcceptedDate: nil,
		// 					CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-18T17:57:40.0278346Z"); return t}()),
		// 					DestinationOwnerID: to.Ptr("b11a05c8-6acc-435e-9a51-2140dea093a5"),
		// 					DestinationTenantID: to.Ptr("45ffe2da-b7a4-460f-9e4c-51afd47b94cb"),
		// 					ExpiresOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-19T17:57:40.0278346Z"); return t}()),
		// 					SourceOwnerEmail: to.Ptr("alice@contso.com"),
		// 					SourceOwnerID: to.Ptr("45ffe2da-b7a4-460f-9e4c-51afd47b94cd"),
		// 					SourceTenantID: to.Ptr("f3a0f89e-12ab-4bcd-8e9f-0123456789ab"),
		// 					Status: to.Ptr(armsubscription.ChangeDirectoryOperationStatus("Initiated")),
		// 					SubscriptionID: to.Ptr("ebe4f8fd-d8b3-4867-bcf4-b2407edd196d"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
