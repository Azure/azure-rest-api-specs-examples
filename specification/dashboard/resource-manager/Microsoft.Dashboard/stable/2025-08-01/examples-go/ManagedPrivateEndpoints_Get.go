package armdashboard_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard/v2"
)

// Generated from example definition: 2025-08-01/ManagedPrivateEndpoints_Get.json
func ExampleManagedPrivateEndpointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdashboard.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedPrivateEndpointsClient().Get(ctx, "myResourceGroup", "myWorkspace", "myMPEName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdashboard.ManagedPrivateEndpointsClientGetResponse{
	// 	ManagedPrivateEndpointModel: &armdashboard.ManagedPrivateEndpointModel{
	// 		Name: to.Ptr("myMPEName"),
	// 		Type: to.Ptr("Microsoft.Dashboard/grafana/managedPrivateEndpoint"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/Microsoft.Dashboard/grafana/myWorkspace/managedPrivateEndpoints/myPrivateEndpointName"),
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armdashboard.ManagedPrivateEndpointModelProperties{
	// 			ConnectionState: &armdashboard.ManagedPrivateEndpointConnectionState{
	// 				Description: to.Ptr("Auto-Approved"),
	// 				Status: to.Ptr(armdashboard.ManagedPrivateEndpointConnectionStatusApproved),
	// 			},
	// 			GroupIDs: []*string{
	// 				to.Ptr("grafana"),
	// 			},
	// 			PrivateLinkResourceID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-000000000000/resourceGroups/xx-rg/providers/Microsoft.Kusto/Clusters/sampleKustoResource"),
	// 			PrivateLinkResourceRegion: to.Ptr("West US"),
	// 			ProvisioningState: to.Ptr(armdashboard.ProvisioningStateSucceeded),
	// 			RequestMessage: to.Ptr("Example Request Message"),
	// 		},
	// 		SystemData: &armdashboard.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.1075056Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armdashboard.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.1974346Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armdashboard.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
