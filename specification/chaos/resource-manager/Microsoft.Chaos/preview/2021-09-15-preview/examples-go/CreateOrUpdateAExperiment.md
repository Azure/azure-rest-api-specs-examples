Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fchaos%2Farmchaos%2Fv0.1.0/sdk/resourcemanager/chaos/armchaos/README.md) on how to add the SDK to your project and authenticate.

```go
package armchaos_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos"
)

// x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2021-09-15-preview/examples/CreateOrUpdateAExperiment.json
func ExampleExperimentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armchaos.NewExperimentsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<experiment-name>",
		armchaos.Experiment{
			TrackedResource: armchaos.TrackedResource{
				Location: to.StringPtr("<location>"),
			},
			Identity: &armchaos.ResourceIdentity{
				Type: armchaos.ResourceIdentityTypeSystemAssigned.ToPtr(),
			},
			Properties: &armchaos.ExperimentProperties{
				Selectors: []*armchaos.Selector{
					{
						Type: armchaos.SelectorTypeList.ToPtr(),
						ID:   to.StringPtr("<id>"),
						Targets: []*armchaos.TargetReference{
							{
								Type: to.StringPtr("<type>"),
								ID:   to.StringPtr("<id>"),
							}},
					}},
				Steps: []*armchaos.Step{
					{
						Name: to.StringPtr("<name>"),
						Branches: []*armchaos.Branch{
							{
								Name: to.StringPtr("<name>"),
								Actions: []armchaos.ActionClassification{
									&armchaos.Action{
										Name: to.StringPtr("<name>"),
										Type: to.StringPtr("<type>"),
									}},
							}},
					}},
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
