package armdashboard_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard/v2"
)

// Generated from example definition: 2024-11-01-preview/Grafana_List.json
func ExampleGrafanaClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdashboard.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGrafanaClient().NewListPager(nil)
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
		// page = armdashboard.GrafanaClientListResponse{
		// 	ManagedGrafanaListResponse: armdashboard.ManagedGrafanaListResponse{
		// 		Value: []*armdashboard.ManagedGrafana{
		// 			{
		// 				Name: to.Ptr("myWorkspace"),
		// 				Type: to.Ptr("Microsoft.Dashboard/grafana"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/grafana/myWorkspace"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armdashboard.ManagedGrafanaProperties{
		// 					APIKey: to.Ptr(armdashboard.APIKeyEnabled),
		// 					AutoGeneratedDomainNameLabelScope: to.Ptr(armdashboard.AutoGeneratedDomainNameLabelScopeTenantReuse),
		// 					DeterministicOutboundIP: to.Ptr(armdashboard.DeterministicOutboundIPEnabled),
		// 					Endpoint: to.Ptr("https://myworkspace-abcdefghijklmnop.wus.grafana.azure.com"),
		// 					EnterpriseConfigurations: &armdashboard.EnterpriseConfigurations{
		// 						MarketplaceAutoRenew: to.Ptr(armdashboard.MarketplaceAutoRenewEnabled),
		// 						MarketplacePlanID: to.Ptr("myPlanId"),
		// 					},
		// 					GrafanaConfigurations: &armdashboard.GrafanaConfigurations{
		// 						Security: &armdashboard.Security{
		// 							CsrfAlwaysCheck: to.Ptr(false),
		// 						},
		// 						SMTP: &armdashboard.SMTP{
		// 							Enabled: to.Ptr(true),
		// 							FromAddress: to.Ptr("test@sendemail.com"),
		// 							FromName: to.Ptr("emailsender"),
		// 							Host: to.Ptr("smtp.sendemail.com:587"),
		// 							SkipVerify: to.Ptr(true),
		// 							StartTLSPolicy: to.Ptr(armdashboard.StartTLSPolicyOpportunisticStartTLS),
		// 							User: to.Ptr("username"),
		// 						},
		// 						Snapshots: &armdashboard.Snapshots{
		// 							ExternalEnabled: to.Ptr(true),
		// 						},
		// 						UnifiedAlertingScreenshots: &armdashboard.UnifiedAlertingScreenshots{
		// 							CaptureEnabled: to.Ptr(false),
		// 						},
		// 						Users: &armdashboard.Users{
		// 							EditorsCanAdmin: to.Ptr(true),
		// 							ViewersCanEdit: to.Ptr(true),
		// 						},
		// 					},
		// 					GrafanaIntegrations: &armdashboard.GrafanaIntegrations{
		// 						AzureMonitorWorkspaceIntegrations: []*armdashboard.AzureMonitorWorkspaceIntegration{
		// 							{
		// 								AzureMonitorWorkspaceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.monitor/accounts/myAzureMonitorWorkspace"),
		// 							},
		// 						},
		// 					},
		// 					GrafanaMajorVersion: to.Ptr("9"),
		// 					GrafanaPlugins: map[string]*armdashboard.GrafanaPlugin{
		// 						"sample-plugin-id": &armdashboard.GrafanaPlugin{
		// 							PluginID: to.Ptr("sample-plugin-id"),
		// 						},
		// 					},
		// 					GrafanaVersion: to.Ptr("9.4.5"),
		// 					OutboundIPs: []*string{
		// 						to.Ptr("192.168.0.1"),
		// 						to.Ptr("192.168.0.2"),
		// 					},
		// 					ProvisioningState: to.Ptr(armdashboard.ProvisioningStateSucceeded),
		// 					PublicNetworkAccess: to.Ptr(armdashboard.PublicNetworkAccessEnabled),
		// 					ZoneRedundancy: to.Ptr(armdashboard.ZoneRedundancyEnabled),
		// 				},
		// 				SKU: &armdashboard.ResourceSKU{
		// 					Name: to.Ptr("Standard"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"Environment": to.Ptr("Dev"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
