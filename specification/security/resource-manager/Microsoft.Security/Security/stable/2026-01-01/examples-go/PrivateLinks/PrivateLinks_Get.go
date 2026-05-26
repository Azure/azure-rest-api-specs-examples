package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2026-01-01/PrivateLinks/PrivateLinks_Get.json
func ExamplePrivateLinksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinksClient().Get(ctx, "rg", armsecurity.PrivateLinkParameters{PrivateLinkName: "spl"}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurity.PrivateLinksClientGetResponse{
	// 	PrivateLinkResource: armsecurity.PrivateLinkResource{
	// 		Name: to.Ptr("spl"),
	// 		Type: to.Ptr("Microsoft.Security/privateLinks"),
	// 		ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg/providers/Microsoft.Security/privateLinks/spl"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armsecurity.PrivateLinkProperties{
	// 			PrivateEndpointConnections: []*armsecurity.PrivateEndpointConnection{
	// 				{
	// 					Name: to.Ptr("pe"),
	// 					Type: to.Ptr("Microsoft.Security/privateLinks/privateEndpointConnections"),
	// 					ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg/providers/Microsoft.Security/privateLinks/spl/privateEndpointConnections/pe"),
	// 					Properties: &armsecurity.PrivateEndpointConnectionProperties{
	// 						GroupIDs: []*string{
	// 							to.Ptr("containers"),
	// 						},
	// 						PrivateEndpoint: &armsecurity.PrivateEndpoint{
	// 							ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg/providers/Microsoft.Network/privateEndpoints/pe"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armsecurity.PrivateLinkServiceConnectionState{
	// 							Description: to.Ptr("Auto-approved by service"),
	// 							ActionsRequired: to.Ptr("None"),
	// 							Status: to.Ptr(armsecurity.PrivateEndpointServiceConnectionStatusApproved),
	// 						},
	// 						ProvisioningState: to.Ptr(armsecurity.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 					},
	// 				},
	// 			},
	// 			PrivateLinkResources: []*armsecurity.PrivateLinkGroupResource{
	// 				{
	// 					Name: to.Ptr("containers"),
	// 					Type: to.Ptr("Microsoft.Security/privateLinks/privateLinkResources"),
	// 					ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg/providers/Microsoft.Security/privateLinks/spl/privateLinkResources/containers"),
	// 					Properties: &armsecurity.PrivateLinkResourceProperties{
	// 						GroupID: to.Ptr("containers"),
	// 						RequiredMembers: []*string{
	// 							to.Ptr("api"),
	// 							to.Ptr("data"),
	// 						},
	// 						RequiredZoneNames: []*string{
	// 							to.Ptr("privatelink.cloud.defender.microsoft.com"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armsecurity.ProvisioningStateSucceeded),
	// 			PublicNetworkAccess: to.Ptr(armsecurity.PublicNetworkAccessDisabled),
	// 		},
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("production"),
	// 			"owner": to.Ptr("security-team"),
	// 			"project": to.Ptr("private-links"),
	// 		},
	// 	},
	// }
}
