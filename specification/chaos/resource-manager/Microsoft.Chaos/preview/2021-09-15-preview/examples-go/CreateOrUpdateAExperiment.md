Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fchaos%2Farmchaos%2Fv0.5.0/sdk/resourcemanager/chaos/armchaos/README.md) on how to add the SDK to your project and authenticate.

```go
package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/chaos/resource-manager/Microsoft.Chaos/preview/2021-09-15-preview/examples/CreateOrUpdateAExperiment.json
func ExampleExperimentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armchaos.NewExperimentsClient("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"exampleRG",
		"exampleExperiment",
		armchaos.Experiment{
			Location: to.Ptr("centraluseuap"),
			Identity: &armchaos.ResourceIdentity{
				Type: to.Ptr(armchaos.ResourceIdentityTypeSystemAssigned),
			},
			Properties: &armchaos.ExperimentProperties{
				Selectors: []*armchaos.Selector{
					{
						Type: to.Ptr(armchaos.SelectorTypeList),
						ID:   to.Ptr("selector1"),
						Targets: []*armchaos.TargetReference{
							{
								Type: to.Ptr("ChaosTarget"),
								ID:   to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM/providers/Microsoft.Chaos/targets/Microsoft-VirtualMachine"),
							}},
					}},
				Steps: []*armchaos.Step{
					{
						Name: to.Ptr("step1"),
						Branches: []*armchaos.Branch{
							{
								Name: to.Ptr("branch1"),
								Actions: []armchaos.ActionClassification{
									&armchaos.Action{
										Name: to.Ptr("urn:csci:provider:providername:Shutdown/1.0"),
										Type: to.Ptr("Continuous"),
									}},
							}},
					}},
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
