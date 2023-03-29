package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/df863270270ad5b54fa8cce71d2c33becee0c097/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-09-02-preview/examples/FleetMembers_Get.json
func ExampleFleetMembersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFleetMembersClient().Get(ctx, "rg1", "fleet-1", "member-1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FleetMember = armcontainerservice.FleetMember{
	// 	Name: to.Ptr("member-1"),
	// 	Type: to.Ptr("Microsoft.ContainerService/fleets/members"),
	// 	ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/fleets/fleet-1/members/member-1"),
	// 	SystemData: &armcontainerservice.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
	// 		CreatedBy: to.Ptr("someUser"),
	// 		CreatedByType: to.Ptr(armcontainerservice.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("someOtherUser"),
	// 		LastModifiedByType: to.Ptr(armcontainerservice.CreatedByTypeUser),
	// 	},
	// 	Etag: to.Ptr("kd30rkdfo49="),
	// 	Properties: &armcontainerservice.FleetMemberProperties{
	// 		ClusterResourceID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/cluster-1"),
	// 		ProvisioningState: to.Ptr(armcontainerservice.FleetMemberProvisioningStateSucceeded),
	// 	},
	// }
}
