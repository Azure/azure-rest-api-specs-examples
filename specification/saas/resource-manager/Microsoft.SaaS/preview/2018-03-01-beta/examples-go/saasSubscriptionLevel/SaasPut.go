package armsaas_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/saas/resource-manager/Microsoft.SaaS/preview/2018-03-01-beta/examples/saasSubscriptionLevel/SaasPut.json
func ExampleSubscriptionLevelClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsaas.NewSubscriptionLevelClient("c825645b-e31b-9cf4-1cee-2aba9e58bc7c", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"my-saas-rg",
		"MyContosoSubscription",
		armsaas.ResourceCreation{
			Name:     to.Ptr("MyContosoSubscription"),
			Location: to.Ptr("global"),
			Properties: &armsaas.CreationProperties{
				OfferID: to.Ptr("contosoOffer"),
				PaymentChannelMetadata: map[string]*string{
					"AzureSubscriptionId": to.Ptr("155af98a-3205-47e7-883b-a2ab9db9f88d"),
				},
				PaymentChannelType: to.Ptr(armsaas.PaymentChannelTypeSubscriptionDelegated),
				PublisherID:        to.Ptr("microsoft-contoso"),
				SaasResourceName:   to.Ptr("MyContosoSubscription"),
				SKUID:              to.Ptr("free"),
				TermID:             to.Ptr("hjdtn7tfnxcy"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
