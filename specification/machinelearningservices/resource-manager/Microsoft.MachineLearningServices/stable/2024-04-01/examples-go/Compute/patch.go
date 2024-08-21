package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Compute/patch.json
func ExampleComputeClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewComputeClient().BeginUpdate(ctx, "testrg123", "workspaces123", "compute123", armmachinelearning.ClusterUpdateParameters{
		Properties: &armmachinelearning.ClusterUpdateProperties{
			Properties: &armmachinelearning.ScaleSettingsInformation{
				ScaleSettings: &armmachinelearning.ScaleSettings{
					MaxNodeCount:                to.Ptr[int32](4),
					MinNodeCount:                to.Ptr[int32](4),
					NodeIdleTimeBeforeScaleDown: to.Ptr("PT5M"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComputeResource = armmachinelearning.ComputeResource{
	// 	Properties: &armmachinelearning.AmlCompute{
	// 		Description: to.Ptr("some compute"),
	// 		ComputeType: to.Ptr(armmachinelearning.ComputeTypeAmlCompute),
	// 		ProvisioningState: to.Ptr(armmachinelearning.ProvisioningStateSucceeded),
	// 	},
	// 	Name: to.Ptr("compute123"),
	// 	Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/computes"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/workspaces123/computes/compute123"),
	// 	Location: to.Ptr("eastus2"),
	// }
}
