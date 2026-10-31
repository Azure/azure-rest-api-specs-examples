package armprivatetrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatetrafficmanager/armprivatetrafficmanager"
)

// Generated from example definition: 2026-10-01-preview/Sites_Get_MaximumSet_Gen.json
func ExampleSitesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatetrafficmanager.NewClientFactory("10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSitesClient().Get(ctx, "rgprivateTrafficManager", "myTopologyMap", "mySite", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armprivatetrafficmanager.SitesClientGetResponse{
	// 	Site: armprivatetrafficmanager.Site{
	// 		Properties: &armprivatetrafficmanager.SiteProperties{
	// 			ProbingGatewayIDs: []*string{
	// 				to.Ptr("/subscriptions/10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA/resourceGroups/rgprivateTrafficManager/providers/Microsoft.Network/probingGateways/myProbingGateway"),
	// 			},
	// 			VirtualNetworkIDs: []*string{
	// 				to.Ptr("/subscriptions/10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA/resourceGroups/rgprivateTrafficManager/providers/Microsoft.Network/virtualNetworks/myVirtualNetwork"),
	// 			},
	// 			ProvisioningState: to.Ptr(armprivatetrafficmanager.ProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA/resourceGroups/rgprivateTrafficManager/providers/Microsoft.Network/topologyMaps/myTopologyMap/sites/mySite"),
	// 		Name: to.Ptr("mySite"),
	// 		Type: to.Ptr("Microsoft.Network/topologyMaps/sites"),
	// 		SystemData: &armprivatetrafficmanager.SystemData{
	// 			CreatedBy: to.Ptr("user@contoso.com"),
	// 			CreatedByType: to.Ptr(armprivatetrafficmanager.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2026, time.February, 18, 10, 5, 58, 138000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("user@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armprivatetrafficmanager.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2026, time.February, 18, 10, 5, 58, 139000000, time.UTC)),
	// 		},
	// 	},
	// }
}
