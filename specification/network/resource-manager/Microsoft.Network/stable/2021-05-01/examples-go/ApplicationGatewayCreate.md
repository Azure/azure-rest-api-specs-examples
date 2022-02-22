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

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ApplicationGatewayCreate.json
func ExampleApplicationGatewaysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewApplicationGatewaysClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<application-gateway-name>",
		armnetwork.ApplicationGateway{
			Location: to.StringPtr("<location>"),
			Identity: &armnetwork.ManagedServiceIdentity{
				Type: armnetwork.ResourceIdentityTypeUserAssigned.ToPtr(),
				UserAssignedIdentities: map[string]*armnetwork.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
					"/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": {},
				},
			},
			Properties: &armnetwork.ApplicationGatewayPropertiesFormat{
				BackendAddressPools: []*armnetwork.ApplicationGatewayBackendAddressPool{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.ApplicationGatewayBackendAddressPoolPropertiesFormat{
							BackendAddresses: []*armnetwork.ApplicationGatewayBackendAddress{
								{
									IPAddress: to.StringPtr("<ipaddress>"),
								},
								{
									IPAddress: to.StringPtr("<ipaddress>"),
								}},
						},
					},
					{
						ID:   to.StringPtr("<id>"),
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.ApplicationGatewayBackendAddressPoolPropertiesFormat{
							BackendAddresses: []*armnetwork.ApplicationGatewayBackendAddress{
								{},
								{}},
						},
					}},
				BackendHTTPSettingsCollection: []*armnetwork.ApplicationGatewayBackendHTTPSettings{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.ApplicationGatewayBackendHTTPSettingsPropertiesFormat{
							CookieBasedAffinity: armnetwork.ApplicationGatewayCookieBasedAffinity("Disabled").ToPtr(),
							Port:                to.Int32Ptr(80),
							RequestTimeout:      to.Int32Ptr(30),
							Protocol:            armnetwork.ApplicationGatewayProtocol("Http").ToPtr(),
						},
					}},
				FrontendIPConfigurations: []*armnetwork.ApplicationGatewayFrontendIPConfiguration{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.ApplicationGatewayFrontendIPConfigurationPropertiesFormat{
							PublicIPAddress: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
						},
					}},
				FrontendPorts: []*armnetwork.ApplicationGatewayFrontendPort{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.ApplicationGatewayFrontendPortPropertiesFormat{
							Port: to.Int32Ptr(443),
						},
					},
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.ApplicationGatewayFrontendPortPropertiesFormat{
							Port: to.Int32Ptr(80),
						},
					}},
				GatewayIPConfigurations: []*armnetwork.ApplicationGatewayIPConfiguration{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.ApplicationGatewayIPConfigurationPropertiesFormat{
							Subnet: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
						},
					}},
				GlobalConfiguration: &armnetwork.ApplicationGatewayGlobalConfiguration{
					EnableRequestBuffering:  to.BoolPtr(true),
					EnableResponseBuffering: to.BoolPtr(true),
				},
				HTTPListeners: []*armnetwork.ApplicationGatewayHTTPListener{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.ApplicationGatewayHTTPListenerPropertiesFormat{
							FrontendIPConfiguration: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							FrontendPort: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							RequireServerNameIndication: to.BoolPtr(false),
							SSLCertificate: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							SSLProfile: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							Protocol: armnetwork.ApplicationGatewayProtocol("Https").ToPtr(),
						},
					},
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.ApplicationGatewayHTTPListenerPropertiesFormat{
							FrontendIPConfiguration: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							FrontendPort: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							Protocol: armnetwork.ApplicationGatewayProtocol("Http").ToPtr(),
						},
					}},
				LoadDistributionPolicies: []*armnetwork.ApplicationGatewayLoadDistributionPolicy{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.ApplicationGatewayLoadDistributionPolicyPropertiesFormat{
							LoadDistributionAlgorithm: armnetwork.ApplicationGatewayLoadDistributionAlgorithm("RoundRobin").ToPtr(),
							LoadDistributionTargets: []*armnetwork.ApplicationGatewayLoadDistributionTarget{
								{
									Name: to.StringPtr("<name>"),
									Properties: &armnetwork.ApplicationGatewayLoadDistributionTargetPropertiesFormat{
										BackendAddressPool: &armnetwork.SubResource{
											ID: to.StringPtr("<id>"),
										},
										WeightPerServer: to.Int32Ptr(40),
									},
								},
								{
									Name: to.StringPtr("<name>"),
									Properties: &armnetwork.ApplicationGatewayLoadDistributionTargetPropertiesFormat{
										BackendAddressPool: &armnetwork.SubResource{
											ID: to.StringPtr("<id>"),
										},
										WeightPerServer: to.Int32Ptr(60),
									},
								}},
						},
					}},
				RequestRoutingRules: []*armnetwork.ApplicationGatewayRequestRoutingRule{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.ApplicationGatewayRequestRoutingRulePropertiesFormat{
							BackendAddressPool: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							BackendHTTPSettings: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							HTTPListener: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							LoadDistributionPolicy: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							Priority: to.Int32Ptr(10),
							RewriteRuleSet: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							RuleType: armnetwork.ApplicationGatewayRequestRoutingRuleType("Basic").ToPtr(),
						},
					},
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.ApplicationGatewayRequestRoutingRulePropertiesFormat{
							HTTPListener: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							Priority: to.Int32Ptr(20),
							RuleType: armnetwork.ApplicationGatewayRequestRoutingRuleType("PathBasedRouting").ToPtr(),
							URLPathMap: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
						},
					}},
				RewriteRuleSets: []*armnetwork.ApplicationGatewayRewriteRuleSet{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.ApplicationGatewayRewriteRuleSetPropertiesFormat{
							RewriteRules: []*armnetwork.ApplicationGatewayRewriteRule{
								{
									Name: to.StringPtr("<name>"),
									ActionSet: &armnetwork.ApplicationGatewayRewriteRuleActionSet{
										RequestHeaderConfigurations: []*armnetwork.ApplicationGatewayHeaderConfiguration{
											{
												HeaderName:  to.StringPtr("<header-name>"),
												HeaderValue: to.StringPtr("<header-value>"),
											}},
										ResponseHeaderConfigurations: []*armnetwork.ApplicationGatewayHeaderConfiguration{
											{
												HeaderName:  to.StringPtr("<header-name>"),
												HeaderValue: to.StringPtr("<header-value>"),
											}},
										URLConfiguration: &armnetwork.ApplicationGatewayURLConfiguration{
											ModifiedPath: to.StringPtr("<modified-path>"),
										},
									},
									Conditions: []*armnetwork.ApplicationGatewayRewriteRuleCondition{
										{
											IgnoreCase: to.BoolPtr(true),
											Negate:     to.BoolPtr(false),
											Pattern:    to.StringPtr("<pattern>"),
											Variable:   to.StringPtr("<variable>"),
										}},
									RuleSequence: to.Int32Ptr(102),
								}},
						},
					}},
				SKU: &armnetwork.ApplicationGatewaySKU{
					Name:     armnetwork.ApplicationGatewaySKUName("Standard_v2").ToPtr(),
					Capacity: to.Int32Ptr(3),
					Tier:     armnetwork.ApplicationGatewayTier("Standard_v2").ToPtr(),
				},
				SSLCertificates: []*armnetwork.ApplicationGatewaySSLCertificate{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.ApplicationGatewaySSLCertificatePropertiesFormat{
							Data:     to.StringPtr("<data>"),
							Password: to.StringPtr("<password>"),
						},
					},
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.ApplicationGatewaySSLCertificatePropertiesFormat{
							KeyVaultSecretID: to.StringPtr("<key-vault-secret-id>"),
						},
					}},
				SSLProfiles: []*armnetwork.ApplicationGatewaySSLProfile{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.ApplicationGatewaySSLProfilePropertiesFormat{
							ClientAuthConfiguration: &armnetwork.ApplicationGatewayClientAuthConfiguration{
								VerifyClientCertIssuerDN: to.BoolPtr(true),
							},
							SSLPolicy: &armnetwork.ApplicationGatewaySSLPolicy{
								CipherSuites: []*armnetwork.ApplicationGatewaySSLCipherSuite{
									armnetwork.ApplicationGatewaySSLCipherSuite("TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256").ToPtr()},
								MinProtocolVersion: armnetwork.ApplicationGatewaySSLProtocol("TLSv1_1").ToPtr(),
								PolicyType:         armnetwork.ApplicationGatewaySSLPolicyType("Custom").ToPtr(),
							},
							TrustedClientCertificates: []*armnetwork.SubResource{
								{
									ID: to.StringPtr("<id>"),
								}},
						},
					}},
				TrustedClientCertificates: []*armnetwork.ApplicationGatewayTrustedClientCertificate{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.ApplicationGatewayTrustedClientCertificatePropertiesFormat{
							Data: to.StringPtr("<data>"),
						},
					}},
				TrustedRootCertificates: []*armnetwork.ApplicationGatewayTrustedRootCertificate{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.ApplicationGatewayTrustedRootCertificatePropertiesFormat{
							Data: to.StringPtr("<data>"),
						},
					},
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.ApplicationGatewayTrustedRootCertificatePropertiesFormat{
							KeyVaultSecretID: to.StringPtr("<key-vault-secret-id>"),
						},
					}},
				URLPathMaps: []*armnetwork.ApplicationGatewayURLPathMap{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.ApplicationGatewayURLPathMapPropertiesFormat{
							DefaultBackendAddressPool: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							DefaultBackendHTTPSettings: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							DefaultLoadDistributionPolicy: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							DefaultRewriteRuleSet: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							PathRules: []*armnetwork.ApplicationGatewayPathRule{
								{
									Name: to.StringPtr("<name>"),
									Properties: &armnetwork.ApplicationGatewayPathRulePropertiesFormat{
										BackendAddressPool: &armnetwork.SubResource{
											ID: to.StringPtr("<id>"),
										},
										BackendHTTPSettings: &armnetwork.SubResource{
											ID: to.StringPtr("<id>"),
										},
										LoadDistributionPolicy: &armnetwork.SubResource{
											ID: to.StringPtr("<id>"),
										},
										Paths: []*string{
											to.StringPtr("/api"),
											to.StringPtr("/v1/api")},
										RewriteRuleSet: &armnetwork.SubResource{
											ID: to.StringPtr("<id>"),
										},
									},
								}},
						},
					}},
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
	log.Printf("Response result: %#v\n", res.ApplicationGatewaysClientCreateOrUpdateResult)
}
```
