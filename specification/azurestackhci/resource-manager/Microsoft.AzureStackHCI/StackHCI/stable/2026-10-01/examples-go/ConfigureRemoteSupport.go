package armazurestackhci_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-10-01/ConfigureRemoteSupport.json
func ExampleClustersClient_BeginConfigureRemoteSupport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginConfigureRemoteSupport(ctx, "test-rg", "mycluster", armazurestackhci.RemoteSupportRequest{
		Properties: &armazurestackhci.RemoteSupportRequestProperties{
			AccessLevel:         to.Ptr(armazurestackhci.AccessLevelDiagnostics),
			ExpirationTimeStamp: to.Ptr(time.Date(2020, time.January, 1, 17, 18, 19, 123456700, time.UTC)),
			RemoteSupportType:   to.Ptr(armazurestackhci.RemoteSupportTypeEnable),
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
	// res = armazurestackhci.ClustersClientConfigureRemoteSupportResponse{
	// 	Cluster: armazurestackhci.Cluster{
	// 		Name: to.Ptr("myCluster"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/clusters"),
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/clusters/myCluster"),
	// 		Location: to.Ptr("East US"),
	// 		Properties: &armazurestackhci.ClusterProperties{
	// 			AADClientID: to.Ptr("24a6e53d-04e5-44d2-b7cc-1b732a847dfc"),
	// 			AADTenantID: to.Ptr("7e589cc1-a8b6-4dff-91bd-5ec0fa18db94"),
	// 			BillingModel: to.Ptr("Trial"),
	// 			CloudID: to.Ptr("a3c0468f-e38e-4dda-ac48-817f620536f0"),
	// 			CloudManagementEndpoint: to.Ptr("https://98294836-31be-4668-aeae-698667faf99b.waconazure.com"),
	// 			DesiredProperties: &armazurestackhci.ClusterDesiredProperties{
	// 				DiagnosticLevel: to.Ptr(armazurestackhci.DiagnosticLevelBasic),
	// 				WindowsServerSubscription: to.Ptr(armazurestackhci.WindowsServerSubscriptionEnabled),
	// 			},
	// 			LastBillingTimestamp: to.Ptr(time.Date(2020, time.March, 12, 8, 12, 55, 231202200, time.UTC)),
	// 			LastSyncTimestamp: to.Ptr(time.Date(2020, time.March, 11, 20, 44, 32, 562512100, time.UTC)),
	// 			ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 			RegistrationTimestamp: to.Ptr(time.Date(2020, time.March, 11, 20, 44, 32, 562512100, time.UTC)),
	// 			RemoteSupportProperties: &armazurestackhci.RemoteSupportProperties{
	// 				AccessLevel: to.Ptr(armazurestackhci.AccessLevelDiagnostics),
	// 				ExpirationTimeStamp: to.Ptr(time.Date(2020, time.January, 1, 17, 18, 19, 123456700, time.UTC)),
	// 				RemoteSupportNodeSettings: []*armazurestackhci.RemoteSupportNodeSettings{
	// 					{
	// 						ArcResourceID: to.Ptr("/subscriptions/<subId>/resourcegroups/<rg>/Microsoft.HybridCompute/machines/arcNode0"),
	// 						ConnectionErrorMessage: to.Ptr("test"),
	// 						ConnectionStatus: to.Ptr("Connected"),
	// 						CreatedAt: to.Ptr(time.Date(2020, time.January, 1, 17, 18, 19, 123456700, time.UTC)),
	// 						State: to.Ptr("Active"),
	// 						TranscriptLocation: to.Ptr("test"),
	// 						UpdatedAt: to.Ptr(time.Date(2020, time.January, 1, 17, 18, 19, 123456700, time.UTC)),
	// 					},
	// 					{
	// 						ArcResourceID: to.Ptr("/subscriptions/<subId>/resourcegroups/<rg>/Microsoft.HybridCompute/machines/arcNode1"),
	// 						ConnectionErrorMessage: to.Ptr("test"),
	// 						ConnectionStatus: to.Ptr("Connected"),
	// 						CreatedAt: to.Ptr(time.Date(2020, time.January, 1, 17, 18, 19, 123456700, time.UTC)),
	// 						State: to.Ptr("Active"),
	// 						TranscriptLocation: to.Ptr("test"),
	// 						UpdatedAt: to.Ptr(time.Date(2020, time.January, 1, 17, 18, 19, 123456700, time.UTC)),
	// 					},
	// 				},
	// 				RemoteSupportSessionDetails: []*armazurestackhci.PerNodeRemoteSupportSession{
	// 					{
	// 						AccessLevel: to.Ptr(armazurestackhci.AccessLevelDiagnostics),
	// 						Duration: to.Ptr[int64](120),
	// 						NodeName: to.Ptr("arcNode0"),
	// 						SessionEndTime: to.Ptr(time.Date(2020, time.January, 1, 17, 18, 19, 123456700, time.UTC)),
	// 						SessionStartTime: to.Ptr(time.Date(2020, time.January, 1, 17, 18, 19, 123456700, time.UTC)),
	// 						TranscriptLocation: to.Ptr("/transcript/location"),
	// 					},
	// 				},
	// 				RemoteSupportType: to.Ptr(armazurestackhci.RemoteSupportTypeEnable),
	// 				RemoteSupportProvisioningState: to.Ptr(armazurestackhci.RemoteSupportProvisioningStateGrantInProgress),
	// 			},
	// 			ReportedProperties: &armazurestackhci.ClusterReportedProperties{
	// 				ClusterID: to.Ptr("a76ac23a-1819-4e82-9410-e3e4ec3d1425"),
	// 				ClusterName: to.Ptr("cluster1"),
	// 				ClusterVersion: to.Ptr("10.0.17777"),
	// 				DiagnosticLevel: to.Ptr(armazurestackhci.DiagnosticLevelBasic),
	// 				ImdsAttestation: to.Ptr(armazurestackhci.ImdsAttestationDisabled),
	// 				LastUpdated: to.Ptr(time.Date(2020, time.March, 11, 19, 24, 42, 194601700, time.UTC)),
	// 				Nodes: []*armazurestackhci.ClusterNode{
	// 					{
	// 						Name: to.Ptr("Node1"),
	// 						CoreCount: to.Ptr[float32](8),
	// 						ID: to.Ptr[float32](1),
	// 						Manufacturer: to.Ptr("Dell Inc."),
	// 						MemoryInGiB: to.Ptr[float32](128),
	// 						Model: to.Ptr("EMC AX740"),
	// 						NodeType: to.Ptr(armazurestackhci.ClusterNodeTypeThirdParty),
	// 						OSName: to.Ptr("Azure Stack HCI"),
	// 						OSVersion: to.Ptr("10.0.17777.1061"),
	// 						SerialNumber: to.Ptr("Q45CZC3"),
	// 						WindowsServerSubscription: to.Ptr(armazurestackhci.WindowsServerSubscriptionEnabled),
	// 					},
	// 				},
	// 			},
	// 			Status: to.Ptr(armazurestackhci.StatusConnectedRecently),
	// 			TrialDaysRemaining: to.Ptr[float32](30),
	// 		},
	// 		SystemData: &armazurestackhci.SystemData{
	// 			CreatedAt: to.Ptr(time.Date(2020, time.January, 1, 17, 18, 19, 123456700, time.UTC)),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2020, time.January, 2, 17, 18, 19, 123456700, time.UTC)),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
