package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/DeviceCapacityRequestPost.json
func ExampleDeviceCapacityCheckClient_BeginCheckResourceCreationFeasibility() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataboxedge.NewDeviceCapacityCheckClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCheckResourceCreationFeasibility(ctx,
		"GroupForEdgeAutomation",
		"testedgedevice",
		armdataboxedge.DeviceCapacityRequestInfo{
			Properties: &armdataboxedge.DeviceCapacityRequestInfoProperties{
				VMPlacementQuery: [][]*string{
					{
						to.Ptr("Standard_D2_v2")}},
			},
		},
		&armdataboxedge.DeviceCapacityCheckClientBeginCheckResourceCreationFeasibilityOptions{CapacityName: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
