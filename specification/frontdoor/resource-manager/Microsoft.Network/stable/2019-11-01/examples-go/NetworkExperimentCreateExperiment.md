Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ffrontdoor%2Farmfrontdoor%2Fv0.1.0/sdk/resourcemanager/frontdoor/armfrontdoor/README.md) on how to add the SDK to your project and authenticate.

```go
package armfrontdoor_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentCreateExperiment.json
func ExampleExperimentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armfrontdoor.NewExperimentsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<experiment-name>",
		armfrontdoor.Experiment{
			Properties: &armfrontdoor.ExperimentProperties{
				Description:  to.StringPtr("<description>"),
				EnabledState: armfrontdoor.StateEnabled.ToPtr(),
				EndpointA: &armfrontdoor.Endpoint{
					Name:     to.StringPtr("<name>"),
					Endpoint: to.StringPtr("<endpoint>"),
				},
				EndpointB: &armfrontdoor.Endpoint{
					Name:     to.StringPtr("<name>"),
					Endpoint: to.StringPtr("<endpoint>"),
				},
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
	log.Printf("Experiment.ID: %s\n", *res.ID)
}
```
