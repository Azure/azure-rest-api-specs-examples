package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ApplicationGatewayListAll.json
func ExampleApplicationGatewaysClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationGatewaysClient().NewListAllPager(nil)
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
		// page = armnetwork.ApplicationGatewaysClientListAllResponse{
		// 	ApplicationGatewayListResult: armnetwork.ApplicationGatewayListResult{
		// 		Value: []*armnetwork.ApplicationGateway{
		// 			{
		// 				Name: to.Ptr("appgw"),
		// 				Type: to.Ptr("Microsoft.Network/applicationGateways"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw"),
		// 				Location: to.Ptr("southcentralus"),
		// 				Properties: &armnetwork.ApplicationGatewayPropertiesFormat{
		// 					AuthenticationCertificates: []*armnetwork.ApplicationGatewayAuthenticationCertificate{
		// 					},
		// 					BackendAddressPools: []*armnetwork.ApplicationGatewayBackendAddressPool{
		// 						{
		// 							Name: to.Ptr("appgwpool"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendAddressPools/appgwpool"),
		// 							Properties: &armnetwork.ApplicationGatewayBackendAddressPoolPropertiesFormat{
		// 								BackendAddresses: []*armnetwork.ApplicationGatewayBackendAddress{
		// 								},
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("appgwpool1"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendAddressPools/appgwpool1"),
		// 							Properties: &armnetwork.ApplicationGatewayBackendAddressPoolPropertiesFormat{
		// 								BackendAddresses: []*armnetwork.ApplicationGatewayBackendAddress{
		// 									{
		// 										IPAddress: to.Ptr("10.0.0.1"),
		// 									},
		// 									{
		// 										IPAddress: to.Ptr("10.0.0.2"),
		// 									},
		// 								},
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							},
		// 						},
		// 					},
		// 					BackendHTTPSettingsCollection: []*armnetwork.ApplicationGatewayBackendHTTPSettings{
		// 						{
		// 							Name: to.Ptr("appgwbhs"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendHttpSettingsCollection/appgwbhs"),
		// 							Properties: &armnetwork.ApplicationGatewayBackendHTTPSettingsPropertiesFormat{
		// 								CookieBasedAffinity: to.Ptr(armnetwork.ApplicationGatewayCookieBasedAffinityDisabled),
		// 								Port: to.Ptr[int32](80),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								RequestTimeout: to.Ptr[int32](30),
		// 								Protocol: to.Ptr(armnetwork.ApplicationGatewayProtocolHTTP),
		// 							},
		// 						},
		// 					},
		// 					EntraJWTValidationConfigs: []*armnetwork.ApplicationGatewayEntraJWTValidationConfig{
		// 						{
		// 							Name: to.Ptr("entraJWTValidationConfig1"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/entraJWTValidationConfigs/entraJWTValidationConfig1"),
		// 							Properties: &armnetwork.ApplicationGatewayEntraJWTValidationConfigPropertiesFormat{
		// 								ClientID: to.Ptr("37293f5a-97b3-451d-b786-f532d711c9ff"),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								TenantID: to.Ptr("70a036f6-8e4d-4615-bad6-149c02e7720d"),
		// 								UnAuthorizedRequestAction: to.Ptr(armnetwork.ApplicationGatewayUnAuthorizedRequestActionDeny),
		// 							},
		// 						},
		// 					},
		// 					FrontendIPConfigurations: []*armnetwork.ApplicationGatewayFrontendIPConfiguration{
		// 						{
		// 							Name: to.Ptr("appgwfip"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendIPConfigurations/appgwfip"),
		// 							Properties: &armnetwork.ApplicationGatewayFrontendIPConfigurationPropertiesFormat{
		// 								PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								PublicIPAddress: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/appgwpip"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					FrontendPorts: []*armnetwork.ApplicationGatewayFrontendPort{
		// 						{
		// 							Name: to.Ptr("appgwfp"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendPorts/appgwfp"),
		// 							Properties: &armnetwork.ApplicationGatewayFrontendPortPropertiesFormat{
		// 								Port: to.Ptr[int32](443),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("appgwfp80"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendPorts/appgwfp80"),
		// 							Properties: &armnetwork.ApplicationGatewayFrontendPortPropertiesFormat{
		// 								Port: to.Ptr[int32](80),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							},
		// 						},
		// 					},
		// 					GatewayIPConfigurations: []*armnetwork.ApplicationGatewayIPConfiguration{
		// 						{
		// 							Name: to.Ptr("appgwipc"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/gatewayIPConfigurations/appgwipc"),
		// 							Properties: &armnetwork.ApplicationGatewayIPConfigurationPropertiesFormat{
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								Subnet: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/virtualNetwork1/subnets/appgwsubnet"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					GlobalConfiguration: &armnetwork.ApplicationGatewayGlobalConfiguration{
		// 						EnableRequestBuffering: to.Ptr(true),
		// 						EnableResponseBuffering: to.Ptr(true),
		// 					},
		// 					HTTPListeners: []*armnetwork.ApplicationGatewayHTTPListener{
		// 						{
		// 							Name: to.Ptr("appgwhl"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/httpListeners/appgwhl"),
		// 							Properties: &armnetwork.ApplicationGatewayHTTPListenerPropertiesFormat{
		// 								FrontendIPConfiguration: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendIPConfigurations/appgwfip"),
		// 								},
		// 								FrontendPort: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendPorts/appgwfp"),
		// 								},
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								RequireServerNameIndication: to.Ptr(false),
		// 								SSLCertificate: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/sslCertificates/sslcert"),
		// 								},
		// 								SSLProfile: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/sslProfiles/sslProfile1"),
		// 								},
		// 								Protocol: to.Ptr(armnetwork.ApplicationGatewayProtocolHTTPS),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("appgwhttplistener"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/httpListeners/appgwhttplistener"),
		// 							Properties: &armnetwork.ApplicationGatewayHTTPListenerPropertiesFormat{
		// 								FrontendIPConfiguration: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendIPConfigurations/appgwfip"),
		// 								},
		// 								FrontendPort: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendPorts/appgwfp80"),
		// 								},
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								Protocol: to.Ptr(armnetwork.ApplicationGatewayProtocolHTTP),
		// 							},
		// 						},
		// 					},
		// 					LoadDistributionPolicies: []*armnetwork.ApplicationGatewayLoadDistributionPolicy{
		// 						{
		// 							Name: to.Ptr("ldp1"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/loadDistributionPolicies/ldp1"),
		// 							Properties: &armnetwork.ApplicationGatewayLoadDistributionPolicyPropertiesFormat{
		// 								LoadDistributionAlgorithm: to.Ptr(armnetwork.ApplicationGatewayLoadDistributionAlgorithmRoundRobin),
		// 								LoadDistributionTargets: []*armnetwork.ApplicationGatewayLoadDistributionTarget{
		// 									{
		// 										Name: to.Ptr("ld11"),
		// 										ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/loadDistributionPolicies/ldp1/loadDistributionTargets/ldt1"),
		// 										Properties: &armnetwork.ApplicationGatewayLoadDistributionTargetPropertiesFormat{
		// 											BackendAddressPool: &armnetwork.SubResource{
		// 												ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendAddressPools/appgwpool"),
		// 											},
		// 											WeightPerServer: to.Ptr[int32](40),
		// 										},
		// 									},
		// 									{
		// 										Name: to.Ptr("ld11"),
		// 										ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/loadDistributionPolicies/ldp1/loadDistributionTargets/ldt1"),
		// 										Properties: &armnetwork.ApplicationGatewayLoadDistributionTargetPropertiesFormat{
		// 											BackendAddressPool: &armnetwork.SubResource{
		// 												ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendAddressPools/appgwpool1"),
		// 											},
		// 											WeightPerServer: to.Ptr[int32](60),
		// 										},
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 					OperationalState: to.Ptr(armnetwork.ApplicationGatewayOperationalStateRunning),
		// 					PrivateEndpointConnections: []*armnetwork.ApplicationGatewayPrivateEndpointConnection{
		// 					},
		// 					PrivateLinkConfigurations: []*armnetwork.ApplicationGatewayPrivateLinkConfiguration{
		// 						{
		// 							Name: to.Ptr("privateLink1"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/privateLinkConfigurations/privateLink1"),
		// 							Properties: &armnetwork.ApplicationGatewayPrivateLinkConfigurationProperties{
		// 								IPConfigurations: []*armnetwork.ApplicationGatewayPrivateLinkIPConfiguration{
		// 									{
		// 										Name: to.Ptr("natNicIpconfig1"),
		// 										ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/privateLinkConfigurations/privateLink1/privateLinkConfigurations/privateLink1/ipConfigurations/natNicIpconfig1"),
		// 										Properties: &armnetwork.ApplicationGatewayPrivateLinkIPConfigurationProperties{
		// 											Primary: to.Ptr(true),
		// 											PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 											ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 											Subnet: &armnetwork.SubResource{
		// 												ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/virtualNetwork1/subnets/appgwsubnet"),
		// 											},
		// 										},
		// 									},
		// 									{
		// 										Name: to.Ptr("natNicIpconfig2"),
		// 										ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/privateLinkConfigurations/privateLink1/privateLinkConfigurations/privateLink1/ipConfigurations/natNicIpconfig2"),
		// 										Properties: &armnetwork.ApplicationGatewayPrivateLinkIPConfigurationProperties{
		// 											PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 											ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 											Subnet: &armnetwork.SubResource{
		// 												ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/virtualNetwork1/subnets/appgwsubnet"),
		// 											},
		// 										},
		// 									},
		// 								},
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							},
		// 						},
		// 					},
		// 					Probes: []*armnetwork.ApplicationGatewayProbe{
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					RequestRoutingRules: []*armnetwork.ApplicationGatewayRequestRoutingRule{
		// 						{
		// 							Name: to.Ptr("appgwrule"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/requestRoutingRules/appgwrule"),
		// 							Properties: &armnetwork.ApplicationGatewayRequestRoutingRulePropertiesFormat{
		// 								BackendAddressPool: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendAddressPools/appgwpool"),
		// 								},
		// 								BackendHTTPSettings: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendHttpSettingsCollection/appgwbhs"),
		// 								},
		// 								EntraJWTValidationConfig: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/entraJWTValidationConfigs/entraJWTValidationConfig1"),
		// 								},
		// 								HTTPListener: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/httpListeners/appgwhl"),
		// 								},
		// 								LoadDistributionPolicy: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/loadDistributionPolicies/ldp1"),
		// 								},
		// 								Priority: to.Ptr[int32](10),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								RewriteRuleSet: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/rewriteRuleSets/rewriteRuleSet1"),
		// 								},
		// 								RuleType: to.Ptr(armnetwork.ApplicationGatewayRequestRoutingRuleTypeBasic),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("appgwPathBasedRule"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/requestRoutingRules/appgwPathBasedRule"),
		// 							Properties: &armnetwork.ApplicationGatewayRequestRoutingRulePropertiesFormat{
		// 								HTTPListener: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/httpListeners/appgwhttplistener"),
		// 								},
		// 								Priority: to.Ptr[int32](20),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								RuleType: to.Ptr(armnetwork.ApplicationGatewayRequestRoutingRuleTypePathBasedRouting),
		// 								URLPathMap: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/urlPathMaps/pathMap1"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					RewriteRuleSets: []*armnetwork.ApplicationGatewayRewriteRuleSet{
		// 						{
		// 							Name: to.Ptr("rewriteRuleSet1"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/rewriteRuleSets/rewriteRuleSet1"),
		// 							Properties: &armnetwork.ApplicationGatewayRewriteRuleSetPropertiesFormat{
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								RewriteRules: []*armnetwork.ApplicationGatewayRewriteRule{
		// 									{
		// 										Name: to.Ptr("Set X-Forwarded-For"),
		// 										ActionSet: &armnetwork.ApplicationGatewayRewriteRuleActionSet{
		// 											RequestHeaderConfigurations: []*armnetwork.ApplicationGatewayHeaderConfiguration{
		// 												{
		// 													HeaderName: to.Ptr("X-Forwarded-For"),
		// 													HeaderValue: to.Ptr("{var_remote-addr}"),
		// 												},
		// 											},
		// 											ResponseHeaderConfigurations: []*armnetwork.ApplicationGatewayHeaderConfiguration{
		// 												{
		// 													HeaderName: to.Ptr("Strict-Transport-Security"),
		// 													HeaderValue: to.Ptr("max-age=31536000"),
		// 												},
		// 											},
		// 											URLConfiguration: &armnetwork.ApplicationGatewayURLConfiguration{
		// 												ModifiedPath: to.Ptr("/abc"),
		// 												ModifiedQueryString: to.Ptr("x=y&a=b"),
		// 												Reroute: to.Ptr(false),
		// 											},
		// 										},
		// 										Conditions: []*armnetwork.ApplicationGatewayRewriteRuleCondition{
		// 											{
		// 												IgnoreCase: to.Ptr(true),
		// 												Negate: to.Ptr(false),
		// 												Pattern: to.Ptr("^Bearer"),
		// 												Variable: to.Ptr("http_req_Authorization"),
		// 											},
		// 										},
		// 										RuleSequence: to.Ptr[int32](102),
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 					SKU: &armnetwork.ApplicationGatewaySKU{
		// 						Name: to.Ptr(armnetwork.ApplicationGatewaySKUNameStandardMedium),
		// 						Capacity: to.Ptr[int32](3),
		// 						Tier: to.Ptr(armnetwork.ApplicationGatewayTierStandard),
		// 					},
		// 					SSLCertificates: []*armnetwork.ApplicationGatewaySSLCertificate{
		// 						{
		// 							Name: to.Ptr("sslcert"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/sslCertificates/sslcert"),
		// 							Properties: &armnetwork.ApplicationGatewaySSLCertificatePropertiesFormat{
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								PublicCertData: to.Ptr("*****"),
		// 							},
		// 						},
		// 					},
		// 					SSLProfiles: []*armnetwork.ApplicationGatewaySSLProfile{
		// 						{
		// 							Name: to.Ptr("sslProfile1"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/sslProfiles/sslProfile1"),
		// 							Properties: &armnetwork.ApplicationGatewaySSLProfilePropertiesFormat{
		// 								ClientAuthConfiguration: &armnetwork.ApplicationGatewayClientAuthConfiguration{
		// 									VerifyClientAuthMode: to.Ptr(armnetwork.ApplicationGatewayClientAuthVerificationModesStrict),
		// 									VerifyClientCertIssuerDN: to.Ptr(true),
		// 									VerifyClientRevocation: to.Ptr(armnetwork.ApplicationGatewayClientRevocationOptionsOCSP),
		// 								},
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								SSLPolicy: &armnetwork.ApplicationGatewaySSLPolicy{
		// 									CipherSuites: []*armnetwork.ApplicationGatewaySSLCipherSuite{
		// 										to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSECDHERSAWITHAES128CBCSHA256),
		// 									},
		// 									MinProtocolVersion: to.Ptr(armnetwork.ApplicationGatewaySSLProtocolTLSv11),
		// 									PolicyType: to.Ptr(armnetwork.ApplicationGatewaySSLPolicyTypeCustom),
		// 								},
		// 								TrustedClientCertificates: []*armnetwork.SubResource{
		// 									{
		// 										ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/trustedClientCertificates/clientcert"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 					TrustedClientCertificates: []*armnetwork.ApplicationGatewayTrustedClientCertificate{
		// 						{
		// 							Name: to.Ptr("clientcert"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/trustedClientCertificates/clientcert"),
		// 							Properties: &armnetwork.ApplicationGatewayTrustedClientCertificatePropertiesFormat{
		// 								ClientCertIssuerDN: to.Ptr("CN=User1, OU=Eng, O=Company Ltd, L=D4, S=Arizona, C=US"),
		// 								Data: to.Ptr("****"),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								ValidatedCertData: to.Ptr("****"),
		// 							},
		// 						},
		// 					},
		// 					URLPathMaps: []*armnetwork.ApplicationGatewayURLPathMap{
		// 						{
		// 							Name: to.Ptr("pathMap1"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/urlPathMaps/pathMap1"),
		// 							Properties: &armnetwork.ApplicationGatewayURLPathMapPropertiesFormat{
		// 								DefaultBackendAddressPool: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendAddressPools/appgwpool"),
		// 								},
		// 								DefaultBackendHTTPSettings: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendHttpSettingsCollection/appgwbhs"),
		// 								},
		// 								DefaultLoadDistributionPolicy: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/loadDistributionPolicies/ldp1"),
		// 								},
		// 								DefaultRewriteRuleSet: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/rewriteRuleSets/rewriteRuleSet1"),
		// 								},
		// 								PathRules: []*armnetwork.ApplicationGatewayPathRule{
		// 									{
		// 										Name: to.Ptr("apiPaths"),
		// 										ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/urlPathMaps/pathMap1/pathRules/apiPaths"),
		// 										Properties: &armnetwork.ApplicationGatewayPathRulePropertiesFormat{
		// 											BackendAddressPool: &armnetwork.SubResource{
		// 												ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendAddressPools/appgwpool"),
		// 											},
		// 											BackendHTTPSettings: &armnetwork.SubResource{
		// 												ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendHttpSettingsCollection/appgwbhs"),
		// 											},
		// 											LoadDistributionPolicy: &armnetwork.SubResource{
		// 												ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/loadDistributionPolicies/ldp1"),
		// 											},
		// 											Paths: []*string{
		// 												to.Ptr("/api"),
		// 												to.Ptr("/v1/api"),
		// 											},
		// 											ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 											RewriteRuleSet: &armnetwork.SubResource{
		// 												ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/rewriteRuleSets/rewriteRuleSet1"),
		// 											},
		// 										},
		// 									},
		// 								},
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
