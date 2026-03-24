package armsubscription_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription/v2"
)

// Generated from example definition: 2025-11-01-preview/acceptOwnershipStatus.json
func ExampleClient_AcceptOwnershipStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsubscription.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().AcceptOwnershipStatus(ctx, "291bba3f-e0a5-47bc-a099-3bdcb2a50a05", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsubscription.ClientAcceptOwnershipStatusResponse{
	// 	AcceptOwnershipStatusResponse: &armsubscription.AcceptOwnershipStatusResponse{
	// 		AcceptOwnershipState: to.Ptr(armsubscription.AcceptOwnershipPending),
	// 		BillingOwner: to.Ptr("abc@test.com"),
	// 		DisplayName: to.Ptr("Test Subscription"),
	// 		ProvisioningState: to.Ptr(armsubscription.ProvisioningAccepted),
	// 		SubscriptionID: to.Ptr("291bba3f-e0a5-47bc-a099-3bdcb2a50a05"),
	// 		SubscriptionTenantID: to.Ptr("6c541ca7-1cab-4ea0-adde-6305e1d534e2"),
	// 		Tags: map[string]*string{
	// 			"tag1": to.Ptr("TagValue1"),
	// 		},
	// 	},
	// }
}
