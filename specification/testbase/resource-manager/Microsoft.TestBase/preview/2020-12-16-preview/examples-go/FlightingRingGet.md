Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ftestbase%2Farmtestbase%2Fv0.2.1/sdk/resourcemanager/testbase/armtestbase/README.md) on how to add the SDK to your project and authenticate.

```go
package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// x-ms-original-file: specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/FlightingRingGet.json
func ExampleFlightingRingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armtestbase.NewFlightingRingsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<test-base-account-name>",
		"<flighting-ring-resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.FlightingRingsClientGetResult)
}
```
