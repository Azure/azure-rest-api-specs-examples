package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Compute/createOrUpdate/AmlCompute.json
func ExampleComputeClient_BeginCreateOrUpdate_updateAAmlCompute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewComputeClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "testrg123", "workspaces123", "compute123", armmachinelearning.ComputeResource{
		Properties: &armmachinelearning.AmlCompute{
			Properties: &armmachinelearning.AmlComputeProperties{
				ScaleSettings: &armmachinelearning.ScaleSettings{
					MaxNodeCount:                to.Ptr[int32](4),
					MinNodeCount:                to.Ptr[int32](4),
					NodeIdleTimeBeforeScaleDown: to.Ptr("PT5M"),
				},
			},
			Description: to.Ptr("some compute"),
			ComputeType: to.Ptr(armmachinelearning.ComputeTypeAmlCompute),
		},
		Location: to.Ptr("eastus"),
	}, nil)
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
