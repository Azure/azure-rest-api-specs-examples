package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/FirewallPolicyPut.json
func ExampleFirewallPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<firewall-policy-name>",
		armnetwork.FirewallPolicy{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"key1": to.Ptr("value1"),
			},
			Properties: &armnetwork.FirewallPolicyPropertiesFormat{
				DNSSettings: &armnetwork.DNSSettings{
					EnableProxy:                 to.Ptr(true),
					RequireProxyForNetworkRules: to.Ptr(false),
					Servers: []*string{
						to.Ptr("30.3.4.5")},
				},
				ExplicitProxySettings: &armnetwork.ExplicitProxySettings{
					EnableExplicitProxy: to.Ptr(true),
					HTTPPort:            to.Ptr[int32](8087),
					HTTPSPort:           to.Ptr[int32](8087),
					PacFile:             to.Ptr("<pac-file>"),
					PacFilePort:         to.Ptr[int32](8087),
				},
				Insights: &armnetwork.FirewallPolicyInsights{
					IsEnabled: to.Ptr(true),
					LogAnalyticsResources: &armnetwork.FirewallPolicyLogAnalyticsResources{
						DefaultWorkspaceID: &armnetwork.SubResource{
							ID: to.Ptr("<id>"),
						},
						Workspaces: []*armnetwork.FirewallPolicyLogAnalyticsWorkspace{
							{
								Region: to.Ptr("<region>"),
								WorkspaceID: &armnetwork.SubResource{
									ID: to.Ptr("<id>"),
								},
							},
							{
								Region: to.Ptr("<region>"),
								WorkspaceID: &armnetwork.SubResource{
									ID: to.Ptr("<id>"),
								},
							}},
					},
					RetentionDays: to.Ptr[int32](100),
				},
				IntrusionDetection: &armnetwork.FirewallPolicyIntrusionDetection{
					Configuration: &armnetwork.FirewallPolicyIntrusionDetectionConfiguration{
						BypassTrafficSettings: []*armnetwork.FirewallPolicyIntrusionDetectionBypassTrafficSpecifications{
							{
								Name:        to.Ptr("<name>"),
								Description: to.Ptr("<description>"),
								DestinationAddresses: []*string{
									to.Ptr("5.6.7.8")},
								DestinationPorts: []*string{
									to.Ptr("*")},
								SourceAddresses: []*string{
									to.Ptr("1.2.3.4")},
								Protocol: to.Ptr(armnetwork.FirewallPolicyIntrusionDetectionProtocolTCP),
							}},
						SignatureOverrides: []*armnetwork.FirewallPolicyIntrusionDetectionSignatureSpecification{
							{
								ID:   to.Ptr("<id>"),
								Mode: to.Ptr(armnetwork.FirewallPolicyIntrusionDetectionStateTypeDeny),
							}},
					},
					Mode: to.Ptr(armnetwork.FirewallPolicyIntrusionDetectionStateTypeAlert),
				},
				SKU: &armnetwork.FirewallPolicySKU{
					Tier: to.Ptr(armnetwork.FirewallPolicySKUTierPremium),
				},
				Snat: &armnetwork.FirewallPolicySNAT{
					PrivateRanges: []*string{
						to.Ptr("IANAPrivateRanges")},
				},
				SQL: &armnetwork.FirewallPolicySQL{
					AllowSQLRedirect: to.Ptr(true),
				},
				ThreatIntelMode: to.Ptr(armnetwork.AzureFirewallThreatIntelModeAlert),
				ThreatIntelWhitelist: &armnetwork.FirewallPolicyThreatIntelWhitelist{
					Fqdns: []*string{
						to.Ptr("*.microsoft.com")},
					IPAddresses: []*string{
						to.Ptr("20.3.4.5")},
				},
				TransportSecurity: &armnetwork.FirewallPolicyTransportSecurity{
					CertificateAuthority: &armnetwork.FirewallPolicyCertificateAuthority{
						Name:             to.Ptr("<name>"),
						KeyVaultSecretID: to.Ptr("<key-vault-secret-id>"),
					},
				},
			},
		},
		&armnetwork.FirewallPoliciesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
