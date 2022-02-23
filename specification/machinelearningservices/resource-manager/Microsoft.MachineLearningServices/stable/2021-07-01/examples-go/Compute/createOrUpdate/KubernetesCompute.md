Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmachinelearningservices%2Farmmachinelearningservices%2Fv0.2.1/sdk/resourcemanager/machinelearningservices/armmachinelearningservices/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Compute/createOrUpdate/KubernetesCompute.json
func ExampleComputeClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmachinelearningservices.NewComputeClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<compute-name>",
		armmachinelearningservices.ComputeResource{
			Properties: &armmachinelearningservices.Kubernetes{
				Description: to.StringPtr("<description>"),
				ComputeType: armmachinelearningservices.ComputeType("Kubernetes").ToPtr(),
				ResourceID:  to.StringPtr("<resource-id>"),
				Properties: &armmachinelearningservices.KubernetesProperties{
					DefaultInstanceType: to.StringPtr("<default-instance-type>"),
					InstanceTypes: map[string]*armmachinelearningservices.InstanceTypeSchema{
						"defaultInstanceType": {
							Resources: &armmachinelearningservices.InstanceTypeSchemaResources{
								Limits: map[string]*string{
									"cpu":            to.StringPtr("1"),
									"memory":         to.StringPtr("4Gi"),
									"nvidia.com/gpu": nil,
								},
								Requests: map[string]*string{
									"cpu":            to.StringPtr("1"),
									"memory":         to.StringPtr("4Gi"),
									"nvidia.com/gpu": nil,
								},
							},
						},
					},
					Namespace: to.StringPtr("<namespace>"),
				},
			},
			Location: to.StringPtr("<location>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ComputeClientCreateOrUpdateResult)
}
```
