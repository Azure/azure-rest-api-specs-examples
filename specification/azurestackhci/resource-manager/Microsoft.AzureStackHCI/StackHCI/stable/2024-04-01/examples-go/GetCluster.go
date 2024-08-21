package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c9b146c38df5f76e2d34a3ef771979805ff4ff73/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/GetCluster.json
func ExampleClustersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClustersClient().Get(ctx, "test-rg", "myCluster", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Cluster = armazurestackhci.Cluster{
	// 	Name: to.Ptr("myCluster"),
	// 	Type: to.Ptr("Microsoft.AzureStackHCI/clusters"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/clusters/myCluster"),
	// 	SystemData: &armazurestackhci.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armazurestackhci.ClusterProperties{
	// 		AADClientID: to.Ptr("24a6e53d-04e5-44d2-b7cc-1b732a847dfc"),
	// 		AADTenantID: to.Ptr("7e589cc1-a8b6-4dff-91bd-5ec0fa18db94"),
	// 		BillingModel: to.Ptr("Trial"),
	// 		CloudID: to.Ptr("a3c0468f-e38e-4dda-ac48-817f620536f0"),
	// 		CloudManagementEndpoint: to.Ptr("https://98294836-31be-4668-aeae-698667faf99b.waconazure.com"),
	// 		ConnectivityStatus: to.Ptr(armazurestackhci.ConnectivityStatusConnected),
	// 		DesiredProperties: &armazurestackhci.ClusterDesiredProperties{
	// 			DiagnosticLevel: to.Ptr(armazurestackhci.DiagnosticLevelBasic),
	// 			WindowsServerSubscription: to.Ptr(armazurestackhci.WindowsServerSubscriptionEnabled),
	// 		},
	// 		IsolatedVMAttestationConfiguration: &armazurestackhci.IsolatedVMAttestationConfiguration{
	// 			AttestationResourceID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.Attestation/attestationProviders/testmaa"),
	// 			AttestationServiceEndpoint: to.Ptr("https://dantestnoauth01.eus.attest.azure.net"),
	// 			RelyingPartyServiceEndpoint: to.Ptr("https://azurestackhci.azurefd.net/eastus"),
	// 		},
	// 		LastBillingTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-12T08:12:55.231Z"); return t}()),
	// 		LastSyncTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-11T20:44:32.562Z"); return t}()),
	// 		LogCollectionProperties: &armazurestackhci.LogCollectionProperties{
	// 			FromDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 			LastLogGenerated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 			LogCollectionSessionDetails: []*armazurestackhci.LogCollectionSession{
	// 				{
	// 					CorrelationID: to.Ptr("a76ab33a-1819-4e82-1212-e3e4ec3d1425"),
	// 					EndTimeCollected: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:25:19.123Z"); return t}()),
	// 					LogCollectionJobType: to.Ptr(armazurestackhci.LogCollectionJobTypeOnDemand),
	// 					LogCollectionStatus: to.Ptr(armazurestackhci.LogCollectionStatusSucceeded),
	// 					LogEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 					LogSize: to.Ptr[int64](1000),
	// 					LogStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 					TimeCollected: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 			}},
	// 			ToDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T17:18:19.123Z"); return t}()),
	// 		},
	// 		ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 		RegistrationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-11T20:44:32.562Z"); return t}()),
	// 		ReportedProperties: &armazurestackhci.ClusterReportedProperties{
	// 			ClusterID: to.Ptr("a76ac23a-1819-4e82-9410-e3e4ec3d1425"),
	// 			ClusterName: to.Ptr("cluster1"),
	// 			ClusterType: to.Ptr(armazurestackhci.ClusterNodeTypeThirdParty),
	// 			ClusterVersion: to.Ptr("10.0.17777"),
	// 			DiagnosticLevel: to.Ptr(armazurestackhci.DiagnosticLevelBasic),
	// 			ImdsAttestation: to.Ptr(armazurestackhci.ImdsAttestationDisabled),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-11T19:24:42.194Z"); return t}()),
	// 			Manufacturer: to.Ptr("Dell Inc."),
	// 			Nodes: []*armazurestackhci.ClusterNode{
	// 				{
	// 					Name: to.Ptr("Node1"),
	// 					CoreCount: to.Ptr[float32](8),
	// 					ID: to.Ptr[float32](1),
	// 					LastLicensingTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-11T19:24:42.194Z"); return t}()),
	// 					Manufacturer: to.Ptr("Dell Inc."),
	// 					MemoryInGiB: to.Ptr[float32](128),
	// 					Model: to.Ptr("EMC AX740"),
	// 					OemActivation: to.Ptr(armazurestackhci.OemActivationDisabled),
	// 					OSName: to.Ptr("Azure Stack HCI"),
	// 					OSVersion: to.Ptr("10.0.17777.1061"),
	// 					SerialNumber: to.Ptr("Q45CZC3"),
	// 					WindowsServerSubscription: to.Ptr(armazurestackhci.WindowsServerSubscriptionEnabled),
	// 				},
	// 				{
	// 					Name: to.Ptr("Node2"),
	// 					CoreCount: to.Ptr[float32](8),
	// 					ID: to.Ptr[float32](2),
	// 					LastLicensingTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-11T19:24:42.194Z"); return t}()),
	// 					Manufacturer: to.Ptr("Dell Inc."),
	// 					MemoryInGiB: to.Ptr[float32](128),
	// 					Model: to.Ptr("EMC AX740"),
	// 					OemActivation: to.Ptr(armazurestackhci.OemActivationDisabled),
	// 					OSName: to.Ptr("Azure Stack HCI"),
	// 					OSVersion: to.Ptr("10.0.17777.1061"),
	// 					SerialNumber: to.Ptr("Q44BSC3"),
	// 					WindowsServerSubscription: to.Ptr(armazurestackhci.WindowsServerSubscriptionEnabled),
	// 				},
	// 				{
	// 					Name: to.Ptr("Node3"),
	// 					CoreCount: to.Ptr[float32](16),
	// 					ID: to.Ptr[float32](3),
	// 					LastLicensingTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-11T19:24:42.194Z"); return t}()),
	// 					Manufacturer: to.Ptr("Dell Inc."),
	// 					MemoryInGiB: to.Ptr[float32](256),
	// 					Model: to.Ptr("EMC AX740"),
	// 					OemActivation: to.Ptr(armazurestackhci.OemActivationDisabled),
	// 					OSName: to.Ptr("Azure Stack HCI"),
	// 					OSVersion: to.Ptr("10.0.17777.1061"),
	// 					SerialNumber: to.Ptr("Q44RFC3"),
	// 					WindowsServerSubscription: to.Ptr(armazurestackhci.WindowsServerSubscriptionEnabled),
	// 			}},
	// 			OemActivation: to.Ptr(armazurestackhci.OemActivationDisabled),
	// 		},
	// 		Status: to.Ptr(armazurestackhci.StatusConnectedRecently),
	// 		TrialDaysRemaining: to.Ptr[float32](30),
	// 	},
	// }
}
