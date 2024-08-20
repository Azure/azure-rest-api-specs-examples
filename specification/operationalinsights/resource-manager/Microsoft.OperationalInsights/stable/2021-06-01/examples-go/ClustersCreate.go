package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4ce13e8353a25125a41bc01705c0a7794dac32a7/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/examples/ClustersCreate.json
func ExampleClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginCreateOrUpdate(ctx, "oiautorest6685", "oiautorest6685", armoperationalinsights.Cluster{
		Location: to.Ptr("australiasoutheast"),
		Tags: map[string]*string{
			"tag1": to.Ptr("val1"),
		},
		SKU: &armoperationalinsights.ClusterSKU{
			Name:     to.Ptr(armoperationalinsights.ClusterSKUNameEnumCapacityReservation),
			Capacity: to.Ptr(armoperationalinsights.CapacityTenHundred),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Cluster = armoperationalinsights.Cluster{
	// 	Name: to.Ptr("oiautorest6685"),
	// 	Type: to.Ptr("Microsoft.OperationalInsights/clusters"),
	// 	ID: to.Ptr("/subscriptions/594038b5-1093-476e-a366-482775671c11/resourcegroups/oiautorest6685/providers/microsoft.operationalinsights/clusters/oiautorest6685"),
	// 	Location: to.Ptr("australiasoutheast"),
	// 	Identity: &armoperationalinsights.Identity{
	// 		Type: to.Ptr(armoperationalinsights.IdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("356d057d-cba5-44dd-8a30-b2e547bc416b"),
	// 		TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 	},
	// 	Properties: &armoperationalinsights.ClusterProperties{
	// 		AssociatedWorkspaces: []*armoperationalinsights.AssociatedWorkspace{
	// 		},
	// 		BillingType: to.Ptr(armoperationalinsights.BillingTypeCluster),
	// 		CapacityReservationProperties: &armoperationalinsights.CapacityReservationProperties{
	// 			LastSKUUpdate: to.Ptr("Thu, 01 Jan 1970 00:00:00 GMT"),
	// 			MinCapacity: to.Ptr[int64](500),
	// 		},
	// 		ClusterID: to.Ptr("1c3f2afc-9ec1-4e74-bfe8-022b83bcd7b1"),
	// 		CreatedDate: to.Ptr("Mon, 13 Jan 2020 14:40:33 GMT"),
	// 		KeyVaultProperties: &armoperationalinsights.KeyVaultProperties{
	// 			KeyName: to.Ptr("aztest2170cert"),
	// 			KeyVaultURI: to.Ptr("https://aztest2170.vault.azure.net"),
	// 			KeyVersion: to.Ptr("654ft6c4e63845cbb50fd6fg51540429"),
	// 		},
	// 		LastModifiedDate: to.Ptr("Sun, 04 Jan 2020 17:10:56 GMT"),
	// 		ProvisioningState: to.Ptr(armoperationalinsights.ClusterEntityStatusSucceeded),
	// 	},
	// 	SKU: &armoperationalinsights.ClusterSKU{
	// 		Name: to.Ptr(armoperationalinsights.ClusterSKUNameEnumCapacityReservation),
	// 		Capacity: to.Ptr(armoperationalinsights.CapacityTenHundred),
	// 	},
	// }
}
