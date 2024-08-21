package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c9b146c38df5f76e2d34a3ef771979805ff4ff73/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/ListClustersByResourceGroup.json
func ExampleClustersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClustersClient().NewListByResourceGroupPager("test-rg", nil)
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
		// page.ClusterList = armazurestackhci.ClusterList{
		// 	Value: []*armazurestackhci.Cluster{
		// 		{
		// 			Name: to.Ptr("myCluster1"),
		// 			Type: to.Ptr("Microsoft.AzureStackHCI/clusters"),
		// 			ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/clusters/myCluster1"),
		// 			SystemData: &armazurestackhci.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armazurestackhci.ClusterProperties{
		// 				AADClientID: to.Ptr("515da1c2-379e-49b4-9975-09e3e40c86be"),
		// 				AADTenantID: to.Ptr("7e589cc1-a8b6-4dff-91bd-5ec0fa18db94"),
		// 				BillingModel: to.Ptr("Trial"),
		// 				CloudID: to.Ptr("91c2b355-4826-4e96-9164-e3f26dcf1cdd"),
		// 				CloudManagementEndpoint: to.Ptr("https://98294836-31be-4668-aeae-698667faf99b.waconazure.com"),
		// 				ConnectivityStatus: to.Ptr(armazurestackhci.ConnectivityStatusNotYetRegistered),
		// 				DesiredProperties: &armazurestackhci.ClusterDesiredProperties{
		// 					DiagnosticLevel: to.Ptr(armazurestackhci.DiagnosticLevelBasic),
		// 					WindowsServerSubscription: to.Ptr(armazurestackhci.WindowsServerSubscriptionEnabled),
		// 				},
		// 				ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
		// 				ReportedProperties: &armazurestackhci.ClusterReportedProperties{
		// 				},
		// 				Status: to.Ptr(armazurestackhci.StatusNotYetRegistered),
		// 				TrialDaysRemaining: to.Ptr[float32](29),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCluster2"),
		// 			Type: to.Ptr("Microsoft.AzureStackHCI/clusters"),
		// 			ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/clusters/myCluster2"),
		// 			SystemData: &armazurestackhci.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armazurestackhci.ClusterProperties{
		// 				AADClientID: to.Ptr("24a6e53d-04e5-44d2-b7cc-1b732a847dfc"),
		// 				AADTenantID: to.Ptr("7e589cc1-a8b6-4dff-91bd-5ec0fa18db94"),
		// 				BillingModel: to.Ptr("Trial"),
		// 				CloudID: to.Ptr("a3c0468f-e38e-4dda-ac48-817f620536f0"),
		// 				CloudManagementEndpoint: to.Ptr("https://98294836-31be-4668-aeae-698667faf99b.waconazure.com"),
		// 				ConnectivityStatus: to.Ptr(armazurestackhci.ConnectivityStatusPartiallyConnected),
		// 				DesiredProperties: &armazurestackhci.ClusterDesiredProperties{
		// 					DiagnosticLevel: to.Ptr(armazurestackhci.DiagnosticLevelBasic),
		// 					WindowsServerSubscription: to.Ptr(armazurestackhci.WindowsServerSubscriptionEnabled),
		// 				},
		// 				LastBillingTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-12T08:12:55.231Z"); return t}()),
		// 				LastSyncTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-11T20:44:32.562Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
		// 				RegistrationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-11T20:44:32.562Z"); return t}()),
		// 				ReportedProperties: &armazurestackhci.ClusterReportedProperties{
		// 					ClusterID: to.Ptr("a76ac23a-1819-4e82-9410-e3e4ec3d1425"),
		// 					ClusterName: to.Ptr("cluster1"),
		// 					ClusterType: to.Ptr(armazurestackhci.ClusterNodeTypeThirdParty),
		// 					ClusterVersion: to.Ptr("10.0.17777"),
		// 					DiagnosticLevel: to.Ptr(armazurestackhci.DiagnosticLevelBasic),
		// 					ImdsAttestation: to.Ptr(armazurestackhci.ImdsAttestationDisabled),
		// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-11T19:24:42.194Z"); return t}()),
		// 					Manufacturer: to.Ptr("Dell Inc."),
		// 					Nodes: []*armazurestackhci.ClusterNode{
		// 						{
		// 							Name: to.Ptr("Node1"),
		// 							CoreCount: to.Ptr[float32](8),
		// 							ID: to.Ptr[float32](0),
		// 							Manufacturer: to.Ptr("Dell Inc."),
		// 							MemoryInGiB: to.Ptr[float32](128),
		// 							Model: to.Ptr("EMC AX740"),
		// 							OemActivation: to.Ptr(armazurestackhci.OemActivationEnabled),
		// 							OSName: to.Ptr("Azure Stack HCI"),
		// 							OSVersion: to.Ptr("10.0.17777.1061"),
		// 							SerialNumber: to.Ptr("Q45CZC3"),
		// 							WindowsServerSubscription: to.Ptr(armazurestackhci.WindowsServerSubscriptionEnabled),
		// 						},
		// 						{
		// 							Name: to.Ptr("Node2"),
		// 							CoreCount: to.Ptr[float32](8),
		// 							ID: to.Ptr[float32](1),
		// 							Manufacturer: to.Ptr("Dell Inc."),
		// 							MemoryInGiB: to.Ptr[float32](128),
		// 							Model: to.Ptr("EMC AX740"),
		// 							OemActivation: to.Ptr(armazurestackhci.OemActivationEnabled),
		// 							OSName: to.Ptr("Azure Stack HCI"),
		// 							OSVersion: to.Ptr("10.0.17777.1061"),
		// 							SerialNumber: to.Ptr("Q44BSC3"),
		// 							WindowsServerSubscription: to.Ptr(armazurestackhci.WindowsServerSubscriptionEnabled),
		// 						},
		// 						{
		// 							Name: to.Ptr("Node3"),
		// 							CoreCount: to.Ptr[float32](16),
		// 							ID: to.Ptr[float32](2),
		// 							Manufacturer: to.Ptr("Dell Inc."),
		// 							MemoryInGiB: to.Ptr[float32](256),
		// 							Model: to.Ptr("EMC AX740"),
		// 							OemActivation: to.Ptr(armazurestackhci.OemActivationDisabled),
		// 							OSName: to.Ptr("Azure Stack HCI"),
		// 							OSVersion: to.Ptr("10.0.17777.1061"),
		// 							SerialNumber: to.Ptr("Q44RFC3"),
		// 							WindowsServerSubscription: to.Ptr(armazurestackhci.WindowsServerSubscriptionEnabled),
		// 					}},
		// 					OemActivation: to.Ptr(armazurestackhci.OemActivationDisabled),
		// 				},
		// 				Status: to.Ptr(armazurestackhci.StatusConnectedRecently),
		// 				TrialDaysRemaining: to.Ptr[float32](30),
		// 			},
		// 	}},
		// }
	}
}
