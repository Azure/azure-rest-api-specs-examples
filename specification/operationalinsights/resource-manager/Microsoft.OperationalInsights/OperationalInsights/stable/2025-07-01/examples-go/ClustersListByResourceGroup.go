package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f0a5127d4e8b1ea6007b0bf9570904d9df860b97/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/ClustersListByResourceGroup.json
func ExampleClustersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClustersClient().NewListByResourceGroupPager("oiautorest6685", nil)
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
		// page.ClusterListResult = armoperationalinsights.ClusterListResult{
		// 	Value: []*armoperationalinsights.Cluster{
		// 		{
		// 			Name: to.Ptr("mycluster"),
		// 			Type: to.Ptr("Microsoft.OperationalInsights/clusters"),
		// 			ID: to.Ptr("/subscriptions/53bc36c5-91e1-4d09-92c9-63b89e571926/resourcegroups/oiautorest6685/providers/microsoft.operationalinsights/clusters/mycluster"),
		// 			Location: to.Ptr("eastus"),
		// 			Identity: &armoperationalinsights.ManagedServiceIdentity{
		// 				Type: to.Ptr(armoperationalinsights.ManagedServiceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("356d057d-cba5-44dd-8a30-b2e547bc416b"),
		// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 			Properties: &armoperationalinsights.ClusterProperties{
		// 				ClusterID: to.Ptr("5b02755b-5bf4-430c-9487-45502a2a7e62"),
		// 				KeyVaultProperties: &armoperationalinsights.KeyVaultProperties{
		// 					KeyName: to.Ptr("aztest2170cert"),
		// 					KeyVaultURI: to.Ptr("https://aztest2170.vault.azure.net"),
		// 					KeyVersion: to.Ptr("654ft6c4e63845cbb50fd6fg51540429"),
		// 				},
		// 				ProvisioningState: to.Ptr(armoperationalinsights.ClusterEntityStatusSucceeded),
		// 			},
		// 			SKU: &armoperationalinsights.ClusterSKU{
		// 				Name: to.Ptr(armoperationalinsights.ClusterSKUNameEnumCapacityReservation),
		// 				Capacity: to.Ptr[int64](1000),
		// 			},
		// 	}},
		// }
	}
}
