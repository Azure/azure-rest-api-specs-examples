package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ab04533261eff228f28e08900445d0edef3eb70c/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/FirewallPolicyDraftGet.json
func ExampleFirewallPolicyDraftsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirewallPolicyDraftsClient().Get(ctx, "rg1", "firewallPolicy", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FirewallPolicyDraft = armnetwork.FirewallPolicyDraft{
	// 	Name: to.Ptr("firewallPolicy"),
	// 	Type: to.Ptr("Microsoft.Network/firewallPolicies"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/firewallPolicy"),
	// 	Properties: &armnetwork.FirewallPolicyDraftProperties{
	// 		DNSSettings: &armnetwork.DNSSettings{
	// 			EnableProxy: to.Ptr(true),
	// 			RequireProxyForNetworkRules: to.Ptr(false),
	// 			Servers: []*string{
	// 				to.Ptr("30.3.4.5")},
	// 			},
	// 			ExplicitProxy: &armnetwork.ExplicitProxySettings{
	// 				EnableExplicitProxy: to.Ptr(true),
	// 				EnablePacFile: to.Ptr(true),
	// 				HTTPPort: to.Ptr[int32](8087),
	// 				HTTPSPort: to.Ptr[int32](8087),
	// 				PacFile: to.Ptr("https://tinawstorage.file.core.windows.net/?sv=2020-02-10&ss=bfqt&srt=sco&sp=rwdlacuptfx&se=2021-06-04T07:01:12Z&st=2021-06-03T23:01:12Z&sip=68.65.171.11&spr=https&sig=Plsa0RRVpGbY0IETZZOT6znOHcSro71LLTTbzquYPgs%3D"),
	// 				PacFilePort: to.Ptr[int32](8087),
	// 			},
	// 			Insights: &armnetwork.FirewallPolicyInsights{
	// 				IsEnabled: to.Ptr(true),
	// 				LogAnalyticsResources: &armnetwork.FirewallPolicyLogAnalyticsResources{
	// 					DefaultWorkspaceID: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/microsoft.operationalinsights/workspaces/defaultWorkspace"),
	// 					},
	// 					Workspaces: []*armnetwork.FirewallPolicyLogAnalyticsWorkspace{
	// 						{
	// 							Region: to.Ptr("westus"),
	// 							WorkspaceID: &armnetwork.SubResource{
	// 								ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/microsoft.operationalinsights/workspaces/workspace1"),
	// 							},
	// 						},
	// 						{
	// 							Region: to.Ptr("eastus"),
	// 							WorkspaceID: &armnetwork.SubResource{
	// 								ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/microsoft.operationalinsights/workspaces/workspace2"),
	// 							},
	// 					}},
	// 				},
	// 				RetentionDays: to.Ptr[int32](100),
	// 			},
	// 			IntrusionDetection: &armnetwork.FirewallPolicyIntrusionDetection{
	// 				Configuration: &armnetwork.FirewallPolicyIntrusionDetectionConfiguration{
	// 					BypassTrafficSettings: []*armnetwork.FirewallPolicyIntrusionDetectionBypassTrafficSpecifications{
	// 						{
	// 							Name: to.Ptr("bypassRule1"),
	// 							Description: to.Ptr("Rule 1"),
	// 							DestinationAddresses: []*string{
	// 								to.Ptr("5.6.7.8")},
	// 								DestinationPorts: []*string{
	// 									to.Ptr("*")},
	// 									SourceAddresses: []*string{
	// 										to.Ptr("1.2.3.4")},
	// 										Protocol: to.Ptr(armnetwork.FirewallPolicyIntrusionDetectionProtocolTCP),
	// 								}},
	// 								SignatureOverrides: []*armnetwork.FirewallPolicyIntrusionDetectionSignatureSpecification{
	// 									{
	// 										ID: to.Ptr("2525004"),
	// 										Mode: to.Ptr(armnetwork.FirewallPolicyIntrusionDetectionStateTypeDeny),
	// 								}},
	// 							},
	// 							Mode: to.Ptr(armnetwork.FirewallPolicyIntrusionDetectionStateTypeAlert),
	// 						},
	// 						Snat: &armnetwork.FirewallPolicySNAT{
	// 							PrivateRanges: []*string{
	// 								to.Ptr("IANAPrivateRanges")},
	// 							},
	// 							SQL: &armnetwork.FirewallPolicySQL{
	// 								AllowSQLRedirect: to.Ptr(true),
	// 							},
	// 							ThreatIntelMode: to.Ptr(armnetwork.AzureFirewallThreatIntelModeAlert),
	// 							ThreatIntelWhitelist: &armnetwork.FirewallPolicyThreatIntelWhitelist{
	// 								Fqdns: []*string{
	// 									to.Ptr("*.microsoft.com")},
	// 									IPAddresses: []*string{
	// 										to.Ptr("20.3.4.5")},
	// 									},
	// 								},
	// 							}
}
