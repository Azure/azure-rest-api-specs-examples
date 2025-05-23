package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet/v2"
)

// Generated from example definition: 2025-03-01/FleetMembers_ListByFleet_MaximumSet_Gen.json
func ExampleFleetMembersClient_NewListByFleetPager_listsTheMembersOfAFleetGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFleetMembersClient().NewListByFleetPager("rgfleets", "fleet1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armcontainerservicefleet.FleetMembersClientListByFleetResponse{
		// 	FleetMemberListResult: armcontainerservicefleet.FleetMemberListResult{
		// 		Value: []*armcontainerservicefleet.FleetMember{
		// 			{
		// 				ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/fleets/fleet-1/members/member-1"),
		// 				Name: to.Ptr("member-1"),
		// 				Type: to.Ptr("Microsoft.ContainerService/fleets/members"),
		// 				SystemData: &armcontainerservicefleet.SystemData{
		// 					CreatedBy: to.Ptr("someUser"),
		// 					CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("someOtherUser"),
		// 					LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
		// 				},
		// 				ETag: to.Ptr("kd30rkdfo49="),
		// 				Properties: &armcontainerservicefleet.FleetMemberProperties{
		// 					ClusterResourceID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/cluster-1"),
		// 					ProvisioningState: to.Ptr(armcontainerservicefleet.FleetMemberProvisioningStateSucceeded),
		// 					Group: to.Ptr("fleet1"),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
