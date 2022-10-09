package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/capacityReservationExamples/CapacityReservationGroup_Update_MaximumSet_Gen.json
func ExampleCapacityReservationGroupsClient_Update_capacityReservationGroupsUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewCapacityReservationGroupsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaaaaa", armcompute.CapacityReservationGroupUpdate{
		Tags: map[string]*string{
			"key5355": to.Ptr("aaa"),
		},
		Properties: &armcompute.CapacityReservationGroupProperties{
			InstanceView: &armcompute.CapacityReservationGroupInstanceView{},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
