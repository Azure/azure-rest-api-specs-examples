package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f0a5127d4e8b1ea6007b0bf9570904d9df860b97/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/ClustersGet.json
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
	// 	ID: to.Ptr("/subscriptions/53bc36c5-91e1-4d09-92c9-63b89e571926/resourcegroups/oiautorest6685/providers/microsoft.operationalinsights/clusters/oiautorest6685"),
	// 	Location: to.Ptr("eastus"),
	// 	Identity: &armoperationalinsights.ManagedServiceIdentity{
	// 		Type: to.Ptr(armoperationalinsights.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("356d057d-cba5-44dd-8a30-b2e547bc416b"),
	// 		TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 	},
	// 	Properties: &armoperationalinsights.ClusterProperties{
	// 		AssociatedWorkspaces: []*armoperationalinsights.AssociatedWorkspace{
	// 			{
	// 				AssociateDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-13T23:40:33.000Z"); return t}()),
	// 				ResourceID: to.Ptr("/subscriptions/53bc36c5-91e1-4d09-92c9-63b89e571926/resourcegroups/testRG/providers/microsoft.operationalinsights/workspaces/testWs1"),
	// 				WorkspaceID: to.Ptr("942bdefd-e6c9-411c-ac69-70ffad564363"),
	// 				WorkspaceName: to.Ptr("testWs1"),
	// 			},
	// 			{
	// 				AssociateDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-13T23:40:33.000Z"); return t}()),
	// 				ResourceID: to.Ptr("/subscriptions/53bc36c5-91e1-4d09-92c9-63b89e571926/resourcegroups/testRG/providers/microsoft.operationalinsights/workspaces/testWs2"),
	// 				WorkspaceID: to.Ptr("c7edb8f8-67f7-41f2-bddb-aecf22507e3f"),
	// 				WorkspaceName: to.Ptr("testWs2"),
	// 		}},
	// 		BillingType: to.Ptr(armoperationalinsights.BillingTypeCluster),
	// 		CapacityReservationProperties: &armoperationalinsights.CapacityReservationProperties{
	// 			LastSKUUpdate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-13T17:40:33.000Z"); return t}()),
	// 			MinCapacity: to.Ptr[int64](100),
	// 		},
	// 		ClusterID: to.Ptr("1c3f2afc-9ec1-4e74-bfe8-022b83bcd7b1"),
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-13T17:40:33.000Z"); return t}()),
	// 		IsAvailabilityZonesEnabled: to.Ptr(true),
	// 		IsDoubleEncryptionEnabled: to.Ptr(true),
	// 		KeyVaultProperties: &armoperationalinsights.KeyVaultProperties{
	// 			KeyName: to.Ptr("aztest2170cert"),
	// 			KeyVaultURI: to.Ptr("https://aztest2170.vault.azure.net"),
	// 			KeyVersion: to.Ptr("654ft6c4e63845cbb50fd6fg51540429"),
	// 		},
	// 		LastModifiedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-13T17:40:33.000Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armoperationalinsights.ClusterEntityStatusSucceeded),
	// 		Replication: &armoperationalinsights.ClusterReplicationProperties{
	// 			CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-13T17:40:33.000Z"); return t}()),
	// 			Enabled: to.Ptr(true),
	// 			LastModifiedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-13T17:40:33.000Z"); return t}()),
	// 			Location: to.Ptr("westus3"),
	// 			ProvisioningState: to.Ptr(armoperationalinsights.ClusterReplicationStateSucceeded),
	// 		},
	// 	},
	// 	SKU: &armoperationalinsights.ClusterSKU{
	// 		Name: to.Ptr(armoperationalinsights.ClusterSKUNameEnumCapacityReservation),
	// 		Capacity: to.Ptr[int64](1000),
	// 	},
	// }
}
