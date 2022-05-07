Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Faad%2Farmaad%2Fv0.4.0/sdk/resourcemanager/aad/armaad/README.md) on how to add the SDK to your project and authenticate.

```go
package armaad_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azureactivedirectory/resource-manager/Microsoft.Aadiam/stable/2020-03-01/examples/AzureADPrivateLinkPolicyCreate.json
func ExamplePrivateLinkForAzureAdClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armaad.NewPrivateLinkForAzureAdClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<policy-name>",
		armaad.PrivateLinkPolicy{
			Name:           to.Ptr("<name>"),
			AllTenants:     to.Ptr(false),
			OwnerTenantID:  to.Ptr("<owner-tenant-id>"),
			ResourceGroup:  to.Ptr("<resource-group>"),
			ResourceName:   to.Ptr("<resource-name>"),
			SubscriptionID: to.Ptr("<subscription-id>"),
			Tenants: []*string{
				to.Ptr("3616657d-1c80-41ae-9d83-2a2776f2c9be"),
				to.Ptr("727b6ef1-18ab-4627-ac95-3f9cd945ed87")},
		},
		&armaad.PrivateLinkForAzureAdClientBeginCreateOptions{ResumeToken: ""})
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
