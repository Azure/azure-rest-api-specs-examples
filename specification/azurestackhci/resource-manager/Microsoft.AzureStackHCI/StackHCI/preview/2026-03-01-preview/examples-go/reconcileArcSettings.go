package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-03-01-preview/reconcileArcSettings.json
func ExampleArcSettingsClient_BeginReconcile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewArcSettingsClient().BeginReconcile(ctx, "test-rg", "myCluster", "default", armazurestackhci.ReconcileArcSettingsRequest{
		Properties: &armazurestackhci.ReconcileArcSettingsRequestProperties{
			ClusterNodes: []*string{
				to.Ptr("/subscriptions/sub1/resourceGroup/res1/providers/Microsoft.HybridCompute/machines/m1"),
				to.Ptr("/subscriptions/sub1/resourceGroup/res1/providers/Microsoft.HybridCompute/machines/m2"),
			},
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
	// res = armazurestackhci.ArcSettingsClientReconcileResponse{
	// 	ArcSetting: &armazurestackhci.ArcSetting{
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/clusters/arcSettings"),
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/clusters/myCluster/arcSettings/default"),
	// 		Properties: &armazurestackhci.ArcSettingProperties{
	// 			AggregateState: to.Ptr(armazurestackhci.ArcSettingAggregateStateCreating),
	// 			ArcInstanceResourceGroup: to.Ptr("ArcInstance-rg"),
	// 			ConnectivityProperties: &armazurestackhci.ArcConnectivityProperties{
	// 				Enabled: to.Ptr(true),
	// 				ServiceConfigurations: []*armazurestackhci.ServiceConfiguration{
	// 					{
	// 						Port: to.Ptr[int64](6516),
	// 						ServiceName: to.Ptr(armazurestackhci.ServiceNameWAC),
	// 					},
	// 				},
	// 			},
	// 			DefaultExtensions: []*armazurestackhci.DefaultExtensionDetails{
	// 				{
	// 					Category: to.Ptr("Telemetry"),
	// 					ConsentTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-01T17:18:19.1234567Z"); return t}()),
	// 				},
	// 				{
	// 					Category: to.Ptr("Supportability"),
	// 				},
	// 			},
	// 			PerNodeDetails: []*armazurestackhci.PerNodeState{
	// 				{
	// 					Name: to.Ptr("Node-1"),
	// 					ArcInstance: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1"),
	// 					ArcNodeServicePrincipalObjectID: to.Ptr("00cc4014-482e-4de9-9932-83415cc75f45"),
	// 					State: to.Ptr(armazurestackhci.NodeArcStateSucceeded),
	// 				},
	// 				{
	// 					Name: to.Ptr("Node-2"),
	// 					ArcInstance: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-2"),
	// 					ArcNodeServicePrincipalObjectID: to.Ptr("00cc4014-482e-4de9-9932-83415cc75f45"),
	// 					State: to.Ptr(armazurestackhci.NodeArcStateSucceeded),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armazurestackhci.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T17:18:19.1234567Z"); return t}()),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-02T17:18:19.1234567Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
