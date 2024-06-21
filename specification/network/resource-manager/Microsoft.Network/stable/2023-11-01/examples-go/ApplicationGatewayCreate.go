package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/ApplicationGatewayCreate.json
func ExampleApplicationGatewaysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationGatewaysClient().BeginCreateOrUpdate(ctx, "rg1", "appgw", armnetwork.ApplicationGateway{
		Location: to.Ptr("eastus"),
		Identity: &armnetwork.ManagedServiceIdentity{
			Type: to.Ptr(armnetwork.ResourceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armnetwork.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
				"/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": {},
			},
		},
		Properties: &armnetwork.ApplicationGatewayPropertiesFormat{
			BackendAddressPools: []*armnetwork.ApplicationGatewayBackendAddressPool{
				{
					Name: to.Ptr("appgwpool"),
					Properties: &armnetwork.ApplicationGatewayBackendAddressPoolPropertiesFormat{
						BackendAddresses: []*armnetwork.ApplicationGatewayBackendAddress{
							{
								IPAddress: to.Ptr("10.0.1.1"),
							},
							{
								IPAddress: to.Ptr("10.0.1.2"),
							}},
					},
				},
				{
					ID:   to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendAddressPools/appgwpool1"),
					Name: to.Ptr("appgwpool1"),
					Properties: &armnetwork.ApplicationGatewayBackendAddressPoolPropertiesFormat{
						BackendAddresses: []*armnetwork.ApplicationGatewayBackendAddress{
							{
								IPAddress: to.Ptr("10.0.0.1"),
							},
							{
								IPAddress: to.Ptr("10.0.0.2"),
							}},
					},
				}},
			BackendHTTPSettingsCollection: []*armnetwork.ApplicationGatewayBackendHTTPSettings{
				{
					Name: to.Ptr("appgwbhs"),
					Properties: &armnetwork.ApplicationGatewayBackendHTTPSettingsPropertiesFormat{
						CookieBasedAffinity: to.Ptr(armnetwork.ApplicationGatewayCookieBasedAffinityDisabled),
						Port:                to.Ptr[int32](80),
						RequestTimeout:      to.Ptr[int32](30),
						Protocol:            to.Ptr(armnetwork.ApplicationGatewayProtocolHTTP),
					},
				}},
			FrontendIPConfigurations: []*armnetwork.ApplicationGatewayFrontendIPConfiguration{
				{
					Name: to.Ptr("appgwfip"),
					Properties: &armnetwork.ApplicationGatewayFrontendIPConfigurationPropertiesFormat{
						PublicIPAddress: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/appgwpip"),
						},
					},
				}},
			FrontendPorts: []*armnetwork.ApplicationGatewayFrontendPort{
				{
					Name: to.Ptr("appgwfp"),
					Properties: &armnetwork.ApplicationGatewayFrontendPortPropertiesFormat{
						Port: to.Ptr[int32](443),
					},
				},
				{
					Name: to.Ptr("appgwfp80"),
					Properties: &armnetwork.ApplicationGatewayFrontendPortPropertiesFormat{
						Port: to.Ptr[int32](80),
					},
				}},
			GatewayIPConfigurations: []*armnetwork.ApplicationGatewayIPConfiguration{
				{
					Name: to.Ptr("appgwipc"),
					Properties: &armnetwork.ApplicationGatewayIPConfigurationPropertiesFormat{
						Subnet: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet/subnets/appgwsubnet"),
						},
					},
				}},
			GlobalConfiguration: &armnetwork.ApplicationGatewayGlobalConfiguration{
				EnableRequestBuffering:  to.Ptr(true),
				EnableResponseBuffering: to.Ptr(true),
			},
			HTTPListeners: []*armnetwork.ApplicationGatewayHTTPListener{
				{
					Name: to.Ptr("appgwhl"),
					Properties: &armnetwork.ApplicationGatewayHTTPListenerPropertiesFormat{
						FrontendIPConfiguration: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendIPConfigurations/appgwfip"),
						},
						FrontendPort: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendPorts/appgwfp"),
						},
						RequireServerNameIndication: to.Ptr(false),
						SSLCertificate: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/sslCertificates/sslcert"),
						},
						SSLProfile: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/sslProfiles/sslProfile1"),
						},
						Protocol: to.Ptr(armnetwork.ApplicationGatewayProtocolHTTPS),
					},
				},
				{
					Name: to.Ptr("appgwhttplistener"),
					Properties: &armnetwork.ApplicationGatewayHTTPListenerPropertiesFormat{
						FrontendIPConfiguration: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendIPConfigurations/appgwfip"),
						},
						FrontendPort: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendPorts/appgwfp80"),
						},
						Protocol: to.Ptr(armnetwork.ApplicationGatewayProtocolHTTP),
					},
				}},
			RequestRoutingRules: []*armnetwork.ApplicationGatewayRequestRoutingRule{
				{
					Name: to.Ptr("appgwrule"),
					Properties: &armnetwork.ApplicationGatewayRequestRoutingRulePropertiesFormat{
						BackendAddressPool: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendAddressPools/appgwpool"),
						},
						BackendHTTPSettings: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendHttpSettingsCollection/appgwbhs"),
						},
						HTTPListener: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/httpListeners/appgwhl"),
						},
						Priority: to.Ptr[int32](10),
						RewriteRuleSet: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/rewriteRuleSets/rewriteRuleSet1"),
						},
						RuleType: to.Ptr(armnetwork.ApplicationGatewayRequestRoutingRuleTypeBasic),
					},
				}},
			RewriteRuleSets: []*armnetwork.ApplicationGatewayRewriteRuleSet{
				{
					Name: to.Ptr("rewriteRuleSet1"),
					Properties: &armnetwork.ApplicationGatewayRewriteRuleSetPropertiesFormat{
						RewriteRules: []*armnetwork.ApplicationGatewayRewriteRule{
							{
								Name: to.Ptr("Set X-Forwarded-For"),
								ActionSet: &armnetwork.ApplicationGatewayRewriteRuleActionSet{
									RequestHeaderConfigurations: []*armnetwork.ApplicationGatewayHeaderConfiguration{
										{
											HeaderName:  to.Ptr("X-Forwarded-For"),
											HeaderValue: to.Ptr("{var_add_x_forwarded_for_proxy}"),
										}},
									ResponseHeaderConfigurations: []*armnetwork.ApplicationGatewayHeaderConfiguration{
										{
											HeaderName:  to.Ptr("Strict-Transport-Security"),
											HeaderValue: to.Ptr("max-age=31536000"),
										}},
									URLConfiguration: &armnetwork.ApplicationGatewayURLConfiguration{
										ModifiedPath: to.Ptr("/abc"),
									},
								},
								Conditions: []*armnetwork.ApplicationGatewayRewriteRuleCondition{
									{
										IgnoreCase: to.Ptr(true),
										Negate:     to.Ptr(false),
										Pattern:    to.Ptr("^Bearer"),
										Variable:   to.Ptr("http_req_Authorization"),
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
					Name: to.Ptr("sslcert"),
					Properties: &armnetwork.ApplicationGatewaySSLCertificatePropertiesFormat{
						Data:     to.Ptr("****"),
						Password: to.Ptr("****"),
					},
				},
				{
					Name: to.Ptr("sslcert2"),
					Properties: &armnetwork.ApplicationGatewaySSLCertificatePropertiesFormat{
						KeyVaultSecretID: to.Ptr("https://kv/secret"),
					},
				}},
			SSLProfiles: []*armnetwork.ApplicationGatewaySSLProfile{
				{
					Name: to.Ptr("sslProfile1"),
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
								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/trustedClientCertificates/clientcert"),
							}},
					},
				}},
			TrustedClientCertificates: []*armnetwork.ApplicationGatewayTrustedClientCertificate{
				{
					Name: to.Ptr("clientcert"),
					Properties: &armnetwork.ApplicationGatewayTrustedClientCertificatePropertiesFormat{
						Data: to.Ptr("****"),
					},
				}},
			TrustedRootCertificates: []*armnetwork.ApplicationGatewayTrustedRootCertificate{
				{
					Name: to.Ptr("rootcert"),
					Properties: &armnetwork.ApplicationGatewayTrustedRootCertificatePropertiesFormat{
						Data: to.Ptr("****"),
					},
				},
				{
					Name: to.Ptr("rootcert1"),
					Properties: &armnetwork.ApplicationGatewayTrustedRootCertificatePropertiesFormat{
						KeyVaultSecretID: to.Ptr("https://kv/secret"),
					},
				}},
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
	// res.ApplicationGateway = armnetwork.ApplicationGateway{
	// 	Name: to.Ptr("appgw"),
	// 	Type: to.Ptr("Microsoft.Network/applicationGateways"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw"),
	// 	Location: to.Ptr("southcentralus"),
	// 	Properties: &armnetwork.ApplicationGatewayPropertiesFormat{
	// 		AuthenticationCertificates: []*armnetwork.ApplicationGatewayAuthenticationCertificate{
	// 		},
	// 		BackendAddressPools: []*armnetwork.ApplicationGatewayBackendAddressPool{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendAddressPools/appgwpool"),
	// 				Name: to.Ptr("appgwpool"),
	// 				Properties: &armnetwork.ApplicationGatewayBackendAddressPoolPropertiesFormat{
	// 					BackendAddresses: []*armnetwork.ApplicationGatewayBackendAddress{
	// 						{
	// 							IPAddress: to.Ptr("10.0.1.1"),
	// 						},
	// 						{
	// 							IPAddress: to.Ptr("10.0.1.2"),
	// 					}},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				},
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendAddressPools/appgwpool1"),
	// 				Name: to.Ptr("appgwpool1"),
	// 				Properties: &armnetwork.ApplicationGatewayBackendAddressPoolPropertiesFormat{
	// 					BackendAddresses: []*armnetwork.ApplicationGatewayBackendAddress{
	// 						{
	// 							IPAddress: to.Ptr("10.0.0.1"),
	// 						},
	// 						{
	// 							IPAddress: to.Ptr("10.0.0.2"),
	// 					}},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				},
	// 		}},
	// 		BackendHTTPSettingsCollection: []*armnetwork.ApplicationGatewayBackendHTTPSettings{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendHttpSettingsCollection/appgwbhs"),
	// 				Name: to.Ptr("appgwbhs"),
	// 				Properties: &armnetwork.ApplicationGatewayBackendHTTPSettingsPropertiesFormat{
	// 					CookieBasedAffinity: to.Ptr(armnetwork.ApplicationGatewayCookieBasedAffinityDisabled),
	// 					Port: to.Ptr[int32](80),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					RequestTimeout: to.Ptr[int32](30),
	// 					Protocol: to.Ptr(armnetwork.ApplicationGatewayProtocolHTTP),
	// 				},
	// 		}},
	// 		FrontendIPConfigurations: []*armnetwork.ApplicationGatewayFrontendIPConfiguration{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendIPConfigurations/appgwfip"),
	// 				Name: to.Ptr("appgwfip"),
	// 				Properties: &armnetwork.ApplicationGatewayFrontendIPConfigurationPropertiesFormat{
	// 					PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					PublicIPAddress: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/appgwpip"),
	// 					},
	// 				},
	// 		}},
	// 		FrontendPorts: []*armnetwork.ApplicationGatewayFrontendPort{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendPorts/appgwfp"),
	// 				Name: to.Ptr("appgwfp"),
	// 				Properties: &armnetwork.ApplicationGatewayFrontendPortPropertiesFormat{
	// 					Port: to.Ptr[int32](443),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				},
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendPorts/appgwfp80"),
	// 				Name: to.Ptr("appgwfp80"),
	// 				Properties: &armnetwork.ApplicationGatewayFrontendPortPropertiesFormat{
	// 					Port: to.Ptr[int32](80),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				},
	// 		}},
	// 		GatewayIPConfigurations: []*armnetwork.ApplicationGatewayIPConfiguration{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/gatewayIPConfigurations/appgwipc"),
	// 				Name: to.Ptr("appgwipc"),
	// 				Properties: &armnetwork.ApplicationGatewayIPConfigurationPropertiesFormat{
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					Subnet: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/virtualNetwork1/subnets/appgwsubnet"),
	// 					},
	// 				},
	// 		}},
	// 		GlobalConfiguration: &armnetwork.ApplicationGatewayGlobalConfiguration{
	// 			EnableRequestBuffering: to.Ptr(true),
	// 			EnableResponseBuffering: to.Ptr(true),
	// 		},
	// 		HTTPListeners: []*armnetwork.ApplicationGatewayHTTPListener{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/httpListeners/appgwhl"),
	// 				Name: to.Ptr("appgwhl"),
	// 				Properties: &armnetwork.ApplicationGatewayHTTPListenerPropertiesFormat{
	// 					FrontendIPConfiguration: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendIPConfigurations/appgwfip"),
	// 					},
	// 					FrontendPort: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendPorts/appgwfp"),
	// 					},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					RequireServerNameIndication: to.Ptr(false),
	// 					SSLCertificate: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/sslCertificates/sslcert"),
	// 					},
	// 					SSLProfile: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/sslProfiles/sslProfile1"),
	// 					},
	// 					Protocol: to.Ptr(armnetwork.ApplicationGatewayProtocolHTTPS),
	// 				},
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/httpListeners/appgwhttplistener"),
	// 				Name: to.Ptr("appgwhttplistener"),
	// 				Properties: &armnetwork.ApplicationGatewayHTTPListenerPropertiesFormat{
	// 					FrontendIPConfiguration: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendIPConfigurations/appgwfip"),
	// 					},
	// 					FrontendPort: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendPorts/appgwfp80"),
	// 					},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					Protocol: to.Ptr(armnetwork.ApplicationGatewayProtocolHTTP),
	// 				},
	// 		}},
	// 		Listeners: []*armnetwork.ApplicationGatewayListener{
	// 		},
	// 		LoadDistributionPolicies: []*armnetwork.ApplicationGatewayLoadDistributionPolicy{
	// 		},
	// 		OperationalState: to.Ptr(armnetwork.ApplicationGatewayOperationalStateRunning),
	// 		PrivateEndpointConnections: []*armnetwork.ApplicationGatewayPrivateEndpointConnection{
	// 		},
	// 		Probes: []*armnetwork.ApplicationGatewayProbe{
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RedirectConfigurations: []*armnetwork.ApplicationGatewayRedirectConfiguration{
	// 		},
	// 		RequestRoutingRules: []*armnetwork.ApplicationGatewayRequestRoutingRule{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/requestRoutingRules/appgwrule"),
	// 				Name: to.Ptr("appgwrule"),
	// 				Properties: &armnetwork.ApplicationGatewayRequestRoutingRulePropertiesFormat{
	// 					BackendAddressPool: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendAddressPools/appgwpool"),
	// 					},
	// 					BackendHTTPSettings: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendHttpSettingsCollection/appgwbhs"),
	// 					},
	// 					HTTPListener: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/httpListeners/appgwhl"),
	// 					},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					RewriteRuleSet: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/rewriteRuleSets/rewriteRuleSet1"),
	// 					},
	// 					RuleType: to.Ptr(armnetwork.ApplicationGatewayRequestRoutingRuleTypeBasic),
	// 				},
	// 		}},
	// 		RewriteRuleSets: []*armnetwork.ApplicationGatewayRewriteRuleSet{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/rewriteRuleSets/rewriteRuleSet1"),
	// 				Name: to.Ptr("rewriteRuleSet1"),
	// 				Properties: &armnetwork.ApplicationGatewayRewriteRuleSetPropertiesFormat{
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					RewriteRules: []*armnetwork.ApplicationGatewayRewriteRule{
	// 						{
	// 							Name: to.Ptr("Set X-Forwarded-For"),
	// 							ActionSet: &armnetwork.ApplicationGatewayRewriteRuleActionSet{
	// 								RequestHeaderConfigurations: []*armnetwork.ApplicationGatewayHeaderConfiguration{
	// 									{
	// 										HeaderName: to.Ptr("X-Forwarded-For"),
	// 										HeaderValue: to.Ptr("{var_remote-addr}"),
	// 								}},
	// 								ResponseHeaderConfigurations: []*armnetwork.ApplicationGatewayHeaderConfiguration{
	// 									{
	// 										HeaderName: to.Ptr("Strict-Transport-Security"),
	// 										HeaderValue: to.Ptr("max-age=31536000"),
	// 								}},
	// 								URLConfiguration: &armnetwork.ApplicationGatewayURLConfiguration{
	// 									ModifiedPath: to.Ptr("/abc"),
	// 									ModifiedQueryString: to.Ptr("x=y&a=b"),
	// 								},
	// 							},
	// 							Conditions: []*armnetwork.ApplicationGatewayRewriteRuleCondition{
	// 								{
	// 									IgnoreCase: to.Ptr(true),
	// 									Negate: to.Ptr(false),
	// 									Pattern: to.Ptr("^Bearer"),
	// 									Variable: to.Ptr("http_req_Authorization"),
	// 							}},
	// 							RuleSequence: to.Ptr[int32](102),
	// 					}},
	// 				},
	// 		}},
	// 		RoutingRules: []*armnetwork.ApplicationGatewayRoutingRule{
	// 		},
	// 		SKU: &armnetwork.ApplicationGatewaySKU{
	// 			Name: to.Ptr(armnetwork.ApplicationGatewaySKUNameStandardMedium),
	// 			Capacity: to.Ptr[int32](3),
	// 			Tier: to.Ptr(armnetwork.ApplicationGatewayTierStandard),
	// 		},
	// 		SSLCertificates: []*armnetwork.ApplicationGatewaySSLCertificate{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/sslCertificates/sslcert"),
	// 				Name: to.Ptr("sslcert"),
	// 				Properties: &armnetwork.ApplicationGatewaySSLCertificatePropertiesFormat{
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					PublicCertData: to.Ptr("*****"),
	// 				},
	// 		}},
	// 		SSLProfiles: []*armnetwork.ApplicationGatewaySSLProfile{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/sslProfiles/sslProfile1"),
	// 				Name: to.Ptr("sslProfile1"),
	// 				Properties: &armnetwork.ApplicationGatewaySSLProfilePropertiesFormat{
	// 					ClientAuthConfiguration: &armnetwork.ApplicationGatewayClientAuthConfiguration{
	// 						VerifyClientCertIssuerDN: to.Ptr(true),
	// 					},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					SSLPolicy: &armnetwork.ApplicationGatewaySSLPolicy{
	// 						CipherSuites: []*armnetwork.ApplicationGatewaySSLCipherSuite{
	// 							to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSECDHERSAWITHAES128CBCSHA256)},
	// 							MinProtocolVersion: to.Ptr(armnetwork.ApplicationGatewaySSLProtocolTLSv11),
	// 							PolicyType: to.Ptr(armnetwork.ApplicationGatewaySSLPolicyTypeCustom),
	// 						},
	// 						TrustedClientCertificates: []*armnetwork.SubResource{
	// 							{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/trustedClientCertificates/clientcert"),
	// 						}},
	// 					},
	// 			}},
	// 			TrustedClientCertificates: []*armnetwork.ApplicationGatewayTrustedClientCertificate{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/trustedClientCertificates/clientcert"),
	// 					Name: to.Ptr("clientcert"),
	// 					Properties: &armnetwork.ApplicationGatewayTrustedClientCertificatePropertiesFormat{
	// 						Data: to.Ptr("****"),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					},
	// 			}},
	// 			URLPathMaps: []*armnetwork.ApplicationGatewayURLPathMap{
	// 			},
	// 		},
	// 	}
}
