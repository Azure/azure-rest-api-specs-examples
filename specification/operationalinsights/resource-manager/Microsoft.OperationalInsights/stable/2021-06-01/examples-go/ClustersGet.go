package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4ce13e8353a25125a41bc01705c0a7794dac32a7/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/examples/ClustersGet.json
func ExampleClustersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClustersClient().Get(ctx, "oiautorest6685", "oiautorest6685", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Cluster = armoperationalinsights.Cluster{
	// 	Name: to.Ptr("TestResourceLock"),
	// 	Type: to.Ptr("Microsoft.OperationalInsights/clusters"),
	// 	ID: to.Ptr("/subscriptions/594038b5-1093-476e-a366-482775671c11/resourcegroups/calbot-rg/providers/microsoft.operationalinsights/clusters/testresourcelock"),
	// 	Location: to.Ptr("eastus"),
	// 	Identity: &armoperationalinsights.Identity{
	// 		Type: to.Ptr(armoperationalinsights.IdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("356d057d-cba5-44dd-8a30-b2e547bc416b"),
	// 		TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 	},
	// 	Properties: &armoperationalinsights.ClusterProperties{
	// 		AssociatedWorkspaces: []*armoperationalinsights.AssociatedWorkspace{
	// 			{
	// 				AssociateDate: to.Ptr("Tue, 07 Jul 2020 07:35:51 GMT"),
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourcegroups/testRG/providers/microsoft.operationalinsights/workspaces/testWs1"),
	// 				WorkspaceID: to.Ptr("942bdefd-e6c9-411c-ac69-70ffad564363"),
	// 				WorkspaceName: to.Ptr("testWs1"),
	// 			},
	// 			{
	// 				AssociateDate: to.Ptr("Mon, 13 Jan 2020 16:03:39 GMT"),
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourcegroups/testRG/providers/microsoft.operationalinsights/workspaces/testWs2"),
	// 				WorkspaceID: to.Ptr("c7edb8f8-67f7-41f2-bddb-aecf22507e3f"),
	// 				WorkspaceName: to.Ptr("testWs2"),
	// 		}},
	// 		BillingType: to.Ptr(armoperationalinsights.BillingTypeCluster),
	// 		CapacityReservationProperties: &armoperationalinsights.CapacityReservationProperties{
	// 			LastSKUUpdate: to.Ptr("Thu, 01 Jan 1970 00:00:00 GMT"),
	// 			MinCapacity: to.Ptr[int64](500),
	// 		},
	// 		ClusterID: to.Ptr("1c3f2afc-9ec1-4e74-bfe8-022b83bcd7b1"),
	// 		CreatedDate: to.Ptr("Mon, 13 Jan 2020 14:40:33 GMT"),
	// 		IsAvailabilityZonesEnabled: to.Ptr(false),
	// 		IsDoubleEncryptionEnabled: to.Ptr(false),
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
