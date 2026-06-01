package armmonitorworkspaces_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitorworkspaces"
)

// Generated from example definition: 2025-10-03/AzureMonitorWorkspaces_ListBySubscription_MaximumSet_Gen.json
func ExampleAzureMonitorWorkspacesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitorworkspaces.NewClientFactory("703362b3-f278-4e4b-9179-c76eaf41ffc2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAzureMonitorWorkspacesClient().NewListBySubscriptionPager(nil)
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
		// page = armmonitorworkspaces.AzureMonitorWorkspacesClientListBySubscriptionResponse{
		// 	AzureMonitorWorkspaceResourceListResult: armmonitorworkspaces.AzureMonitorWorkspaceResourceListResult{
		// 		Value: []*armmonitorworkspaces.AzureMonitorWorkspaceResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Monitor/accounts/myAzureMonitorWorkspace"),
		// 				Name: to.Ptr("myAzureMonitorWorkspace"),
		// 				Type: to.Ptr("Microsoft.Monitor/accounts"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armmonitorworkspaces.AzureMonitorWorkspace{
		// 					AccountID: to.Ptr("2df515bf-c3ce-4920-84d4-1d9d16542d9f"),
		// 					Metrics: &armmonitorworkspaces.AzureMonitorWorkspaceMetrics{
		// 						PrometheusQueryEndpoint: to.Ptr("https://myAzureMonitorWorkspace-v8hx.eastus.prometheus.monitor.azure.com"),
		// 						InternalID: to.Ptr("mac_2df515bf-c3ce-4920-84d4-1d9d16542d9f"),
		// 						EnableAccessUsingResourcePermissions: to.Ptr(true),
		// 					},
		// 					ProvisioningState: to.Ptr(armmonitorworkspaces.ResourceProvisioningStateSucceeded),
		// 					DefaultIngestionSettings: &armmonitorworkspaces.AzureMonitorWorkspaceDefaultIngestionSettings{
		// 						DataCollectionRuleResourceID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/MA_myAzureMonitorWorkspace_eastus_managed/providers/Microsoft.Insights/dataCollectionRules/myAzureMonitorWorkspace"),
		// 						DataCollectionEndpointResourceID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/MA_myAzureMonitorWorkspace_eastus_managed/providers/Microsoft.Insights/dataCollectionEndpoints/myAzureMonitorWorkspace"),
		// 						DataCollectionRuleImmutableID: to.Ptr("603362b3-f278-4e4b-9179-c76eaf41ffc2"),
		// 						IngestionEndpoints: &armmonitorworkspaces.IngestionEndpoints{
		// 							Metrics: to.Ptr("https://myAzureMonitorWorkspace-gb69.eastus-1.metrics.ingest.monitor.azure.com"),
		// 						},
		// 					},
		// 					PrivateEndpointConnections: []*armmonitorworkspaces.PrivateEndpointConnection{
		// 						{
		// 							ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Monitor/accounts/myAzureMonitorWorkspace/privateEndpointConnections/myPrivateEndpointConnection"),
		// 							Name: to.Ptr("myPrivateEndpointConnection"),
		// 							Type: to.Ptr("Microsoft.Monitor/accounts/privateEndpointConnections"),
		// 							Properties: &armmonitorworkspaces.PrivateEndpointConnectionProperties{
		// 								ProvisioningState: to.Ptr(armmonitorworkspaces.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 								PrivateEndpoint: &armmonitorworkspaces.PrivateEndpoint{
		// 									ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpointConnection"),
		// 								},
		// 								GroupIDs: []*string{
		// 									to.Ptr("prometheusMetrics"),
		// 								},
		// 								PrivateLinkServiceConnectionState: &armmonitorworkspaces.PrivateLinkServiceConnectionState{
		// 									Status: to.Ptr(armmonitorworkspaces.PrivateEndpointServiceConnectionStatusApproved),
		// 									ActionsRequired: to.Ptr("None"),
		// 									Description: to.Ptr("jawmectradlpxxaalglrydamehym"),
		// 								},
		// 							},
		// 							SystemData: &armmonitorworkspaces.SystemData{
		// 								CreatedBy: to.Ptr("user1"),
		// 								CreatedByType: to.Ptr(armmonitorworkspaces.CreatedByTypeUser),
		// 								CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T12:34:56.1234567Z"); return t}()),
		// 								LastModifiedBy: to.Ptr("user2"),
		// 								LastModifiedByType: to.Ptr(armmonitorworkspaces.CreatedByTypeUser),
		// 								LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T12:34:56.1234567Z"); return t}()),
		// 							},
		// 						},
		// 					},
		// 					PublicNetworkAccess: to.Ptr(armmonitorworkspaces.PublicNetworkAccessEnabled),
		// 				},
		// 				SystemData: &armmonitorworkspaces.SystemData{
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armmonitorworkspaces.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T12:34:56.1234567Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armmonitorworkspaces.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T12:34:56.1234567Z"); return t}()),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key4981": to.Ptr("akpkhqbugamcavvmdqevahsnqebh"),
		// 				},
		// 				Identity: &armmonitorworkspaces.ManagedServiceIdentity{
		// 					Type: to.Ptr(armmonitorworkspaces.ManagedServiceIdentityTypeSystemAssigned),
		// 					PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				},
		// 				Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/herResourceGroup/providers/Microsoft.Monitor/accounts/herAzureMonitorWorkspace"),
		// 				Name: to.Ptr("herAzureMonitorWorkspace"),
		// 				Type: to.Ptr("Microsoft.Monitor/accounts"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armmonitorworkspaces.AzureMonitorWorkspace{
		// 					AccountID: to.Ptr("823220c6-0415-44d8-bfb2-d5c1c9ea1172"),
		// 					Metrics: &armmonitorworkspaces.AzureMonitorWorkspaceMetrics{
		// 						PrometheusQueryEndpoint: to.Ptr("https://herAzureMonitorWorkspace-xywz.westus.prometheus.monitor.azure.com"),
		// 						InternalID: to.Ptr("mac_823220c6-0415-44d8-bfb2-d5c1c9ea1172"),
		// 						EnableAccessUsingResourcePermissions: to.Ptr(true),
		// 					},
		// 					ProvisioningState: to.Ptr(armmonitorworkspaces.ResourceProvisioningStateSucceeded),
		// 					DefaultIngestionSettings: &armmonitorworkspaces.AzureMonitorWorkspaceDefaultIngestionSettings{
		// 						DataCollectionRuleResourceID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/MA_herAzureMonitorWorkspace_westus_managed/providers/Microsoft.Insights/dataCollectionRules/herAzureMonitorWorkspace"),
		// 						DataCollectionEndpointResourceID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/MA_herAzureMonitorWorkspace_westus_managed/providers/Microsoft.Insights/dataCollectionEndpoints/herAzureMonitorWorkspace"),
		// 						DataCollectionRuleImmutableID: to.Ptr("903362b3-f278-4e4b-9179-c76eaf41ffc2"),
		// 						IngestionEndpoints: &armmonitorworkspaces.IngestionEndpoints{
		// 							Metrics: to.Ptr("https://herAzureMonitorWorkspace-gb69.westus-1.metrics.ingest.monitor.azure.com"),
		// 						},
		// 					},
		// 					PrivateEndpointConnections: []*armmonitorworkspaces.PrivateEndpointConnection{
		// 						{
		// 							ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/herResourceGroup/providers/Microsoft.Monitor/accounts/herAzureMonitorWorkspace/privateEndpointConnections/herPrivateEndpointConnection"),
		// 							Name: to.Ptr("herPrivateEndpointConnection"),
		// 							Type: to.Ptr("Microsoft.Monitor/accounts/privateEndpointConnections"),
		// 							Properties: &armmonitorworkspaces.PrivateEndpointConnectionProperties{
		// 								ProvisioningState: to.Ptr(armmonitorworkspaces.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 								PrivateEndpoint: &armmonitorworkspaces.PrivateEndpoint{
		// 									ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/herResourceGroup/providers/Microsoft.Network/privateEndpoints/herPrivateEndpointConnection"),
		// 								},
		// 								GroupIDs: []*string{
		// 									to.Ptr("prometheusMetrics"),
		// 								},
		// 								PrivateLinkServiceConnectionState: &armmonitorworkspaces.PrivateLinkServiceConnectionState{
		// 									Status: to.Ptr(armmonitorworkspaces.PrivateEndpointServiceConnectionStatusApproved),
		// 									ActionsRequired: to.Ptr("None"),
		// 									Description: to.Ptr("ujtgmztnkqnqzcawmbbmr"),
		// 								},
		// 							},
		// 							SystemData: &armmonitorworkspaces.SystemData{
		// 								CreatedBy: to.Ptr("user1"),
		// 								CreatedByType: to.Ptr(armmonitorworkspaces.CreatedByTypeUser),
		// 								CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T12:34:56.1234567Z"); return t}()),
		// 								LastModifiedBy: to.Ptr("user2"),
		// 								LastModifiedByType: to.Ptr(armmonitorworkspaces.CreatedByTypeUser),
		// 								LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T12:34:56.1234567Z"); return t}()),
		// 							},
		// 						},
		// 					},
		// 					PublicNetworkAccess: to.Ptr(armmonitorworkspaces.PublicNetworkAccessEnabled),
		// 				},
		// 				SystemData: &armmonitorworkspaces.SystemData{
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armmonitorworkspaces.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T12:34:56.1234567Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armmonitorworkspaces.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T12:34:56.1234567Z"); return t}()),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key583": to.Ptr("deudtkpbrajjhz"),
		// 				},
		// 				Identity: &armmonitorworkspaces.ManagedServiceIdentity{
		// 					Type: to.Ptr(armmonitorworkspaces.ManagedServiceIdentityTypeSystemAssigned),
		// 					PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 					TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				},
		// 				Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
