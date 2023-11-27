package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/SiteGet.json
func ExampleSitesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSitesClient().Get(ctx, "rg1", "testSite", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Site = armhybridnetwork.Site{
	// 	Name: to.Ptr("testSite"),
	// 	Type: to.Ptr("Microsoft.HybridNetwork/sites"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HybridNetwork/sites/testSite"),
	// 	SystemData: &armhybridnetwork.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armhybridnetwork.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armhybridnetwork.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("westUs2"),
	// 	Properties: &armhybridnetwork.SitePropertiesFormat{
	// 		Nfvis: []armhybridnetwork.NFVIsClassification{
	// 			&armhybridnetwork.AzureCoreNFVIDetails{
	// 				Name: to.Ptr("nfvi1"),
	// 				NfviType: to.Ptr(armhybridnetwork.NFVITypeAzureCore),
	// 				Location: to.Ptr("westUs2"),
	// 			},
	// 			&armhybridnetwork.AzureArcK8SClusterNFVIDetails{
	// 				Name: to.Ptr("nfvi2"),
	// 				NfviType: to.Ptr(armhybridnetwork.NFVITypeAzureArcKubernetes),
	// 				CustomLocationReference: &armhybridnetwork.ReferencedResource{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/testResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/testCustomLocation1"),
	// 				},
	// 			},
	// 			&armhybridnetwork.AzureOperatorNexusClusterNFVIDetails{
	// 				Name: to.Ptr("nfvi3"),
	// 				NfviType: to.Ptr(armhybridnetwork.NFVITypeAzureOperatorNexus),
	// 				CustomLocationReference: &armhybridnetwork.ReferencedResource{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/testResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/testCustomLocation2"),
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armhybridnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}
