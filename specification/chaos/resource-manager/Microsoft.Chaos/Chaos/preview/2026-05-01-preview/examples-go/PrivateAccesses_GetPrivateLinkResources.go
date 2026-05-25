package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-05-01-preview/PrivateAccesses_GetPrivateLinkResources.json
func ExamplePrivateAccessesClient_GetPrivateLinkResources() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateAccessesClient().GetPrivateLinkResources(ctx, "myResourceGroup", "myPrivateAccess", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armchaos.PrivateAccessesClientGetPrivateLinkResourcesResponse{
	// 	PrivateLinkResourceListResult: armchaos.PrivateLinkResourceListResult{
	// 		Value: []*armchaos.PrivateLinkResource{
	// 			{
	// 				Name: to.Ptr("agents"),
	// 				Type: to.Ptr("Microsoft.Chaos/privateAccesses/privateLinkResources"),
	// 				ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/myResourceGroup/providers/Microsoft.Chaos/privateAccesses/myPrivateAccess/privateLinkResources/agents"),
	// 				Location: to.Ptr("centraluseuap"),
	// 				Properties: &armchaos.PrivateLinkResourceProperties{
	// 					GroupID: to.Ptr("agents"),
	// 					RequiredMembers: []*string{
	// 						to.Ptr("privateAccess_1"),
	// 					},
	// 					RequiredZoneNames: []*string{
	// 						to.Ptr("privatelink.agents.core.windows.net"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
