package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet/v3"
)

// Generated from example definition: 2026-03-02-preview/ClusterMeshProfiles_ListByFleet.json
func ExampleClusterMeshProfilesClient_NewListByFleetPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClusterMeshProfilesClient().NewListByFleetPager("rgfleets", "fleet1", nil)
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
		// page = armcontainerservicefleet.ClusterMeshProfilesClientListByFleetResponse{
		// 	ClusterMeshProfileListResult: armcontainerservicefleet.ClusterMeshProfileListResult{
		// 		Value: []*armcontainerservicefleet.ClusterMeshProfile{
		// 			{
		// 				Properties: &armcontainerservicefleet.ClusterMeshProfileProperties{
		// 					ProvisioningState: to.Ptr(armcontainerservicefleet.ClusterMeshProfileProvisioningStateSucceeded),
		// 					MemberSelector: &armcontainerservicefleet.MemberSelector{
		// 						ByLabel: to.Ptr("env=production"),
		// 					},
		// 					Status: &armcontainerservicefleet.ClusterMeshProfileStatus{
		// 						State: to.Ptr(armcontainerservicefleet.ClusterMeshStateConnected),
		// 						LastAppliedMemberSelector: &armcontainerservicefleet.MemberSelector{
		// 							ByLabel: to.Ptr("env=production"),
		// 						},
		// 						LastOperationID: to.Ptr("12345678-1234-1234-1234-123456789012"),
		// 					},
		// 				},
		// 				ETag: to.Ptr("\"EtagValue\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/fleets/myFleet/clusterMeshProfiles/clustermeshprofile1"),
		// 				Name: to.Ptr("clustermeshprofile1"),
		// 				Type: to.Ptr("Microsoft.ContainerService/fleets/clusterMeshProfiles"),
		// 				SystemData: &armcontainerservicefleet.SystemData{
		// 					CreatedBy: to.Ptr("someUser"),
		// 					CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("someOtherUser"),
		// 					LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("http://nextlink.contoso.com"),
		// 	},
		// }
	}
}
