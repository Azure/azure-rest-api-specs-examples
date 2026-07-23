package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v3"
)

// Generated from example definition: 2026-06-16-preview/gateway/Gateway_ListByResourceGroup.json
func ExampleGatewaysClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("ffd506c8-3415-42d3-9612-fdb423fb17df", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGatewaysClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page = armhybridcompute.GatewaysClientListByResourceGroupResponse{
		// 	GatewaysListResult: armhybridcompute.GatewaysListResult{
		// 		Value: []*armhybridcompute.Gateway{
		// 			{
		// 				Name: to.Ptr("{gatewayName}"),
		// 				Type: to.Ptr("Microsoft.HybridCompute/gateways"),
		// 				ID: to.Ptr("/subscriptions/ffd506c8-3415-42d3-9612-fdb423fb17df/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/gateways/{gatewayName}"),
		// 				Location: to.Ptr("eastus2euap"),
		// 				Properties: &armhybridcompute.GatewayProperties{
		// 					AllowedFeatures: []*string{
		// 						to.Ptr("*"),
		// 					},
		// 					GatewayEndpoint: to.Ptr("https://uniqueValue.contoso.com"),
		// 					GatewayID: to.Ptr("<generated Guid>"),
		// 					GatewayType: to.Ptr(armhybridcompute.GatewayTypePublic),
		// 					GatewayBypass: []*string{
		// 						to.Ptr("contoso.com"),
		// 						to.Ptr("internal.corp.net"),
		// 					},
		// 					ProvisioningState: to.Ptr(armhybridcompute.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
