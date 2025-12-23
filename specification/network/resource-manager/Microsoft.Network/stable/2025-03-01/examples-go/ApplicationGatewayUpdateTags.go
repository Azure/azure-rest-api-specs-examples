package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ApplicationGatewayUpdateTags.json
func ExampleApplicationGatewaysClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationGatewaysClient().UpdateTags(ctx, "rg1", "AppGw", armnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplicationGateway = armnetwork.ApplicationGateway{
	// 	Name: to.Ptr("AppGw"),
	// 	Type: to.Ptr("Microsoft.Network/applicationGateways"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/AppGw"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armnetwork.ApplicationGatewayPropertiesFormat{
	// 		AuthenticationCertificates: []*armnetwork.ApplicationGatewayAuthenticationCertificate{
	// 		},
	// 		BackendAddressPools: []*armnetwork.ApplicationGatewayBackendAddressPool{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/AppGw/backendAddressPools/Pool01"),
	// 				Name: to.Ptr("Pool01"),
	// 				Properties: &armnetwork.ApplicationGatewayBackendAddressPoolPropertiesFormat{
	// 					BackendAddresses: []*armnetwork.ApplicationGatewayBackendAddress{
	// 						{
	// 							IPAddress: to.Ptr("10.10.10.1"),
	// 						},
	// 						{
	// 							IPAddress: to.Ptr("10.10.10.2"),
	// 						},
	// 						{
	// 							IPAddress: to.Ptr("10.10.10.3"),
	// 					}},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				},
	// 		}},
	// 		BackendHTTPSettingsCollection: []*armnetwork.ApplicationGatewayBackendHTTPSettings{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/AppGw/backendHttpSettingsCollection/PoolSetting01"),
	// 				Name: to.Ptr("PoolSetting01"),
	// 				Properties: &armnetwork.ApplicationGatewayBackendHTTPSettingsPropertiesFormat{
	// 					CookieBasedAffinity: to.Ptr(armnetwork.ApplicationGatewayCookieBasedAffinityDisabled),
	// 					PickHostNameFromBackendAddress: to.Ptr(false),
	// 					Port: to.Ptr[int32](80),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					RequestTimeout: to.Ptr[int32](30),
	// 					Protocol: to.Ptr(armnetwork.ApplicationGatewayProtocolHTTP),
	// 				},
	// 		}},
	// 		FrontendIPConfigurations: []*armnetwork.ApplicationGatewayFrontendIPConfiguration{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/AppGw/frontendIPConfigurations/FrontEndConfig01"),
	// 				Name: to.Ptr("FrontEndConfig01"),
	// 				Properties: &armnetwork.ApplicationGatewayFrontendIPConfigurationPropertiesFormat{
	// 					PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					PublicIPAddress: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/publicIp1"),
	// 					},
	// 				},
	// 		}},
	// 		FrontendPorts: []*armnetwork.ApplicationGatewayFrontendPort{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/AppGw/frontendPorts/FrontEndPort01"),
	// 				Name: to.Ptr("FrontEndPort01"),
	// 				Properties: &armnetwork.ApplicationGatewayFrontendPortPropertiesFormat{
	// 					Port: to.Ptr[int32](80),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				},
	// 		}},
	// 		GatewayIPConfigurations: []*armnetwork.ApplicationGatewayIPConfiguration{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/AppGw/gatewayIPConfigurations/GatewayIp01"),
	// 				Name: to.Ptr("GatewayIp01"),
	// 				Properties: &armnetwork.ApplicationGatewayIPConfigurationPropertiesFormat{
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					Subnet: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet/subnets/subnet1"),
	// 					},
	// 				},
	// 		}},
	// 		HTTPListeners: []*armnetwork.ApplicationGatewayHTTPListener{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/AppGw/httpListeners/listener1"),
	// 				Name: to.Ptr("listener1"),
	// 				Properties: &armnetwork.ApplicationGatewayHTTPListenerPropertiesFormat{
	// 					FrontendIPConfiguration: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/AppGw/frontendIPConfigurations/FrontEndConfig01"),
	// 					},
	// 					FrontendPort: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/AppGw/frontendPorts/FrontEndPort01"),
	// 					},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					RequireServerNameIndication: to.Ptr(false),
	// 					Protocol: to.Ptr(armnetwork.ApplicationGatewayProtocolHTTP),
	// 				},
	// 		}},
	// 		OperationalState: to.Ptr(armnetwork.ApplicationGatewayOperationalStateRunning),
	// 		Probes: []*armnetwork.ApplicationGatewayProbe{
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RedirectConfigurations: []*armnetwork.ApplicationGatewayRedirectConfiguration{
	// 		},
	// 		RequestRoutingRules: []*armnetwork.ApplicationGatewayRequestRoutingRule{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/AppGw/requestRoutingRules/Rule01"),
	// 				Name: to.Ptr("Rule01"),
	// 				Properties: &armnetwork.ApplicationGatewayRequestRoutingRulePropertiesFormat{
	// 					BackendAddressPool: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/AppGw/backendAddressPools/Pool01"),
	// 					},
	// 					BackendHTTPSettings: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/AppGw/backendHttpSettingsCollection/PoolSetting01"),
	// 					},
	// 					HTTPListener: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/AppGw/httpListeners/listener1"),
	// 					},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					RuleType: to.Ptr(armnetwork.ApplicationGatewayRequestRoutingRuleTypeBasic),
	// 				},
	// 		}},
	// 		ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		SKU: &armnetwork.ApplicationGatewaySKU{
	// 			Name: to.Ptr(armnetwork.ApplicationGatewaySKUNameStandardSmall),
	// 			Capacity: to.Ptr[int32](2),
	// 			Tier: to.Ptr(armnetwork.ApplicationGatewayTierStandard),
	// 		},
	// 		SSLCertificates: []*armnetwork.ApplicationGatewaySSLCertificate{
	// 		},
	// 		URLPathMaps: []*armnetwork.ApplicationGatewayURLPathMap{
	// 		},
	// 	},
	// }
}
