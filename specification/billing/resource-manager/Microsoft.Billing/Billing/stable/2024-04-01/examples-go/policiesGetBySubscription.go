package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/policiesGetBySubscription.json
func ExamplePoliciesClient_GetBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPoliciesClient().GetBySubscription(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbilling.PoliciesClientGetBySubscriptionResponse{
	// 	SubscriptionPolicy: armbilling.SubscriptionPolicy{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.Billing/policies"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Billing/policies/default"),
	// 		Properties: &armbilling.SubscriptionPolicyProperties{
	// 			Policies: []*armbilling.PolicySummary{
	// 				{
	// 					Name: to.Ptr("ViewCharges"),
	// 					PolicyType: to.Ptr(armbilling.PolicyTypeSystemControlled),
	// 					Scope: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000"),
	// 					Value: to.Ptr("Allowed"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
