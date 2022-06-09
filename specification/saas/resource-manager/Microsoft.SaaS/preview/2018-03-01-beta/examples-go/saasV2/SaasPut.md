```go
package armsaas_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/saas/resource-manager/Microsoft.SaaS/preview/2018-03-01-beta/examples/saasV2/SaasPut.json
func ExampleClient_BeginCreateResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsaas.NewClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateResource(ctx,
		armsaas.ResourceCreation{
			Properties: &armsaas.CreationProperties{
				OfferID: to.Ptr("microsofthealthcarebot"),
				PaymentChannelMetadata: map[string]*string{
					"AzureSubscriptionId": to.Ptr("155af98a-3205-47e7-883b-a2ab9db9f88d"),
				},
				PaymentChannelType: to.Ptr(armsaas.PaymentChannelTypeSubscriptionDelegated),
				PublisherID:        to.Ptr("microsoft-hcb"),
				SaasResourceName:   to.Ptr("testRunnerFromArm"),
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsaas%2Farmsaas%2Fv0.5.0/sdk/resourcemanager/saas/armsaas/README.md) on how to add the SDK to your project and authenticate.
