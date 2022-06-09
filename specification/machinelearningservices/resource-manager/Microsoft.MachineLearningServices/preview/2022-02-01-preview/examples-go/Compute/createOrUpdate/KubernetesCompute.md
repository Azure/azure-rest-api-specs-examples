```go
package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Compute/createOrUpdate/KubernetesCompute.json
func ExampleComputeClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewComputeClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"testrg123",
		"workspaces123",
		"compute123",
		armmachinelearning.ComputeResource{
			Properties: &armmachinelearning.Kubernetes{
				Description: to.Ptr("some compute"),
				ComputeType: to.Ptr(armmachinelearning.ComputeTypeKubernetes),
				ResourceID:  to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/Microsoft.ContainerService/managedClusters/compute123-56826-c9b00420020b2"),
				Properties: &armmachinelearning.KubernetesProperties{
					DefaultInstanceType: to.Ptr("defaultInstanceType"),
					InstanceTypes: map[string]*armmachinelearning.InstanceTypeSchema{
						"defaultInstanceType": {
							Resources: &armmachinelearning.InstanceTypeSchemaResources{
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmachinelearning%2Farmmachinelearning%2Fv2.0.0-beta.1/sdk/resourcemanager/machinelearning/armmachinelearning/README.md) on how to add the SDK to your project and authenticate.
