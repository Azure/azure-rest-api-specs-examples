package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualNetworkGatewayConnectionGet.json
func ExampleVirtualNetworkGatewayConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworkGatewayConnectionsClient().Get(ctx, "rg1", "connS2S", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetworkGatewayConnection = armnetwork.VirtualNetworkGatewayConnection{
	// 	Name: to.Ptr("connS2S"),
	// 	Type: to.Ptr("Microsoft.Network/connections"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/connections/connS2S"),
	// 	Location: to.Ptr("centralus"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 	Properties: &armnetwork.VirtualNetworkGatewayConnectionPropertiesFormat{
	// 		AuthenticationType: to.Ptr(armnetwork.ConnectionAuthenticationTypeCertificate),
	// 		CertificateAuthentication: &armnetwork.CertificateAuthentication{
	// 			InboundAuthCertificateChain: []*string{
	// 				to.Ptr("MIIC+TCCAeGgAwIBAgIQFOJUqDaxV5xJcKpTKO..."),
	// 				to.Ptr("MIIC+TCCAeGgAwIBAgIQPJerInitNblK7yBgkqh...")},
	// 				InboundAuthCertificateSubjectName: to.Ptr("CN=rootCert.com"),
	// 				OutboundAuthCertificate: to.Ptr("https://customerKv.vault.azure.net/Certificates/outBoundcert/Version"),
	// 			},
	// 			ConnectionMode: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionModeDefault),
	// 			ConnectionProtocol: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionProtocolIKEv2),
	// 			ConnectionStatus: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionStatusConnecting),
	// 			ConnectionType: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionTypeIPsec),
	// 			DpdTimeoutSeconds: to.Ptr[int32](30),
	// 			EgressBytesTransferred: to.Ptr[int64](0),
	// 			EgressNatRules: []*armnetwork.SubResource{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule2"),
	// 			}},
	// 			EnableBgp: to.Ptr(false),
	// 			GatewayCustomBgpIPAddresses: []*armnetwork.GatewayCustomBgpIPAddressIPConfiguration{
	// 				{
	// 					CustomBgpIPAddress: to.Ptr("169.254.21.1"),
	// 					IPConfigurationID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/default"),
	// 				},
	// 				{
	// 					CustomBgpIPAddress: to.Ptr("169.254.21.3"),
	// 					IPConfigurationID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/ActiveActive"),
	// 			}},
	// 			IngressBytesTransferred: to.Ptr[int64](0),
	// 			IngressNatRules: []*armnetwork.SubResource{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule1"),
	// 			}},
	// 			IPSecPolicies: []*armnetwork.IPSecPolicy{
	// 			},
	// 			LocalNetworkGateway2: &armnetwork.LocalNetworkGateway{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/localNetworkGateways/localgw"),
	// 				Properties: &armnetwork.LocalNetworkGatewayPropertiesFormat{
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			RoutingWeight: to.Ptr[int32](0),
	// 			TrafficSelectorPolicies: []*armnetwork.TrafficSelectorPolicy{
	// 			},
	// 			TunnelProperties: []*armnetwork.VirtualNetworkGatewayConnectionTunnelProperties{
	// 				{
	// 					BgpPeeringAddress: to.Ptr("10.78.1.17"),
	// 					TunnelIPAddress: to.Ptr("10.78.1.5"),
	// 				},
	// 				{
	// 					BgpPeeringAddress: to.Ptr("10.78.1.20"),
	// 					TunnelIPAddress: to.Ptr("10.78.1.7"),
	// 			}},
	// 			UseLocalAzureIPAddress: to.Ptr(false),
	// 			UsePolicyBasedTrafficSelectors: to.Ptr(false),
	// 			VirtualNetworkGateway1: &armnetwork.VirtualNetworkGateway{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw"),
	// 				Properties: &armnetwork.VirtualNetworkGatewayPropertiesFormat{
	// 				},
	// 			},
	// 		},
	// 	}
}
