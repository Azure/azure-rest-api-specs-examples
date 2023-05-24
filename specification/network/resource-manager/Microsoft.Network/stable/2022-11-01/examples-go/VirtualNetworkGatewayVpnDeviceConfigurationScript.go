package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/VirtualNetworkGatewayVpnDeviceConfigurationScript.json
func ExampleVirtualNetworkGatewaysClient_VPNDeviceConfigurationScript() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworkGatewaysClient().VPNDeviceConfigurationScript(ctx, "rg1", "vpngw", armnetwork.VPNDeviceScriptParameters{
		DeviceFamily:    to.Ptr("ISR"),
		FirmwareVersion: to.Ptr("IOS 15.1 (Preview)"),
		Vendor:          to.Ptr("Cisco"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Value = "! Microsoft Corporation\r\n! ---------------------------------------------------------------------------------------------------------------------\r\n! Sample VPN tunnel configuration template for IOS-based devices\r\n!\r\n! This configuration template applies to Cisco VPN devices running IOS 15.1 or beyond (ISR or ASR)\r\n!\r\n\r\n\r\n		\r\n\r\n! ---------------------------------------------------------------------------------------------------------------------\r\n! ACL rules\r\n!\r\n! Some VPN devices require explicit ACL rules to allow cross-premises traffic:\r\n!\r\n! 1. Allow traffic between on premises address ranges and VNet address ranges\r\n! 2. Allow IKE traffic (UDP:500) between on premises VPN devices and Azure VPN gateway\r\n! 3. Allow IPsec traffic (Proto:ESP) between on premises VPN devices and Azure VPN gateway\r\n!\r\n		\r\naccess-list 101 permit ip 10.1.0.0 0.0.255.255 10.0.0.0 0.0.255.255\r\n\r\n! ---------------------------------------------------------------------------------------------------------------------\r\n! Internet Key Exchange (IKE) configuration\r\n!\r\n! This section specifies the authentication, encryption, hashing, and Diffie-Hellman group parameters for IKE\r\n! main mode or phase 1\r\n!\r\n\r\ncrypto ikev2 proposal SwaggerS2S-proposal\r\n  encryption DES3\r\n  integrity  SHA384\r\n  group      DHGroup24\r\n  lifetime   3600\r\n exit\r\n\r\ncrypto ikev2 policy SwaggerS2S-policy\r\n  proposal SwaggerS2S-proposal\r\n  exit\r\n\r\ncrypto ikev2 keyring SwaggerBranch-keyring\r\n		\r\n		peer 52.173.199.254\r\n		address        52.173.199.254\r\n		pre-shared-key lALEHuppeopJmA94exRNiRr2QzuZ6lOsvzu5IlJUEA6LthbTc8g5MTT86MCsGNMzGkTAaLuLnEJoD1Cn4cIlr94qKZm9drsgllzWvsPNezS71stAkaW1Bb7h6GBnDlDP\r\n  exit\r\n\r\ncrypto ikev2 profile  SwaggerS2S-profile\r\n  match address  local 10.3.0.0\r\n	match identity remote address 52.173.199.254 255.255.255.255\r\n		\r\n  authentication remote pre-share\r\n  authentication local  pre-share\r\n  keyring        SwaggerBranch-keyring\r\n  exit\r\n\r\n! ---------------------------------------------------------------------------------------------------------------------\r\n! IPsec configuration\r\n!\r\n! This section specifies encryption, authentication, tunnel mode properties for the Phase 2 negotiation\r\n!\r\ncrypto ipsec transform-set SwaggerS2S-TransformSet DES3 DES3\r\n mode  tunnel\r\n exit\r\n\r\n! ---------------------------------------------------------------------------------------------------------------------\r\n! Crypto map configuration\r\n!\r\n! This section defines a crypto profile that binds the cross-premises network traffic to the IPsec and IKE\r\n! policy profiles for this connection. Then defines the VTI (virtual tunnel interface) with the crypto\r\n! profile. A random interface number (tunnel 1) was used with a random link local address (169.254.0.1/28)\r\n! for the tunnel interface. If either selection is already used in the VPN device, please select another\r\n! interface number or address.  The only requirement is that they must not overlap with another interface\r\n! on the same VPN device.\r\n!\r\ncrypto ipsec profile SwaggerS2S-IPsecProfile\r\n  set transform-set  SwaggerS2S-TransformSet\r\n  set ikev2-profile  SwaggerS2S-profile\r\n  set pfs            None\r\n  set security-association lifetime 3600\r\n exit\r\n\r\n\r\nint tunnel 52.173.199.254\r\n  ip address 169.254.0.1 255.255.255.252\r\n  ip tcp adjust-mss 1350\r\n  tunnel source 10.3.0.0\r\n  tunnel mode ipsec ipv4\r\n  tunnel destination 52.173.199.254\r\n  tunnel protection ipsec profile SwaggerS2S-IPsecProfile\r\n  exit\r\n\r\n	ip route 10.0.0.0 255.255.0.0 tunnel 52.173.199.254 "
}
