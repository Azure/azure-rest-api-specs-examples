Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Foep%2Farmoep%2Fv0.3.0/sdk/resourcemanager/oep/armoep/README.md) on how to add the SDK to your project and authenticate.

```go
package armoep_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oep/armoep"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/oep/resource-manager/Microsoft.OpenEnergyPlatform/preview/2021-06-01-preview/examples/Locations_CheckNameAvailability.json
func ExampleLocationsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armoep.NewLocationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CheckNameAvailability(ctx,
		armoep.CheckNameAvailabilityRequest{
			Name: to.Ptr("<name>"),
			Type: to.Ptr("<type>"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
