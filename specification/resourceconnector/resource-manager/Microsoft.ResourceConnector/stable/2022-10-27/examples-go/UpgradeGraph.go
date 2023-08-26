package armresourceconnector_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourceconnector/armresourceconnector"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5f700acd3d094d8eedca381932f2e7916afd2e55/specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/UpgradeGraph.json
func ExampleAppliancesClient_GetUpgradeGraph() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourceconnector.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAppliancesClient().GetUpgradeGraph(ctx, "testresourcegroup", "appliance01", "stable", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UpgradeGraph = armresourceconnector.UpgradeGraph{
	// 	Name: to.Ptr("stable"),
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ResourceConnector/appliances/appliance01"),
	// 	Properties: &armresourceconnector.UpgradeGraphProperties{
	// 		ApplianceVersion: to.Ptr("1.0.0"),
	// 		SupportedVersions: []*armresourceconnector.SupportedVersion{
	// 			{
	// 				Metadata: &armresourceconnector.SupportedVersionMetadata{
	// 					CatalogVersion: &armresourceconnector.SupportedVersionCatalogVersion{
	// 						Name: to.Ptr("cloudop-product-information"),
	// 						Data: &armresourceconnector.SupportedVersionCatalogVersionData{
	// 							Audience: to.Ptr("stable"),
	// 							Catalog: to.Ptr("arc-appliance-stable-catalogs-ext"),
	// 							Offer: to.Ptr("arcappliance"),
	// 							Version: to.Ptr("0.1.5.11115"),
	// 						},
	// 						Namespace: to.Ptr("cloudop-system"),
	// 					},
	// 				},
	// 				Version: to.Ptr("1.0.1"),
	// 		}},
	// 	},
	// }
}
