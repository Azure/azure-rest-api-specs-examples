package armsaas_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/saas/resource-manager/Microsoft.SaaS/preview/2018-03-01-beta/examples/saasSubscriptionLevel/SaasUnsubscribe.json
func ExampleSubscriptionLevelClient_BeginUpdateToUnsubscribed() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsaas.NewSubscriptionLevelClient("c825645b-e31b-9cf4-1cee-2aba9e58bc7c", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdateToUnsubscribed(ctx,
		"my-saas-rg",
		"MyContosoSubscription",
		armsaas.DeleteOptions{
			Feedback:        to.Ptr("No longer need this SaaS"),
			ReasonCode:      to.Ptr[float32](0),
			UnsubscribeOnly: to.Ptr(true),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
