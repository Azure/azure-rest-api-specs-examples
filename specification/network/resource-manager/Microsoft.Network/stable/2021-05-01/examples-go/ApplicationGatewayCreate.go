package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ApplicationGatewayCreate.json
func ExampleApplicationGatewaysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewApplicationGatewaysClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<application-gateway-name>",
		armnetwork.ApplicationGateway{
			Location: to.Ptr("<location>"),
			Identity: &armnetwork.ManagedServiceIdentity{
				Type: to.Ptr(armnetwork.ResourceIdentityTypeUserAssigned),
				UserAssignedIdentities: map[string]*armnetwork.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
					"/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": {},
				},
			},
			Properties: &armnetwork.ApplicationGatewayPropertiesFormat{
				BackendAddressPools: []*armnetwork.ApplicationGatewayBackendAddressPool{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ApplicationGatewayBackendAddressPoolPropertiesFormat{
							BackendAddresses: []*armnetwork.ApplicationGatewayBackendAddress{
								{
									IPAddress: to.Ptr("<ipaddress>"),
								},
								{
									IPAddress: to.Ptr("<ipaddress>"),
								}},
						},
					},
					{
						ID:   to.Ptr("<id>"),
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ApplicationGatewayBackendAddressPoolPropertiesFormat{
							BackendAddresses: []*armnetwork.ApplicationGatewayBackendAddress{
								{},
								{}},
						},
					}},
				BackendHTTPSettingsCollection: []*armnetwork.ApplicationGatewayBackendHTTPSettings{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ApplicationGatewayBackendHTTPSettingsPropertiesFormat{
							CookieBasedAffinity: to.Ptr(armnetwork.ApplicationGatewayCookieBasedAffinityDisabled),
							Port:                to.Ptr[int32](80),
							RequestTimeout:      to.Ptr[int32](30),
							Protocol:            to.Ptr(armnetwork.ApplicationGatewayProtocolHTTP),
						},
					}},
				FrontendIPConfigurations: []*armnetwork.ApplicationGatewayFrontendIPConfiguration{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ApplicationGatewayFrontendIPConfigurationPropertiesFormat{
							PublicIPAddress: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
						},
					}},
				FrontendPorts: []*armnetwork.ApplicationGatewayFrontendPort{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ApplicationGatewayFrontendPortPropertiesFormat{
							Port: to.Ptr[int32](443),
						},
					},
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ApplicationGatewayFrontendPortPropertiesFormat{
							Port: to.Ptr[int32](80),
						},
					}},
				GatewayIPConfigurations: []*armnetwork.ApplicationGatewayIPConfiguration{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ApplicationGatewayIPConfigurationPropertiesFormat{
							Subnet: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
						},
					}},
				GlobalConfiguration: &armnetwork.ApplicationGatewayGlobalConfiguration{
					EnableRequestBuffering:  to.Ptr(true),
					EnableResponseBuffering: to.Ptr(true),
				},
				HTTPListeners: []*armnetwork.ApplicationGatewayHTTPListener{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ApplicationGatewayHTTPListenerPropertiesFormat{
							FrontendIPConfiguration: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							FrontendPort: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							RequireServerNameIndication: to.Ptr(false),
							SSLCertificate: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							SSLProfile: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							Protocol: to.Ptr(armnetwork.ApplicationGatewayProtocolHTTPS),
						},
					},
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ApplicationGatewayHTTPListenerPropertiesFormat{
							FrontendIPConfiguration: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							FrontendPort: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							Protocol: to.Ptr(armnetwork.ApplicationGatewayProtocolHTTP),
						},
					}},
				LoadDistributionPolicies: []*armnetwork.ApplicationGatewayLoadDistributionPolicy{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ApplicationGatewayLoadDistributionPolicyPropertiesFormat{
							LoadDistributionAlgorithm: to.Ptr(armnetwork.ApplicationGatewayLoadDistributionAlgorithmRoundRobin),
							LoadDistributionTargets: []*armnetwork.ApplicationGatewayLoadDistributionTarget{
								{
									Name: to.Ptr("<name>"),
									Properties: &armnetwork.ApplicationGatewayLoadDistributionTargetPropertiesFormat{
										BackendAddressPool: &armnetwork.SubResource{
											ID: to.Ptr("<id>"),
										},
										WeightPerServer: to.Ptr[int32](40),
									},
								},
								{
									Name: to.Ptr("<name>"),
									Properties: &armnetwork.ApplicationGatewayLoadDistributionTargetPropertiesFormat{
										BackendAddressPool: &armnetwork.SubResource{
											ID: to.Ptr("<id>"),
										},
										WeightPerServer: to.Ptr[int32](60),
									},
								}},
						},
					}},
				RequestRoutingRules: []*armnetwork.ApplicationGatewayRequestRoutingRule{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ApplicationGatewayRequestRoutingRulePropertiesFormat{
							BackendAddressPool: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							BackendHTTPSettings: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							HTTPListener: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							LoadDistributionPolicy: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							Priority: to.Ptr[int32](10),
							RewriteRuleSet: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							RuleType: to.Ptr(armnetwork.ApplicationGatewayRequestRoutingRuleTypeBasic),
						},
					},
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ApplicationGatewayRequestRoutingRulePropertiesFormat{
							HTTPListener: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							Priority: to.Ptr[int32](20),
							RuleType: to.Ptr(armnetwork.ApplicationGatewayRequestRoutingRuleTypePathBasedRouting),
							URLPathMap: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
						},
					}},
				RewriteRuleSets: []*armnetwork.ApplicationGatewayRewriteRuleSet{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ApplicationGatewayRewriteRuleSetPropertiesFormat{
							RewriteRules: []*armnetwork.ApplicationGatewayRewriteRule{
								{
									Name: to.Ptr("<name>"),
									ActionSet: &armnetwork.ApplicationGatewayRewriteRuleActionSet{
										RequestHeaderConfigurations: []*armnetwork.ApplicationGatewayHeaderConfiguration{
											{
												HeaderName:  to.Ptr("<header-name>"),
												HeaderValue: to.Ptr("<header-value>"),
											}},
										ResponseHeaderConfigurations: []*armnetwork.ApplicationGatewayHeaderConfiguration{
											{
												HeaderName:  to.Ptr("<header-name>"),
												HeaderValue: to.Ptr("<header-value>"),
											}},
										URLConfiguration: &armnetwork.ApplicationGatewayURLConfiguration{
											ModifiedPath: to.Ptr("<modified-path>"),
										},
									},
									Conditions: []*armnetwork.ApplicationGatewayRewriteRuleCondition{
										{
											IgnoreCase: to.Ptr(true),
											Negate:     to.Ptr(false),
											Pattern:    to.Ptr("<pattern>"),
											Variable:   to.Ptr("<variable>"),
										}},
									RuleSequence: to.Ptr[int32](102),
								}},
						},
					}},
				SKU: &armnetwork.ApplicationGatewaySKU{
					Name:     to.Ptr(armnetwork.ApplicationGatewaySKUNameStandardV2),
					Capacity: to.Ptr[int32](3),
					Tier:     to.Ptr(armnetwork.ApplicationGatewayTierStandardV2),
				},
				SSLCertificates: []*armnetwork.ApplicationGatewaySSLCertificate{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ApplicationGatewaySSLCertificatePropertiesFormat{
							Data:     to.Ptr("<data>"),
							Password: to.Ptr("<password>"),
						},
					},
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ApplicationGatewaySSLCertificatePropertiesFormat{
							KeyVaultSecretID: to.Ptr("<key-vault-secret-id>"),
						},
					}},
				SSLProfiles: []*armnetwork.ApplicationGatewaySSLProfile{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ApplicationGatewaySSLProfilePropertiesFormat{
							ClientAuthConfiguration: &armnetwork.ApplicationGatewayClientAuthConfiguration{
								VerifyClientCertIssuerDN: to.Ptr(true),
							},
							SSLPolicy: &armnetwork.ApplicationGatewaySSLPolicy{
								CipherSuites: []*armnetwork.ApplicationGatewaySSLCipherSuite{
									to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSECDHERSAWITHAES128CBCSHA256)},
								MinProtocolVersion: to.Ptr(armnetwork.ApplicationGatewaySSLProtocolTLSv11),
								PolicyType:         to.Ptr(armnetwork.ApplicationGatewaySSLPolicyTypeCustom),
							},
							TrustedClientCertificates: []*armnetwork.SubResource{
								{
									ID: to.Ptr("<id>"),
								}},
						},
					}},
				TrustedClientCertificates: []*armnetwork.ApplicationGatewayTrustedClientCertificate{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ApplicationGatewayTrustedClientCertificatePropertiesFormat{
							Data: to.Ptr("<data>"),
						},
					}},
				TrustedRootCertificates: []*armnetwork.ApplicationGatewayTrustedRootCertificate{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ApplicationGatewayTrustedRootCertificatePropertiesFormat{
							Data: to.Ptr("<data>"),
						},
					},
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ApplicationGatewayTrustedRootCertificatePropertiesFormat{
							KeyVaultSecretID: to.Ptr("<key-vault-secret-id>"),
						},
					}},
				URLPathMaps: []*armnetwork.ApplicationGatewayURLPathMap{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ApplicationGatewayURLPathMapPropertiesFormat{
							DefaultBackendAddressPool: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							DefaultBackendHTTPSettings: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							DefaultLoadDistributionPolicy: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							DefaultRewriteRuleSet: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							PathRules: []*armnetwork.ApplicationGatewayPathRule{
								{
									Name: to.Ptr("<name>"),
									Properties: &armnetwork.ApplicationGatewayPathRulePropertiesFormat{
										BackendAddressPool: &armnetwork.SubResource{
											ID: to.Ptr("<id>"),
										},
										BackendHTTPSettings: &armnetwork.SubResource{
											ID: to.Ptr("<id>"),
										},
										LoadDistributionPolicy: &armnetwork.SubResource{
											ID: to.Ptr("<id>"),
										},
										Paths: []*string{
											to.Ptr("/api"),
											to.Ptr("/v1/api")},
										RewriteRuleSet: &armnetwork.SubResource{
											ID: to.Ptr("<id>"),
										},
									},
								}},
						},
					}},
			},
		},
		&armnetwork.ApplicationGatewaysClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
