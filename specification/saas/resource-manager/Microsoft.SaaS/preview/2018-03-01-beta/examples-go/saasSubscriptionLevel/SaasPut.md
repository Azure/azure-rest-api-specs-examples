Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsaas%2Farmsaas%2Fv0.4.0/sdk/resourcemanager/saas/armsaas/README.md) on how to add the SDK to your project and authenticate.

```go
package armsaas_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/saas/resource-manager/Microsoft.SaaS/preview/2018-03-01-beta/examples/saasSubscriptionLevel/SaasPut.json
func ExampleSubscriptionLevelClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsaas.NewSubscriptionLevelClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armsaas.ResourceCreation{
			Name:     to.Ptr("<name>"),
			Location: to.Ptr("<location>"),
			Properties: &armsaas.CreationProperties{
				OfferID: to.Ptr("<offer-id>"),
				PaymentChannelMetadata: map[string]*string{
					"AzureSubscriptionId": to.Ptr("155af98a-3205-47e7-883b-a2ab9db9f88d"),
				},
				PaymentChannelType: to.Ptr(armsaas.PaymentChannelTypeSubscriptionDelegated),
				PublisherID:        to.Ptr("<publisher-id>"),
				SaasResourceName:   to.Ptr("<saas-resource-name>"),
				SKUID:              to.Ptr("<skuid>"),
				TermID:             to.Ptr("<term-id>"),
			},
		},
		&armsaas.SubscriptionLevelClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
