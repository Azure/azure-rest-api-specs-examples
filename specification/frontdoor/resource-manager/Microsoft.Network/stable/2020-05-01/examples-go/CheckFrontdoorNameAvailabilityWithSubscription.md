Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ffrontdoor%2Farmfrontdoor%2Fv0.1.0/sdk/resourcemanager/frontdoor/armfrontdoor/README.md) on how to add the SDK to your project and authenticate.

```go
package armfrontdoor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/CheckFrontdoorNameAvailabilityWithSubscription.json
func ExampleFrontDoorNameAvailabilityWithSubscriptionClient_Check() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armfrontdoor.NewFrontDoorNameAvailabilityWithSubscriptionClient("<subscription-id>", cred, nil)
	_, err = client.Check(ctx,
		armfrontdoor.CheckNameAvailabilityInput{
			Name: to.StringPtr("<name>"),
			Type: armfrontdoor.ResourceTypeMicrosoftNetworkFrontDoorsFrontendEndpoints.ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
