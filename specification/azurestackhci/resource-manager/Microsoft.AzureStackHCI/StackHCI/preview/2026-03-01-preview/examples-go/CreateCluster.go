package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-03-01-preview/CreateCluster.json
func ExampleClustersClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClustersClient().Create(ctx, "test-rg", "myCluster", armazurestackhci.Cluster{
		Identity: &armazurestackhci.ManagedServiceIdentity{
			Type: to.Ptr(armazurestackhci.ManagedServiceIdentityTypeSystemAssigned),
		},
		Kind:     to.Ptr("AzureLocal"),
		Location: to.Ptr("East US"),
		Properties: &armazurestackhci.ClusterProperties{
			AADClientID:             to.Ptr("24a6e53d-04e5-44d2-b7cc-1b732a847dfc"),
			AADTenantID:             to.Ptr("7e589cc1-a8b6-4dff-91bd-5ec0fa18db94"),
			CloudManagementEndpoint: to.Ptr("https://98294836-31be-4668-aeae-698667faf99b.waconazure.com"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armazurestackhci.ClustersClientCreateResponse{
	// 	Cluster: &armazurestackhci.Cluster{
	// 		Name: to.Ptr("myCluster"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/clusters"),
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/clusters/myCluster"),
	// 		Identity: &armazurestackhci.ManagedServiceIdentity{
	// 			Type: to.Ptr(armazurestackhci.ManagedServiceIdentityTypeSystemAssigned),
	// 			PrincipalID: to.Ptr("87a834db-2e45-409e-911b-e16a44562ec3"),
	// 			TenantID: to.Ptr("7e589cc1-a8b6-4dff-91bd-5ec0fa18db94"),
	// 		},
	// 		Kind: to.Ptr("AzureLocal"),
	// 		Location: to.Ptr("East US"),
	// 		Properties: &armazurestackhci.ClusterProperties{
	// 			AADClientID: to.Ptr("24a6e53d-04e5-44d2-b7cc-1b732a847dfc"),
	// 			AADTenantID: to.Ptr("7e589cc1-a8b6-4dff-91bd-5ec0fa18db94"),
	// 			BillingModel: to.Ptr("Trial"),
	// 			CloudID: to.Ptr("a3c0468f-e38e-4dda-ac48-817f620536f0"),
	// 			CloudManagementEndpoint: to.Ptr("https://98294836-31be-4668-aeae-698667faf99b.waconazure.com"),
	// 			ConnectivityStatus: to.Ptr(armazurestackhci.ConnectivityStatusNotYetRegistered),
	// 			DesiredProperties: &armazurestackhci.ClusterDesiredProperties{
	// 				DiagnosticLevel: to.Ptr(armazurestackhci.DiagnosticLevelBasic),
	// 				WindowsServerSubscription: to.Ptr(armazurestackhci.WindowsServerSubscriptionDisabled),
	// 			},
	// 			IsManagementCluster: to.Ptr(true),
	// 			ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 			ReportedProperties: &armazurestackhci.ClusterReportedProperties{
	// 			},
	// 			ServiceEndpoint: to.Ptr("https://azurestackhci.azurefd.net/eastus"),
	// 			Status: to.Ptr(armazurestackhci.StatusNotYetRegistered),
	// 			TrialDaysRemaining: to.Ptr[float32](30),
	// 		},
	// 		SystemData: &armazurestackhci.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
