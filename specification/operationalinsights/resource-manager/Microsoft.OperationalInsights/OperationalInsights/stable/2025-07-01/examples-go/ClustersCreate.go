package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v3"
)

// Generated from example definition: 2025-07-01/ClustersCreate.json
func ExampleClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("53bc36c5-91e1-4d09-92c9-63b89e571926", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginCreateOrUpdate(ctx, "oiautorest6685", "oiautorest6685", armoperationalinsights.Cluster{
		Location: to.Ptr("eastus"),
		SKU: &armoperationalinsights.ClusterSKU{
			Name:     to.Ptr(armoperationalinsights.ClusterSKUNameEnumCapacityReservation),
			Capacity: to.Ptr[int64](100),
		},
		Tags: map[string]*string{
			"tag1": to.Ptr("val1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armoperationalinsights.ClustersClientCreateOrUpdateResponse{
	// 	Cluster: armoperationalinsights.Cluster{
	// 		Name: to.Ptr("TestResourceLock"),
	// 		Type: to.Ptr("Microsoft.OperationalInsights/clusters"),
	// 		ID: to.Ptr("/subscriptions/53bc36c5-91e1-4d09-92c9-63b89e571926/resourcegroups/oiautorest6685/providers/microsoft.operationalinsights/clusters/oiautorest6685"),
	// 		Identity: &armoperationalinsights.ManagedServiceIdentity{
	// 			Type: to.Ptr(armoperationalinsights.ManagedServiceIdentityTypeSystemAssigned),
	// 			PrincipalID: to.Ptr("356d057d-cba5-44dd-8a30-b2e547bc416b"),
	// 			TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armoperationalinsights.ClusterProperties{
	// 			BillingType: to.Ptr(armoperationalinsights.BillingTypeCluster),
	// 			ClusterID: to.Ptr("1c3f2afc-9ec1-4e74-bfe8-022b83bcd7b1"),
	// 			CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-13T17:40:33Z"); return t}()),
	// 			IsAvailabilityZonesEnabled: to.Ptr(true),
	// 			IsDoubleEncryptionEnabled: to.Ptr(true),
	// 			LastModifiedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-13T17:40:33Z"); return t}()),
	// 			ProvisioningState: to.Ptr(armoperationalinsights.ClusterEntityStatusSucceeded),
	// 		},
	// 		SKU: &armoperationalinsights.ClusterSKU{
	// 			Name: to.Ptr(armoperationalinsights.ClusterSKUNameEnumCapacityReservation),
	// 			Capacity: to.Ptr[int64](100),
	// 		},
	// 	},
	// }
}
