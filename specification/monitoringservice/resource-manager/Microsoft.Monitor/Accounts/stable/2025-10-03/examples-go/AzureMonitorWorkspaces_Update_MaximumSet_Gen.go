package armmonitorworkspaces_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitorworkspaces"
)

// Generated from example definition: 2025-10-03/AzureMonitorWorkspaces_Update_MaximumSet_Gen.json
func ExampleAzureMonitorWorkspacesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitorworkspaces.NewClientFactory("703362b3-f278-4e4b-9179-c76eaf41ffc2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAzureMonitorWorkspacesClient().Update(ctx, "rgazuremonitorworkspace", "myAzureMonitorWorkspace", armmonitorworkspaces.AzureMonitorWorkspaceResourceUpdate{
		Tags: map[string]*string{},
		Properties: &armmonitorworkspaces.AzureMonitorWorkspace{
			PublicNetworkAccess: to.Ptr(armmonitorworkspaces.PublicNetworkAccessEnabled),
			Metrics: &armmonitorworkspaces.AzureMonitorWorkspaceMetrics{
				EnableAccessUsingResourcePermissions: to.Ptr(true),
			},
		},
		Identity: &armmonitorworkspaces.ManagedServiceIdentity{
			Type: to.Ptr(armmonitorworkspaces.ManagedServiceIdentityTypeSystemAssigned),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmonitorworkspaces.AzureMonitorWorkspacesClientUpdateResponse{
	// 	AzureMonitorWorkspaceResource: armmonitorworkspaces.AzureMonitorWorkspaceResource{
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armmonitorworkspaces.AzureMonitorWorkspace{
	// 			AccountID: to.Ptr("2df515bf-c3ce-4920-84d4-1d9d16542d9f"),
	// 			Metrics: &armmonitorworkspaces.AzureMonitorWorkspaceMetrics{
	// 				PrometheusQueryEndpoint: to.Ptr("https://myAzureMonitorWorkspace-v8hx.eastus.prometheus.monitor.azure.com"),
	// 				InternalID: to.Ptr("mac_2df515bf-c3ce-4920-84d4-1d9d16542d9f"),
	// 				EnableAccessUsingResourcePermissions: to.Ptr(true),
	// 			},
	// 			ProvisioningState: to.Ptr(armmonitorworkspaces.ResourceProvisioningStateSucceeded),
	// 			DefaultIngestionSettings: &armmonitorworkspaces.AzureMonitorWorkspaceDefaultIngestionSettings{
	// 				DataCollectionRuleResourceID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/MA_myAzureMonitorWorkspace_eastus_managed/providers/Microsoft.Insights/dataCollectionRules/myAzureMonitorWorkspace"),
	// 				DataCollectionEndpointResourceID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/MA_myAzureMonitorWorkspace_eastus_managed/providers/Microsoft.Insights/dataCollectionEndpoints/myAzureMonitorWorkspace"),
	// 				DataCollectionRuleImmutableID: to.Ptr("503362b3-f278-4e4b-9179-c76eaf41ffc2"),
	// 				IngestionEndpoints: &armmonitorworkspaces.IngestionEndpoints{
	// 					Metrics: to.Ptr("https://myAzureMonitorWorkspace-gb69.eastus-1.metrics.ingest.monitor.azure.com"),
	// 				},
	// 			},
	// 			PrivateEndpointConnections: []*armmonitorworkspaces.PrivateEndpointConnection{
	// 				{
	// 					ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Monitor/accounts/myAzureMonitorWorkspace/privateEndpointConnections/myPrivateEndpointConnection"),
	// 					Name: to.Ptr("myPrivateEndpointConnection"),
	// 					Type: to.Ptr("Microsoft.Monitor/accounts/privateEndpointConnections"),
	// 					Properties: &armmonitorworkspaces.PrivateEndpointConnectionProperties{
	// 						ProvisioningState: to.Ptr(armmonitorworkspaces.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 						PrivateEndpoint: &armmonitorworkspaces.PrivateEndpoint{
	// 							ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpointConnection"),
	// 						},
	// 						GroupIDs: []*string{
	// 							to.Ptr("prometheusMetrics"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armmonitorworkspaces.PrivateLinkServiceConnectionState{
	// 							Status: to.Ptr(armmonitorworkspaces.PrivateEndpointServiceConnectionStatusApproved),
	// 							ActionsRequired: to.Ptr("None"),
	// 							Description: to.Ptr("jawmectradlpxxaalglrydamehym"),
	// 						},
	// 					},
	// 					SystemData: &armmonitorworkspaces.SystemData{
	// 						CreatedBy: to.Ptr("user1"),
	// 						CreatedByType: to.Ptr(armmonitorworkspaces.CreatedByTypeUser),
	// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T12:34:56.1234567Z"); return t}()),
	// 						LastModifiedBy: to.Ptr("user2"),
	// 						LastModifiedByType: to.Ptr(armmonitorworkspaces.CreatedByTypeUser),
	// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T12:34:56.1234567Z"); return t}()),
	// 					},
	// 				},
	// 			},
	// 			PublicNetworkAccess: to.Ptr(armmonitorworkspaces.PublicNetworkAccessEnabled),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1693": to.Ptr("xkjnypgoxx"),
	// 			"key4981": to.Ptr("akpkhqbugamcavvmdqevahsnqebh"),
	// 		},
	// 		Identity: &armmonitorworkspaces.ManagedServiceIdentity{
	// 			Type: to.Ptr(armmonitorworkspaces.ManagedServiceIdentityTypeSystemAssigned),
	// 			PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Monitor/accounts/myAzureMonitorWorkspace"),
	// 		Name: to.Ptr("myAzureMonitorWorkspace"),
	// 		Type: to.Ptr("Microsoft.Monitor/accounts"),
	// 		SystemData: &armmonitorworkspaces.SystemData{
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armmonitorworkspaces.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T12:34:56.1234567Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armmonitorworkspaces.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T12:34:56.1234567Z"); return t}()),
	// 		},
	// 		Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	},
	// }
}
