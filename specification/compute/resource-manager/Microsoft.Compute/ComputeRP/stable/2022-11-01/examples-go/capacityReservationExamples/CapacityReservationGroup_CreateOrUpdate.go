package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/19f98c9f526f8db961f172276dd6d6882a86ed86/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/capacityReservationExamples/CapacityReservationGroup_CreateOrUpdate.json
func ExampleCapacityReservationGroupsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCapacityReservationGroupsClient().CreateOrUpdate(ctx, "myResourceGroup", "myCapacityReservationGroup", armcompute.CapacityReservationGroup{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"department": to.Ptr("finance"),
		},
		Zones: []*string{
			to.Ptr("1"),
			to.Ptr("2")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CapacityReservationGroup = armcompute.CapacityReservationGroup{
	// 	Name: to.Ptr("myCapacityReservationGroup"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/CapacityReservationGroups/myCapacityReservationGroup"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"department": to.Ptr("finance"),
	// 		"owner": to.Ptr("myCompany"),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1"),
	// 		to.Ptr("2")},
	// 	}
}
