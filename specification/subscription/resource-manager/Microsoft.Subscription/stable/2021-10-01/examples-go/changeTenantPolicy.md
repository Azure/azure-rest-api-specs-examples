Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsubscription%2Farmsubscription%2Fv0.2.0/sdk/resourcemanager/subscription/armsubscription/README.md) on how to add the SDK to your project and authenticate.

```go
package armsubscription_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
)

// x-ms-original-file: specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/changeTenantPolicy.json
func ExamplePolicyClient_AddUpdatePolicyForTenant() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsubscription.NewPolicyClient(cred, nil)
	res, err := client.AddUpdatePolicyForTenant(ctx,
		armsubscription.PutTenantPolicyRequestProperties{
			BlockSubscriptionsIntoTenant:    to.BoolPtr(true),
			BlockSubscriptionsLeavingTenant: to.BoolPtr(true),
			ExemptedPrincipals: []*string{
				to.StringPtr("e879cf0f-2b4d-5431-109a-f72fc9868693"),
				to.StringPtr("9792da87-c97b-410d-a97d-27021ba09ce6")},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PolicyClientAddUpdatePolicyForTenantResult)
}
```
