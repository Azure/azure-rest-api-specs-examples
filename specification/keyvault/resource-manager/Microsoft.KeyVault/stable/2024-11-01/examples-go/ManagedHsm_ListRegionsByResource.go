package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ee1eec42dcc710ff88db2d1bf574b2f9afe3d654/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2024-11-01/examples/ManagedHsm_ListRegionsByResource.json
func ExampleMHSMRegionsClient_NewListByResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMHSMRegionsClient().NewListByResourcePager("sample-group", "sample-mhsm", nil)
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
		// page.MHSMRegionsListResult = armkeyvault.MHSMRegionsListResult{
		// 	Value: []*armkeyvault.MHSMGeoReplicatedRegion{
		// 		{
		// 			Name: to.Ptr("sample-region1"),
		// 			IsPrimary: to.Ptr(true),
		// 			ProvisioningState: to.Ptr(armkeyvault.GeoReplicationRegionProvisioningStateSucceeded),
		// 		},
		// 		{
		// 			Name: to.Ptr("sample-region2"),
		// 			IsPrimary: to.Ptr(false),
		// 			ProvisioningState: to.Ptr(armkeyvault.GeoReplicationRegionProvisioningStateSucceeded),
		// 	}},
		// }
	}
}
