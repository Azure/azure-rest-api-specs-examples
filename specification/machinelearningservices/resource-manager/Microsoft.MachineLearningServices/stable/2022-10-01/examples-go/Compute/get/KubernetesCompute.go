package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aafb0944f7ab936e8cfbad8969bd5eb32263fb4f/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Compute/get/KubernetesCompute.json
func ExampleComputeClient_Get_getAKubernetesCompute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComputeClient().Get(ctx, "testrg123", "workspaces123", "compute123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComputeResource = armmachinelearning.ComputeResource{
	// 	Properties: &armmachinelearning.Kubernetes{
	// 		Description: to.Ptr("some compute"),
	// 		ComputeType: to.Ptr(armmachinelearning.ComputeTypeKubernetes),
	// 		CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T22:00:00.000Z"); return t}()),
	// 		IsAttachedCompute: to.Ptr(true),
	// 		ModifiedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T22:00:00.000Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armmachinelearning.ProvisioningStateSucceeded),
	// 		ResourceID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/Microsoft.ContainerService/managedClusters/compute123-56826-c9b00420020b2"),
	// 		Properties: &armmachinelearning.KubernetesProperties{
	// 			DefaultInstanceType: to.Ptr("defaultInstanceType"),
	// 			ExtensionInstanceReleaseTrain: to.Ptr("stable"),
	// 			InstanceTypes: map[string]*armmachinelearning.InstanceTypeSchema{
	// 				"defaultInstanceType": &armmachinelearning.InstanceTypeSchema{
	// 					Resources: &armmachinelearning.InstanceTypeSchemaResources{
	// 						Limits: map[string]*string{
	// 							"cpu": to.Ptr("1"),
	// 							"memory": to.Ptr("4Gi"),
	// 							"nvidia.com/gpu": nil,
	// 						},
	// 						Requests: map[string]*string{
	// 							"cpu": to.Ptr("1"),
	// 							"memory": to.Ptr("4Gi"),
	// 							"nvidia.com/gpu": nil,
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Namespace: to.Ptr("default"),
	// 		},
	// 	},
	// 	Name: to.Ptr("compute123"),
	// 	Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/computes"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/workspaces123/computes/compute123"),
	// 	Location: to.Ptr("eastus"),
	// }
}
