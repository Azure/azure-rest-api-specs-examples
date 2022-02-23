Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsaas%2Farmsaas%2Fv0.2.1/sdk/resourcemanager/saas/armsaas/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/saas/resource-manager/Microsoft.SaaS/preview/2018-03-01-beta/examples/saasSubscriptionLevel/SaasPut.json
func ExampleSubscriptionLevelClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsaas.NewSubscriptionLevelClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armsaas.ResourceCreation{
			Name:     to.StringPtr("<name>"),
			Location: to.StringPtr("<location>"),
			Properties: &armsaas.CreationProperties{
				OfferID: to.StringPtr("<offer-id>"),
				PaymentChannelMetadata: map[string]*string{
					"AzureSubscriptionId": to.StringPtr("155af98a-3205-47e7-883b-a2ab9db9f88d"),
				},
				PaymentChannelType: armsaas.PaymentChannelType("SubscriptionDelegated").ToPtr(),
				PublisherID:        to.StringPtr("<publisher-id>"),
				SaasResourceName:   to.StringPtr("<saas-resource-name>"),
				SKUID:              to.StringPtr("<skuid>"),
				TermID:             to.StringPtr("<term-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SubscriptionLevelClientCreateOrUpdateResult)
}
```
