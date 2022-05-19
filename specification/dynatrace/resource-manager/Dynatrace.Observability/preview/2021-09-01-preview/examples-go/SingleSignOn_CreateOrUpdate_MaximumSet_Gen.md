Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdynatrace%2Farmdynatrace%2Fv0.1.0/sdk/resourcemanager/dynatrace/armdynatrace/README.md) on how to add the SDK to your project and authenticate.

```go
package armdynatrace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dynatrace/armdynatrace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/SingleSignOn_CreateOrUpdate_MaximumSet_Gen.json
func ExampleSingleSignOnClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdynatrace.NewSingleSignOnClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myMonitor",
		"default",
		armdynatrace.SingleSignOnResource{
			Properties: &armdynatrace.SingleSignOnProperties{
				AADDomains: []*string{
					to.Ptr("mpliftrdt20210811outlook.onmicrosoft.com")},
				EnterpriseAppID:   to.Ptr("00000000-0000-0000-0000-000000000000"),
				ProvisioningState: to.Ptr(armdynatrace.ProvisioningStateAccepted),
				SingleSignOnState: to.Ptr(armdynatrace.SingleSignOnStatesEnable),
				SingleSignOnURL:   to.Ptr("https://www.dynatrace.io"),
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
