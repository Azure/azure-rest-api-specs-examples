Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flogic%2Farmlogic%2Fv0.3.0/sdk/resourcemanager/logic/armlogic/README.md) on how to add the SDK to your project and authenticate.

```go
package armlogic_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_Patch.json
func ExampleIntegrationServiceEnvironmentsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogic.NewIntegrationServiceEnvironmentsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group>",
		"<integration-service-environment-name>",
		armlogic.IntegrationServiceEnvironment{
			Tags: map[string]*string{
				"tag1": to.StringPtr("value1"),
			},
			SKU: &armlogic.IntegrationServiceEnvironmentSKU{
				Name:     armlogic.IntegrationServiceEnvironmentSKUName("Developer").ToPtr(),
				Capacity: to.Int32Ptr(0),
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
	log.Printf("Response result: %#v\n", res.IntegrationServiceEnvironmentsClientUpdateResult)
}
```
