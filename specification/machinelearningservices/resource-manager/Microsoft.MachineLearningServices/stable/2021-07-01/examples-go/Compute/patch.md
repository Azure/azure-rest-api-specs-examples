Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmachinelearningservices%2Farmmachinelearningservices%2Fv0.4.0/sdk/resourcemanager/machinelearningservices/armmachinelearningservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armmachinelearningservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearningservices/armmachinelearningservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Compute/patch.json
func ExampleComputeClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewComputeClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<compute-name>",
		armmachinelearningservices.ClusterUpdateParameters{
			Properties: &armmachinelearningservices.ClusterUpdateProperties{
				Properties: &armmachinelearningservices.ScaleSettingsInformation{
					ScaleSettings: &armmachinelearningservices.ScaleSettings{
						MaxNodeCount:                to.Ptr[int32](4),
						MinNodeCount:                to.Ptr[int32](4),
						NodeIdleTimeBeforeScaleDown: to.Ptr("<node-idle-time-before-scale-down>"),
					},
				},
			},
		},
		&armmachinelearningservices.ComputeClientBeginUpdateOptions{ResumeToken: ""})
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
