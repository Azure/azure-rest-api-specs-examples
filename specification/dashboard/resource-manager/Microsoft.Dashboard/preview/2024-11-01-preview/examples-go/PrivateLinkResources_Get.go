package armdashboard_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard/v2"
)

// Generated from example definition: 2024-11-01-preview/PrivateLinkResources_Get.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdashboard.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().Get(ctx, "myResourceGroup", "myWorkspace", "grafana", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdashboard.PrivateLinkResourcesClientGetResponse{
	// 	PrivateLinkResource: &armdashboard.PrivateLinkResource{
	// 		Name: to.Ptr("grafana"),
	// 		Type: to.Ptr("Microsoft.Dashboard/grafana/PrivateLinkResources"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/Microsoft.Dashboard/grafana/myWorkspace/privateLinkResources/grafana"),
	// 		Properties: &armdashboard.PrivateLinkResourceProperties{
	// 			GroupID: to.Ptr("grafana"),
	// 			RequiredMembers: []*string{
	// 				to.Ptr("grafana"),
	// 			},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("grafana-test.azure.com"),
	// 			},
	// 		},
	// 	},
	// }
}
