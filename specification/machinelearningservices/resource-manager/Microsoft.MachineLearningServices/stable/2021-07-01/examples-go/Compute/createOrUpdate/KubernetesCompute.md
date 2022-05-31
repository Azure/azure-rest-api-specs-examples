Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmachinelearningservices%2Farmmachinelearningservices%2Fv1.0.1/sdk/resourcemanager/machinelearningservices/armmachinelearningservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armmachinelearningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearningservices/armmachinelearningservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Compute/createOrUpdate/KubernetesCompute.json
func ExampleComputeClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewComputeClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"testrg123",
		"workspaces123",
		"compute123",
		armmachinelearningservices.ComputeResource{
			Properties: &armmachinelearningservices.Kubernetes{
				Description: to.Ptr("some compute"),
				ComputeType: to.Ptr(armmachinelearningservices.ComputeTypeKubernetes),
				ResourceID:  to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/Microsoft.ContainerService/managedClusters/compute123-56826-c9b00420020b2"),
				Properties: &armmachinelearningservices.KubernetesProperties{
					DefaultInstanceType: to.Ptr("defaultInstanceType"),
					InstanceTypes: map[string]*armmachinelearningservices.InstanceTypeSchema{
						"defaultInstanceType": {
							Resources: &armmachinelearningservices.InstanceTypeSchemaResources{
								Limits: map[string]*string{
									"cpu":            to.Ptr("1"),
									"memory":         to.Ptr("4Gi"),
									"nvidia.com/gpu": nil,
								},
								Requests: map[string]*string{
									"cpu":            to.Ptr("1"),
									"memory":         to.Ptr("4Gi"),
									"nvidia.com/gpu": nil,
								},
							},
						},
					},
					Namespace: to.Ptr("default"),
				},
			},
			Location: to.Ptr("eastus"),
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
