Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.3.1/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/FirewallPolicyPut.json
func ExampleFirewallPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewFirewallPoliciesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<firewall-policy-name>",
		armnetwork.FirewallPolicy{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"key1": to.StringPtr("value1"),
			},
			Properties: &armnetwork.FirewallPolicyPropertiesFormat{
				DNSSettings: &armnetwork.DNSSettings{
					EnableProxy:                 to.BoolPtr(true),
					RequireProxyForNetworkRules: to.BoolPtr(false),
					Servers: []*string{
						to.StringPtr("30.3.4.5")},
				},
				ExplicitProxySettings: &armnetwork.ExplicitProxySettings{
					EnableExplicitProxy: to.BoolPtr(true),
					HTTPPort:            to.Int32Ptr(8087),
					HTTPSPort:           to.Int32Ptr(8087),
					PacFile:             to.StringPtr("<pac-file>"),
					PacFilePort:         to.Int32Ptr(8087),
				},
				Insights: &armnetwork.FirewallPolicyInsights{
					IsEnabled: to.BoolPtr(true),
					LogAnalyticsResources: &armnetwork.FirewallPolicyLogAnalyticsResources{
						DefaultWorkspaceID: &armnetwork.SubResource{
							ID: to.StringPtr("<id>"),
						},
						Workspaces: []*armnetwork.FirewallPolicyLogAnalyticsWorkspace{
							{
								Region: to.StringPtr("<region>"),
								WorkspaceID: &armnetwork.SubResource{
									ID: to.StringPtr("<id>"),
								},
							},
							{
								Region: to.StringPtr("<region>"),
								WorkspaceID: &armnetwork.SubResource{
									ID: to.StringPtr("<id>"),
								},
							}},
					},
					RetentionDays: to.Int32Ptr(100),
				},
				IntrusionDetection: &armnetwork.FirewallPolicyIntrusionDetection{
					Configuration: &armnetwork.FirewallPolicyIntrusionDetectionConfiguration{
						BypassTrafficSettings: []*armnetwork.FirewallPolicyIntrusionDetectionBypassTrafficSpecifications{
							{
								Name:        to.StringPtr("<name>"),
								Description: to.StringPtr("<description>"),
								DestinationAddresses: []*string{
									to.StringPtr("5.6.7.8")},
								DestinationPorts: []*string{
									to.StringPtr("*")},
								SourceAddresses: []*string{
									to.StringPtr("1.2.3.4")},
								Protocol: armnetwork.FirewallPolicyIntrusionDetectionProtocol("TCP").ToPtr(),
							}},
						SignatureOverrides: []*armnetwork.FirewallPolicyIntrusionDetectionSignatureSpecification{
							{
								ID:   to.StringPtr("<id>"),
								Mode: armnetwork.FirewallPolicyIntrusionDetectionStateType("Deny").ToPtr(),
							}},
					},
					Mode: armnetwork.FirewallPolicyIntrusionDetectionStateType("Alert").ToPtr(),
				},
				SKU: &armnetwork.FirewallPolicySKU{
					Tier: armnetwork.FirewallPolicySKUTier("Premium").ToPtr(),
				},
				Snat: &armnetwork.FirewallPolicySNAT{
					PrivateRanges: []*string{
						to.StringPtr("IANAPrivateRanges")},
				},
				SQL: &armnetwork.FirewallPolicySQL{
					AllowSQLRedirect: to.BoolPtr(true),
				},
				ThreatIntelMode: armnetwork.AzureFirewallThreatIntelMode("Alert").ToPtr(),
				ThreatIntelWhitelist: &armnetwork.FirewallPolicyThreatIntelWhitelist{
					Fqdns: []*string{
						to.StringPtr("*.microsoft.com")},
					IPAddresses: []*string{
						to.StringPtr("20.3.4.5")},
				},
				TransportSecurity: &armnetwork.FirewallPolicyTransportSecurity{
					CertificateAuthority: &armnetwork.FirewallPolicyCertificateAuthority{
						Name:             to.StringPtr("<name>"),
						KeyVaultSecretID: to.StringPtr("<key-vault-secret-id>"),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.FirewallPoliciesClientCreateOrUpdateResult)
}
```
