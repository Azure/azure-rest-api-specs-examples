package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet/v3"
)

// Generated from example definition: 2026-03-02-preview/FleetMembers_Get.json
func ExampleFleetMembersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFleetMembersClient().Get(ctx, "rgfleets", "fleet1", "member1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerservicefleet.FleetMembersClientGetResponse{
	// 	FleetMember: armcontainerservicefleet.FleetMember{
	// 		Properties: &armcontainerservicefleet.FleetMemberProperties{
	// 			ClusterResourceID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/cluster-1"),
	// 			Group: to.Ptr("fleet1"),
	// 			ProvisioningState: to.Ptr(armcontainerservicefleet.FleetMemberProvisioningStateSucceeded),
	// 			Labels: map[string]*string{
	// 				"env": to.Ptr("production"),
	// 				"fleet.azure.com/cluster-name": to.Ptr("cluster1"),
	// 				"fleet.azure.com/location": to.Ptr("East US"),
	// 				"fleet.azure.com/member-name": to.Ptr("member1"),
	// 				"fleet.azure.com/resource-group": to.Ptr("rgfleets"),
	// 				"fleet.azure.com/subscription-id": to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			},
	// 			Status: &armcontainerservicefleet.FleetMemberStatus{
	// 				LastOperationID: to.Ptr("12345678-1234-1234-1234-123456789012"),
	// 			},
	// 			MeshProperties: &armcontainerservicefleet.MeshProperties{
	// 				CiliumProperties: &armcontainerservicefleet.CiliumProperties{
	// 					ID: to.Ptr[int32](1),
	// 					Name: to.Ptr("cilium-1"),
	// 				},
	// 				Status: &armcontainerservicefleet.MeshMemberStatus{
	// 					State: to.Ptr(armcontainerservicefleet.MeshMemberStateConnected),
	// 					LastOperationID: to.Ptr("abcdef78-1234-1234-1234-123456789012"),
	// 				},
	// 				ClusterMeshProfileResourceID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/fleets/fleet/clusterMeshProfiles/clusterMeshProfile-1"),
	// 			},
	// 		},
	// 		ETag: to.Ptr("kd30rkdfo49="),
	// 		ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/fleets/fleet-1/members/member-1"),
	// 		Name: to.Ptr("member-1"),
	// 		Type: to.Ptr("Microsoft.ContainerService/fleets/members"),
	// 		SystemData: &armcontainerservicefleet.SystemData{
	// 			CreatedBy: to.Ptr("someUser"),
	// 			CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("someOtherUser"),
	// 			LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
	// 		},
	// 	},
	// }
}
