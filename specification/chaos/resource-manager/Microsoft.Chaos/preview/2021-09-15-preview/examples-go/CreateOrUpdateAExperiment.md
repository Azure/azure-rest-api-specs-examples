Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fchaos%2Farmchaos%2Fv0.4.0/sdk/resourcemanager/chaos/armchaos/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/chaos/resource-manager/Microsoft.Chaos/preview/2021-09-15-preview/examples/CreateOrUpdateAExperiment.json
func ExampleExperimentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armchaos.NewExperimentsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<experiment-name>",
		armchaos.Experiment{
			Location: to.Ptr("<location>"),
			Identity: &armchaos.ResourceIdentity{
				Type: to.Ptr(armchaos.ResourceIdentityTypeSystemAssigned),
			},
			Properties: &armchaos.ExperimentProperties{
				Selectors: []*armchaos.Selector{
					{
						Type: to.Ptr(armchaos.SelectorTypeList),
						ID:   to.Ptr("<id>"),
						Targets: []*armchaos.TargetReference{
							{
								Type: to.Ptr("<type>"),
								ID:   to.Ptr("<id>"),
							}},
					}},
				Steps: []*armchaos.Step{
					{
						Name: to.Ptr("<name>"),
						Branches: []*armchaos.Branch{
							{
								Name: to.Ptr("<name>"),
								Actions: []armchaos.ActionClassification{
									&armchaos.Action{
										Name: to.Ptr("<name>"),
										Type: to.Ptr("<type>"),
									}},
							}},
					}},
			},
		},
		&armchaos.ExperimentsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
