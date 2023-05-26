package armsubscriptions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4f4073bdb028bc84bc3e6405c1cbaf8e89b83caf/specification/resources/resource-manager/Microsoft.Resources/stable/2022-12-01/examples/GetSubscription.json
func ExampleClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsubscriptions.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Get(ctx, "291bba3f-e0a5-47bc-a099-3bdcb2a50a05", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Subscription = armsubscriptions.Subscription{
	// 	AuthorizationSource: to.Ptr("Bypassed"),
	// 	DisplayName: to.Ptr("Example Subscription"),
	// 	ID: to.Ptr("/subscriptions/291bba3f-e0a5-47bc-a099-3bdcb2a50a05"),
	// 	ManagedByTenants: []*armsubscriptions.ManagedByTenant{
	// 		{
	// 			TenantID: to.Ptr("8f70baf1-1f6e-46a2-a1ff-238dac1ebfb7"),
	// 	}},
	// 	State: to.Ptr(armsubscriptions.SubscriptionStateEnabled),
	// 	SubscriptionID: to.Ptr("291bba3f-e0a5-47bc-a099-3bdcb2a50a05"),
	// 	SubscriptionPolicies: &armsubscriptions.SubscriptionPolicies{
	// 		LocationPlacementID: to.Ptr("Internal_2014-09-01"),
	// 		QuotaID: to.Ptr("Internal_2014-09-01"),
	// 		SpendingLimit: to.Ptr(armsubscriptions.SpendingLimitOff),
	// 	},
	// 	Tags: map[string]*string{
	// 		"tagKey1": to.Ptr("tagValue1"),
	// 		"tagKey2": to.Ptr("tagValue2"),
	// 	},
	// 	TenantID: to.Ptr("31c75423-32d6-4322-88b7-c478bdde4858"),
	// }
}
