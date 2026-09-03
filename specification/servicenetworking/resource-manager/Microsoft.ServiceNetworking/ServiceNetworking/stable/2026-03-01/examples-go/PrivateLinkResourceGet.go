package armservicenetworking_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicenetworking/armservicenetworking"
)

// Generated from example definition: 2026-03-01/PrivateLinkResourceGet.json
func ExamplePrivateLinkResourcesInterfaceClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicenetworking.NewClientFactory("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesInterfaceClient().Get(ctx, "rg1", "tc1", "fe1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicenetworking.PrivateLinkResourcesInterfaceClientGetResponse{
	// 	PrivateLinkResource: armservicenetworking.PrivateLinkResource{
	// 		Name: to.Ptr("fe1"),
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ServiceNetworking/trafficControllers/tc1/privateLinkResources/fe1"),
	// 		Type: to.Ptr("Microsoft.ServiceNetworking/trafficControllers/privateLinkResources"),
	// 		Properties: &armservicenetworking.PrivateLinkResourceProperties{
	// 			GroupID: to.Ptr("fe1"),
	// 			RequiredMembers: []*string{
	// 				to.Ptr("fe1"),
	// 			},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("privatelink.alb.azure.com"),
	// 			},
	// 		},
	// 	},
	// }
}
